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

func (h *StackOverflowFormatter) Format(path value.Path, title value.Title, url *value.URL) (value.Title, *value.URL, error) {
	if path.MatchString(`^/questions/\d+/.+$`) {
		parts, err := title.DisassembleIntoParts(`^[^-]+ - (.+) - Stack Overflow$`)
		if err != nil {
			return "", nil, fmt.Errorf("title.DisassembleIntoParts: %w", err)
		}
		newTitle, err := parts.Assemble("%s - Stack Overflow", 0)
		if err != nil {
			return "", nil, fmt.Errorf("parts.Assemble: %w", err)
		}
		return newTitle, url, nil
	}

	return title, url, nil
}
