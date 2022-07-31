package database

import (
	"errors"
	"log"

	"github.com/wojbog/ToDoList/models"
	"gorm.io/gorm"
)

//error variables
var (
	CanNotAdd                 = errors.New("can not add TODO")
	CanNotDelete              = errors.New("can not delete TODO ")
	CanNotUpdate              = errors.New("can not update TODO")
	CanNotGet                 = errors.New("can not get TODOs")
	CanNotGetSpecific         = errors.New("TODO does not exist")
	CanNotGetSpecificInterval = errors.New("For this interval todos does not exist")
)

type Repository interface {
	GetTodo(id int) (models.ToDo, error)
	AddToDo(m models.ToDo) (uint, error)
	UpdateToDo(m models.ToDo) (uint, error)
	DeleteTodo(id int) error
	GetAllTodos() ([]models.ToDo, int64, error)
	UpdateCompleteToDo(id models.ToDo, m models.ToDoComp) (uint, error)
	GetIncTodo(firstDate, secondDate string) ([]models.ToDo, int64, error)
}

//database instance
type repo struct {
	DB *gorm.DB
}

//repo constructor
func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

//add TDO to database
func (todoDB *repo) AddToDo(m models.ToDo) (uint, error) {
	if result := todoDB.DB.Create(&m); result.Error != nil {
		log.Println(result.Error)
		return 0, CanNotAdd
	} else {
		log.Println("successfully TODO created")
		return m.ID, nil
	}
}

func (todoDB *repo) UpdateToDo(m models.ToDo) (uint, error) {
	if result := todoDB.DB.Save(&m); result.Error != nil {
		log.Println(result.Error)
		return 0, CanNotUpdate
	} else {
		log.Println("successfully TODO Updated")
		return m.ID, nil
	}
}

func (todoDB *repo) UpdateCompleteToDo(id models.ToDo, m models.ToDoComp) (uint, error) {
	if result := todoDB.DB.Model(&id).Update("Complete", m.Complete); result.Error != nil {
		log.Println(result.Error)
		return 0, CanNotUpdate
	} else {
		log.Println("successfully TODO Updated")
		return id.ID, nil
	}
}

func (todoDB *repo) GetAllTodos() ([]models.ToDo, int64, error) {
	m := []models.ToDo{}
	if result := todoDB.DB.Find(&m); result.Error != nil {
		log.Println(result.Error)
		return []models.ToDo{}, 0, CanNotGet
	} else {
		return m, result.RowsAffected, nil
	}
}

func (todoDB *repo) GetTodo(id int) (models.ToDo, error) {
	m := models.ToDo{}
	if result := todoDB.DB.Find(&m, id); result.Error != nil {
		log.Println(result.Error)
		return models.ToDo{}, CanNotGet
	} else if result.RowsAffected == 0 {
		return models.ToDo{}, CanNotGetSpecific
	} else {
		return m, nil
	}

}

func (todoDB *repo) DeleteTodo(id int) error {
	m := models.ToDo{}
	if result := todoDB.DB.Delete(&m, id); result.Error != nil {
		log.Println(result.Error)
		return CanNotDelete
	} else {
		log.Println("successfully TODO Deleted")
		return nil
	}

}

func (todoDB *repo) GetIncTodo(firstDate, secondDate string) ([]models.ToDo, int64, error) {
	m := []models.ToDo{}
	if result := todoDB.DB.Where("expiry BETWEEN ? AND ?", firstDate, secondDate).Find(&m); result.Error != nil {
		log.Println(result.Error)
		return []models.ToDo{}, 0, CanNotGet
	} else if result.RowsAffected == 0 {
		return []models.ToDo{}, 0, CanNotGetSpecificInterval
	} else {
		return m, result.RowsAffected, nil
	}
}
