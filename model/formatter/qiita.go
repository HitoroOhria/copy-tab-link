package formatter

import (
	"fmt"
	"regexp"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type QiitaFormatter struct{}

func (h *QiitaFormatter) Name() string {
	return "Qiita"
}

func (h *QiitaFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("qiita.com")
}

func (h *QiitaFormatter) Format(path value.Path, title string) (string, error) {
	if path.MatchString(`^/[^/]+/items/[a-f0-9]+$`) {
		re := regexp.MustCompile(`^(.+?) #.+ - Qiita$`)
		matches := re.FindStringSubmatch(title)
		if len(matches) < 2 {
			return "", fmt.Errorf("qiita title format not matched")
		}

		return fmt.Sprintf("%s - Qiita", matches[1]), nil
	}

	return title, nil
}
