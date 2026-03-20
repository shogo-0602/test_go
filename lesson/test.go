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

func map_syudy() {
	// マップはキーと値のペアを格納するデータ構造です。
	// pythonの辞書と似ています。
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println("マップの要素数は", len(m))
	fmt.Println(m)
	fmt.Println(m["apple"])

	// マップに要素を追加することができます。
	m["orange"] = 300
	fmt.Println("orengeを追加")
	fmt.Println(m)

	// マップの要素を更新することができます。
	m["apple"] = 150
	fmt.Println("appleの値を更新")
	fmt.Println(m)

	// マップの要素を削除することができます。
	delete(m, "banana")
	fmt.Println("bananaを削除")
	fmt.Println(m)

	// マップの要素を取得することができます。
	applePrice, ok := m["apple"]
	if ok {
		fmt.Println("appleの値は", applePrice)
		fmt.Println(ok)
	} else {
		fmt.Println("appleはマップに存在しません")
	}

	// make関数を使用してマップを作成することもできます。
	m2 := make(map[string]int)
	m2["chocolate"] = 400
	fmt.Println(m2)
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
	map_syudy()

	var num int = 55
	var text string = "Hello World\n %d"

	fmt.Printf(text, num)
}
