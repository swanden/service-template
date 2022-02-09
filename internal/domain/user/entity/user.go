package entity

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID
	Email    string
	Password string
}
