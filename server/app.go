package main

import (
	"github.com/CMOISDEAD/todo_go_svelte/database"
	"github.com/CMOISDEAD/todo_go_svelte/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Post("/create", routes.CreateTask)
	app.Get("/getTask/:id", routes.GetTask)
	app.Get("/getAll", routes.GetAllTasks)
	app.Delete("/removeTask/:id", routes.RemoveTask)
	app.Put("/updateTask/:id", routes.UpdateTask)
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	setUpRoutes(app)
	
	app.Use(cors.New())

	app.Listen(":8080")
}
