## インストール

`$ go get -u golang.org/x/net/http2`

## オレオレ証明書の作成

`$ openssl genrsa 2048 > server.key`
`$ openssl req -new -key server.key > server.csr`
`$ openssl x509 -days 3650 -req -signkey server.key < server.csr > server.crt`

参考
- [Go1.5でHTTP2サーバーとクライアント - Qiita](https://qiita.com/koki_cheese/items/35c3fad6f1eb8458eafd)

## コピーシェルスクリプト実行

実行権限を与える
`$ chmod u+x test.sh`

実行
`$ ./copy.sh`

参考
- [ファイルのコピーを連番で繰り返し行うときのメモ【シェルスクリプト】 – Nishiaki's Log](http://nishiaki.probo.jp/2009/01/blog-post_25.html)
