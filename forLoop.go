package main

import (
	"fmt"
)

func fizzBuzz() {
	for i := 0; i <= 20; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzBuzz")

		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("Hii Harsh lets GO")
	for i := 0; i < 21; i++ {
		if i == 19 {
			break
		}
		if i == 3 {

			continue
		}
		fmt.Println(i)
	}
	//like while
	i := 0
	for i < 21 {
		if i == 19 {
			break
		}
		if i == 3 {
			i++
			continue
		}
		fmt.Println(i)
		i++

	}
	//like infinite
	j := 1
	for {
		if j == 19 {
			break
		}
		if j == 3 {
			j++
			continue
		}
		fmt.Println(j)
		j++
	}
	fizzBuzz()
}
