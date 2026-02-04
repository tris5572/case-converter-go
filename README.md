# case-converter-go

文字列のケース変換を行うコマンドラインツールです。

## 機能

- 入力文字列を以下のスタイルへ変換
  - camelCase, PascalCase, snake_case, kebab-case, lowercase, UPPERCASE, UPPER_SNAKE_CASE, Train-Case
- コマンドライン引数または標準入力（stdin）から入力を受け取る
- `-`（単一ダッシュ）を引数の位置に書くと、その位置を stdin の各行で置換
- `--help` でヘルプを表示

## インストール

Go がインストールされていれば簡単に実行できます。

```bash
# 実行（開発中）
go run . --camel "Hello World"

# ビルドして実行ファイルを作成
go build -o case-converter-go
./case-converter-go --snake "Hello World"

# （任意）インストール
go install
```

## 使い方

### 基本

```bash
# 引数から変換
./case-converter-go --camel "Hello World"
# 標準入力から変換（引数が空）
echo "hello world" | ./case-converter-go --camel
# '-' をその場で stdin の行に置き換える
echo -e "a\nb" | ./case-converter-go before - after
```

### オプション（抜粋）

- `-h`, `--help` — ヘルプを表示
- `-` — stdin からの行で置換（または引数が空のときに stdin を読み込む）
- `--camel` / `camel` — camelCase
- `--pascal` / `pascal` — PascalCase
- `--snake` / `snake_case` — snake_case
- `--kebab` / `kebab-case` — kebab-case
- `--lower` / `lowercase` — lowercase
- `--upper` / `uppercase` — UPPERCASE
- `--upper-snake` / `upper_snake_case` — UPPER_SNAKE_CASE
- `--train` / `train-case` — Train-Case

> 注意: フラグは後勝ちで、複数のスタイル指定がある場合、最後の指定が適用されます。

## テスト

プロジェクトのテストは以下で実行できます:

```bash
go test ./...
```
