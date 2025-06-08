package formatter

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type ConfluenceFormatter struct{}

func (h *ConfluenceFormatter) Name() string {
	return "Confluence"
}

func (h *ConfluenceFormatter) Match(u *url.URL) bool {
	return strings.HasSuffix(u.Host, "atlassian.net") && strings.HasPrefix(u.Path, "/wiki/")
}

func (h *ConfluenceFormatter) Format(u *url.URL, title string) (string, error) {
	re := regexp.MustCompile(`^(.+?) - .+ - Confluence$`)
	matches := re.FindStringSubmatch(title)
	if len(matches) < 2 {
		return "", fmt.Errorf("confluence title format not matched")
	}

	return fmt.Sprintf("%s - Confluence", matches[1]), nil
}
