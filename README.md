# Copy Tab Link

アクティブなブラウザのタイトルとリンクを取得し、Markdown形式でクリップボードにコピーする。
動作環境の対象は macOS である。

# Development

## Run

次のコマンドでプログラムを実行することができる。

```shell
$ go run ./...

$ task run
```

対象ブラウザを次のいずれかで指定することができる。

```shell
$ sed -i '' 's/^BROWSER_NAME=.*/BROWSER_NAME="Chrome"/' .env && go run ./...

$ go run ./... -browser-name=Chrome

$ task run BROWSER_NAME=Chrome
```

## Check title of url

URL から取得されるタイトルを知りたい場合、次のコマンドで確認できる。

```shell
$ task dev:title URL='https://example.com'
```
