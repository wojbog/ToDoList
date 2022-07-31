package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"testing"

	"github.com/wojbog/ToDoList/models"
	"github.com/wojbog/ToDoList/service"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func (t *MockService) CreateToDo(m models.ToDo) (uint, error) {
	if m.Title == "title1" {
		return 0, service.BadValidation
	} else if m.Title == "title2" {
		return 0, errors.New("error new")
	} else {
		return 0, nil
	}
}

func (t *MockService) UpdateToDo(m models.ToDo) (uint, error) {
	if m.Title == "title1" {
		return 0, service.BadValidation
	} else if m.Title == "title2" {
		return 0, errors.New("error new")
	} else {
		return 0, nil
	}

}
func (t *MockService) GetAllTodos() ([]models.ToDo, int64, error) {

	return []models.ToDo{}, 0, nil
}

func (t *MockService) GetTodo(id int) (models.ToDo, error) {
	if id == 1 {
		return models.ToDo{}, errors.New("error")
	} else {
		return models.ToDo{}, nil
	}
}
func (t *MockService) DeleteTodo(id int) error {
	if id == 1 {
		return errors.New("error")
	} else {
		return nil
	}
}
func (t *MockService) SetCompTodo(id models.ToDo, m models.ToDoComp) (uint, error) {
	if id.ID == 1 {
		return 0, service.BadValidation
	} else if id.ID == 2 {
		return 0, errors.New("error new")
	} else {
		return 0, nil
	}
}
func (t *MockService) GetIncTodo(mode int) ([]models.ToDo, int64, error) {
	return []models.ToDo{}, 0, nil
}

type MockService struct {
}

func CreateMockService() *MockService {
	return &MockService{}
}

