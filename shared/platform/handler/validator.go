package handler

import "github.com/asaskevich/govalidator"

func Validator(body interface{}) error {
	_, err := govalidator.ValidateStruct(body)

	return err
}
