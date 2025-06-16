package model_test

import (
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model"
	"github.com/HitoroOhria/copy_tab_link/model/value"
	"github.com/stretchr/testify/assert"
)

func TestTab_FormatForEachSite(t *testing.T) {
	type fields struct {
		Title string
		URL   *value.URL
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
				Title: value.Title("golang/go"),
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
				Title: value.Title("cmd/cgo: fails with gcc 4.4.1 #1"),
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
				Title: value.Title("net/url: Fixed url parsing with invalid slashes. #9219"),
				URL:   parseURL(t, "https://github.com/golang/go/pull/9219"),
			},
			wantErr: nil,
		},
		{
			name: "qiita.com であり、記事である場合、タイトルのみを残すこと",
			fields: fields{
				Title: "ドメインの.comとか.jpってなに？ #FQDN - Qiita",
				URL:   parseURL(t, "https://qiita.com/miyuki_samitani/items/1667128245b14ae6e421"),
			},
			want: &model.Tab{
				Title: value.Title("ドメインの.comとか.jpってなに？ - Qiita"),
				URL:   parseURL(t, "https://qiita.com/miyuki_samitani/items/1667128245b14ae6e421"),
			},
			wantErr: nil,
		},
		{
			name: "stackoverflow.com であり、記事である場合、タイトルのみを残すこと",
			fields: fields{
				Title: "python - How to `assert_called_with` an object instance? - Stack Overflow",
				URL:   parseURL(t, "https://stackoverflow.com/questions/79656866/how-to-assert-called-with-an-object-instance"),
			},
			want: &model.Tab{
				Title: value.Title("How to `assert_called_with` an object instance? - Stack Overflow"),
				URL:   parseURL(t, "https://stackoverflow.com/questions/79656866/how-to-assert-called-with-an-object-instance"),
			},
			wantErr: nil,
		},
		{
			name: "zenn.com であり、記事である場合、Zenn を付与すること",
			fields: fields{
				Title: "【初心者歓迎】第２回 AI Agent Hackathon、開催決定！",
				URL:   parseURL(t, "https://zenn.dev/hackathons/google-cloud-japan-ai-hackathon-vol2"),
			},
			want: &model.Tab{
				Title: value.Title("【初心者歓迎】第２回 AI Agent Hackathon、開催決定！ - Zenn"),
				URL:   parseURL(t, "https://zenn.dev/hackathons/google-cloud-japan-ai-hackathon-vol2"),
			},
			wantErr: nil,
		},
		{
			name: "example.atlassian.net であり、Confluence であり、ページである場合、ページタイトルのみを残すこと",
			fields: fields{
				Title: "設計ドキュメント - EXAMPLE - 開発チーム - Confluence",
				URL:   parseURL(t, "https://example.atlassian.net/wiki/spaces/EXAMPLE/pages/1"),
			},
			want: &model.Tab{
				Title: value.Title("設計ドキュメント - Confluence"),
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
				Title: value.Title("下北沢 焼とりダービーのご予約 | 食べログ"),
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
				Title: value.Title("下北沢 肉バル Bon | 食べログ"),
				URL:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13188119/"),
			},
			wantErr: nil,
		},
		{
			name: "tabelog.com であり、コース一一覧ページである場合、タイトルに店名のみを残し、リンクを店舗トップページにすること",
			fields: fields{
				Title: "コース一覧 : 下北沢 焼とりダービー - 下北沢/焼き鳥 | 食べログ",
				URL:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13283195/party/"),
			},
			want: &model.Tab{
				Title: value.Title("下北沢 焼とりダービー | 食べログ"),
				URL:   parseURL(t, "https://tabelog.com/tokyo/A1318/A131802/13283195/"),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tab := &model.Tab{
				Title: value.Title(tt.fields.Title),
				URL:   tt.fields.URL,
			}

			tab.SetFormatterForTest()
			tt.want.SetFormatterForTest()

			err := tab.FormatForEachSite()
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, tab)
			} else {
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			}
		})
	}
}
