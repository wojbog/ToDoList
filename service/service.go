package service

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/wojbog/ToDoList/models"
	database "github.com/wojbog/ToDoList/repository"
)

//error variuables
var (
	BadValidation = errors.New("data must be a valid")
)

type ServiceFunc interface {
	CreateToDo(m models.ToDo) (uint, error)
	UpdateToDo(m models.ToDo) (uint, error)
	GetAllTodos() ([]models.ToDo, int64, error)
	GetTodo(id int) (models.ToDo, error)
	DeleteTodo(id int) error
	SetCompTodo(id models.ToDo, m models.ToDoComp) (uint, error)
	GetIncTodo(mode int) ([]models.ToDo, int64, error)
}

//service instance
type Service struct {
	db database.Repository
}

//ServiceDB constructor
func NewService(db database.Repository) *Service {
	return &Service{db}
}

//create TODO func
//validation TODO struct
//parma: model ToDo
//return model id and error
func (s *Service) CreateToDo(m models.ToDo) (uint, error) {
	validate := validator.New()
	if err := validate.Struct(m); err != nil {
		return 0, BadValidation
	}
	id, err := s.db.AddToDo(m)
	return id, err
}

//Update TODO func
//validation TODO struct
//param: model ToDo
//return model id and error
func (s *Service) UpdateToDo(m models.ToDo) (uint, error) {
	validate := validator.New()
	if err := validate.Struct(m); err != nil {
		return 0, BadValidation
	}
	id, err := s.db.UpdateToDo(m)
	return id, err
}

//get all Todos service
//return array of Todos, count, error
func (s *Service) GetAllTodos() ([]models.ToDo, int64, error) {
	m, count, err := s.db.GetAllTodos()

	return m, count, err
}

//get one Todo
//param: model id
//return model TODo, error
func (s *Service) GetTodo(id int) (models.ToDo, error) {
	m, err := s.db.GetTodo(id)

	return m, err
}

//delete one Todo
//param: model id
//return error
func (s *Service) DeleteTodo(id int) error {
	err := s.db.DeleteTodo(id)

	return err
}

//set complete ToDo
//param: model ToDo with id, model ToDoComp
//return model id, error
func (s *Service) SetCompTodo(id models.ToDo, m models.ToDoComp) (uint, error) {
	validate := validator.New()
	if err := validate.Struct(m); err != nil {
		return 0, BadValidation
	}
	idd, err := s.db.UpdateCompleteToDo(id, m)

	return idd, err
}

//get Incoming ToDo
//param: mode: 0 - today, 1 - tomorrow, 2 - current week, default - today
//return array of models, count, error
func (s *Service) GetIncTodo(mode int) ([]models.ToDo, int64, error) {
	layout := "2006-01-02"
	t := time.Now()
	var firstDate, secondDate string
	switch mode {
	case 0:
		firstDate = t.Format(layout)
		secondDate = t.AddDate(0, 0, 1).Format(layout)
	case 1:
		t = t.AddDate(0, 0, 1)
		firstDate = t.Format(layout)
		secondDate = t.AddDate(0, 0, 1).Format(layout)
	case 2:
		tn := t.AddDate(0, 0, 8)
		firstDate = t.Format(layout)
		secondDate = tn.Format(layout)
	default:
		firstDate = t.Format(layout)
		secondDate = t.AddDate(0, 0, 1).Format(layout)
	}
	m, count, err := s.db.GetIncTodo(firstDate, secondDate)

	return m, count, err
}
