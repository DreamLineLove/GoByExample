package Sort

import (
	"fmt"
	"sort"
)

func primitives() {
	floats := []float64{9.086, 7.680, 77.58, 45.0}
	fmt.Println(floats)
	sort.Float64s(floats)
	fmt.Println(floats)
}
