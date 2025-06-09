package formatter

import (
	"fmt"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type GitHubFormatter struct{}

func (h *GitHubFormatter) Name() string {
	return "GitHub"
}

func (h *GitHubFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("github.com")
}

func (h *GitHubFormatter) Format(path value.Path, title value.Title, url *value.URL) (value.Title, *value.URL, error) {
	// リポジトリルートの場合: "golang/go: The Go programming language" -> "golang/go"
	if path.MatchString(`^/[^/]+/[^/]+/?$`) {
		return title.ReplaceAllString(`^(.+): .+$`, "$1"), url, nil
	}
	// Issue の場合: "cmd/cgo: fails with gcc 4.4.1 · Issue #1 · golang/go" -> "fails with gcc 4.4.1 #1"
	if path.MatchString(`/issues/\d+$`) {
		parts, err := title.DisassembleIntoParts(`^.+: (.+) · Issue #(\d+) · .+$`)
		if err != nil {
			return "", nil, fmt.Errorf("title.DisassembleIntoParts: %w", err)
		}
		newTitle, err := parts.Assemble("%s #%s", 0, 1)
		if err != nil {
			return "", nil, fmt.Errorf("parts.Assemble: %w", err)
		}

		return newTitle, url, nil
	}
	// PR の場合: "net/url: Fixed url parsing with invalid slashes. by odeke-em · Pull Request #9219 · golang/go" -> "Fixed url parsing with invalid slashes. #9219"
	if path.MatchString(`/pull/\d+$`) {
		parts, err := title.DisassembleIntoParts(`^(.+) by .+ · Pull Request #(\d+) · .+$`)
		if err != nil {
			return "", nil, fmt.Errorf("title.DisassembleIntoParts: %w", err)
		}
		newTitle, err := parts.Assemble("%s #%s", 0, 1)
		if err != nil {
			return "", nil, fmt.Errorf("parts.Assemble: %w", err)
		}

		return newTitle, url, nil
	}

	return title, url, nil
}
