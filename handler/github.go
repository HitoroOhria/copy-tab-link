package handler

import (
	"fmt"
	"net/url"
	"regexp"
)

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
