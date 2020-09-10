# 概要
Docker の仕組みを実装を通して学ぶために、
最小限の機能を持つ Linux コンテナを Golang で作成します。


# 実行方法
コンテナの作成は Linux カーネルの機能に依存するため、 Linux 以外の環境をご利用の場合は、
Docker コンテナや VM の中で `main.go` を実行する必要があります。

```bash
# 適当な Docker イメージに含まれる root filesystem を `rootfs` ディレクトリにエクスポート
docker export $(docker create busybox) | tar -C rootfs -xvf -

# `main.go` を実行するための Docker コンテナを起動
docker-compose run app

# Docker コンテナ内で `main.go` を実行し、自前のコンテナを作成
go run main.go run <コマンド> <引数>
```
