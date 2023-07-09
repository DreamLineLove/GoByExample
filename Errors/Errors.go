package Errors

import (
	"errors"
	"fmt"
)

func Errors() {
	simpleError := errors.New("This should return an error")
	fmt.Println(simpleError)

	if message, ok := simpleError.(error); ok {
		fmt.Println("This was an error indeed!\n\tMessage:", message)
	}

	simpleErrorf := fmt.Errorf("this is a simple Errorf without any contained errors")
	fmt.Println(simpleErrorf)

	err1 := errors.New("w1")
	errorF := fmt.Errorf("outer, and inner: %v", err1)
	fmt.Println(errorF)

	// The following logic returns nil.
	inner := errors.Unwrap(errorF)
	fmt.Println(inner)

	err1 = errors.New("w1 again")
	errorF = fmt.Errorf("outer, and inner: %w", err1)
	inner = errors.Unwrap(errorF)
	fmt.Println(inner)
}
