package Structs_Methods_Interfaces

import "fmt"

type teacher interface {
	calculate_DOB() int
}

func print_DOB(t teacher) {
	fmt.Println("the teacher's birthday is", t.calculate_DOB(), "!")
}
