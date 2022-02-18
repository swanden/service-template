package errors

import (
	"github.com/pkg/errors"
	"github.com/swanden/service-template/internal/domain"
)

var AlreadyExistsUser = errors.Wrap(domain.Error, "user already exists")
