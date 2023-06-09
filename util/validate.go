package util

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"reflect"
)

var validate = validator.New()

func ValidateStruct(ctx context.Context, entity interface{}) error {
	val := reflect.ValueOf(entity)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("not a struct")
	}

	err := validate.Struct(entity)
	if err != nil {
		var errors gqlerror.List

		for _, errValidation := range err.(validator.ValidationErrors) {
			errors = append(errors, &gqlerror.Error{
				Path:    graphql.GetPath(ctx),
				Message: fmt.Sprintf("valdation error on field '%s'", errValidation.Field()),
				Extensions: map[string]interface{}{
					"failed_field": errValidation.StructNamespace(),
					"tag":          errValidation.Tag(),
					"value":        errValidation.Param(),
					"field":        errValidation.Field(),
					"errors":       errValidation.Error(),
					"actual_tag":   errValidation.ActualTag(),
				},
			})
		}

		return errors
	}

	return nil
}
