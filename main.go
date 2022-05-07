package main

import (
	"fmt"

	"example.com/m/stud1"
	"example.com/m/stud2"
)


func main(){

	// function call package stud1
	stud1.SayHello()	
	
	// function call package stud2
	fmt.Println(stud2.Add(4, 5))
	fmt.Println(stud2.Add(234, 12))
}