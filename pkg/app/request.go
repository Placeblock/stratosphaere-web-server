package app

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func MarkErrors(errors validator.ValidationErrors) {
	for _, err := range errors {
		fmt.Println(err.Error())
	}
}
