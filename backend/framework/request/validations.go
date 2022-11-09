package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationInterface interface {
	UniqueField(validator.FieldLevel) bool
}

func ValidationMessages(err *validator.ValidationErrors, fields *map[string]string, messages *map[string]string) map[string][]string {
	vMessages := make(map[string][]string)

	for _, er := range *err {
		nsp := er.Namespace()
		tag := er.Tag()
		msg, ok := (*messages)[nsp+"."+tag]
		key := (*fields)[nsp]
		switch tag {
		case "required":
		case "required_without":
			if !ok {
				msg = fmt.Sprintf("Please provide %s", key)
			}
		case "email":
			if !ok {
				msg = fmt.Sprintf("Please provide valid email for %s", key)
			}
		}
		vMessages[key] = append(vMessages[key], msg)
	}

	return vMessages
}
