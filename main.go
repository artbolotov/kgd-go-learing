package main

import (
	"fmt"

	"example.com/m/part1"
	"example.com/m/part2"
	"example.com/m/part3"
	"example.com/m/part4"
	"example.com/m/part5"
)


func main(){

	// function call package part1
	part1.SayHello()	
	
	// function call package part2
	fmt.Println(part2.Add(4, 5))
	fmt.Println(part2.Add(234, 12))
	fmt.Println(part2.Multiply(10, 112))

	// function call package part3
	fmt.Println(part3.Repeat("a"))
	fmt.Println(part3.SumAllNumbers(1, 3, 2000, -4, -543))
	fmt.Println(part3.SumPositiveNumbers(1, 3, 2000, -4))

	// function call package part4
	arr := [5]int{2, 5, 6, 12, 8} //array
	fmt.Println(part4.Sum(arr[:]))

	slice := []int {1, 2, 3, 4, 5, 6} //slice
	fmt.Println(part4.Sum(slice[:]))

	fmt.Println(part4.SumAllTails(arr[:]))

	// function call package part5
	rectangle := part5.Rectangle{12, 6}
	fmt.Println(part5.Perimeter(rectangle))

	fmt.Println(part5.Rectangle.Area(rectangle))
}