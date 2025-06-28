package model_test

import (
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model"
	"github.com/HitoroOhria/copy_tab_link/model/value"
	"github.com/stretchr/testify/assert"
)

func TestTab_FormatForEachSite_Integration(t *testing.T) {
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
			name: "適切なフォーマッターが選択され、GitHub の整形が実行される",
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
			name: "適切なフォーマッターが選択され、Qiita の整形が実行される",
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
			name: "該当するフォーマッターがない場合、元のタイトルとURLが保持される",
			fields: fields{
				Title: "Example Domain",
				URL:   parseURL(t, "https://example.com"),
			},
			want: &model.Tab{
				Title: value.Title("Example Domain"),
				URL:   parseURL(t, "https://example.com"),
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
