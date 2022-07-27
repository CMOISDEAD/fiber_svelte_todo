package routes

import (
	"github.com/CMOISDEAD/todo_go_svelte/database"
	"github.com/CMOISDEAD/todo_go_svelte/models"
	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Create(&task)

	return c.Status(200).JSON(task)
}

func RemoveTask(c *fiber.Ctx) error {
	task := []models.Task{}

	database.DBConn.Delete(&task, c.Params("id"))

	return c.Status(200).JSON(task)
}

func GetAllTasks(c *fiber.Ctx) error {
	tasks := []models.Task{}

	database.DBConn.Find(&tasks)

	return c.Status(200).JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	task := []models.Task{}

	if err := database.DBConn.First(&task, c.Params("id")); err != nil {
		return c.Status(404).JSON(task)
	}

	return c.Status(200).JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	task := []models.Task{}
	newTask := new(models.Task)

	database.DBConn.First(&task, c.Params("id"))

	if err := c.BodyParser(newTask); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Model(&task).Select("Title", "Description", "Completed").Updates(models.Task{Title: newTask.Title, Description: newTask.Description, Completed: newTask.Completed})

	return c.Status(400).JSON("Updated")
}
