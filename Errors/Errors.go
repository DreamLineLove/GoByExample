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

	if feedback, err := simple_error_return("haha", true); err != nil {
		fmt.Println(feedback, err)
	} else {
		fmt.Println("error not found! we will proceed to other logic!")
	}

	if feedback, err := simple_error_return("haha", false); err != nil {
		fmt.Println(feedback, err)
	} else {
		fmt.Println("error not found! we will proceed to other logic!")
	}

	var err error

	if err = complex_error_return(true); err != nil {
		fmt.Println("error returned!", err)
		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println("proceeding withut any errors!")
	}

	if err = complex_error_return(false); err != nil {
		fmt.Println("error returned!", err)
	} else {
		fmt.Println("proceeding withut any errors!")
	}

	error2 := errors.New("error2")
	error1 := fmt.Errorf("error1, inner: %w", error2)
	outerError := fmt.Errorf("outer. inner: %w", error1)
	fmt.Println(errors.Unwrap(outerError))
}

func simple_error_return(message string, should_return bool) (string, error) {
	if should_return {
		return "error was returned: ", errors.New(message)
	} else {
		return "error was not returned: ", nil
	}
}

func return_error(name string) error {
	return errors.New(name)
}

func complex_error_return(should_return bool) error {
	if should_return {
		database := return_error("database error!")
		fetch := return_error("fetch failed!")
		fileOpen := return_error("unable to open file!")
		return fmt.Errorf("outer error, inner errors: %w %w %w", database, fetch, fileOpen)
	} else {
		return nil
	}
}
