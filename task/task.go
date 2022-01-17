package task

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"

	"todolist/database"
)

type Task struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

type TaskFromUI struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetAllTasks(c *fiber.Ctx) {
	db := database.DBConn
	var tasks []Task
	db.Find(&tasks)

	c.Status(200).JSON(tasks)
}

func GetOneTaskByID(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var task Task
	db.First(&task, id)
	if task.Title == "" {
		return
	}
	c.Send(task)
}

func InsertTask(c *fiber.Ctx) {
	taskfromui := new(TaskFromUI)
	c.BodyParser(taskfromui)

	db := database.DBConn
	var task Task
	task.Title = taskfromui.Title
	task.Body = taskfromui.Body
	db.Create(&task)
	c.Status(200).JSON(task)
}

func UpdateTasksByTitle(c *fiber.Ctx) {
	taskfromui := new(TaskFromUI)
	c.BodyParser(taskfromui)

	db := database.DBConn
	var task Task
	db.First(&task, "id = ?", taskfromui.ID)
	if task.Title == "" {
		return
	}
	task.Title = taskfromui.Title
	task.Body = taskfromui.Body
	db.Updates(task)
	c.Status(200).Send(task)
}

func DeleteTaskByID(c *fiber.Ctx) {
	taskfromui := new(TaskFromUI)
	c.BodyParser(taskfromui)

	db := database.DBConn
	var task Task
	db.First(&task, "id = ?", taskfromui.ID)
	if task.Title == "" {
		return
	}
	db.Delete(&task)
	c.Status(200).Send("Deleted")
}
