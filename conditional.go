package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hii Harsh lets GO")
	n := 75
	if n%5 == 0 {
		fmt.Println("Divided by 5")
	} else {
		fmt.Println("Not Divided by 5")
	}

	if n%5 == 0 && n%3 == 0 {
		fmt.Println("fizzBuZZ")
	} else if n%5 == 0 || n%3 == 0 {
		fmt.Println("fizzz")
	} else {
		fmt.Println("none")
	}

	switch n {
	case 75:
		fmt.Println("75")
	case 100:
		fmt.Println("100")
	default:
		fmt.Println("Default")

	}

	switch {
	case n%5 == 0 && n%3 == 0:
		fmt.Println("fizzBuZZ")
	case n%5 == 0 || n%3 == 0:
		fmt.Println("fizzz")
	default:
		fmt.Println("Default")
	}
}
