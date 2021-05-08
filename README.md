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

# 注意ポイント
db, err := gorm.Open("mysql", "kouropa:password@tcp(dockerMySQL:3306)/golang_db?parseTime=true") 
//rootだとアクセスできなかった。あとホスト名はcomposeのコンテナネーム.parsetimeは時間の取得のために追加

router.LoadHTMLGlob("docker/go/views/*.html") //ここのパス書き換えたら動いた

dockerのイメージがalpineの時は/bin/shだけどgolangの時は/bash/ 



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

# 今のところ
①mysqlとgoをdocker Composeでコンテナを立てて連携

②modファイルの設定などを行なってginをインストール。ginはmodがないと使えないエラーがでた。Dockerfileでinitするあたりが正しいのかはこれから調べる必要があるがとりあえず使えるようになった。

③mysqlのCRUD操作ををgoで書いた。DBとの連携の部分でDockerComposeの設定をいろいろ変えた。ポイントはどのユーザにどの権限があるのか、とコンテナ名を記述すること、その辺を変更するとアクセスできるようになった。


# これからの予定

TOdoアプリとしてローカルで動くようにする。見栄えはあまり気にしない。CRUD処理が行えて、フロントで値を表示する練習。

AWSにデプロイする。dockerとAWSのサービスがあるので、それを使ったデプロイを行う

nginxを使ってリバースプロキシを設定する。

APIを利用したアプリにする。例えば、スーモからデータを撮ってきて表示するなど

クーバネティスを組み込む

circlCIを組み込む

AWSの設定をより細かく設定し、実際の運用に近いように構成する。

セキュリティ対策




