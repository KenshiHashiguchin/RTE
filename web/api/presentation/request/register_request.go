package request

import (
	"github.com/go-playground/validator/v10"
)

type Register struct {
	ReceivedAddress string `validate:"required"`
	Name            string `validate:"required,max=30"`
	Tel             string `validate:"required,max=12"`
	MerchantAddress string `validate:"required"`
	Deposit         uint   `validate:"required,min=0"`
	CancelableTime  uint   `validate:"required,min=0"`
}

func (r *Register) Validate() (ok bool, error map[string]string) {
	result := make(map[string]string)
	validate := validator.New()
	err := validate.Struct(*r)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		for i := range errors {
			if errors[i].Tag() == "required" {
				result["ReceivedAddress"] = "入力必須です。"
				ok = false
				continue
			}
			switch errors[i].StructField() {
			case "Name":
				switch errors[i].Tag() {
				case "max":
					result["Name"] = "30字以内で入力してください。"
				}
			case "Tel":
				switch errors[i].Tag() {
				case "max":
					result["ReceivedAddress"] = "20字以内で入力してください。"
				}
			case "Deposit":
				switch errors[i].Tag() {
				case "min":
					result["Deposit"] = "0以上の数字を入力してください。"
				}
			case "CancelableTime":
				switch errors[i].Tag() {
				case "min":
					result["CancelableTime"] = "0以上の数字を入力してください。"
				}
			}
		}
	}

	if validateAddress(r.ReceivedAddress) != true {
		result["Address"] = "不正なアドレス形式です"
	}

	if len(result) > 0 {
		return false, result
	}

	return true, nil
}
