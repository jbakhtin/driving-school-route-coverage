package main

import (
	"errors"
	"fmt"
)

func main() {
	var ErrorDuplicateKey error = errors.New("duplicate key value violates unique constraint")
	var ErrorTest error = errors.New("duplicate key value violates unique constraint")

	fmt.Println(ErrorDuplicateKey.Error() == ErrorTest.Error())
}
