package formatter

import (
	"github.com/HitoroOhria/copy_tab_link/model/value"
)

type TabFormatter interface {
	Name() string
	Match(domain value.Domain) bool
	Format(path value.Path, title value.Title) (value.Title, error)
}

var AllFormatters = []TabFormatter{
	&GitHubFormatter{},
	&QiitaFormatter{},
	&StackOverflowFormatter{},
	&ZennFormatter{},
	&AtlassianFormatter{},
	&TabelogFormatter{},
}
