# mysqlとdockerの環境構築

docker-compose up -d --build # コンテナを作成しスタートする

# コンテナに入る

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
#db

mysql> SHOW GRANTS FOR user;
+------------------------------------------------------+
| Grants for user@%                                    |
+------------------------------------------------------+
| GRANT USAGE ON *.* TO 'user'@'%'                     |
| GRANT ALL PRIVILEGES ON `sample\_db`.* TO 'user'@'%' |
+------------------------------------------------------+
2 rows in set (0.00 sec)

goland_db/に関する権限がuserにはなかった。ルートでsqlに入るとあった
| Database           |
+--------------------+
| information_schema |
| golang_db          |
| mysql              |
| performance_schema |
| sample_db          |
| sys                |
+--------------------+
6 rows in set (0.08 sec)

+---------------+-------------+
| user          | host        |
+---------------+-------------+
| kouropa       | %           |
| root          | dockermysql |
| mysql.session | localhost   |
| mysql.sys     | localhost   |
| root          | localhost   |
| user          | localhost   |
+---------------+-------------+
％はローカルホスト以外からのどのホストからでもアクセス可能。
OpenできてもCRUDができないケースもある。注意するところはコンテナネームとユーザの持ってるDBに対する権限を変更する。できればDockerfileで権限を変更したい。



