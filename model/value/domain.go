package value

import (
	"strings"
)

// Domain は URL のドメイン
type Domain string

func NewDomain(u *URL) Domain {
	return Domain(u.Hostname())
}

func (d Domain) string() string {
	return string(d)
}

// MatchAsFQDN は FQDN レベルで合致するか判定する
func (d Domain) MatchAsFQDN(target string) bool {
	return d.string() == target
}

// MatchAsServer はサーバードメインレベルで合致するか判定する
// 例えば "foo.example.com" の場合、"example.com" と合致するかを判定する
func (d Domain) MatchAsServer(target string) bool {
	domains := strings.Split(d.string(), ".")
	serverDomains := domains[1:]
	serverDomain := strings.Join(serverDomains, ".")

	return serverDomain == target
}
