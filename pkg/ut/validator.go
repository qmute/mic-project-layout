package ut

import (
	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
)

func ValidStruct(obj interface{}) error {
	if err := validator.New().Struct(obj); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
