/*
Goの制御構造の1つにループがあります。
ループは、繰り返し処理を行うための構文です。
Goでは、for文があります。
while文はありませんが、for文を使って同様のことができます。
*/

package widget

import (
	"fmt"
	"time"
)

func Loop() {
	fmt.Println("\n\nループの例")
	// for文は、繰り返し処理を行うための構文です。
	// for文は、初期化文、条件式、後処理文の3つの部分から構成されます。
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// go言語はコンパイルするため、処理が速いです。
	// そのため、ループの回数が多くても処理が速いです。
	i := 0
	now := time.Now()
	for i < 100000 {
		fmt.Println(i)
		i++
		if i%2 == 0 {
			fmt.Println(i, "は偶数")
		} else {
			fmt.Println(i, "は奇数")
		}
	}
	finish := time.Now()
	//　処理時間を計測することができます。（秒）
	fmt.Printf("処理時間: %f秒\n", finish.Sub(now).Seconds())

	// for文は、配列やスライス、マップなどのコレクションを繰り返し処理するための構文も提供しています。
	n := []int{1, 2, 3, 4, 5}
	for index, value := range n {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	// for文とif文を組み合わせて、特定の条件を満たす要素だけを処理することもできます。
	for _, value := range n {
		if value%2 == 0 {
			fmt.Printf("%dは偶数\n", value)
		} else {
			fmt.Printf("%dは奇数\n", value)
		}

		if value == 4 {
			fmt.Println("valueは4です")
			break // ループを終了する
		}
	}
}
