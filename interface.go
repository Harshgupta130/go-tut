package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type square struct{
	Length float64
}
type circle struct{
	Radius float64
}
func(s square) Area()float64{
	return s.Length*s.Length
}
func (c circle) Area()float64{
	return math.Pi*c.Radius*c.Radius
}

func sumShape(shapes []shape)float64{
	total:=0.0

	for _,shape:=range shapes{
		total+=shape.Area()
	}
	return total
}

type shape interface{
	Area()float64
}



func main(){


	s:=square{11.2}
	fmt.Println(s.Area())

	c:=circle{4.4}
	fmt.Println(c.Area())
	shapes:=[]shape{s,c}
	fmt.Println(sumShape(shapes))



	
}