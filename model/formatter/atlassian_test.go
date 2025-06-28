package formatter_test

import (
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model/formatter"
	"github.com/HitoroOhria/copy_tab_link/model/value"
	"github.com/stretchr/testify/assert"
)

func TestAtlassianFormatter_Match(t *testing.T) {
	fmtr := &formatter.AtlassianFormatter{}

	tests := []struct {
		name   string
		domain string
		want   bool
	}{
		{
			name:   "atlassian.netドメインにマッチする",
			domain: "example.atlassian.net",
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

func TestAtlassianFormatter_Format(t *testing.T) {
	fmtr := &formatter.AtlassianFormatter{}

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
			name: "example.atlassian.net であり、Confluence であり、ページである場合、ページタイトルのみを残すこと",
			args: args{
				path:  value.Path("/wiki/spaces/EXAMPLE/pages/1"),
				title: value.Title("設計ドキュメント - EXAMPLE - 開発チーム - Confluence"),
				url:   parseURL(t, "https://example.atlassian.net/wiki/spaces/EXAMPLE/pages/1"),
			},
			want: want{
				title: value.Title("設計ドキュメント - Confluence"),
				url:   parseURL(t, "https://example.atlassian.net/wiki/spaces/EXAMPLE/pages/1"),
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

func TestAtlassianFormatter_Name(t *testing.T) {
	fmtr := &formatter.AtlassianFormatter{}
	assert.Equal(t, "Atlassian", fmtr.Name())
}
