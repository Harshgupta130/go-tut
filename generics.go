package main

import (
	"fmt"
)

type ordered interface{
	int|float64|string
}
func min[T ordered](values []T)(T,error){
	if len(values)==0{ 
		var zero T
		return zero,fmt.Errorf("emtpty list")
	}
	m:=values[0]

	for _,v:=range values[1:]{
		if v<m{
			m=v
		}
	}

	return m,nil
}
func main(){
	fmt.Println(min([]int{12,3,755,5}))
	fmt.Println(min([]string{"a","v","d"}))
}