package main

import (
	"fmt"
	"time"
)

func main(){
	ch1,ch2:=make(chan int),make(chan int)

	go func ()  {
		ch2<-33
	}()

	select{
	case val:=<-ch1:
		fmt.Printf("Got %d from ch1\n",val)
	case val:=<-ch2:
		fmt.Printf("Got %d from ch2\n",val)
	}


	out:=make(chan int)

	go func(){
		time.Sleep(100*time.Millisecond)
		out<-44
	}()

	select{
	case val:=<-out:
		fmt.Printf("Got %d from out\n",val)
	case <-time.After(20*time.Millisecond):
		fmt.Println("timeout")
	}
}