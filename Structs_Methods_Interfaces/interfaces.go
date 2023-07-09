package Structs_Methods_Interfaces

import "fmt"

type teacher interface {
	calculate_DOB() int
	list_details() (string, int, string)
}

func Print_teacher_details(name string, t teacher) {
	fmt.Print("Teacher ", name, "'s birthday is ", t.calculate_DOB(), "!\n")
}

func print_teacher_details_all(t teacher) {
	Name, age, other_info := t.list_details()
	fmt.Println(Name)
	fmt.Println(age)
	fmt.Println(other_info)
}
