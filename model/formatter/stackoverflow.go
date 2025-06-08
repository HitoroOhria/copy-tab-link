package formatter

import (
	"fmt"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type StackOverflowFormatter struct{}

func (h *StackOverflowFormatter) Name() string {
	return "Stack Overflow"
}

func (h *StackOverflowFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("stackoverflow.com")
}

func (h *StackOverflowFormatter) Format(path value.Path, title value.Title) (value.Title, error) {
	if path.MatchString(`^/questions/\d+/.+$`) {
		matches := title.FindStringSubmatch(`^[^-]+ - (.+) - Stack Overflow$`)
		if len(matches) < 2 {
			return value.Title(""), fmt.Errorf("stack overflow title format not matched")
		}

		return value.NewTitle(fmt.Sprintf("%s - Stack Overflow", matches[1])), nil
	}

	return title, nil
}
