package main

import (
	//"fmt"
	"log"
	"os/exec"
	"strings"

	// "log"
	// "net/http"

	// "github.com/kgd-go-learing/hellogo"
	// "github.com/kgd-go-learing/part10"
	// "github.com/kgd-go-learing/integersgo"
	// "github.com/kgd-go-learing/iterationgo"
	// "github.com/kgd-go-learing/part4"
	// "github.com/kgd-go-learing/part5"
	// "github.com/kgd-go-learing/part6"
	// "github.com/kgd-go-learing/part7"
	// "github.com/kgd-go-learing/part9"
	//"github.com/kgd-go-learing/part8"
)

//go:generate go run gen/generate_readme.go

func main() {

	defer checkGenerated()

	// // function call package part1
	// hellogo.SayHello()
	// fmt.Println("***end part 1***")

	// // function call package part2
	// fmt.Println(part2.Add(4, 5))
	// fmt.Println(part2.Add(234, 12))
	// fmt.Println(part2.Multiply(10, 112))
	// fmt.Println("***end part 2***")

	// // function call package part3
	// fmt.Println(part3.Repeat("a"))
	// fmt.Println(part3.SumAllNumbers(1, 3, 2000, -4, -543))
	// fmt.Println(part3.SumPositiveNumbers(1, 3, 2000, -4))
	// fmt.Println("***end part 3***")

	// // function call package part4
	// arr := [5]int{2, 5, 6, 12, 8} //array
	// fmt.Println(part4.Sum(arr[:]))

	// slice := []int{1, 2, 3, 4, 5, 6} //slice
	// fmt.Println(part4.Sum(slice[:]))

	// fmt.Println(part4.SumAllTails(arr[:]))

	// fmt.Println("***end part 4***")

	// // function call package part5
	// rectangle := part5.Rectangle{12, 6}
	// fmt.Println(part5.Perimeter(rectangle))

	// fmt.Println(part5.Rectangle.Area(rectangle))
	// fmt.Println(part5.Shape.Area(part5.Triangle{9.0, 10.0}))

	// fmt.Println("***end part 5***")

	// // function call package part6

	// //BTC
	// wallet := part6.Wallet{}

	// wallet.Deposit(part6.Bitcoin(10))
	// fmt.Println(wallet.Balance().String())

	// wallet.Deposit(part6.Bitcoin(20))
	// fmt.Println(wallet.Balance().String())

	// withdraw := wallet.Withdraw(part6.Bitcoin(50))
	// if withdraw == nil {
	// 	fmt.Println(wallet.Balance().String())
	// } else {
	// 	fmt.Println(withdraw)
	// }

	// //USD
	// dollar := part6.WalletDol{}

	// dollar.Insert(part6.Dollar(100000))
	// fmt.Println(dollar.BalanceDollar().StringDollar())

	// fmt.Printf("Текущий балланс %d USD\n", dollar.Change(&wallet))

	// fmt.Println("***end part 6***")

	// // function call package part7
	// dictionary := part7.Dictionary{}
	// word := "test"
	// definition := "this is just a test"
	// dictionary.Add(word, definition)

	// for key, value := range dictionary {
	// 	fmt.Printf("key: %s \nvalue: %s\n", key, value)
	// }

	// fmt.Println("***end part 7***")

	// // function call package part8
	// // part8.Greet(os.Stdout, "Artem")
	// // log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(part8.MyGreeterHandler)))
	// fmt.Println("***end part 8***")

	// // function call package part9

	// part9.Mocking()

	// fmt.Println("\n***end part 9***")

	// fmt.Println(part10.CheckWebsite("http://yandex.com"))
	// fmt.Println("***end part 10***")
}

func checkGenerated(){

	version, err := exec.Command("golangci-lint", "--version").Output()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(version), "golangci-lint has version") {
		log.Fatal("please, install golangci-lint")
	}

	version, err = exec.Command("git", "version").Output()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(version), "git version") {
		log.Fatal("please, install git")
	}

	log.Println("Everything installed correct")
}