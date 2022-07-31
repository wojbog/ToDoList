package server

import (
	"github.com/gofiber/fiber/v2"

	"github.com/wojbog/ToDoList/handlers"
	"github.com/wojbog/ToDoList/service"
)

func Routing(app *fiber.App, s *service.Service) {

	// add TODO endpoint
	app.Post("/todo/", handlers.CreateToDo(s))

	// get all TODOs endpoint
	app.Get("/todo/", handlers.GetTodos(s))

	// update specific TDO endpoint
	app.Put("/todo/:id", handlers.UpdateTodo(s))

	// get specific TODO endpoint
	app.Get("/todo/:id", handlers.GetTodo(s))

	// delete specific TODO endpoint
	app.Delete("/todo/:id", handlers.DeleteTodo(s))

	// set Complete TODO endpoint
	app.Patch("/todo/:id", handlers.SetCompleteTodo(s))

	// set done to TODO endpoint
	app.Patch("/done/todo/:id", handlers.SetDoneTodo(s))

	// get TODOs for today endpoint
	app.Get("/today/todo/", handlers.GetTodoToday(s))

	// get TODOs for tomorrow endpoint
	app.Get("/nextday/todo/", handlers.GetTodoNextDay(s))

	//get TODOs for current week endpoint
	app.Get("/currentweek/todo/", handlers.GetTodoWeek(s))

}
