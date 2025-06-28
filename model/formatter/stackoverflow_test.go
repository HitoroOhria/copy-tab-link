package formatter_test

import (
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model/formatter"
	"github.com/HitoroOhria/copy_tab_link/model/value"
	"github.com/stretchr/testify/assert"
)

func TestStackOverflowFormatter_Match(t *testing.T) {
	fmtr := &formatter.StackOverflowFormatter{}

	tests := []struct {
		name   string
		domain string
		want   bool
	}{
		{
			name:   "stackoverflow.comにマッチする",
			domain: "stackoverflow.com",
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

func TestStackOverflowFormatter_Format(t *testing.T) {
	fmtr := &formatter.StackOverflowFormatter{}

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
			name: "stackoverflow.com であり、記事である場合、タイトルのみを残すこと",
			args: args{
				path:  value.Path("/questions/79656866/how-to-assert-called-with-an-object-instance"),
				title: value.Title("python - How to `assert_called_with` an object instance? - Stack Overflow"),
				url:   parseURL(t, "https://stackoverflow.com/questions/79656866/how-to-assert-called-with-an-object-instance"),
			},
			want: want{
				title: value.Title("How to `assert_called_with` an object instance? - Stack Overflow"),
				url:   parseURL(t, "https://stackoverflow.com/questions/79656866/how-to-assert-called-with-an-object-instance"),
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

func TestStackOverflowFormatter_Name(t *testing.T) {
	fmtr := &formatter.StackOverflowFormatter{}
	assert.Equal(t, "Stack Overflow", fmtr.Name())
}
