package main

import (
	"fmt"

	"todolist/database"

	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"todolist/task"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World")
	c.Render("index.html", fiber.Map{})
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/task", task.GetAllTasks)
	app.Get("/api/v1/task/:id", task.GetOneTaskByID)
	app.Post("/api/v1/task", task.InsertTask)
	app.Put("/api/v1/task", task.UpdateTasksByTitle)
	app.Delete("/api/v1/task", task.DeleteTaskByID)
}

const (
	HOST = "0.0.0.0"
	PORT = 5432

	DB_USER     = "postgres"
	DB_PASSWORD = "123456"
	DB_NAME     = "todo"
)

func initializeDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, DB_USER, DB_PASSWORD, DB_NAME)

	database.DBConn, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("connected")
	// create table
	database.DBConn.AutoMigrate(&task.Task{})
	fmt.Println("migrate")
}

func main() {
	app := fiber.New()

	initializeDatabase()

	setupRoutes(app)
	app.Static("/todolist", "./public")

	app.Listen(":3000")

	//TODO
	//defer database.DBConn.Close()
}
