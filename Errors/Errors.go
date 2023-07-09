package Errors

import (
	"errors"
	"fmt"
)

func Errors() {
	simpleError := errors.New("This should return an error")
	fmt.Println(simpleError)
}
