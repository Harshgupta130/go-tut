package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hii Harsh lets GO")
	// x := 8
	// // var y int
	// // y = 9
	// var x float64
	// y:=1.0
	// x=2.0
	x, y := 1.05, 2.8
	fmt.Printf("x=%v, Type of x is %T\n", x, x)
	fmt.Printf("y=%v, Type of y is %T\n", y, y)
	fmt.Printf("average of x and y is %v\n", (x+y)/2.0)
	mean := (x + y) / 2.0
	fmt.Printf("mean=%v,type of %T\n", mean, mean)
	// var x float64
	// y=1.0
	// x=2.0

}
