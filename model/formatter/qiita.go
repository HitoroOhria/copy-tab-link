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

func (h *QiitaFormatter) Format(path value.Path, title value.Title, url *value.URL) (value.Title, *value.URL, error) {
	if path.MatchString(`^/[^/]+/items/[a-f0-9]+$`) {
		parts := title.DisassembleIntoParts(`^(.+?) #.+ - Qiita$`)
		newTitle, err := parts.Assemble("%s - Qiita", 0)
		return newTitle, url, err
	}

	return title, url, nil
}
