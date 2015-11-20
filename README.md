シンプルサーバー
------------------

## やること

- 擬似ログイン機能
- ログインしていない場合にリクエスト弾く

## ビルド方法

[gb](https://getgb.io/examples/getting-started/)が必要。

```
$ go get github.com/constabulary/gb/...
$ go get github.com/constabulary/gb/cmd/gb-vendor
```

プロジェクトディレクトリで

```
$ gb vendor restore
$ gb build all
$ ./bin/sampleserver
```

で起動します。