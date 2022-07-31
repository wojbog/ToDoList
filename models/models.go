package models

import (
	"time"

	"gorm.io/gorm"
)

//TODO struct
type ToDo struct {
	gorm.Model
	Title       string    `json:"title" validate:"required,alphanum"`
	Description string    `json:"description" validate:"required,alphanum"`
	Complete    int       `json:"complete" gorm:"default:0" validate:"gte=0,lte=100"`
	Expiry      time.Time `json:"expiry" validate:"required"`
}

type ToDoComp struct {
	Complete int `validate:"gte=0,lte=100,required"`
}
