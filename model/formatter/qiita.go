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

func (h *QiitaFormatter) Format(path value.Path, title value.Title) (value.Title, error) {
	if path.MatchString(`^/[^/]+/items/[a-f0-9]+$`) {
		matches := title.FindStringSubmatch(`^(.+?) #.+ - Qiita$`)
		if len(matches) < 2 {
			return value.Title(""), fmt.Errorf("qiita title format not matched")
		}

		return value.NewTitle(fmt.Sprintf("%s - Qiita", matches[1])), nil
	}

	return title, nil
}
