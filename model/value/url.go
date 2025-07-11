package value

import (
	"fmt"
	"net/url"
	"regexp"
)

type URL struct {
	u *url.URL
}

func NewURL(rawURL string) (*URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("url.Parse: %w", err)
	}

	return &URL{u: u}, nil
}

func (u *URL) duplicate() *URL {
	newURL, _ := NewURL(u.String())
	return newURL
}

func (u *URL) updatePath(rawURL string) error {
	newURL, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("url.Parse: %w", err)
	}

	u.u = newURL
	return nil
}

func (u *URL) String() string {
	return u.u.String()
}

func (u *URL) Hostname() string {
	return u.u.Hostname()
}

func (u *URL) Path() string {
	return u.u.Path
}

func (u *URL) GetQueryParam(key string) string {
	return u.u.Query().Get(key)
}

func (u *URL) RemoveLastPath() (*URL, error) {
	newURL := u.duplicate()

	path := newURL.Path()
	// 末尾のスラッシュを削除
	path = regexp.MustCompile(`/+$`).ReplaceAllString(path, "")

	// 最後のパスセグメントを削除
	re := regexp.MustCompile(`/[^/]+$`)
	newPath := re.ReplaceAllString(path, "") + "/"

	// 新しいURLを構築
	newURLString := newURL.u.Scheme + "://" + newURL.u.Host + newPath
	err := newURL.updatePath(newURLString)
	if err != nil {
		return nil, fmt.Errorf("newURL.updatePath: %w", err)
	}

	return newURL, nil
}

func (u *URL) ExtractAmazonASIN() (string, bool) {
	asinRegex := regexp.MustCompile(`/dp/([A-Z0-9]{10})`)
	matches := asinRegex.FindStringSubmatch(u.String())
	if len(matches) > 1 {
		return matches[1], true
	}
	return "", false
}

const baseAmazonDPURL = "https://www.amazon.co.jp/dp/"

func (u *URL) CreateAmazonShortURL(asin string) (*URL, error) {
	return NewURL(baseAmazonDPURL + asin)
}
