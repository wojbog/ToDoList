package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/wojbog/ToDoList/models"
	database "github.com/wojbog/ToDoList/repository"
	"github.com/wojbog/ToDoList/server"
	"github.com/wojbog/ToDoList/service"
)

func runMigration(db *gorm.DB) {
	var p models.ToDo
	db.AutoMigrate(&p)
}

func main() {
	app := fiber.New()

	databaseName := os.Getenv("MARIADB_DATABASE")
	if databaseName == "" {
		log.Fatal("NO DATABASE NAME")
	}

	databaseUser := os.Getenv("MARIADB_USER")
	if databaseUser == "" {
		log.Fatal("NO DATABASE USER")
	}
	databasePassword := os.Getenv("MARIADB_PASSWORD")
	if databasePassword == "" {
		log.Fatal("NO DATABASE PASSWORD")
	}

	//connect to database
	dsn := databaseUser + ":" + databasePassword + "@tcp(db:3306)/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("CANNOT CONNECT TO DB")
	} else {
		log.Println("connect to DB")
	}

	runMigration(db)

	repo := database.NewRepository(db)
	s := service.NewService(repo)

	server.Routing(app, s)

	app.Listen(":8000")
}
