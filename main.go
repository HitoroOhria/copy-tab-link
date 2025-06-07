package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const browserAppName = "Brave Browser"
const copyCommand = "pbcopy"

func main() {
	// ブラウザのタイトルを取得
	title, err := getBrowserTitle(browserAppName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "タイトル取得エラー: %v\n", err)
		os.Exit(1)
	}

	// ブラウザのURLを取得
	url, err := getBrowserURL(browserAppName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "URL取得エラー: %v\n", err)
		os.Exit(1)
	}

	// タイトルを編集
	tab, err := NewTab(title, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Tab作成エラー: %v\n", err)
		os.Exit(1)
	}
	tab.RemoveTabNumber()

	// Markdown形式でクリップボードにコピー
	markdownLink := tab.MarkdownLink()
	err = copyToClipboard(markdownLink)
	if err != nil {
		fmt.Fprintf(os.Stderr, "クリップボードコピーエラー: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("リンクをクリップボードにコピーしました")
}

// getBrowserTitle はブラウザのアクティブタブのタイトルを取得する
func getBrowserTitle(appName string) (string, error) {
	script := fmt.Sprintf(`tell application "%s" to get title of active tab of front window`, appName)
	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("cmd.Output: %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}

// getBrowserURL はブラウザのアクティブタブのURLを取得する
func getBrowserURL(appName string) (string, error) {
	script := fmt.Sprintf(`tell application "%s" to get URL of active tab of front window`, appName)
	cmd := exec.Command("osascript", "-e", script)

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("cmd.Output: %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}

// copyToClipboard はテキストをクリップボードにコピーする
func copyToClipboard(text string) error {
	cmd := exec.Command(copyCommand)
	cmd.Stdin = strings.NewReader(text)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("cmd.Run: %w", err)
	}

	return nil
}
