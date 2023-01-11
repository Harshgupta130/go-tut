package main

import(
	"fmt"
)



func main(){
	vals:=[]int{1,2,3}
	// v:=vals[10]//wil cause a panic

	// fmt.Println(v)
	
	
	//panic("oops creating a panic")

	v,err:=safeValue(vals,10)
	if err!=nil{
		fmt.Println("error : ",err)
	}
	fmt.Println("v: " ,v)

}

func safeValue(vals []int,index int)(n int,err error){
	defer func(){
		if e:=recover();e!=nil{
			err=fmt.Errorf("%v",e)
		}
	}()

	return vals[index], nil
}