package main

import "fmt"

func max(l []int) int {
	j := l[0]
	for _, i := range l[1:] {
		if j < i {
			j = i
		}
	}
	return j
}

func main() {

	loops := []string{"harsh", "adi", "shreya"}

	fmt.Println(loops)
	fmt.Println(len(loops))
	fmt.Println(loops[1])
	fmt.Println(loops[:])
	fmt.Println(loops[1:])

	fmt.Printf("%s is %T\n", loops, loops)

	for i := 0; i < len(loops); i++ {
		fmt.Println(loops[i])
	}

	for i := range loops {
		fmt.Println(i)
	}
	for i, name := range loops {
		fmt.Printf("%s at %v\n", name, i)
	}
	for _, name := range loops {
		fmt.Println(name)
	}

	loops = append(loops, "ravi")
	fmt.Println(loops)
	loops = append(loops, "mritu")
	fmt.Println(loops)
	l := []int{5, 10, 20, 3, 99, 4, 88, 100, 5, 0}
	fmt.Println(max(l))

}
