package formatter

import (
	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type AtlassianFormatter struct{}

func (h *AtlassianFormatter) Name() string {
	return "Atlassian"
}

func (h *AtlassianFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsServer("atlassian.net")
}

func (h *AtlassianFormatter) Format(path value.Path, title value.Title) (value.Title, error) {
	// Confluence の場合: "設計ドキュメント - EXAMPLE - 開発チーム - Confluence" -> "設計ドキュメント - Confluence"
	if path.MatchString(`/wiki/.+`) {
		parts := title.DisassembleIntoParts(`^(.+?) - .+ - Confluence$`)
		return parts.Assemble("%s - Confluence", 0)
	}

	return title, nil
}
