package requests

import "github.com/go-playground/validator/v10"

type CreateMemberRequestValidation struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func (r *CreateMemberRequestValidation) Validate() error {
	return validator.New().Struct(r)
}
