package formatter_test

import (
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model/formatter"
	"github.com/HitoroOhria/copy_tab_link/model/value"
	"github.com/stretchr/testify/assert"
)

func TestQiitaFormatter_Match(t *testing.T) {
	fmtr := &formatter.QiitaFormatter{}

	tests := []struct {
		name   string
		domain string
		want   bool
	}{
		{
			name:   "qiita.comにマッチする",
			domain: "qiita.com",
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

func TestQiitaFormatter_Format(t *testing.T) {
	fmtr := &formatter.QiitaFormatter{}

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
			name: "qiita.com であり、記事である場合、タイトルのみを残すこと",
			args: args{
				path:  value.Path("/miyuki_samitani/items/1667128245b14ae6e421"),
				title: value.Title("ドメインの.comとか.jpってなに？ #FQDN - Qiita"),
				url:   parseURL(t, "https://qiita.com/miyuki_samitani/items/1667128245b14ae6e421"),
			},
			want: want{
				title: value.Title("ドメインの.comとか.jpってなに？ - Qiita"),
				url:   parseURL(t, "https://qiita.com/miyuki_samitani/items/1667128245b14ae6e421"),
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

func TestQiitaFormatter_Name(t *testing.T) {
	fmtr := &formatter.QiitaFormatter{}
	assert.Equal(t, "Qiita", fmtr.Name())
}
