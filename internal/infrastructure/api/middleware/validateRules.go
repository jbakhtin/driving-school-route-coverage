package middleware

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func LineString(fl validator.FieldLevel) bool {
	coordinates := fl.Field().Interface().([][]float64)
	// Проверяем, что в LineString есть хотя бы две точки
	if len(coordinates) < 2 {
		return false
	}

	// Проверяем формат каждой координаты (широта, долгота)
	coordRegex := regexp.MustCompile(`^-?\d+(\.\d+)?,-?\d+(\.\d+)?$`)
	for _, c := range coordinates {
		if !coordRegex.MatchString(fmt.Sprintf("%f,%f", c[0], c[1])) {
			return false
		}
	}
	return true
}