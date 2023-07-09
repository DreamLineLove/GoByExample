package Structs_Methods_Interfaces

import "time"

func (g goethe_teacher) calculate_DOB() int {
	return time.Now().Year() - g.age
}

func (g goethe_teacher) list_details() (string, int, string) {
	return g.Name, g.age, g.highest_qualification
}

func (b british_council_teacher) calculate_DOB() int {
	return time.Now().Year() - b.age
}

func (b british_council_teacher) list_details() (string, int, string) {
	return b.Name, b.age, b.location
}
