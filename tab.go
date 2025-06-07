package main

import (
	"fmt"
	"net/url"
	"regexp"
)

type Tab struct {
	Title string
	URL   *url.URL
}

func NewTab(title string, rawURL string) (*Tab, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("url.Parse: %w", err)
	}

	return &Tab{
		Title: title,
		URL:   u,
	}, nil
}

// RemoveTabNumber はタイトルからタブ番号を除去する
// "Chrome Show Tab Numbers" Extension で表示される付与される "1. Google" の番号を対象とする
func (t *Tab) RemoveTabNumber() {
	re := regexp.MustCompile(`^[0-9]\. `)
	removed := re.ReplaceAllString(t.Title, "")

	t.Title = removed
}

func (t *Tab) FormatTitleForEachSite() error {
	formatted := t.Title
	hostname := t.URL.Hostname()

	switch hostname {
	case "github.com":
		re := regexp.MustCompile(`^(.+): .+$`)
		replaced := re.ReplaceAllString(t.Title, "$1")

		formatted = replaced
	}

	t.Title = formatted
	return nil
}

// MarkdownLink は [text](url) 形式の文字列を生成する
func (t *Tab) MarkdownLink() string {
	return fmt.Sprintf("[%s](%s)", t.Title, t.URL)
}