func TestCreateToDo(t *testing.T) {

	s := CreateMockService()

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		m            models.ToDo
	}{
		{
			description:  "bad validation error",
			route:        "/todo/",
			expectedCode: 400,
			m:            models.ToDo{Title: "title1", Description: "opis"},
		},
		{
			description:  "Server Error",
			route:        "/todo/",
			expectedCode: 500,
			m:            models.ToDo{Title: "title2", Description: "opis"},
		},
		{
			description:  "created, no error",
			route:        "/todo/",
			expectedCode: 201,
			m:            models.ToDo{Title: "title3", Description: "opis"},
		},
	}

	app := fiber.New()

	app.Post("/todo/", CreateToDo(s))

	var buf bytes.Buffer
	for _, test := range tests {
		err := json.NewEncoder(&buf).Encode(test.m)

		if err != nil {
			log.Fatal(err)
		}
		req, err := http.NewRequest("POST", test.route, &buf)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
	req, err := http.NewRequest("POST", "/todo/", nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	resp, _ := app.Test(req)
	assert.Equal(t, 400, resp.StatusCode, "bad body")

}

func TestGetTodos(t *testing.T) {
	s := CreateMockService()

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get all ToDos",
			route:        "/todo/",
			expectedCode: 200,
		},
	}

	app := fiber.New()
	for _, test := range tests {
		app.Get("/todo/", GetTodos(s))
		req, err := http.NewRequest("GET", test.route, nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		resp, _ := app.Test(req)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetTodo(t *testing.T) {
	s := CreateMockService()

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		id           string
	}{
		{
			description:  "bad id, id as a string",
			route:        "/todo/",
			expectedCode: 400,
			id:           "hello",
		},
		{
			description:  "todo does not exist",
			route:        "/todo/",
			expectedCode: 400,
			id:           "1",
		},
		{
			description:  "get succesfully",
			route:        "/todo/",
			expectedCode: 200,
			id:           "2",
		},
	}

	app := fiber.New()

	app.Get("/todo/:id", GetTodo(s))

	for _, test := range tests {

		req, err := http.NewRequest("GET", test.route+test.id, nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		resp, _ := app.Test(req)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestDeleteTodo(t *testing.T) {
	s := CreateMockService()

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		id           string
	}{
		{
			description:  "bad id, id as a string",
			route:        "/todo/",
			expectedCode: 400,
			id:           "hello",
		},
		{
			description:  "todo does not exist",
			route:        "/todo/",
			expectedCode: 400,
			id:           "1",
		},
		{
			description:  "get succesfully",
			route:        "/todo/",
			expectedCode: 204,
			id:           "2",
		},
	}

	app := fiber.New()

	app.Delete("/todo/:id", DeleteTodo(s))

	for _, test := range tests {

		req, err := http.NewRequest("DELETE", test.route+test.id, nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		resp, _ := app.Test(req)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetTodoToday(t *testing.T) {
	s := CreateMockService()

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get for today",
			route:        "/today/todo/",
			expectedCode: 200,
		},
	}

	app := fiber.New()
	app.Get("/today/todo/", GetTodoToday(s))
	for _, test := range tests {

		req, err := http.NewRequest("GET", test.route, nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		resp, _ := app.Test(req)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetTodoNextDay(t *testing.T) {
	s := CreateMockService()

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get for tomorrow",
			route:        "/nextday/todo/",
			expectedCode: 200,
		},
	}

	app := fiber.New()
	app.Get("/nextday/todo/", GetTodoNextDay(s))
	for _, test := range tests {

		req, err := http.NewRequest("GET", test.route, nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		resp, _ := app.Test(req)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetTodoWeek(t *testing.T) {
	s := CreateMockService()

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get for current week",
			route:        "/currentweek/todo/",
			expectedCode: 200,
		},
	}

	app := fiber.New()
	app.Get("/currentweek/todo/", GetTodoWeek(s))
	for _, test := range tests {

		req, err := http.NewRequest("GET", test.route, nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		resp, _ := app.Test(req)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestUpdateTodo(t *testing.T) {
	s := CreateMockService()

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		m            models.ToDo
		id           string
	}{
		{
			description:  "bad id",
			route:        "/todo/",
			expectedCode: 400,
			id:           "hello",
		},
		{
			description:  "bad request",
			route:        "/todo/",
			expectedCode: 400,
			m:            models.ToDo{Title: "title1", Description: "opis"},
			id:           "1",
		},
		{
			description:  "Server Error",
			route:        "/todo/",
			expectedCode: 500,
			m:            models.ToDo{Title: "title2", Description: "opis"},
			id:           "1",
		},
		{
			description:  "updated succesfully",
			route:        "/todo/",
			expectedCode: 200,
			m:            models.ToDo{Title: "title3", Description: "opis"},
			id:           "1",
		},
	}

	app := fiber.New()

	app.Put("/todo/:id", UpdateTodo(s))

	var buf bytes.Buffer
	for _, test := range tests {
		err := json.NewEncoder(&buf).Encode(test.m)

		if err != nil {
			log.Fatal(err)
		}
		req, err := http.NewRequest("PUT", test.route+test.id, &buf)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
	req, err := http.NewRequest("PUT", "/todo/2", nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	resp, _ := app.Test(req)
	assert.Equal(t, 400, resp.StatusCode, "bad body")

}

func TestSetCompleteTodo(t *testing.T) {
	s := CreateMockService()

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		m            models.ToDoComp
		id           string
	}{
		{
			description:  "bad id",
			route:        "/todo/",
			expectedCode: 400,
			id:           "hello",
		},
		{
			description:  "bad request",
			route:        "/todo/",
			expectedCode: 400,
			m:            models.ToDoComp{Complete: 15},
			id:           "1",
		},
		{
			description:  "Server Error",
			route:        "/todo/",
			expectedCode: 500,
			m:            models.ToDoComp{Complete: 15},
			id:           "2",
		},
		{
			description:  "set Complete succesfully",
			route:        "/todo/",
			expectedCode: 200,
			m:            models.ToDoComp{Complete: 15},
			id:           "3",
		},
	}

	app := fiber.New()

	app.Patch("/todo/:id", SetCompleteTodo(s))

	var buf bytes.Buffer
	for _, test := range tests {
		err := json.NewEncoder(&buf).Encode(test.m)

		if err != nil {
			log.Fatal(err)
		}
		req, err := http.NewRequest("PATCH", test.route+test.id, &buf)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
	req, err := http.NewRequest("PATCH", "/todo/7", nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	resp, _ := app.Test(req)
	assert.Equal(t, 400, resp.StatusCode, "bad body")
}

func TestSetDoneTodo(t *testing.T) {
	s := CreateMockService()

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		m            models.ToDoComp
		id           string
	}{
		{
			description:  "bad id",
			route:        "/done/todo/",
			expectedCode: 400,
			id:           "hello",
		},
		{
			description:  "bad request",
			route:        "/done/todo/",
			expectedCode: 400,
			m:            models.ToDoComp{Complete: 15},
			id:           "1",
		},
		{
			description:  "Server Error",
			route:        "/done/todo/",
			expectedCode: 500,
			m:            models.ToDoComp{Complete: 15},
			id:           "2",
		},
		{
			description:  "set Done succesfully",
			route:        "/done/todo/",
			expectedCode: 200,
			m:            models.ToDoComp{Complete: 15},
			id:           "3",
		},
	}

	app := fiber.New()

	app.Patch("/done/todo/:id", SetDoneTodo(s))

	for _, test := range tests {
		req, err := http.NewRequest("PATCH", test.route+test.id, nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
