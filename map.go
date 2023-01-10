package main

import (
	"fmt"
	"strings"
)

func occurence(s string) map[string]int {
	words := strings.Fields(s)
	count := map[string]int{}
	for _, i := range words {
		count[strings.ToLower(i)]++
	}
	return count
}

func main() {
	m := map[string]float64{
		"amzn": 123.4,
		"gogl": 856.9,
		"fab":  765.0,
	}
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["amzn"])

	fmt.Println(m["tesla"])
	m["tesla"] = 444.4
	fmt.Println(m)
	value, ok := m["tesla"]
	if !ok {
		fmt.Println("tesla not found")
	} else {
		fmt.Println(value)
	}

	delete(m, "tesla")
	fmt.Println(m)
	value, k := m["tesla"]
	if !k {
		fmt.Println("tesla not found")
	} else {
		fmt.Println(value)
	}

	for key := range m {
		fmt.Println(key)
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
	for _, value := range m {
		fmt.Println(value)
	}

	fmt.Println(occurence("hii its harsh nice to meet you"))

}
