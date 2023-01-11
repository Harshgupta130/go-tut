package main

import (
	"fmt"
	"time"
)

func main(){

	ch:=make(chan int)
	//ch <- 33//send it will panic as their goroutine are at sleep

	go func(){
		ch<-33
	}()
	val:=<-ch//recieve
	fmt.Println("got ",val)//it 

fmt.Println("-------")
	const count=3
	go func(){
		for i:=0;i<count;i++{
			fmt.Printf("sending %d\n",i)
			ch<-i
			time.Sleep(time.Second)
		}
	}()

	for i:=0;i<count;i++{
		val:=<-ch
		fmt.Printf("recieved %d\n",val)
	}

	fmt.Println("--------")
	go func(){
		for i:=0;i<count;i++{
			fmt.Printf("sending %d\n",i)
			ch<-i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i:=range ch{
		
		fmt.Printf("recieved %d\n",i)
	}
	fmt.Println("---------")

	ch2:=make(chan int,1)
	ch2<-18
	val1:=<-ch2
	fmt.Println(val1)


}