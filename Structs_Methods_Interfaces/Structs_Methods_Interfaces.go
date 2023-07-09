package Structs_Methods_Interfaces

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
}
