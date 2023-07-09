package main

import (
	"fmt"
	"strings"
)

func write_title(title string) {
	fmt.Println("\t", strings.ToUpper(title))
}

func new_line() {
	fmt.Println()
}
