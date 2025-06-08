package formatter

import (
	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type ZennFormatter struct{}

func (h *ZennFormatter) Name() string {
	return "Zenn"
}

func (h *ZennFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("zenn.dev")
}

func (h *ZennFormatter) Format(path value.Path, title value.Title) (value.Title, error) {
	// 記事の場合: "【初心者歓迎】第２回 AI Agent Hackathon、開催決定！" -> "【初心者歓迎】第２回 AI Agent Hackathon、開催決定！ - Zenn"
	if path.MatchString(`^/[^/]+/.+$`) {
		return title.AddSuffix(" - Zenn"), nil
	}

	return title, nil
}
