package main

import "fmt"

func bazz() {
	fmt.Printf("Hey!!")
}

// init はどこに書いても一番最初に実行される
func init() {
	fmt.Printf("Yeah!!!\n\n")
}

func list() {
	// これは配列。サイズが固定されている。
	var numpy [3]int
	numpy[0] = 5
	numpy[1] = 55
	numpy[2] = 2
	fmt.Println("配列の要素数は", len(numpy))
	fmt.Println(numpy)
	// 配列はサイズが固定されているため、要素を追加することができません。
	numpy[0] = 5
	numpy[1] = 55
	numpy[2] = 2
	fmt.Println(numpy)
}

func slice() {
	// これはスライス。サイズが可変。
	n := []int{1, 2, 3, 4, 5}
	fmt.Println("スライスの要素数は", len(n))
	fmt.Println(n)

	// スライスはサイズが可変であるため、要素を追加することができます。
	n = append(n, 6)
	fmt.Println("スライスの要素数は", len(n))
	fmt.Println(n)
}

func dateframe() {
	// 2次元配列 5列の配列を作成(空の要素は0で埋められる)
	var dateframe = [][5]int{
		{1, 2, 3, 4},
		{6, 7, 8, 9, 10},
	}
	fmt.Println("2次元配列の要素数は", len(dateframe))
	fmt.Println(dateframe)

	// 2次元配列に要素を追加することができます。
	dateframe = append(dateframe, [5]int{11, 12, 13, 14, 15})
	fmt.Println("2次元配列の要素数は", len(dateframe))
	fmt.Println(dateframe)
}

func main() {
	bazz()
	list()
	slice()
	dateframe()

	var num int = 55
	var text string = "Hello World\n %d"

	fmt.Printf(text, num)
}
