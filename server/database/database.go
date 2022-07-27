package database

import (
  "log"
	"os"

	"github.com/CMOISDEAD/todo_go_svelte/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var (
  DBConn *gorm.DB 
)

func ConnectDB() {
  
	dsn := "camilo:Imnotafraid31@/todo_go?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.AutoMigrate(&models.Task{})
	DBConn = db
}
