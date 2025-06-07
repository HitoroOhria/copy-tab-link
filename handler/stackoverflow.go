package handler

import (
	"fmt"
	"net/url"
	"regexp"
)

type StackOverflowHandler struct{}

func (h *StackOverflowHandler) Name() string {
	return "Stack Overflow"
}

func (h *StackOverflowHandler) Match(u *url.URL) bool {
	return u.Host == "stackoverflow.com"
}

func (h *StackOverflowHandler) Handle(u *url.URL, title string) (string, error) {
	if regexp.MustCompile(`^/questions/\d+/.+$`).MatchString(u.Path) {
		re := regexp.MustCompile(`^[^-]+ - (.+) - Stack Overflow$`)
		matches := re.FindStringSubmatch(title)
		if len(matches) < 2 {
			return "", fmt.Errorf("stack overflow title format not matched")
		}

		return fmt.Sprintf("%s - Stack Overflow", matches[1]), nil
	}

	return title, nil
}
