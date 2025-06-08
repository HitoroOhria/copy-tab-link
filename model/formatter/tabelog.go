package formatter

import (
	"fmt"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type TabelogFormatter struct{}

func (h *TabelogFormatter) Name() string {
	return "Tabelog"
}

func (h *TabelogFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("tabelog.com")
}

func (h *TabelogFormatter) Format(path value.Path, title value.Title) (value.Title, error) {
	// 店舗ページの場合: /地域/A地域番号/A地域番号/店舗ID/
	if path.MatchString(`^/[^/]+/A\d{4}/A\d{6}/\d+/?$`) {
		// 括弧ありの場合: "下北沢 肉バル Bon （ボン【旧店名】ワイン食堂 馬肉de Bon）のご予約 - 下北沢/バル | 食べログ" -> "下北沢 肉バル Bon | 食べログ"
		matches := title.FindStringSubmatch(`^([^（]+\S)\s*（.*?）.*? \| 食べログ$`)
		if len(matches) >= 2 {
			return value.NewTitle(fmt.Sprintf("%s | 食べログ", matches[1])), nil
		}

		// 括弧なしの場合: "下北沢 焼とりダービーのご予約 - 下北沢/焼き鳥 | 食べログ" -> "下北沢 焼とりダービーのご予約 | 食べログ"
		matches = title.FindStringSubmatch(`^(.+?)のご予約 - .+ \| 食べログ$`)
		if len(matches) >= 2 {
			return value.NewTitle(fmt.Sprintf("%sのご予約 | 食べログ", matches[1])), nil
		}

		return value.Title(""), fmt.Errorf("tabelog title format not matched")
	}

	return title, nil
}
