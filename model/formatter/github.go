package formatter

import (
	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type GitHubFormatter struct{}

func (h *GitHubFormatter) Name() string {
	return "GitHub"
}

func (h *GitHubFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("github.com")
}

func (h *GitHubFormatter) Format(path value.Path, title value.Title) (value.Title, error) {
	// リポジトリルートの場合: "golang/go: The Go programming language" -> "golang/go"
	if path.MatchString(`^/[^/]+/[^/]+/?$`) {
		return title.ReplaceAllString(`^(.+): .+$`, "$1"), nil
	}
	// Issue の場合: "cmd/cgo: fails with gcc 4.4.1 · Issue #1 · golang/go" -> "fails with gcc 4.4.1 #1"
	if path.MatchString(`/issues/\d+$`) {
		parts := title.DisassembleIntoParts(`^.+: (.+) · Issue #(\d+) · .+$`)
		return parts.Assemble("%s #%s", 0, 1)
	}
	// PR の場合: "net/url: Fixed url parsing with invalid slashes. by odeke-em · Pull Request #9219 · golang/go" -> "Fixed url parsing with invalid slashes. #9219"
	if path.MatchString(`/pull/\d+$`) {
		parts := title.DisassembleIntoParts(`^.+: (.+) by .+ · Pull Request #(\d+) · .+$`)
		return parts.Assemble("%s #%s", 0, 1)
	}

	return title, nil
}
