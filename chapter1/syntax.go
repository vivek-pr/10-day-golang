package main
import "fmt"

func arthimetic(){
	var a = 12
	var b = 18
	var c = a + b
	fmt.Println(c)
	var d = a-b
	fmt.Println(d)

	var e = a*b
	fmt.Println(e)

	var f = a/b
	fmt.Println(f)
}

func comparision(){
	var a = 72;
	var b = 42;

	var c = a > b;
	fmt.Println(c)

	var d = a<b;
	fmt.Println(d)

	var e = a==b;
	fmt.Println(e)
}
func main(){
	arthimetic()
	comparision()
}

