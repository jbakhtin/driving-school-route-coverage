package main

import (
	"errors"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/apperror"
)

func main() {
	appError := apperror.New(nil, "test apperror", "000", "developer message test", nil)

	test := errors.New("message test")

	temp := errors.As(appError, &appError)

	fmt.Println(temp)
	fmt.Println(test)
	fmt.Println(appError)
}
