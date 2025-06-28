# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## アーキテクチャ

このプロジェクトは、macOSでアクティブなブラウザのタイトルとURLを取得し、Markdown形式でクリップボードにコピーするGoアプリケーションです。

### ディレクトリ構造
- `cmd/main.go`: エントリーポイント。AppleScriptを使用してブラウザから情報を取得
- `model/`: コアビジネスロジック
  - `tab.go`: Tab構造体とタイトル編集ロジック
  - `value/`: ドメイン、パス、タイトル、URLの値オブジェクト
  - `formatter/`: サイト別のタイトル整形処理

### 主要コンポーネント
- **Tab**: タイトルとURLのペアを管理し、サイト固有の整形を実行
- **TabFormatter**: 各サイト（GitHub、Qiita、Stack Overflow、Zenn、Atlassian、Tabelog）に特化したフォーマッター
- **Value Objects**: ドメイン、パス、タイトル、URLの型安全な表現

## 開発コマンド

### ビルドとテスト
- ビルド: `task build`
- 実行: `task run` （デフォルトブラウザ: Brave Browser）
- テスト: `task test`
- 特定テスト実行: `task test:run NAME=<テスト名>`

### コード品質
- フォーマット: `task fmt`
- Lint: `task vet`

### デバッグ
- ブラウザタイトル取得テスト: `task dev:title URL=<URL>`

## Taskfile設定
- デフォルトブラウザ: "Brave Browser"
- Go version: 1.24.3 (mise.tomlで管理)
- テストフレームワーク: testify

## 依存関係
- Go 1.24.3
- github.com/stretchr/testify (テスト用)
- mise (Go版管理)
- AppleScript（macOSブラウザ操作）