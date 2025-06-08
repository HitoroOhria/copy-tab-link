package formatter

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type StackOverflowFormatter struct{}

func (h *StackOverflowFormatter) Name() string {
	return "Stack Overflow"
}

func (h *StackOverflowFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("stackoverflow.com")
}

func (h *StackOverflowFormatter) Format(u *url.URL, title string) (string, error) {
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
