package formatter

import (
	"fmt"

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
	if path.MatchString(`^/[^/]+/items/[a-f0-9]+/?$`) {
		parts, err := title.DisassembleIntoParts(`^(.+?) #.+ - Qiita$`)
		if err != nil {
			return "", nil, fmt.Errorf("title.DisassembleIntoParts: %w", err)
		}
		newTitle, err := parts.Assemble("%s - Qiita", 0)
		if err != nil {
			return "", nil, fmt.Errorf("parts.Assemble: %w", err)
		}

		return newTitle, url, nil
	}

	return title, url, nil
}
