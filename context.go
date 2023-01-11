package main

import (
	"context"
	"fmt"
	"time"
)

type Bid struct {
	adurl string
	price float64
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)

	return Bid{
		adurl: "https://adsRus.com/12",
		price: 0.09,
	}

}

var defaultBid = Bid{
	adurl: "https://adsRus.com/default",
	price: 0.02,
}

func findBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) //buffered to avoid goroutine leak

	go func() {
		ch <- bestBid(url)
	}()

	select {
	case Bid1 := <-ch:
		return Bid1
	case <-ctx.Done():
		return defaultBid
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	url := "https://httpl.cat/23"
	Bid1 := findBid(ctx, url)
	fmt.Println(Bid1)

	fmt.Println("----------")

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	url1 := "https://httpl.cat/2387"
	Bid1:= findBid(ctx, url1)
	fmt.Println(Bid1)
}
