package request

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type Token struct {
	Address string `json:"address" validate:"required"`
}

func (t *Token) Validate() (ok bool, error map[string]string) {
	result := make(map[string]string)
	validate := validator.New()
	err := validate.Struct(*t)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		for i := range errors {
			if errors[i].StructField() == "Address" {
				result["Address"] = "アドレスを入力してください"
				return false, result
			}
		}
	}

	if validateAddress(t.Address) != true {
		result["Address"] = "不正なアドレス形式です"
		return false, result
	}

	return true, nil
}

// ウォレットアドレス検証
func validateAddress(address string) bool {
	match, _ := regexp.MatchString(`^0x[0-9a-fA-F]{40}$`, address)
	return match
}
