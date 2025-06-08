package formatter

import (
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
		parts := title.DisassembleIntoParts(`^[^-]+ - (.+) - Stack Overflow$`)
		return parts.Assemble("%s - Stack Overflow", 0)
	}

	return title, nil
}
