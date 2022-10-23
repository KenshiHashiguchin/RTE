package request

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Reserve struct {
	MerchantId  uint      `validate:"required"`
	Surname     string    `validate:"required,max=30"`
	Firstname   string    `validate:"required,max=30"`
	Phonenumber string    `validate:"required,max=12"`
	Number      uint      `validate:"required,min=1,max=100"`
	Time        time.Time `validate:"required,min=1,max=100"`
}

func (r *Reserve) Validate() (ok bool, error map[string]string) {
	result := make(map[string]string)
	validate := validator.New()
	err := validate.Struct(*r)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		for i := range errors {
			if errors[i].Tag() == "required" {
				result["ReceivedAddress"] = "入力必須です。"
				continue
			}
			switch errors[i].StructField() {
			case "Surname":
				switch errors[i].Tag() {
				case "max":
					result["Surname"] = "30字以内で入力してください。"
				}
			case "Firstname":
				switch errors[i].Tag() {
				case "max":
					result["Surname"] = "30字以内で入力してください。"
				}
			case "Phonenumber":
				switch errors[i].Tag() {
				case "max":
					result["Phonenumber"] = "12字以内で入力してください。"
				}
			case "Number":
				switch errors[i].Tag() {
				case "max":
					result["Name"] = "100人以上の予約はできません。"
				}
			case "Time":
				switch errors[i].Tag() {
				case "min":
					result["Time"] = "0以上の数字を入力してください。"
				}
			}
		}
	}

	// time
	if r.Time.After(time.Now()) {
		result["Time"] = "不正な日付です。"
	}

	return true, nil
}
