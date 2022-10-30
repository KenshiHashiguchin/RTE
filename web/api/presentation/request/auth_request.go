package request

import (
	"github.com/go-playground/validator/v10"
)

type Auth struct {
	Address   string `json:"address" validate:"required"`
	Signature string `json:"signature" validate:"required"`
}

func (t *Auth) Validate() (ok bool, error map[string]string) {
	result := make(map[string]string)
	validate := validator.New()
	err := validate.Struct(*t)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) > 0 {
			result["Address"] = "不正な署名です"
			return false, result
		}
	}

	return true, nil
}
