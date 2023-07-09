package Structs_Methods_Interfaces

import (
	"fmt"
	"time"
)

type goethe_teacher struct {
	year_of_birth         int
	highest_qualification string
}

func (g goethe_teacher) calculate_age() int {
	return time.Now().Year() - g.year_of_birth
}

type british_council_teacher struct {
	year_of_birth int
	location      string
}

func (b british_council_teacher) calculate_age() int {
	return time.Now().Year() - b.year_of_birth
}

type teacher interface {
	calculate_age() int
}

func print_age(t teacher) {
	fmt.Println("the teacher is", t.calculate_age(), "years old!")
}
