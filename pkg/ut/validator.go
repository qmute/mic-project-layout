package ut

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func ValidStruct(obj interface{}) error {
	if err := validator.New().Struct(obj); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
