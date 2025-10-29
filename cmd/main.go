package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/HitoroOhria/copy_tab_link/model"
)

func getArgs() string {
	defaultValue := os.Getenv("BROWSER_NAME")
	browserAppName := flag.String("browser-name", defaultValue, "Browser app name.")
	flag.Parse()

	return *browserAppName
}

// アクティブなブラウザのタイトルとリンクを取得し、Markdown形式でクリップボードにコピーする
// 動作環境の対象は macOS である
func main() {
	browserAppName := getArgs()

	// ブラウザのタイトルを取得
	title, err := getBrowserTitle(browserAppName)
	if err != nil {
		handleError(err, "failed to get title")
		return
	}

	// ブラウザのURLを取得
	url, err := getBrowserURL(browserAppName)
	if err != nil {
		handleError(err, "failed to get url")
		return
	}

	// タイトルを編集
	tab, err := model.NewTab(title, url)
	if err != nil {
		handleError(err, "failed to new Tab struct")
		return
	}
	tab.RemoveTabNumber()
	err = tab.FormatForEachSite()
	if err != nil {
		handleError(err, "failed to format title")
		return
	}

	// Markdown形式でクリップボードにコピー
	markdownLink := tab.MarkdownLink()
	err = copyToClipboard(markdownLink)
	if err != nil {
		handleError(err, "failed to copy to clipboard")
		return
	}

	fmt.Println("success to copy title to clipboard")
}

func handleError(err error, msg string) {
	_, printErr := fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
	if printErr != nil {
		log.Fatalf("fmt.Fprintf: %v\n%s: %v\n", printErr, msg, err)
	}

	os.Exit(1)
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
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(text)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("cmd.Run: %w", err)
	}

	return nil
}
