package entity

import "github.com/google/uuid"

type ID struct {
	Value string `gorm:"column:id;type:varchar(255)"`
}

func NewID(value string) *ID {
	return &ID{value}
}

func NextID() *ID {
	return &ID{uuid.NewString()}
}
