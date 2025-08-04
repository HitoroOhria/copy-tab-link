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
	// パターン1: "Amazon.co.jp: " で始まるタイトル（物理本・Kindle）
	if parts, err := title.DisassembleIntoParts(`^Amazon\.co\.jp: (.+?) (?:eBook )?:.*$`); err == nil {
		if formattedTitle, err := parts.Assemble("%s", 0); err == nil {
			return formattedTitle
		}
	}

	// パターン2: " | " で区切られたタイトル（通常の商品）
	if parts, err := title.DisassembleIntoParts(`^(.+?) \|.*$`); err == nil {
		if formattedTitle, err := parts.Assemble("%s", 0); err == nil {
			return formattedTitle
		}
	}

	return title
}
