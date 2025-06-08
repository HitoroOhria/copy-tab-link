package model

import (
	"fmt"

	"github.com/HitoroOhria/copy_tab_link/model/formatter"
	"github.com/HitoroOhria/copy_tab_link/model/value"
)

// Tab はタイトルを編集するためのタイトルと URL のセット
// 内部にハンドラーを保持し、サイトごとのタイトルの整形を行う
type Tab struct {
	Title value.Title
	URL   *value.URL

	formatters []formatter.TabFormatter
}

func NewTab(title string, rawURL string) (*Tab, error) {
	t := value.NewTitle(title)

	u, err := value.NewURL(rawURL)
	if err != nil {
		return nil, fmt.Errorf("value.NewURL: %w", err)
	}

	return &Tab{
		Title:      t,
		URL:        u,
		formatters: formatter.AllFormatters,
	}, nil
}

// RemoveTabNumber はタイトルからタブ番号を除去する
// "Chrome Show Tab Numbers" Extension で表示される付与される "1. Google" の番号を対象とする
func (t *Tab) RemoveTabNumber() {
	t.Title = t.Title.ReplaceAllString(`^[0-9]\. `, "")
}

// FormatTitleForEachSite はサイトに応じてタイトルを整形する
// 関数の仕様はテストを参照してください
func (t *Tab) FormatTitleForEachSite() error {
	domain := value.NewDomain(t.URL)
	path := value.NewPath(t.URL)

	for _, h := range t.formatters {
		if !h.Match(domain) {
			continue
		}

		formatted, newURL, err := h.Format(path, t.Title, t.URL)
		if err != nil {
			return fmt.Errorf("formatter.Format: name = %s, title = %s, url = %s: %w", h.Name(), t.Title, t.URL, err)
		}

		t.Title = formatted
		t.URL = newURL
		return nil
	}

	return nil
}

// MarkdownLink は [text](url) 形式の文字列を生成する
func (t *Tab) MarkdownLink() string {
	return fmt.Sprintf("[%s](%s)", t.Title, t.URL)
}

func (t *Tab) SetFormatterForTest() {
	t.formatters = formatter.AllFormatters
}
