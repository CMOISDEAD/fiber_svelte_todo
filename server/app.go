package main

import (
	"github.com/CMOISDEAD/todo_go_svelte/database"
	"github.com/CMOISDEAD/todo_go_svelte/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	jwtware "github.com/gofiber/jwt/v3"
)

func setUpAuthRoutes(app *fiber.App) {
	app.Post("/register", routes.Register)
	app.Post("/login", routes.Login)
}

func setUpUserRoutes(app *fiber.App) {
	app.Post("/create", routes.CreateTask)
	app.Get("/getTask/:id", routes.GetTask)
	app.Get("/getAll/:id", routes.GetAllTasks)
	app.Delete("/removeTask/:id", routes.RemoveTask)
	app.Put("/updateTask/:id", routes.UpdateTask)
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	setUpAuthRoutes(app)
	
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	
	setUpUserRoutes(app)
	
	app.Use(cors.New())

	app.Listen(":8080")
}
