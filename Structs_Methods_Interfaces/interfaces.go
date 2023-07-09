package Structs_Methods_Interfaces

import "fmt"

type teacher interface {
	calculate_DOB() int
}

func Print_teacher_details(name string, t teacher) {
	fmt.Print("Teacher ", name, "'s birthday is ", t.calculate_DOB(), "!\n")
}
