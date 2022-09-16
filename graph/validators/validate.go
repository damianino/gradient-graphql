package validators

import (
	"github.com/go-playground/validator/v10"
)

func Validator() *validator.Validate {
	v := validator.New()

	_ = v.RegisterValidation("imgsize", func(fl validator.FieldLevel) bool {
		size := fl.Field().Int()
		return size > 0 && size <= 4096
	})
	_ = v.RegisterValidation("comment", func(fl validator.FieldLevel) bool {
		com := fl.Field().String()
		return len(com) < 300 && len(com) > 0
	})
	_ = v.RegisterValidation("name", func(fl validator.FieldLevel) bool {
		name := fl.Field().String()
		return len(name) < 128 && len(name) > 0
	})
	_ = v.RegisterValidation("stops", func(fl validator.FieldLevel) bool {

		return true
	})

	return v
}
