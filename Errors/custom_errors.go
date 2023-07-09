package Errors

import (
	"fmt"
	"strconv"
)

type QueryError struct {
	query string
	load  int
}

func (qe QueryError) Error() string {
	return "query: " + qe.query + " | load:" + strconv.Itoa(qe.load)
}

func Custom_errors() {
	yup := "foobar"
	nope := "djlfasdfj"
	var errOR error

	if result, err := query(yup); err != nil {
		errOR = err
	} else {
		fmt.Println("results:", result)
	}

	if result, err := query(nope); err != nil {
		errOR = err
	} else {
		fmt.Println("results:", result)
	}

	errorf := fmt.Errorf("outer error, inner: %w", errOR)
	fmt.Println(errorf)
}

func query(term string) (string, error) {
	if term == "foobar" {
		return "it should be foobar", nil
	} else {
		return "", &QueryError{query: term, load: 555}
	}
}
