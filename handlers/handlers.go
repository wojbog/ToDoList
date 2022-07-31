package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wojbog/ToDoList/models"
	"github.com/wojbog/ToDoList/service"
)

//add TODO handler
//return model id if successful created else error
func CreateToDo(s service.ServiceFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		todo := new(models.ToDo)

		//parse JSON to TODO struct
		if err := c.BodyParser(todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"error": service.BadValidation.Error(),
				})
		}

		//pass TODO to service
		id, err := s.CreateToDo(*todo)

		//error handler
		if err != nil {
			if err == service.BadValidation {
				return c.Status(fiber.StatusBadRequest).JSON(
					&fiber.Map{
						"error": err.Error(),
					})
			} else {
				return c.Status(fiber.StatusInternalServerError).JSON(
					&fiber.Map{
						"error": err.Error(),
					})
			}
		} else {
			return c.Status(fiber.StatusCreated).JSON(
				&fiber.Map{
					"id": id,
				})
		}
	}

}

//get all TODOs handler
//return count, array of ToDos
func GetTodos(s service.ServiceFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if m, count, err := s.GetAllTodos(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				&fiber.Map{
					"error": err.Error(),
				})
		} else {
			return c.Status(fiber.StatusOK).JSON(
				&fiber.Map{
					"count": count,
					"todo":  m,
				})
		}
	}
}

//get one ToDO handler
//param: model id
//return todo if exist else error
func GetTodo(s service.ServiceFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if id, err := c.ParamsInt("id"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"error": "todo does not exist",
				})
		} else {
			if m, err := s.GetTodo(id); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(
					&fiber.Map{
						"error": err.Error(),
					})
			} else {
				return c.Status(fiber.StatusOK).JSON(
					&fiber.Map{
						"todo": m,
					})
			}
		}
	}
}

//delete ToDO handler
//param: model id
//return: status 204 if deleted successfully or error otherwise
func DeleteTodo(s service.ServiceFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if id, err := c.ParamsInt("id"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"error": "todo does not exist",
				})
		} else {
			if err := s.DeleteTodo(id); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(
					&fiber.Map{
						"error": err.Error(),
					})
			} else {
				return c.SendStatus(fiber.StatusNoContent)
			}
		}
	}
}

//get ToDos for today handler
//return: count, array of ToDos
func GetTodoToday(s service.ServiceFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if m, count, err := s.GetIncTodo(0); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"error": err.Error(),
				})
		} else {
			return c.Status(fiber.StatusOK).JSON(
				&fiber.Map{
					"count": count,
					"todo":  m,
				})
		}
	}
}

//get ToDos for tomorrow handler
//return: count, array of ToDos
func GetTodoNextDay(s service.ServiceFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if m, count, err := s.GetIncTodo(1); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"error": err.Error(),
				})
		} else {
			return c.Status(fiber.StatusOK).JSON(
				&fiber.Map{
					"count": count,
					"todo":  m,
				})
		}
	}
}

//get ToDos for tomorrow handler
//return: count, array of ToDos
func GetTodoWeek(s service.ServiceFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if m, count, err := s.GetIncTodo(2); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"error": err.Error(),
				})
		} else {
			return c.Status(fiber.StatusOK).JSON(
				&fiber.Map{
					"count": count,
					"todo":  m,
				})
		}
	}
}

//update ToDo handler
//param: model id
//return: model id
func UpdateTodo(s service.ServiceFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if id, err := c.ParamsInt("id"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"error": "todo does not exist",
				})
		} else {
			todo := new(models.ToDo)

			//parse JSON to TODO struct
			if err := c.BodyParser(todo); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(
					&fiber.Map{
						"error": service.BadValidation.Error(),
					})
			}

			//pass TODO to service
			todo.ID = uint(id)
			id, err := s.UpdateToDo(*todo)

			//error handler
			if err != nil {
				if err == service.BadValidation {
					return c.Status(fiber.StatusBadRequest).JSON(
						&fiber.Map{
							"error": err.Error(),
						})
				} else {
					return c.Status(fiber.StatusInternalServerError).JSON(
						&fiber.Map{
							"error": err.Error(),
						})
				}
			} else {
				return c.Status(fiber.StatusOK).JSON(
					&fiber.Map{
						"id": id,
					})
			}
		}
	}
}

//set complete to ToDO handler
//param: model id
//return model id
func SetCompleteTodo(s service.ServiceFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if id, err := c.ParamsInt("id"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"error": "todo does not exist",
				})
		} else {
			todoComp := new(models.ToDoComp)
			todoID := new(models.ToDo)

			//parse JSON to TODO struct
			if err := c.BodyParser(todoComp); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(
					&fiber.Map{
						"error": service.BadValidation.Error(),
					})
			}

			//pass TODO to service
			todoID.ID = uint(id)
			id, err := s.SetCompTodo(*todoID, *todoComp)

			//error handler
			if err != nil {
				if err == service.BadValidation {
					return c.Status(fiber.StatusBadRequest).JSON(
						&fiber.Map{
							"error": err.Error(),
						})
				} else {
					return c.Status(fiber.StatusInternalServerError).JSON(
						&fiber.Map{
							"error": err.Error(),
						})
				}
			} else {
				return c.Status(fiber.StatusOK).JSON(
					&fiber.Map{
						"id": id,
					})
			}
		}
	}
}

//set ToDO as Done handler
//parma: model id
//return: model id
func SetDoneTodo(s service.ServiceFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if id, err := c.ParamsInt("id"); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				&fiber.Map{
					"error": "todo does not exist",
				})
		} else {
			todoID := new(models.ToDo)
			todoComp := new(models.ToDoComp)

			//pass TODO to service
			todoID.ID = uint(id)
			todoComp.Complete = 100
			id, err := s.SetCompTodo(*todoID, *todoComp)

			//error handler
			if err != nil {
				if err == service.BadValidation {
					return c.Status(fiber.StatusBadRequest).JSON(
						&fiber.Map{
							"error": err.Error(),
						})
				} else {
					return c.Status(fiber.StatusInternalServerError).JSON(
						&fiber.Map{
							"error": err.Error(),
						})
				}
			} else {
				return c.Status(fiber.StatusOK).JSON(
					&fiber.Map{
						"id": id,
					})
			}
		}
	}
}
