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

// FormatTitleForEachSite はサイトに応じてタイトルを整形する
// 関数の仕様はテストを参照してください
func (t *Tab) FormatTitleForEachSite() error {
	formatted := t.Title
	hostname := t.URL.Hostname()

	switch hostname {
	case "github.com":
		if regexp.MustCompile(`/issues/\d+$`).MatchString(t.URL.Path) {
			// Issue の場合: "cmd/cgo: fails with gcc 4.4.1 · Issue #1 · golang/go" -> "fails with gcc 4.4.1 #1"
			re := regexp.MustCompile(`^.+: (.+) · Issue #(\d+) · .+$`)
			matches := re.FindStringSubmatch(t.Title)
			if len(matches) < 3 {
				return fmt.Errorf("GitHub issue title format not matched: %s", t.Title)
			}
			formatted = fmt.Sprintf("%s #%s", matches[1], matches[2])
		} else {
			// リポジトリルートの場合: "golang/go: The Go programming language" -> "golang/go"
			re := regexp.MustCompile(`^(.+): .+$`)
			replaced := re.ReplaceAllString(t.Title, "$1")
			formatted = replaced
		}
	}

	t.Title = formatted
	return nil
}

// MarkdownLink は [text](url) 形式の文字列を生成する
func (t *Tab) MarkdownLink() string {
	return fmt.Sprintf("[%s](%s)", t.Title, t.URL)
}
