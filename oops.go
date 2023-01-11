package main

import (
	"fmt"
	"time"
)

type badg struct {
	Id      string
	Amt     float64
	Country string
	Time    time.Time
}
func newBadg(id string,amt float64,country string,timex time.Time)(*badg,error){
	if id ==""{
		return nil,fmt.Errorf("empty id")

	}
	if amt<=0{
		return nil,fmt.Errorf("balance is less")
	}
	if country==""{
		return nil,fmt.Errorf("empty country")

	}
	if timex.Before(time.Now()){
		return nil,fmt.Errorf("bad date")

	}
	b:=badg{
		Id:id,
		Amt:amt,
		Country: country,
		Time: timex,
	}
	return &b,nil
}
func (b badg) current()time.Duration{
	return b.Time.Sub(time.Now().UTC())
}
//we pass pointer of reciver when we do changes in object

func (b *badg)Newamt(v float64){
	b.Amt+=v
}
func main() {

	b := badg{"amzn", 444.5, "usa", time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b)
	fmt.Printf("%#v\n", b) //%#v for more details
	println(b.Id)

	b1 := badg{Id: "googl", Amt: 234.9}
	fmt.Println(b1)
	fmt.Printf("%#v\n",b1)


	var b3 badg
	fmt.Println(b3)

	//everyting starts with capital is accessible from outside package


	fmt.Println(b.current())
	b.Newamt(10.5)
	fmt.Println(b)


	b4,err:=newBadg("tesla",55.3,"usa",time.Now().Add(2*24*time.Hour))
	if err!=nil{
		fmt.Println(err)
	}else{
	fmt.Printf("%#v\n",b4)}
}
