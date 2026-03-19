package main

import "fmt"

func bazz() {
	fmt.Printf("Hey!!")
}

// init はどこに書いても一番最初に実行される
func init() {
	fmt.Printf("Yeah!!!\n\n")
}

func main() {
	bazz()
	var num int = 55
	var text string = "Hello World\n %d"
	fmt.Printf(text, num)
}
