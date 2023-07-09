package main

import (
	"fmt"

	"github.com/DreamLineLove/GoByExample/Hello_World"
	"github.com/DreamLineLove/GoByExample/Structs_Methods_Interfaces"
)

func main() {
	write_title("Hello World")
	Hello_World.Hello_World()
	fmt.Println(Hello_World.Username)
	new_line()

	write_title("Structs, Methods, and Interfaces combined")
	Structs_Methods_Interfaces.Structs_Methods_Interfaces()
	Structs_Methods_Interfaces.Print_teacher_details(Structs_Methods_Interfaces.Tr_eiphyo.Name, Structs_Methods_Interfaces.Tr_eiphyo)
	Structs_Methods_Interfaces.Print_teacher_details(Structs_Methods_Interfaces.Tr_helen.Name, Structs_Methods_Interfaces.Tr_helen)
	new_line()
}
