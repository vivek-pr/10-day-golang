package main

import "fmt"

func decision_making(){
	x := 12
	if x >10{
		fmt.Println("x is greater than 10")
	}else if x == 10{
		fmt.Println("x is equal to 10")
	}else{
		fmt.Println("x is less than 10")
	}
}

func switchControl(){
	x := "I am groot"
	switch x{
	case "I am groot":
		fmt.Println("I am groot")
	case "I am batista":
		fmt.Println("I am batista")
	case "I am goku":
		fmt.Println("I am goku")
	default:
		fmt.Println("I am not groot")
	}
}

func whileLoop(){
	var i = 0;
	
	for i< 5{
		fmt.Println("this is while loop", i)
	}
}

func forLoop(){

	for i:=0; i<10; i++{
		fmt.Println("this is for loop", i)
	}
}

func main(){
	fmt.Println("Welcome to learn control structure")
	decision_making()
	switchControl()
	forLoop()

}