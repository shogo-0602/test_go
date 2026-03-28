package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

/*
このファイルは、GO言語のゴルーチンとチャネルの基本的な使用方法を示すサンプルコードです。
ゴルーチンは、軽量なスレッドのようなもので、関数を並行して実行することができます。
チャネルは、ゴルーチン間でデータをやり取りするための通信手段です。
*/

type Logger struct {
	logging chan string
}

var logger = Logger{
	logging: make(chan string),
}

func worker(id string, ch chan<- string) {
	for i := 0; i < 5; i++ {
		num := i + 1
		id_string := fmt.Sprintf("%d番目", num)
		message := fmt.Sprintf("Worker %s: iteration %s", id, id_string)
		ch <- message // チャネルにメッセージを送信
	}
}

// loggerLoop: logger.logging チャネルからメッセージを受け取り、標準出力に出力。
// ゴルーチン間の通信を適切に処理するため、専用のリスナーゴルーチンが必要。
func loggerLoop() {
	// チャネルからメッセージを受信し続けるループ。
	// チャネルが閉じられるまで、受信したメッセージを標準出力に表示します。
	for message := range logger.logging {
		fmt.Println(message)
	}
}

func loggerSetup() {
	ops := &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format("2006/01/02 15:04:05"))
			}
			return a
		},
	}

	logfile, err := os.OpenFile("goroutine.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}
	defer logfile.Close()

	// ロガーを初期化。標準出力・標準エラー・ログファイルの3箇所に出力。
	multi_writer := io.MultiWriter(os.Stdout, os.Stderr, logfile)
	logger := slog.New(slog.NewJSONHandler(multi_writer, ops))
	logger.Info("logger セットアップ完了")

}

func init() {
	loggerSetup()
}

func conditionalBranch() {
	// チャネルを使用して、条件分岐の例を実行するゴルーチンを起動。
	ch := make(chan string)
	fmt.Println("条件分岐のゴルーチンを起動します。")
	go func() {
		x := 10
		if x > 5 {
			ch <- "xは5より大きい"
		} else if x == 5 {
			ch <- "xは5と等しい"
		} else {
			ch <- "xは5より小さい"
		}
		close(ch) // チャネルを閉じて、受信側に終了を通知。
	}()
	fmt.Println("条件分岐のゴルーチンが実行されました。")

	// チャネルからメッセージを受信して表示。
	for message := range ch {
		fmt.Println("条件分岐の結果:", message)
	}

}

func main() {
	// logger.logging チャネルを処理するゴルーチンを起動。
	// リスナーがないとチャネルへの送信がブロックされ、デッドロックが発生するため必須。
	go loggerLoop()

	// ワーカーゴルーチンを使用して、複数の関数を並行して実行します。
	ch := make(chan string)
	for i := 1; i <= 3; i++ {
		// ゴルーチンを起動して、worker関数を実行。
		// goキーワードで軽量スレッドとしてworker関数を並行実行。
		var number_workers string = fmt.Sprintf("%d号機", i)
		go worker(number_workers, ch)
	}

	// チャネルからメッセージを受信して、ロガーチャネルに送信。
	for i := 0; i < 15; i++ {
		// <-演算子でチャネルからメッセージを受信。
		// worker側の処理が完了するまでブロック。
		message := <-ch
		logger.logging <- message
	}

	logger.logging <- "全てのワーカーが完了しました。"
	// workerゴルーチンの完了を待つため、少し遅延を入れてからチャネルをclose。
	// (実運用ではsync.WaitGroupを使用が推奨)
	close(logger.logging)

	fmt.Println("メインゴルーチンが終了します。")
	fmt.Println("goroutineの条件分岐を実行します。")
	conditionalBranch() // 条件分岐の例を実行するゴルーチンを起動。
	fmt.Println("条件分岐のゴルーチンが起動されました。")

}
