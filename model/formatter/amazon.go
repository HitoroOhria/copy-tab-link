package formatter

import (
	"strings"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type AmazonFormatter struct{}

func (f *AmazonFormatter) Name() string {
	return "Amazon"
}

func (f *AmazonFormatter) Match(domain value.Domain) bool {
	return string(domain) == "amazon.co.jp" || string(domain) == "www.amazon.co.jp"
}

func (f *AmazonFormatter) Format(path value.Path, title value.Title, url *value.URL) (value.Title, *value.URL, error) {
	// タイトルの整形
	formattedTitle := f.formatTitle(title)

	// URLを短縮（/dp/{ASIN}の形式に変換）
	if asin, found := url.ExtractAmazonASIN(); found {
		shortURL, err := url.CreateAmazonShortURL(asin)
		if err != nil {
			return formattedTitle, url, err
		}

		return formattedTitle, shortURL, nil
	}

	// ASINが見つからない場合は元のURLを使用
	return formattedTitle, url, nil
}

func (f *AmazonFormatter) formatTitle(title value.Title) value.Title {
	titleStr := string(title)

	// "Amazon.co.jp: " のプレフィックスを削除
	if strings.HasPrefix(titleStr, "Amazon.co.jp: ") {
		titleStr = strings.TrimPrefix(titleStr, "Amazon.co.jp: ")
	}

	// " | " 以降を削除
	if idx := strings.Index(titleStr, " |"); idx != -1 {
		titleStr = titleStr[:idx]
	}

	// " eBook :" 以降を削除（Kindle版の場合）
	if idx := strings.Index(titleStr, " eBook :"); idx != -1 {
		titleStr = titleStr[:idx]
	}

	// " :" 以降を削除（物理本の場合）
	if idx := strings.Index(titleStr, " :"); idx != -1 {
		titleStr = titleStr[:idx]
	}

	return value.Title(titleStr)
}
