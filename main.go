package main

import (
	"fmt"	
)

const russian = "Russian"
//const english = "English"
const englishHelloPrefix = "Hello, "
const russianHelloPrefix = "Привет, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == russian {
		return russianHelloPrefix + name 
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){
	switch language{
	case russian:
		return russianHelloPrefix
	default:
		return englishHelloPrefix			
	}
	
}

func main() {
	fmt.Println(Hello("Артем", russian)) //
}
