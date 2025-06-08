package formatter

import (
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
		if title.Contains("（") {
			parts := title.DisassembleIntoParts(`^([^（]+\S)\s*（.*?）.*? \| 食べログ$`)
			return parts.Assemble("%s | 食べログ", 0)
		}

		// 括弧なしの場合: "下北沢 焼とりダービーのご予約 - 下北沢/焼き鳥 | 食べログ" -> "下北沢 焼とりダービーのご予約 | 食べログ"
		parts := title.DisassembleIntoParts(`^(.+?)のご予約 - .+ \| 食べログ$`)
		return parts.Assemble("%sのご予約 | 食べログ", 0)
	}

	return title, nil
}
