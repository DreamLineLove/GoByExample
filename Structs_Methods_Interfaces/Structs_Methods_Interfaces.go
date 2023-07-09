package Structs_Methods_Interfaces

func Structs_Methods_Interfaces() {
	tr_eiphyo := goethe_teacher{
		highest_qualification: "Grunes Diplom",
		age:                   29,
	}
	tr_helen := british_council_teacher{
		location: "Yangon",
		age:      29,
	}
	print_DOB(tr_eiphyo)
	print_DOB(tr_helen)
}
