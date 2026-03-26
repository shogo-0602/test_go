/*
条件分岐は、条件に応じて異なるコードを実行するための構文です。
Goでは、if文、switch文、select文があります。
*/

package widget

import (
	"fmt"
)

func ConditionalBranch() {
	fmt.Println("\n\n条件分岐の例")
	x := 10

	// if文は、条件が真の場合にコードを実行するための構文です。
	if x > 5 {
		fmt.Println("xは5より大きい")
	} else if x == 5 {
		fmt.Println("xは5と等しい")
	} else {
		fmt.Println("xは5より小さい")
	}

	// switch文は、複数の条件を評価するための構文です。
	switch x {
	case 1:
		fmt.Println("xは1")
	case 10:
		fmt.Println("xは10")
	default:
		fmt.Println("xは1でも10でもない")
	}

	/*
		select文は、複数のチャネル操作を待ち受けるための構文です。
		select文は、チャネルからのデータの受信や、
		チャネルへのデータの送信を待ち受けることができます。
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)
	data := 42
	select {
	case <-ch1:
		fmt.Println("ch1からデータを受信")
	case ch2 <- data:
		fmt.Println("ch2にデータを送信")
	default:
		fmt.Println("どのチャネルも操作されなかった")
	}

}
