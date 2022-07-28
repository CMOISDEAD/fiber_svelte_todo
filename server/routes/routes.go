package routes

import (
	"time"

	"github.com/CMOISDEAD/todo_go_svelte/database"
	"github.com/CMOISDEAD/todo_go_svelte/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	dbuser := []models.User{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if result := database.DBConn.Where(&models.User{
		Name: user.Name,
	}).First(&dbuser); result.Error == nil {
		return c.Status(400).JSON("name is already taked")
	}

	database.DBConn.Create(&user)

	return c.Status(200).JSON(user)
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)
	dbuser := []models.User{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if result := database.DBConn.Where(&models.User{
		Name:     user.Name,
		Password: user.Password,
	}).First(&dbuser); result.Error != nil {
		return c.Status(404).JSON("User not found")
	}

	claims := jwt.MapClaims{
		"id":   &dbuser[0].ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(fiber.Map{
		"token":     t,
		"user_info": &dbuser,
	})
}

// Restricted routes

func CreateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	id := claims["id"].(float64)

	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	task.UserID = id

	database.DBConn.Create(&task)

	return c.Status(200).JSON(fiber.Map{
		"name":      name,
		"task_info": task,
	})
}

func RemoveTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	task := []models.Task{}

	database.DBConn.Delete(&task, c.Params("id"))

	return c.Status(200).JSON(fiber.Map{
		"name":      name,
		"task_info": "Task with id has been deleted",
	})
}

func GetAllTasks(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	tasks := []models.Task{}

	database.DBConn.Where("user_id = ?", c.Params("id")).Find(&tasks)

	return c.Status(200).JSON(fiber.Map{
		"name":             name,
		"tasks_collection": tasks,
	})
}

func GetTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	id := claims["id"].(float64)

	task := []models.Task{}

	if result := database.DBConn.First(&task, c.Params("id")); result.Error != nil {
		return c.Status(404).JSON("Task dont found")
	}

	return c.Status(200).JSON(fiber.Map{
		"user_id":   id,
		"user_name": name,
		"task_info": task,
	})
}

func UpdateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	task := []models.Task{}
	newTask := new(models.Task)

	database.DBConn.First(&task, c.Params("id"))

	if err := c.BodyParser(newTask); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Model(&task).Select("Title", "Description", "Completed").Updates(models.Task{
		Title:       newTask.Title,
		Description: newTask.Description,
		Completed:   newTask.Completed,
	})

	return c.Status(400).JSON(fiber.Map{
		"name":      name,
		"task_info": "Task updated",
	})
}
