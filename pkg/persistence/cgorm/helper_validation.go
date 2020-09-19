package cgorm

import "github.com/TudorHulban/GinCRM/pkg/validator"

func validateStruct(data interface{}) error {
	return validator.GetValidator().Struct(data)
}
