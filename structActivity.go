package main
import(
	"fmt"
)

type square struct{
	X int
	Y int
	Length int
}
func NewSquare(x int,y int,length int)(*square,error){
if length<=0{
	return nil,fmt.Errorf("cant create square")
}
s:=square{
	X:x,
	Y: y,
	Length: length,
}
return &s,nil
}

func (b *square) move(dx int,dy int){

	b.X+=dx
	b.Y+=dy
}

func (s *square) square()int{
return s.Length*s.Length
}
func main(){

	s,err:=NewSquare(1,1,22)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(s)
	}
	fmt.Println(s.square())
	s.move(3,4)
	fmt.Printf("%+v\n")
	fmt.Println(s)

}