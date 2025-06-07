package handler

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type TitleFormattingHandler interface {
	Name() string
	Match(u *url.URL) bool
	Handle(u *url.URL, title string) (string, error)
}

var AllHandlers = []TitleFormattingHandler{
	&GitHubHandler{},
	&ConfluenceHandler{},
	&TabelogHandler{},
}

type GitHubHandler struct{}

func (h *GitHubHandler) Name() string {
	return "GitHub"
}

func (h *GitHubHandler) Match(u *url.URL) bool {
	return u.Host == "github.com"
}

func (h *GitHubHandler) Handle(u *url.URL, title string) (string, error) {
	// リポジトリルートの場合: "golang/go: The Go programming language" -> "golang/go"
	if regexp.MustCompile(`^/[^/]+/[^/]+/?$`).MatchString(u.Path) {
		re := regexp.MustCompile(`^(.+): .+$`)
		replaced := re.ReplaceAllString(title, "$1")

		return replaced, nil
	}
	// Issue の場合: "cmd/cgo: fails with gcc 4.4.1 · Issue #1 · golang/go" -> "fails with gcc 4.4.1 #1"
	if regexp.MustCompile(`/issues/\d+$`).MatchString(u.Path) {
		re := regexp.MustCompile(`^.+: (.+) · Issue #(\d+) · .+$`)
		matches := re.FindStringSubmatch(title)
		if len(matches) < 3 {
			return "", fmt.Errorf("GitHub issue title format not matched")
		}
		replaced := fmt.Sprintf("%s #%s", matches[1], matches[2])

		return replaced, nil
	}
	// PR の場合: "net/url: Fixed url parsing with invalid slashes. by odeke-em · Pull Request #9219 · golang/go" -> "Fixed url parsing with invalid slashes. #9219"
	if regexp.MustCompile(`/pull/\d+$`).MatchString(u.Path) {
		re := regexp.MustCompile(`^.+: (.+) by .+ · Pull Request #(\d+) · .+$`)
		matches := re.FindStringSubmatch(title)
		if len(matches) < 3 {
			return "", fmt.Errorf("GitHub PR title format not matched")
		}
		replaced := fmt.Sprintf("%s #%s", matches[1], matches[2])

		return replaced, nil
	}

	return title, nil
}

type ConfluenceHandler struct{}

func (h *ConfluenceHandler) Name() string {
	return "Confluence"
}

func (h *ConfluenceHandler) Match(u *url.URL) bool {
	return strings.HasSuffix(u.Host, "atlassian.net") && strings.HasPrefix(u.Path, "/wiki/")
}

func (h *ConfluenceHandler) Handle(u *url.URL, title string) (string, error) {
	re := regexp.MustCompile(`^(.+?) - .+ - Confluence$`)
	matches := re.FindStringSubmatch(title)
	if len(matches) < 2 {
		return "", fmt.Errorf("confluence title format not matched")
	}

	return fmt.Sprintf("%s - Confluence", matches[1]), nil
}

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
