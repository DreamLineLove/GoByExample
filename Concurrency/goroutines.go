package Concurrency

import (
	"fmt"
	"time"
)

func goroutines() {
	func(age int) {
		fmt.Println("year of birth:", time.Now().Year()-age)
	}(18)
}
