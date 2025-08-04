package formatter_test

import (
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model/formatter"
	"github.com/HitoroOhria/copy_tab_link/model/value"
	"github.com/stretchr/testify/assert"
)

func TestAmazonFormatter_Match(t *testing.T) {
	fmtr := &formatter.AmazonFormatter{}

	tests := []struct {
		name   string
		domain string
		want   bool
	}{
		{
			name:   "amazon.co.jpにマッチする",
			domain: "amazon.co.jp",
			want:   true,
		},
		{
			name:   "www.amazon.co.jpにマッチする",
			domain: "www.amazon.co.jp",
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

func TestAmazonFormatter_Format(t *testing.T) {
	fmtr := &formatter.AmazonFormatter{}

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
			name: "amazon.co.jp であり、商品ページである場合、商品名のみを残し、URL を短縮すること",
			args: args{
				path:  value.Path("/初めてのGo言語-―他言語プログラマーのためのイディオマティックGo実践ガイド-Jon-Bodner/dp/4814400047/ref=sr_1_5"),
				title: value.Title("初めてのGo言語 ―他言語プログラマーのためのイディオマティックGo実践ガイド | Jon Bodner, 武舎 広幸 |本 | 通販 | Amazon"),
				url:   parseURL(t, "https://www.amazon.co.jp/初めてのGo言語-―他言語プログラマーのためのイディオマティックGo実践ガイド-Jon-Bodner/dp/4814400047/ref=sr_1_5?__mk_ja_JP=カタカナ&crid=3JSO08O084J74&dib=eyJ2IjoiMSJ9.Kome8eBCZysR72wgx3sO3cPjPDiCUOiLWOpVm6XlXfk75W5ZmkNqqcJDPTjipCoQ__vO3wCh4dSJbAZ-vZK_PUdUAZR3coRs5zFIU6LDLfPpN10QzDJy55hwWPDW0uMzlPE6Zi3IbmBo4BWwJfdVDr9_24Jar313Wz1niHNQiRAAHtEpJfiZve2tj7CRUIPAedWScmCKmANbAuO8XJaxx_p-8lHPiJ1UXo8vro_eWNc4YfdOCu9EwSp6z2SSyJsehzESPjE-1diPIixV4FnBUnTmistmz8dITo0mY45ZSuc.ApSD2ueW7gLmuGUi94jbkwwqFsAqAHjQ6Lz-diZqr44&dib_tag=se&keywords=go+本&qid=1751089169&sprefix=go+本,aps,171&sr=8-5"),
			},
			want: want{
				title: value.Title("初めてのGo言語 ―他言語プログラマーのためのイディオマティックGo実践ガイド"),
				url:   parseURL(t, "https://www.amazon.co.jp/dp/4814400047"),
			},
			wantErr: false,
		},
		{
			name: "amazon.co.jp であり、Kindle の商品ページである場合、商品名のみを残し、URL を短縮すること",
			args: args{
				path:  value.Path("/エンタープライズアプリケーションアーキテクチャパターン-マーチン・ファウラー-ebook/dp/B01B5MX2O2/ref=tmm_kin_swatch_0"),
				title: value.Title("Amazon.co.jp: エンタープライズアプリケーションアーキテクチャパターン eBook : マーチン・ファウラー, 株式会社テクノロジックアート, 長瀬嘉秀, 長瀬嘉秀: Kindleストア"),
				url:   parseURL(t, "https://www.amazon.co.jp/エンタープライズアプリケーションアーキテクチャパターン-マーチン・ファウラー-ebook/dp/B01B5MX2O2/ref=tmm_kin_swatch_0?_encoding=UTF8&dib_tag=se&dib=eyJ2IjoiMSJ9.F869wwPCVh0DUDmFEGeb9wF23FDIlwAmutnU_HnpS25GhQ_rJUwqUVcHilAoiqCyXZfx0GH4kK0Nt4Pccsues4-Mpmejryw9OvZeKcI5ecWTJ0qDUMZNPKs-23vYi3k5qEKcraS0RkEom-5YgOX_biUmTVknS1TuXPzc-FpX9aRUnYZMN_-_SsQy8M-PBIFI_TUD8EMgWtBLszriQADb-M3o1nrfPlSq4RTEuNZjeCn_dFsFVmHIdbbVxSZ9fuUAKVM9IhLfzwZ0Wkij4Vh7YOi7q4uy88-JsEBjUcZS6FY.Www2pEz5FCZeQFlIhkw0Ip6WxglYcNRGBs15G4YIP-M&qid=1754283348&sr=8-1"),
			},
			want: want{
				title: value.Title("エンタープライズアプリケーションアーキテクチャパターン"),
				url:   parseURL(t, "https://www.amazon.co.jp/dp/B01B5MX2O2"),
			},
			wantErr: false,
		},
		{
			name: "amazon.co.jp であり、物理本の商品ページである場合、商品名のみを残し、URL を短縮すること",
			args: args{
				path:  value.Path("https://www.amazon.co.jp/エンタープライズ-アプリケーションアーキテクチャパターン-Object-Oriented-SELECTION/dp/4798105538/ref=tmm_other_meta_binding_swatch_0"),
				title: value.Title("Amazon.co.jp: エンタープライズ アプリケーションアーキテクチャパターン (Object Oriented SELECTION) : マーチン・ファウラー, 長瀬 嘉秀, 株式会社 テクノロジックアート: 本"),
				url:   parseURL(t, "https://www.amazon.co.jp/エンタープライズ-アプリケーションアーキテクチャパターン-Object-Oriented-SELECTION/dp/4798105538/ref=tmm_other_meta_binding_swatch_0?_encoding=UTF8&dib_tag=se&dib=eyJ2IjoiMSJ9.F869wwPCVh0DUDmFEGeb9wF23FDIlwAmutnU_HnpS25GhQ_rJUwqUVcHilAoiqCyXZfx0GH4kK0Nt4Pccsues4-Mpmejryw9OvZeKcI5ecWTJ0qDUMZNPKs-23vYi3k5qEKcraS0RkEom-5YgOX_biUmTVknS1TuXPzc-FpX9aRUnYZMN_-_SsQy8M-PBIFI_TUD8EMgWtBLszriQADb-M3o1nrfPlSq4RTEuNZjeCn_dFsFVmHIdbbVxSZ9fuUAKVM9IhLfzwZ0Wkij4Vh7YOi7q4uy88-JsEBjUcZS6FY.Www2pEz5FCZeQFlIhkw0Ip6WxglYcNRGBs15G4YIP-M&qid=1754283348&sr=8-1"),
			},
			want: want{
				title: value.Title("エンタープライズアプリケーションアーキテクチャパターン (Object Oriented SELECTION)"),
				url:   parseURL(t, "https://www.amazon.co.jp/dp/B01B5MX2O2"),
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

func TestAmazonFormatter_Name(t *testing.T) {
	fmtr := &formatter.AmazonFormatter{}
	assert.Equal(t, "Amazon", fmtr.Name())
}
