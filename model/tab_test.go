package model_test

import (
	"net/url"
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model"
	"github.com/stretchr/testify/assert"
)

func TestTab_FormatTitleForEachSite(t *testing.T) {
	type fields struct {
		Title string
		URL   *url.URL
	}

	tests := []struct {
		name    string
		fields  fields
		want    *model.Tab
		wantErr error
	}{
		{
			name: "github.com であり、リポジトリルートである場合、パッケージ名のみを残すこと",
			fields: fields{
				Title: "golang/go: The Go programming language",
				URL:   parseURL(t, "https://github.com/golang/go"),
			},
			want: &model.Tab{
				Title: "golang/go",
				URL:   parseURL(t, "https://github.com/golang/go"),
			},
			wantErr: nil,
		},
		{
			name: "github.com であり、Issue である場合、Issue タイトルと番号のみを残すこと",
			fields: fields{
				Title: "cmd/cgo: fails with gcc 4.4.1 · Issue #1 · golang/go",
				URL:   parseURL(t, "https://github.com/golang/go/issues/1"),
			},
			want: &model.Tab{
				Title: "fails with gcc 4.4.1 #1",
				URL:   parseURL(t, "https://github.com/golang/go/issues/1"),
			},
			wantErr: nil,
		},
		{
			name: "github.com であり、PR である場合、PR タイトルと番号のみを残すこと",
			fields: fields{
				Title: "net/url: Fixed url parsing with invalid slashes. by odeke-em · Pull Request #9219 · golang/go",
				URL:   parseURL(t, "https://github.com/golang/go/pull/9219"),
			},
			want: &model.Tab{
				Title: "Fixed url parsing with invalid slashes. #9219",
				URL:   parseURL(t, "https://github.com/golang/go/pull/9219"),
			},
			wantErr: nil,
		},
		{
			name: "example.atlassian.net/wiki (Confluence) であり、ページである場合、ページタイトルのみを残すこと",
			fields: fields{
				Title: "設計ドキュメント - EXAMPLE - 開発チーム - Confluence",
				URL:   parseURL(t, "https://example.atlassian.net/wiki/spaces/EXAMPLE/pages/1"),
			},
			want: &model.Tab{
				Title: "設計ドキュメント - Confluence",
				URL:   parseURL(t, "https://example.atlassian.net/wiki/spaces/EXAMPLE/pages/1"),
			},
			wantErr: nil,
		},
		{
			name: "tabelog.com であり、店舗ページトップである場合、店名のみを残すこと",
			fields: fields{
				Title: "下北沢 焼とりダービーのご予約 - 下北沢/焼き鳥 | 食べログ",
				URL:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13283195/"),
			},
			want: &model.Tab{
				Title: "下北沢 焼とりダービーのご予約 | 食べログ",
				URL:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13283195/"),
			},
			wantErr: nil,
		},
		{
			name: "tabelog.com であり、店舗ページトップであり、店名に（）ありの場合、店名のみを残すこと",
			fields: fields{
				Title: "下北沢 肉バル Bon （ボン【旧店名】ワイン食堂 馬肉de Bon）のご予約 - 下北沢/バル | 食べログ",
				URL:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13188119/"),
			},
			want: &model.Tab{
				Title: "下北沢 肉バル Bon | 食べログ",
				URL:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13188119/"),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tab := &model.Tab{
				Title: tt.fields.Title,
				URL:   tt.fields.URL,
			}

			tab.SetHandlerForTest()
			tt.want.SetHandlerForTest()

			err := tab.FormatTitleForEachSite()
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, tab)
			} else {
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			}
		})
	}
}
