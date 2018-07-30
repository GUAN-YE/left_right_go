package main
import "fmt"


func main() {
	a := 60
	b := 80
	var f,g,h,c ,d,e int
	c = a*b
	d = a+b
	e = b-a
	f = b/a
	g = a<<2
	h = b>>2
	fmt.Println("a*b=",c)
	fmt.Println("a+b=",d)
	fmt.Println("b-a=",e)
	fmt.Println("b/a=",f)
	fmt.Println("a<<2=",g)
	fmt.Println("b>>2=",h)
}