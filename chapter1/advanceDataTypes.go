package main

import "fmt"

func arrayEx(){
	var arrayFirst [5]int = [5]int{1,2,3,4,5}
	fmt.Println(arrayFirst)
}

func sliceEx(){
	var sliceFirst []int = []int{}
	sliceFirst = append(sliceFirst, 1,2,3,4,5)
	fmt.Println(sliceFirst)

	var sliceSecond []int
	sliceSecond = append(sliceSecond, 1)
	sliceSecond = append(sliceSecond, 2)
	sliceSecond = append(sliceSecond, 3)
	fmt.Println(sliceSecond)
}

func mapEx(){
	var mappFirst map[string]string = map[string]string{
		"first": "1",
		"second": "2",
		"third": "3",
	}
	fmt.Println(mappFirst)

	var mappSecond map[string]string
	mappSecond = make(map[string]string)
	mappSecond["first"] = "1"
	mappSecond["second"] = "2"
	mappSecond["third"] = "3"

	fmt.Println(mappSecond)
}

func main(){

	fmt.Println("welcome in learning advanced data types ex:- Array, Slice, map")
	arrayEx()
	sliceEx()
	mapEx()

}