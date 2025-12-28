package main

import (
	"fmt"
)

func main() {
	x := NewMatrix(2, 2)
	x.Data = []float64{0, 1, 2, 3} // toy
	fmt.Println(sigmoid(x))
}
