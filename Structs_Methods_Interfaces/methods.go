package Structs_Methods_Interfaces

import "time"

func (g goethe_teacher) calculate_DOB() int {
	return time.Now().Year() - g.age
}

func (b british_council_teacher) calculate_DOB() int {
	return time.Now().Year() - b.age
}
