package value

import "regexp"

type Title string

func NewTitle(str string) Title {
	return Title(str)
}

func (t Title) string() string {
	return string(t)
}

func (t Title) ReplaceAllString(re string, repl string) Title {
	replaced := regexp.MustCompile(re).ReplaceAllString(t.string(), repl)
	return NewTitle(replaced)
}

func (t Title) FindStringSubmatch(re string) []string {
	return regexp.MustCompile(re).FindStringSubmatch(t.string())
}
