# go、gin,mysql環境をdockerで構築して簡単なCRUDを実装。awsでデプロイ
ー使い方ー

$ git clone this ripository

$ docker-compose up -d --build 

$ docker-comopse exec go bash

$ go run docker/go/main.go

実際に動いているところ
http://3.113.12.37/

ec2にssh接続

ssh -i go-keypair.pem ec2-user@3.113.12.37

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



=======

# docker-mysql-gin

# 今のところ
①mysqlとgoをdocker Composeでコンテナを立てて連携

②modファイルの設定などを行なってginをインストール。ginはmodがないと使えないエラーがでた。Dockerfileでinitするあたりが正しいのかはこれから調べる必要があるがとりあえず使えるようになった。

③mysqlのCRUD操作ををgoで書いた。DBとの連携の部分でDockerComposeの設定をいろいろ変えた。ポイントはどのユーザにどの権限があるのか、とコンテナ名を記述すること、その辺を変更するとアクセスできるようになった。

④ec2インスタンスを作成し、dockerなどをインストールしてインスタンス内で
$ docker-comose up -d --build
$ go run docker/go/main.go

⑤route53とacmの設定を行った


# これからの予定

ロードバランサを使う


つぶやきを投稿するだけでなく何か追加する

クーバネティスを組み込む

circlCIとかでGitHubにプッシュすると自動でデプロイするように組み込む


AWSの設定をより細かく設定し、実際の運用に近いように構成する。

セキュリティ対策





