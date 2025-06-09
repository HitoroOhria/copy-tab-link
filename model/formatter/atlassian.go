package formatter

import (
	"fmt"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type AtlassianFormatter struct{}

func (h *AtlassianFormatter) Name() string {
	return "Atlassian"
}

func (h *AtlassianFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsServer("atlassian.net")
}

func (h *AtlassianFormatter) Format(path value.Path, title value.Title, url *value.URL) (value.Title, *value.URL, error) {
	// Confluence の場合: "設計ドキュメント - EXAMPLE - 開発チーム - Confluence" -> "設計ドキュメント - Confluence"
	if path.MatchString(`/wiki/.+`) {
		parts, err := title.DisassembleIntoParts(`^(.+?) - .+ - Confluence$`)
		if err != nil {
			return "", nil, fmt.Errorf("title.DisassembleIntoParts: %w", err)
		}
		newTitle, err := parts.Assemble("%s - Confluence", 0)
		if err != nil {
			return "", nil, fmt.Errorf("parts.Assemble: %w", err)
		}
		return newTitle, url, nil
	}

	return title, url, nil
}
