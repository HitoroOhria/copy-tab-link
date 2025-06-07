package model

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/HitoroOhria/copy_tab_link/handler"
)

// Tab はタイトルを編集するためのタイトルと URL のセット
// 内部にハンドラーを保持し、サイトごとのタイトルの整形を行う
type Tab struct {
	Title string
	URL   *url.URL

	handlers []handler.TitleFormattingHandler
}

func NewTab(title string, rawURL string) (*Tab, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("url.Parse: %w", err)
	}

	return &Tab{
		Title:    title,
		URL:      u,
		handlers: handler.AllHandlers,
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
	for _, h := range t.handlers {
		if !h.Match(t.URL) {
			continue
		}

		formatted, err := h.Handle(t.URL, t.Title)
		if err != nil {
			return fmt.Errorf("handler.Handle: name = %s, title = %s, url = %s: %w", h.Name(), t.Title, t.URL, err)
		}

		t.Title = formatted
		return nil
	}

	return nil
}

// MarkdownLink は [text](url) 形式の文字列を生成する
func (t *Tab) MarkdownLink() string {
	return fmt.Sprintf("[%s](%s)", t.Title, t.URL)
}

func (t *Tab) SetHandlerForTest() {
	t.handlers = handler.AllHandlers
}
