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

func (h *TabelogFormatter) Format(path value.Path, title value.Title, url *value.URL) (value.Title, *value.URL, error) {
	// 店舗ページの場合: /地域/A地域番号/A地域番号/店舗ID/
	if path.MatchString(`^/[^/]+/A\d{4}/A\d{6}/\d+/?$`) {
		// 括弧ありの場合: "下北沢 肉バル Bon （ボン【旧店名】ワイン食堂 馬肉de Bon）のご予約 - 下北沢/バル | 食べログ" -> "下北沢 肉バル Bon | 食べログ"
		if title.Contains("（") {
			parts, err := title.DisassembleIntoParts(`^([^（]+\S)\s*（.*?）.*? \| 食べログ$`)
			if err != nil {
				return "", nil, fmt.Errorf("title.DisassembleIntoParts: %w", err)
			}
			newTitle, err := parts.Assemble("%s | 食べログ", 0)
			if err != nil {
				return "", nil, fmt.Errorf("parts.Assemble: %w", err)
			}

			return newTitle, url, nil
		}

		// 括弧なしの場合: "下北沢 焼とりダービーのご予約 - 下北沢/焼き鳥 | 食べログ" -> "下北沢 焼とりダービーのご予約 | 食べログ"
		parts, err := title.DisassembleIntoParts(`^(.+?)のご予約 - .+ \| 食べログ$`)
		if err != nil {
			return "", nil, fmt.Errorf("title.DisassembleIntoParts: %w", err)
		}
		newTitle, err := parts.Assemble("%sのご予約 | 食べログ", 0)
		if err != nil {
			return "", nil, fmt.Errorf("parts.Assemble: %w", err)
		}

		return newTitle, url, nil
	}
	// コース一覧ページの場合: /地域/A地域番号/A地域番号/店舗ID/party/
	if path.MatchString(`^/[^/]+/A\d{4}/A\d{6}/\d+/party/?$`) {
		// "コース一覧 : 下北沢 焼とりダービー - 下北沢/焼き鳥 | 食べログ" -> "下北沢 焼とりダービー | 食べログ"
		parts, err := title.DisassembleIntoParts(`^コース一覧 : (.+?) - .+ \| 食べログ$`)
		if err != nil {
			return "", nil, fmt.Errorf("title.DisassembleIntoParts: %w", err)
		}
		newTitle, err := parts.Assemble("%s | 食べログ", 0)
		if err != nil {
			return "", nil, fmt.Errorf("parts.Assemble: %w", err)
		}

		// URLから /party/ を削除して店舗トップページにリダイレクト
		newURL, err := url.RemoveLastPath()
		if err != nil {
			return "", nil, fmt.Errorf("url.RemoveLastPath: %w", err)
		}

		return newTitle, newURL, nil
	}

	return title, url, nil
}
