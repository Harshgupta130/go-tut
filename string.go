package main

import "fmt"

func evenEndeddigit() int {
	count := 0
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			n := a * b
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count += 1
			}
		}
	}
	return count
}

func main() {
	book := "The wheather is good."
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Printf("book[0]=%v type id %T", book[0], book[0])

	fmt.Println(book[:5])
	fmt.Println(book[4:])
	fmt.Println(book[4:9])
	fmt.Println("t" + book[1:])

	poem := `
	hii iam harsh
	go goa gone
	`

	//String is immutable in go
	//book="harsh" not possible
	//its unicode u can use like emoji and else in string
	fmt.Println(poem)
	fmt.Println("the ½ happy ")
	fmt.Println("result of evenended ", evenEndeddigit())
}
