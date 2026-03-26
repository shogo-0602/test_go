package main

import (
	"errors"
	"io"
	"log"
	"log/slog"
	"os"
)

type FileProcessor struct {
	filename  string
	data_init string
}

// グローバル変数を定義します。これで複数の関数でファイル名と初期データを共有できます。
// 　構造体を使用して、ファイル名と初期データを定義します。
var file_data FileProcessor
var file string
var data string
var logger *slog.Logger
var logFile *os.File

func loggerSetup() {
	// ログファイルの名前を定義します。
	logfile := "app.log"

	// slog.HandlerOptionsを使用して、ログの出力形式や内容をカスタマイズします。
	// AddSource: trueに設定すると、ログにソースコードの位置が追加されます。
	ops := &slog.HandlerOptions{
		AddSource: true, // ソースコードの位置をログに追加します。
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// ログのタイムスタンプを人間が読みやすい形式に変換します。
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format("2006/01/02 15:04:05"))
			}
			// ログレベルを文字列に変換します。
			if a.Key == slog.LevelKey {
				a.Value = slog.StringValue(a.Value.String())
			}
			return a
		},
	}

	// ログファイルを開きます。存在しない場合は作成されます。
	var err error
	/*
		os.O_CREATE: ファイルが存在しない場合に新しいファイルを作成します。
		os.O_WRONLY: ファイルを開くときに書き込み専用で開きます。
		os.O_APPEND: ファイルに書き込むときに、既存の内容の末尾に追加します。
		0644: ファイルのパーミッションを設定します。ここでは、所有者が読み書きでき、グループとその他のユーザーが読み取りのみできるように設定しています。
	*/
	logFile, err = os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("ログファイルのオープンに失敗: %v", err)
	}

	// ロガーを初期化します。ここでは、標準出力と標準エラーの両方にログを出力するように設定しています。
	/*
		os.Stdout: 標準出力にログを出力します。
		os.Stderr: 標準エラーにログを出力します。
		logFile: ログファイルにログを出力します。
		ops: slog.HandlerOptionsで定義したオプションを使用して、ログの出力形式や内容をカスタマイズします。
	*/
	multi_writer := io.MultiWriter(os.Stdout, os.Stderr, logFile)
	logger = slog.New(slog.NewJSONHandler(multi_writer, ops))
}

func init() {
	loggerSetup()

	// 構造体を使用して、ファイル名と初期データを定義します。
	file_data := FileProcessor{
		filename:  "dt/test.csv",
		data_init: "id,name,age\n1,Alice,30\n2,Bob,25\n",
	}

	// ファイル名と初期データを変数に格納します。
	file = file_data.filename
	data = file_data.data_init

	// ファイルの状態を確認します。
	_, err := os.Stat(file)

	if err != nil {
		logger.Error("ファイルの状態の確認に失敗", "file", file, "error", err)
		// エラーがファイルの存在に関連しているかを確認します。
		if errors.Is(err, os.ErrNotExist) {
			logger.Info("ファイルが存在しません。ファイルを作成します。", "file", file)
			// ファイルを作成する処理をここに追加します。
			f, err := os.Create(file)
			if err != nil {
				logger.Error("ファイルの作成に失敗", "file", file, "error", err)
				return
			}

			// 作成したファイルに初期値を書き込みます。
			_, err = f.WriteString(data)
			if err != nil {
				logger.Error("ファイルへの書き込みに失敗", "file", file, "error", err)
			}
			f.Close()

		} else {
			logger.Error("ファイルの状態の確認に失敗", "file", file, "error", err)
		}
	} else {
		logger.Info("ファイルが既に存在しています。", "file", file)
	}
}

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// ファイルの内容を読み取る処理
	return io.ReadAll(f)
}

func main() {
	//　GO言語では、エラー処理が重要で関数ごとにエラーを返すことが一般的です。
	// 戻り値は、値とエラーの2つを返すことが多いです。

	f, err := os.Open(file_data.filename)

	// エラーが発生した場合の処理
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logger.Info("ファイルが存在しません", "file", file_data.filename)
		} else {
			logger.Error("ファイルのオープンに失敗", "file", file_data.filename, "error", err)
		}
	}

	// エラーが発生しなかった場合の処理
	defer f.Close() // ファイルを閉じる処理

	// ファイルを読み取る処理
	data, err := io.ReadAll(f)
	if err != nil {
		logger.Error("ファイルの読み込みに失敗", "file", file_data.filename, "error", err)
	} else {
		logger.Info("ファイルの内容", "file", file_data.filename, "data", string(data))
	}

	// 存在しないファイルを開く例
	data, err = readFile("dt/nonexistent.csv")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logger.Info("ファイルが存在しません", "file", "dt/nonexistent.csv")
		} else {
			logger.Error("nonexistent.csvのオープンに失敗", "file", "dt/nonexistent.csv", "error", err)
			return
		}
	}

	// ここには到達しないはずですが、もし到達した場合はファイルの内容を表示します。
	logger.Info("nonexistent.csvの内容: ", "data", string(data))

	// ログファイルをクローズします。
	logFile.Close()
}
