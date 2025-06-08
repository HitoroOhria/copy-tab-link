package formatter

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type ConfluenceFormatter struct{}

func (h *ConfluenceFormatter) Name() string {
	return "Confluence"
}

func (h *ConfluenceFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsServer("atlassian.net")
}

func (h *ConfluenceFormatter) Format(u *url.URL, title string) (string, error) {
	if strings.HasPrefix(u.Path, "/wiki/") {
		re := regexp.MustCompile(`^(.+?) - .+ - Confluence$`)
		matches := re.FindStringSubmatch(title)
		if len(matches) < 2 {
			return "", fmt.Errorf("confluence title format not matched")
		}

		return fmt.Sprintf("%s - Confluence", matches[1]), nil
	}

	return title, nil
}
