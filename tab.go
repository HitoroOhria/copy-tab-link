package main

import (
	"fmt"
	"regexp"
)

type Tab struct {
	Title string
	URL   string
}

func NewTab(title string, url string) *Tab {
	return &Tab{
		Title: title,
		URL:   url,
	}
}

// RemoveTabNumber はタイトルからタブ番号を除去する
// "Chrome Show Tab Numbers" Extension で表示される付与される "1. Google" の番号を対象とする
func (t *Tab) RemoveTabNumber() {
	re := regexp.MustCompile(`^[0-9]\. `)
	removed := re.ReplaceAllString(t.Title, "")

	t.Title = removed
}

// MarkdownLink は [text](url) 形式の文字列を生成する
func (t *Tab) MarkdownLink() string {
	return fmt.Sprintf("[%s](%s)", t.Title, t.URL)
}
