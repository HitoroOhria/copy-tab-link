package handler

import (
	"net/url"
)

type TitleFormattingHandler interface {
	Name() string
	Match(u *url.URL) bool
	Handle(u *url.URL, title string) (string, error)
}

var AllHandlers = []TitleFormattingHandler{
	&GitHubHandler{},
	&QiitaHandler{},
	&ConfluenceHandler{},
	&TabelogHandler{},
}
