package main

import (
	"context"
	"fmt"
	"time"
)


type bid struct{
	adurl string
	price float64
}

func bestBid(url string)bid{
time.Sleep(20*time.Millisecond)

return bid{
adurl: "https://adsRus.com/12",
price: 0.09,
}

}

var defaultBid=bid{
	adurl: "https://adsRus.com/default",
	price: 0.02,
	
}

func findBid(ctx context.Context,url string)bid{
	ch:=make(chan bid ,1)//buffered to avoid goroutine leak

	go func(){
		ch<-bestBid(url)
	}()

	select{
	case bid1:=<-ch:
		return bid1
	case <-ctx.Done():
		return defaultBid
	}
}

func main(){
	ctx,cancel:=context.WithTimeout(context.Background(),100*time.Millisecond)
	defer cancel()
	url:="https://httpl.cat/23"
	bid1:=findBid(ctx,url)
	fmt.Println(bid1)

	fmt.Println("----------")

	ctx,cancel=context.WithTimeout(context.Background(),10*time.Millisecond)
	defer cancel()
	url1:="https://httpl.cat/2387"
	bid2:=findBid(ctx,url1)
	fmt.Println(bid2)
}