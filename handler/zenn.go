package handler

import (
	"fmt"
	"net/url"
	"regexp"
)

type ZennHandler struct{}

func (h *ZennHandler) Name() string {
	return "Zenn"
}

func (h *ZennHandler) Match(u *url.URL) bool {
	return u.Host == "zenn.dev"
}

func (h *ZennHandler) Handle(u *url.URL, title string) (string, error) {
	if regexp.MustCompile(`^/[^/]+/.+$`).MatchString(u.Path) {
		return fmt.Sprintf("%s - Zenn", title), nil
	}

	return title, nil
}
