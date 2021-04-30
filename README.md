# mysqlとdockerの環境構築

docker-compose up -d # コンテナを作成しスタートする

docker exec -it mysqlgo_db_1　bash　

もしくは
docker-compose  exec  db  bash


mysql -u user -p # ログインする password で入れる。


# コンテナ外からログインする
mysql -u user -h  127.0.0.1 -ppassword

#データベースの中を見る

show databases;

use sample_db;

select from users;
# docker-mysql-gin

goのコンテナに入る

docker-compose exec go /bin/sh

＃RUN コマンドではmodファイルに書き込みがされないのでコンテナに入ってginを入れる。
go get -u github.com/gin-gonic/gin 
