package formatter

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type ZennFormatter struct{}

func (h *ZennFormatter) Name() string {
	return "Zenn"
}

func (h *ZennFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("zenn.dev")
}

func (h *ZennFormatter) Format(u *url.URL, title string) (string, error) {
	if regexp.MustCompile(`^/[^/]+/.+$`).MatchString(u.Path) {
		return fmt.Sprintf("%s - Zenn", title), nil
	}

	return title, nil
}
