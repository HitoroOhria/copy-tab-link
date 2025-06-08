package formatter

import (
	"net/url"
)

type TabFormatter interface {
	Name() string
	Match(u *url.URL) bool
	Format(u *url.URL, title string) (newTitle string, err error)
}

var AllFormatters = []TabFormatter{
	&GitHubFormatter{},
	&QiitaFormatter{},
	&StackOverflowFormatter{},
	&ZennFormatter{},
	&ConfluenceFormatter{},
	&TabelogFormatter{},
}
