package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {

	fmt.Println(add(10, 45))
	div, mod := divmod(15, 4)
	fmt.Printf("div=%v\nmod=%v\n", div, mod)

	values := []int{1, 2, 3, 4}
	fmt.Println(values)
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	double(val)
	fmt.Println(val)

	//now using pointer

	doublePtr(&val)
	fmt.Println(val)

	//error return from function

	s, err := errorsqrt(78.0)
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	} else {
		fmt.Println(s)
	}

	s, err = errorsqrt(-65)
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	} else {
		fmt.Println(s)
	}
	//defer is kind a destructor
	worker()

	ctype, err := contentType("https://likedin.com")

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(ctype)
	}
}
func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (float64, int) {
	return float64(a) / float64(b), a % b
}

func doubleAt(values []int, i int) {
	values[i] *= 2
}
func double(val int) {
	val *= 2
}

func doublePtr(v *int) {
	*v *= 2
}

func errorsqrt(a float64) (float64, error) {
	if a < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%f)", a)
	}
	return math.Sqrt(a), nil
}

func worker() {
	r1, err := acciqure("Harsh")
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	defer release(r1)

	r2, err := acciqure("Ravi")
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	defer release(r2)

	fmt.Println("Worker")

}

func acciqure(a string) (string, error) {
	return a, nil
}

func release(a string) {
	fmt.Printf("Cleaning up %s\n", a)
}

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("can't fint content-type header")
	}
	return ctype, nil

}
