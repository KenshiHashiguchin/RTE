package request

import (
	"github.com/go-playground/validator/v10"
)

type Reserve struct {
	PaymentId       string `json:"payment_id" validate:"required"`
	MerchantAddress string `json:"merchant_address" validate:"required"`
	Surname         string `json:"surname" validate:"required,max=30"`
	Firstname       string `json:"firstname" validate:"required,max=30"`
	Phonenumber     string `json:"phonenumber" validate:"required,max=20"`
	Number          uint   `json:"number" validate:"required,min=1,max=100"`
	Time            uint   `json:"time" validate:"required"`
}

func (r *Reserve) Validate() (ok bool, error map[string]string) {
	result := make(map[string]string)
	validate := validator.New()
	err := validate.Struct(*r)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		for i := range errors {
			if errors[i].Tag() == "required" {
				result[errors[i].StructField()] = "入力必須です。"
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

	if validateAddress(r.MerchantAddress) != true {
		result["MerchantAddress"] = "不正なアドレス形式です"
	}

	// time
	//time := time.Unix(int64(r.Time), 0)
	//if time.After(time.No) {
	//	result["Time"] = "不正な日付です。"
	//}

	return true, nil
}
