package Structs_Methods_Interfaces

import (
	"fmt"
)

var Tr_eiphyo goethe_teacher
var Tr_helen british_council_teacher

func Structs_Methods_Interfaces() {
	Tr_eiphyo = goethe_teacher{
		Name:                  "Ei Phyo",
		highest_qualification: "Grunes Diplom",
		age:                   29,
	}
	Tr_helen = british_council_teacher{
		Name:     "Helen",
		location: "Yangon",
		age:      28,
	}

	one_time_struct := struct {
		official_name string
		like_it       bool
	}{
		official_name: "Anonymous structs or something like that",
		like_it:       true,
	}
	fmt.Println("concept:", one_time_struct.official_name)
	if one_time_struct.like_it {
		fmt.Println("I kinda like them")
	} else {
		fmt.Println("I kinda don't like them")
	}

}
