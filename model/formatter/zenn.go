package formatter

import (
	"fmt"
	"net/url"
	"regexp"
)

type ZennFormatter struct{}

func (h *ZennFormatter) Name() string {
	return "Zenn"
}

func (h *ZennFormatter) Match(u *url.URL) bool {
	return u.Host == "zenn.dev"
}

func (h *ZennFormatter) Format(u *url.URL, title string) (string, error) {
	if regexp.MustCompile(`^/[^/]+/.+$`).MatchString(u.Path) {
		return fmt.Sprintf("%s - Zenn", title), nil
	}

	return title, nil
}
