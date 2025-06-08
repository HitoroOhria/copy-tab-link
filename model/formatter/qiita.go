package formatter

import (
	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type QiitaFormatter struct{}

func (h *QiitaFormatter) Name() string {
	return "Qiita"
}

func (h *QiitaFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("qiita.com")
}

func (h *QiitaFormatter) Format(path value.Path, title value.Title) (value.Title, error) {
	if path.MatchString(`^/[^/]+/items/[a-f0-9]+$`) {
		parts := title.DisassembleIntoParts(`^(.+?) #.+ - Qiita$`)
		return parts.Assemble("%s - Qiita", 0)
	}

	return title, nil
}
