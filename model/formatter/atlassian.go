package formatter

import (
	"fmt"
	"regexp"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type AtlassianFormatter struct{}

func (h *AtlassianFormatter) Name() string {
	return "Atlassian"
}

func (h *AtlassianFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsServer("atlassian.net")
}

func (h *AtlassianFormatter) Format(path value.Path, title string) (string, error) {
	// Confluence の場合: "設計ドキュメント - EXAMPLE - 開発チーム - Confluence" -> "設計ドキュメント - Confluence"
	if path.MatchString(`/wiki/.+`) {
		re := regexp.MustCompile(`^(.+?) - .+ - Confluence$`)
		matches := re.FindStringSubmatch(title)
		if len(matches) < 2 {
			return "", fmt.Errorf("confluence title format not matched")
		}

		return fmt.Sprintf("%s - Confluence", matches[1]), nil
	}

	return title, nil
}
