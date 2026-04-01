package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func init() {
	fmt.Printf("Yeah!!!\n\n")
}

func test_gui() {
	fmt.Println("アプリを作成します。")
	// Fyneを使用するには、まずアプリケーションを作成します。
	my_app := app.New() // Newは新しいアプリケーションを作成する関数です。

	// ウィンドウを作成します。
	my_window := my_app.NewWindow("Hello")

	// ウィンドウのサイズを設定します。
	my_window.Resize(fyne.NewSize(400, 300))

	// ウィンドウをスクリーン中央に配置します。
	my_window.CenterOnScreen()

	fmt.Println("ウィンドウを作成し、サイズと位置を設定しました。内容を設定してウィンドウを表示します。")

	// ウィンドウの内容を設定します。
	my_window.SetContent( // VBoxは垂直にウィジェットを配置するコンテナです。
		container.NewVBox( // NewLabelはテキストを表示するウィジェットです。
			widget.NewLabel("Hello Fyne!"), // NewButtonはボタンを作成するウィジェットです。
			widget.NewButton("クリックして!", func() { // ボタンがクリックされたときに実行される関数を定義します。
				// サブウィンドウを作成します。
				sub_window := my_app.NewWindow("Sub Window")      // NewWindowは新しいウィンドウを作成する関数です。
				sub_window.SetContent(widget.NewLabel("サブウィンドウ")) // サブウィンドウの内容を設定します。
				sub_window.Resize(fyne.NewSize(200, 150))         // サブウィンドウのサイズを設定します。
				sub_window.Show()                                 // サブウィンドウを表示します。
				sub_window.CenterOnScreen()                       // サブウィンドウをスクリーン中央に配置します。
				sub_window.SetContent(
					// VBoxは垂直にウィジェットを配置するコンテナです。
					container.NewVBox(
						widget.NewLabel("This is a sub window!"), // NewLabelはテキストを表示するウィジェットです。
						widget.NewButton("閉じる", func() { // NewButtonはボタンを作成するウィジェットです。
							sub_window.Close() // サブウィンドウを閉じます。
						}),
					))
			}),
		),
	)

	fmt.Println("内容を設定しました。ウィンドウを表示してアプリを実行します。")
	my_window.Show() // ウィンドウを表示します。
	fmt.Println("ウィンドウを表示しました。アプリのイベントループを開始します。")
	my_app.Run() // アプリケーションのイベントループを開始します。

}

func main() {
	// FyneはGoでGUIアプリケーションを作成するためのライブラリです。
	test_gui()
}
