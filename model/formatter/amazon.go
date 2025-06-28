package formatter

import (
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
	// 商品名のみを抽出（" | " 以降を削除）
	formattedTitle := title.TrimAfter(" |")

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
