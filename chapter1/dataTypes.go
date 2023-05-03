package main

import "fmt"

func stringExample(){
	var firstEx string = "Hello"
	fmt.Println("To decalare a string we use var firstEx string = \"Hello\"")
	fmt.Println("The value of firstEx is ",firstEx)
}

func intExample(){
	var firstEx int = 10
	fmt.Println("To declare an int we use var firstEx int = 10");
	fmt.Println("The value of firstEx is ",firstEx)
}

func floatExample(){
	var firstEx float64 = 10.5
	fmt.Println("To declare an float we use var firstEx float64 = 10.5");
	fmt.Println("The value of firstEx is ",firstEx)
}

func booleanExample(){
	var firstEx bool = true
	fmt.Println("To declare an boolean we use var firstEx bool = true");
		fmt.Println("The value of firstEx is ",firstEx)
}


func main(){
	fmt.Println("Welcome to the world of data type. We will be laerning about string, int, float, boolean")
	stringExample()
	intExample()
	floatExample()
	booleanExample()
}