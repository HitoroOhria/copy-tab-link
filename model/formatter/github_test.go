package formatter_test

import (
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model/formatter"
	"github.com/HitoroOhria/copy_tab_link/model/value"
	"github.com/stretchr/testify/assert"
)

func parseURL(t *testing.T, urlStr string) *value.URL {
	t.Helper()
	url, err := value.NewURL(urlStr)
	if err != nil {
		t.Fatalf("failed to parse URL: %v", err)
	}
	return url
}

func TestGitHubFormatter_Match(t *testing.T) {
	fmtr := &formatter.GitHubFormatter{}

	tests := []struct {
		name   string
		domain string
		want   bool
	}{
		{
			name:   "github.comにマッチする",
			domain: "github.com",
			want:   true,
		},
		{
			name:   "他のドメインにはマッチしない",
			domain: "example.com",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			domain := value.Domain(tt.domain)
			got := fmtr.Match(domain)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGitHubFormatter_Format(t *testing.T) {
	fmtr := &formatter.GitHubFormatter{}

	type args struct {
		path  value.Path
		title value.Title
		url   *value.URL
	}

	type want struct {
		title value.Title
		url   *value.URL
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "github.com であり、リポジトリルートである場合、パッケージ名のみを残すこと",
			args: args{
				path:  value.Path("/golang/go"),
				title: value.Title("golang/go: The Go programming language"),
				url:   parseURL(t, "https://github.com/golang/go"),
			},
			want: want{
				title: value.Title("golang/go"),
				url:   parseURL(t, "https://github.com/golang/go"),
			},
			wantErr: false,
		},
		{
			name: "github.com であり、Issue である場合、Issue タイトルと番号のみを残すこと",
			args: args{
				path:  value.Path("/golang/go/issues/1"),
				title: value.Title("cmd/cgo: fails with gcc 4.4.1 · Issue #1 · golang/go"),
				url:   parseURL(t, "https://github.com/golang/go/issues/1"),
			},
			want: want{
				title: value.Title("cmd/cgo: fails with gcc 4.4.1 #1"),
				url:   parseURL(t, "https://github.com/golang/go/issues/1"),
			},
			wantErr: false,
		},
		{
			name: "github.com であり、PR である場合、PR タイトルと番号のみを残すこと",
			args: args{
				path:  value.Path("/golang/go/pull/9219"),
				title: value.Title("net/url: Fixed url parsing with invalid slashes. by odeke-em · Pull Request #9219 · golang/go"),
				url:   parseURL(t, "https://github.com/golang/go/pull/9219"),
			},
			want: want{
				title: value.Title("net/url: Fixed url parsing with invalid slashes. #9219"),
				url:   parseURL(t, "https://github.com/golang/go/pull/9219"),
			},
			wantErr: false,
		},
		{
			name: "github.com であり、プロジェクトのサイド Issue ビューである場合、Issue タイトルと番号のみを残すこと",
			args: args{
				path:  value.Path("/orgs/golang/projects/5/views/1"),
				title: value.Title("net/url: Fixed url parsing with invalid slashes. · golang/go"),
				url:   parseURL(t, "https://github.com/orgs/golang/projects/5/views/1?pane=issue&itemId=116474225&issue=golang%7Cgo%7C1"),
			},
			want: want{
				title: value.Title("net/url: Fixed url parsing with invalid slashes. #1"),
				url:   parseURL(t, "https://github.com/orgs/golang/projects/5/views/1?pane=issue&itemId=116474225&issue=golang%7Cgo%7C1"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTitle, gotURL, err := fmtr.Format(tt.args.path, tt.args.title, tt.args.url)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.title, gotTitle)
			assert.Equal(t, tt.want.url.String(), gotURL.String())
		})
	}
}

func TestGitHubFormatter_Name(t *testing.T) {
	fmtr := &formatter.GitHubFormatter{}
	assert.Equal(t, "GitHub", fmtr.Name())
}
