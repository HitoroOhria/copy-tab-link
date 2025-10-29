# Copy Tab Link

アクティブなブラウザのタイトルとリンクを取得し、Markdown形式でクリップボードにコピーする。
動作環境の対象は macOS である。

# Usage

## Build

アプリケーションをビルドする。

```shell
$ task build
```

## Run

環境変数をエクスポートして、アプリケーションを実行する。


```shell
$ export BROWSER_NAME=Chrome
$ ./copy-tab-link
```

もしくは、引数を指定してアプリケーションを実行する。

```shell
$ ./copy-tab-link -browser-name=Chrome
```

# Development

## Run

次のいずれかのコマンドでプログラムを実行することができる。

```shell
$ go run ./...

$ task run
```

次のいずれかの方法で対象ブラウザを指定することができる。

```shell
$ go run ./... -browser-name=Chrome

$ task run BROWSER_NAME=Chrome

$ sed -i '' 's/^BROWSER_NAME=.*/BROWSER_NAME=Chrome/' .env && task run
```

## Check title of url

URL から取得されるタイトルを知りたい場合、次のコマンドで確認できる。

```shell
$ task dev:title URL='https://example.com'
```
