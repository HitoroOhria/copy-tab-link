package formatter_test

import (
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model/formatter"
	"github.com/HitoroOhria/copy_tab_link/model/value"
	"github.com/stretchr/testify/assert"
)

func TestZennFormatter_Match(t *testing.T) {
	fmtr := &formatter.ZennFormatter{}

	tests := []struct {
		name   string
		domain string
		want   bool
	}{
		{
			name:   "zenn.devにマッチする",
			domain: "zenn.dev",
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

func TestZennFormatter_Format(t *testing.T) {
	fmtr := &formatter.ZennFormatter{}

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
			name: "zenn.com であり、記事である場合、Zenn を付与すること",
			args: args{
				path:  value.Path("/hackathons/google-cloud-japan-ai-hackathon-vol2"),
				title: value.Title("【初心者歓迎】第２回 AI Agent Hackathon、開催決定！"),
				url:   parseURL(t, "https://zenn.dev/hackathons/google-cloud-japan-ai-hackathon-vol2"),
			},
			want: want{
				title: value.Title("【初心者歓迎】第２回 AI Agent Hackathon、開催決定！ - Zenn"),
				url:   parseURL(t, "https://zenn.dev/hackathons/google-cloud-japan-ai-hackathon-vol2"),
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

func TestZennFormatter_Name(t *testing.T) {
	fmtr := &formatter.ZennFormatter{}
	assert.Equal(t, "Zenn", fmtr.Name())
}
