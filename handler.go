package main

import (
	"fmt"
	"regexp"
)

func (t *Tab) handleGitHub() (string, error) {
	// リポジトリルートの場合: "golang/go: The Go programming language" -> "golang/go"
	if regexp.MustCompile(`^/[^/]+/[^/]+/?$`).MatchString(t.URL.Path) {
		re := regexp.MustCompile(`^(.+): .+$`)
		replaced := re.ReplaceAllString(t.Title, "$1")

		return replaced, nil
	}
	// Issue の場合: "cmd/cgo: fails with gcc 4.4.1 · Issue #1 · golang/go" -> "fails with gcc 4.4.1 #1"
	if regexp.MustCompile(`/issues/\d+$`).MatchString(t.URL.Path) {
		re := regexp.MustCompile(`^.+: (.+) · Issue #(\d+) · .+$`)
		matches := re.FindStringSubmatch(t.Title)
		if len(matches) < 3 {
			return "", fmt.Errorf("GitHub issue title format not matched: title = %s", t.Title)
		}
		replaced := fmt.Sprintf("%s #%s", matches[1], matches[2])

		return replaced, nil
	}
	// PR の場合: "net/url: Fixed url parsing with invalid slashes. by odeke-em · Pull Request #9219 · golang/go" -> "Fixed url parsing with invalid slashes. #9219"
	if regexp.MustCompile(`/pull/\d+$`).MatchString(t.URL.Path) {
		re := regexp.MustCompile(`^.+: (.+) by .+ · Pull Request #(\d+) · .+$`)
		matches := re.FindStringSubmatch(t.Title)
		if len(matches) < 3 {
			return "", fmt.Errorf("GitHub PR title format not matched: title = %s", t.Title)
		}
		replaced := fmt.Sprintf("%s #%s", matches[1], matches[2])

		return replaced, nil
	}

	return t.Title, nil
}

func (t *Tab) handleTabelog() (string, error) {
	// 店舗ページの場合: /地域/A地域番号/A地域番号/店舗ID/
	if regexp.MustCompile(`^/[^/]+/A\d{4}/A\d{6}/\d+/?$`).MatchString(t.URL.Path) {
		// 括弧ありの場合: "下北沢 肉バル Bon （ボン【旧店名】ワイン食堂 馬肉de Bon）のご予約 - 下北沢/バル | 食べログ" -> "下北沢 肉バル Bon | 食べログ"
		re := regexp.MustCompile(`^([^（]+\S)\s*（.*?）.*? \| 食べログ$`)
		matches := re.FindStringSubmatch(t.Title)
		if len(matches) >= 2 {
			return fmt.Sprintf("%s | 食べログ", matches[1]), nil
		}

		// 括弧なしの場合: "下北沢 焼とりダービーのご予約 - 下北沢/焼き鳥 | 食べログ" -> "下北沢 焼とりダービーのご予約 | 食べログ"
		re = regexp.MustCompile(`^(.+?)のご予約 - .+ \| 食べログ$`)
		matches = re.FindStringSubmatch(t.Title)
		if len(matches) >= 2 {
			return fmt.Sprintf("%sのご予約 | 食べログ", matches[1]), nil
		}

		return "", fmt.Errorf("tabelog title format not matched: title = %s", t.Title)
	}

	return t.Title, nil
}
