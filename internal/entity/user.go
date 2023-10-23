package entity

import (
	"github.com/google/uuid"
)

type User struct {
	UserID      uuid.UUID 	`json:"user_id" db:"user_id" validate:"omitempty"`
	Name   string    				`json:"name" db:"name" validate:"required,lte=30"`
	Surname    string    		`json:"surname" db:"surname" validate:"required,lte=30"`
	Patronymic     string   `json:"patronymic" db:"patronymic"`
	Age         int       	`json:"age" db:"age" validate:"required"`
	Gender      string    	`json:"gender" db:"gender" validate:"required"`
	Nationality string    	`json:"nationality" db:"nationality" validate:"required"`
}

type RichData struct{
		Age int
		Gender string
		Nationality string
	}
