シンプルサーバー
------------------

## やること

- 擬似ログイン機能(POST /login)
- ログインしていない場合にリクエスト弾く(GET /dashboard/json)

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
