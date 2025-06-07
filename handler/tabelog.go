package handler

import (
	"fmt"
	"net/url"
	"regexp"
)

type TabelogHandler struct{}

func (h *TabelogHandler) Name() string {
	return "Tabelog"
}

func (h *TabelogHandler) Match(u *url.URL) bool {
	return u.Host == "tabelog.com"
}

func (h *TabelogHandler) Handle(u *url.URL, title string) (string, error) {
	// 店舗ページの場合: /地域/A地域番号/A地域番号/店舗ID/
	if regexp.MustCompile(`^/[^/]+/A\d{4}/A\d{6}/\d+/?$`).MatchString(u.Path) {
		// 括弧ありの場合: "下北沢 肉バル Bon （ボン【旧店名】ワイン食堂 馬肉de Bon）のご予約 - 下北沢/バル | 食べログ" -> "下北沢 肉バル Bon | 食べログ"
		re := regexp.MustCompile(`^([^（]+\S)\s*（.*?）.*? \| 食べログ$`)
		matches := re.FindStringSubmatch(title)
		if len(matches) >= 2 {
			return fmt.Sprintf("%s | 食べログ", matches[1]), nil
		}

		// 括弧なしの場合: "下北沢 焼とりダービーのご予約 - 下北沢/焼き鳥 | 食べログ" -> "下北沢 焼とりダービーのご予約 | 食べログ"
		re = regexp.MustCompile(`^(.+?)のご予約 - .+ \| 食べログ$`)
		matches = re.FindStringSubmatch(title)
		if len(matches) >= 2 {
			return fmt.Sprintf("%sのご予約 | 食べログ", matches[1]), nil
		}

		return "", fmt.Errorf("tabelog title format not matched")
	}

	return title, nil
}
