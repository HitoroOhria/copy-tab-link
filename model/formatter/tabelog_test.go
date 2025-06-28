package formatter_test

import (
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model/formatter"
	"github.com/HitoroOhria/copy_tab_link/model/value"
	"github.com/stretchr/testify/assert"
)

func TestTabelogFormatter_Match(t *testing.T) {
	fmtr := &formatter.TabelogFormatter{}

	tests := []struct {
		name   string
		domain string
		want   bool
	}{
		{
			name:   "tabelog.comにマッチする",
			domain: "tabelog.com",
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

func TestTabelogFormatter_Format(t *testing.T) {
	fmtr := &formatter.TabelogFormatter{}

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
			name: "tabelog.com であり、店舗ページトップである場合、店名のみを残すこと",
			args: args{
				path:  value.Path("/tokyo/A1318/A131802/13283195/"),
				title: value.Title("下北沢 焼とりダービーのご予約 - 下北沢/焼き鳥 | 食べログ"),
				url:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13283195/"),
			},
			want: want{
				title: value.Title("下北沢 焼とりダービーのご予約 | 食べログ"),
				url:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13283195/"),
			},
			wantErr: false,
		},
		{
			name: "tabelog.com であり、店舗ページトップであり、店名に（）ありの場合、店名のみを残すこと",
			args: args{
				path:  value.Path("/tokyo/A1318/A131802/13188119/"),
				title: value.Title("下北沢 肉バル Bon （ボン【旧店名】ワイン食堂 馬肉de Bon）のご予約 - 下北沢/バル | 食べログ"),
				url:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13188119/"),
			},
			want: want{
				title: value.Title("下北沢 肉バル Bon | 食べログ"),
				url:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13188119/"),
			},
			wantErr: false,
		},
		{
			name: "tabelog.com であり、コース一覧ページである場合、タイトルに店名のみを残し、リンクを店舗トップページにすること",
			args: args{
				path:  value.Path("/tokyo/A1318/A131802/13283195/party/"),
				title: value.Title("コース一覧 : 下北沢 焼とりダービー - 下北沢/焼き鳥 | 食べログ"),
				url:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13283195/party/"),
			},
			want: want{
				title: value.Title("下北沢 焼とりダービー | 食べログ"),
				url:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13283195/"),
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

func TestTabelogFormatter_Name(t *testing.T) {
	fmtr := &formatter.TabelogFormatter{}
	assert.Equal(t, "Tabelog", fmtr.Name())
}
