package main

import (
	"fmt"

	"example.com/m/stud1"
	"example.com/m/stud2"
	"example.com/m/stud3"
)


func main(){

	// function call package stud1
	stud1.SayHello()	
	
	// function call package stud2
	fmt.Println(stud2.Add(4, 5))
	fmt.Println(stud2.Add(234, 12))
	fmt.Println(stud2.Multiply(10, 112))

	// function call package stud3
	fmt.Println(stud3.Repeat("a"))
	fmt.Println(stud3.SumAllNumbers(1, 3, 2000, -4, -543))
	fmt.Println(stud3.SumPositiveNumbers(1, 3, 2000, -4))
}