package value

import (
	"regexp"
)

// Path は URL のパス
type Path string

func NewPath(u *URL) Path {
	return Path(u.Path())
}

func (p Path) string() string {
	return string(p)
}

func (p Path) MatchString(re string) bool {
	return regexp.MustCompile(re).MatchString(p.string())
}
