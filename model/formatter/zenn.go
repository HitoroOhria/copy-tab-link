package formatter

import (
	"fmt"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type ZennFormatter struct{}

func (h *ZennFormatter) Name() string {
	return "Zenn"
}

func (h *ZennFormatter) Match(domain value.Domain) bool {
	return domain.MatchAsFQDN("zenn.dev")
}

func (h *ZennFormatter) Format(path value.Path, title string) (string, error) {
	if path.MatchString(`^/[^/]+/.+$`) {
		return fmt.Sprintf("%s - Zenn", title), nil
	}

	return title, nil
}
