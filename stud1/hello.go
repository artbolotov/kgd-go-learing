package stud1

import (
	"fmt"
)

const russian = "Russian"
//const english = "English"
const englishHelloPrefix = "Hello, "
const russianHelloPrefix = "Привет, "

func hello(name string, languages string) string {
	if name == "" {
		name = "World"
	}
	if languages == russian {
		return russianHelloPrefix + name 
	}
	return greetingPrefix(languages) + name
}

func greetingPrefix(languages string) (prefix string){
	switch languages{
	case russian:
		return russianHelloPrefix
	default:
		return englishHelloPrefix			
	}
	
}

func SayHello() {
	fmt.Println(hello("Артем", russian)) //
}
