package main_test

import (
    "net/url"
    "testing"

    "github.com/HitoroOhria/copy_tab_link"
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
        want    *main.Tab
        wantErr error
    }{
        {
            name: "github.com であり、リポジトリルートである場合、パッケージ名のみを残すこと",
            fields: fields{
                Title: "golang/go: The Go programming language",
                URL:   parseURL(t, "https://github.com/golang/go"),
            },
            want: &main.Tab{
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
            want: &main.Tab{
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
            want: &main.Tab{
                Title: "Fixed url parsing with invalid slashes. #9219",
                URL:   parseURL(t, "https://github.com/golang/go/pull/9219"),
            },
            wantErr: nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tab := &main.Tab{
                Title: tt.fields.Title,
                URL:   tt.fields.URL,
            }

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
