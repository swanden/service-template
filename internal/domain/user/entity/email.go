package entity

import (
	"github.com/pkg/errors"
	"github.com/swanden/service-template/internal/domain"
	"regexp"
)

type Email struct {
	Value string `gorm:"column:email;type:varchar(255)"`
}

func NewEmail(value string) (*Email, error) {
	if !isEmailValid(value) {
		return nil, errors.Wrap(domain.Error, "invalid email")
	}

	return &Email{Value: value}, nil
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}

	return emailRegex.MatchString(e)
}
