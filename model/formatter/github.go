package formatter

import (
	"fmt"
	"regexp"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type GitHubFormatter struct{}

func (h *GitHubFormatter) Name() string {
	return "GitHub"
}

func (h *GitHubFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("github.com")
}

func (h *GitHubFormatter) Format(path value.Path, title string) (string, error) {
	// リポジトリルートの場合: "golang/go: The Go programming language" -> "golang/go"
	if path.MatchString(`^/[^/]+/[^/]+/?$`) {
		re := regexp.MustCompile(`^(.+): .+$`)
		replaced := re.ReplaceAllString(title, "$1")

		return replaced, nil
	}
	// Issue の場合: "cmd/cgo: fails with gcc 4.4.1 · Issue #1 · golang/go" -> "fails with gcc 4.4.1 #1"
	if path.MatchString(`/issues/\d+$`) {
		re := regexp.MustCompile(`^.+: (.+) · Issue #(\d+) · .+$`)
		matches := re.FindStringSubmatch(title)
		if len(matches) < 3 {
			return "", fmt.Errorf("GitHub issue title format not matched")
		}
		replaced := fmt.Sprintf("%s #%s", matches[1], matches[2])

		return replaced, nil
	}
	// PR の場合: "net/url: Fixed url parsing with invalid slashes. by odeke-em · Pull Request #9219 · golang/go" -> "Fixed url parsing with invalid slashes. #9219"
	if path.MatchString(`/pull/\d+$`) {
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
