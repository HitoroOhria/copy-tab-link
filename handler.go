package main

import (
	"fmt"
	"regexp"
)

func (t *Tab) handleGitHub() (string, error) {
	// リポジトリルートの場合: "golang/go: The Go programming language" -> "golang/go"
	if regexp.MustCompile(`^/[^/]+/[^/]+/?$`).MatchString(t.URL.Path) {
		re := regexp.MustCompile(`^(.+): .+$`)
		if re.MatchString(t.Title) {
			replaced := re.ReplaceAllString(t.Title, "$1")
			return replaced, nil
		}
		return t.Title, nil
	}
	// Issue の場合: "cmd/cgo: fails with gcc 4.4.1 · Issue #1 · golang/go" -> "fails with gcc 4.4.1 #1"
	if regexp.MustCompile(`/issues/\d+$`).MatchString(t.URL.Path) {
		re := regexp.MustCompile(`^.+: (.+) · Issue #(\d+) · .+$`)
		matches := re.FindStringSubmatch(t.Title)
		if len(matches) < 3 {
			return "", fmt.Errorf("GitHub issue title format not matched: %s", t.Title)
		}
		replaced := fmt.Sprintf("%s #%s", matches[1], matches[2])

		return replaced, nil
	}
	// PR の場合: "net/url: Fixed url parsing with invalid slashes. by odeke-em · Pull Request #9219 · golang/go" -> "Fixed url parsing with invalid slashes. #9219"
	if regexp.MustCompile(`/pull/\d+$`).MatchString(t.URL.Path) {
		re := regexp.MustCompile(`^.+: (.+) by .+ · Pull Request #(\d+) · .+$`)
		matches := re.FindStringSubmatch(t.Title)
		if len(matches) < 3 {
			return "", fmt.Errorf("GitHub PR title format not matched: %s", t.Title)
		}
		replaced := fmt.Sprintf("%s #%s", matches[1], matches[2])

		return replaced, nil
	}

	return t.Title, nil
}
