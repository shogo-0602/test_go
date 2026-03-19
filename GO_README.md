# Go Language Workspace

このワークスペースは、Go言語の学習および開発用環境です。
Windowsのセキュリティポリシー制限を回避し、デバッグを可能にするための特殊な構成を採用しています。

## 1. 開発環境の構成
- **OS**: Windows 11
- **Go Version**: 1.16+
- **Toolbox (重要)**: `C:\Program Files\GoTools\bin`
  - セキュリティブロックを回避するため、デバッガ（dlv.exe）および各種ツールをここに手動で配置しています。

## 2. VS Code の必須設定
デバッグ（F5）を正常に動作させるため、`.vscode/settings.json` に以下のパス指定が必要です。

```json
{
    "go.alternateTools": {
        "dlv": "C:\\Program Files\\GoTools\\bin\\dlv.exe"
    }
}
```

## 3. 基本操作コマンド
プログラムの実行 (実行ファイルの生成なし)
```
go run <ファイル名>.go
```

バイナリのビルド (Windows実行用 .exe の生成)
```
go build -o app.exe <ファイル名>.go
./app.exe
```

#### デバッグ手順
1. ソースコードの行番号左をクリックし、**ブレークポイント（赤点）**を設置。
2. F5キー を押してデバッグ開始。
3. 変数パネルやデバッグコンソールで状態を確認。

## 4. トラブルシューティング
- spawn UNKNOWN エラー: dlv.exe の実行がOSにブロックされています。ファイルが C:\Program Files\GoTools\bin にあるか確認してください。
- go.mod がない: ターミナルで go mod init lesson を実行して初期化してください。

---

### 2. Go_Guide.md
Python経験者であるあなたに向けた、Go言語の要点をまとめた「技術解説書」です。

```markdown
# Go Language Technical Guide (for Python Developers)

Goは Google によって設計された、**「シンプルさ」と「安全性」**を追求した静的型付けのコンパイル言語です。

## 1. Pythonとの決定的な違い
- **コンパイルが必要**: 実行前にソースコードを機械語（.exe）に変換するため、実行速度が極めて高速です。
- **静的型付け**: 変数の型を実行前に決定します。型の間違いをコンパイル時に検知できるため、バグが混入しにくい設計です。
- **未使用変数の禁止**: 宣言して使っていない変数やインポートがあるとコンパイルエラーになります。

## 2. 変数宣言のスタイル
```go
// 1. 基本形 (型を明示)
var count int = 10

// 2. 短縮形 (型推論 / 関数内でのみ使用可能)
// Pythonの変数宣言に近い感覚で使えます
message := "Hello Go!"
```

## 3.「Verb（動詞）」

| Verb | 意味 | 渡すべきデータの型 |
| :--- | :--- | :--- |
| **%d** | 10進数の整数 | `int` |
| **%s** | 文字列 | `string` |
| **%f** | 浮動小数点（小数） | `float64` |
| **%t** | 真偽値 (true/false) | `bool` |
| **%v** | 「いい感じに」出す (value) | 何でもOK |
| **%T** | そのデータの「型」を出す | 何でもOK |
## 4. 制御構文のルール
- if文: 条件式にカッコ () は不要です。
- for文: Goには while がありません。すべてのループを for で記述します。

```go
// 基本的なループ
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## 5.構造体 (Structs)
Pythonの「クラス」に近いものですが、よりシンプルです。

```go
type Player struct {
    Name  string
    Level int
}

p := Player{Name: "Gopher", Level: 10}
fmt.Printf("%+v\n", p) // フィールド名付きで中身を表示
```

## 6.学習のロードマップ
1. スライス: Pythonのリストに近い動的配列。append 関数をマスターする。
2. エラーハンドリング: if err != nil という書き方を覚える。
3. ポインタ: メモリを効率よく扱うための「アドレス」の概念を理解する。

