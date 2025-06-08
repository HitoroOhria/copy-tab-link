package formatter

import (
	"fmt"
	"net/url"
	"regexp"
)

type QiitaFormatter struct{}

func (h *QiitaFormatter) Name() string {
	return "Qiita"
}

func (h *QiitaFormatter) Match(u *url.URL) bool {
	return u.Host == "qiita.com"
}

func (h *QiitaFormatter) Format(u *url.URL, title string) (string, error) {
	if regexp.MustCompile(`^/[^/]+/items/[a-f0-9]+$`).MatchString(u.Path) {
		re := regexp.MustCompile(`^(.+?) #.+ - Qiita$`)
		matches := re.FindStringSubmatch(title)
		if len(matches) < 2 {
			return "", fmt.Errorf("qiita title format not matched")
		}

		return fmt.Sprintf("%s - Qiita", matches[1]), nil
	}

	return title, nil
}
