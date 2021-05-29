# goとMySQL環境をdockerで構築して簡単なCRUDを実装しawsで公開した
ー使い方ー

$ git clone this ripository

$ docker-compose up -d --build 

$ docker-comopse exec go bash

$ go run docker/go/main.go

実際に動いているところ
http://3.113.12.37/

http://kouropa.com


![image](https://user-images.githubusercontent.com/36202603/120076007-9c26e700-c0de-11eb-9036-23dd5f693968.png)



ec2にssh接続

$ ssh -i go-keypair.pem ec2-user@3.113.12.37

# AWS RDS

EC2のセキュリティグループと、RDSのセキュリティグループのインバウンドルールを変更。mysqlクライアントをインストール、バージョンを合わせる。

$ mysql -u kouropa -h RDSのエンドポイント　-p

# circleCI

初回は普通にdocker-compose up -d --buildできたはずだが、できなくなった


下記のコマンドを実行するとテキストファイルが開かれる
$sudo visudo
下記の指定を
Defaults    secure_path = /sbin:/bin:/usr/sbin:/usr/bin
下記に変更する

Defaults        env_keep +="PATH"


たまにコンテナのイメージとかを削除しないといけないらしい
Service 'db' failed to build: Error processing tar file(exit status 1): write /var/lib/mysql/ib_logfile1: no space left on device

$ docker system prune


# RDS
ただdockerMySQLのところをRDSのエンドポイント名で置き換えるだけ
最初はdatabaseがないので中にはいって
$ create database golang_db;
日本語を使えるようにする

use golang_db;
show variables like 'char%'
で確認。latin1にdatabaseがなってたら変更する



# AWS ACM
Aレコードの値＝IPのレコードを消して、新しくAタイプのレコードを追加。値はエイリアスからアプリケーションロードバランサーー東京リージョンーロードバランサーの選択をする。とHTTPSでアクセスできた。

cnameタイプで値をKouropa、名前をルート５３で設定したDNS名にするとhttpsでアクセス可能
httpでアクセスしてもリダイレクトされるように設定済み。リスナーのルール設定



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
ポイントはどのユーザにどの権限があるのか、とコンテナ名を記述すること、その辺を変更するとアクセスできるようになった。

ginはmodがないと使えないエラーがでた。Dockerfileでinitするあたりが正しいのかはこれから調べる必要があるがとりあえず使えるようになった。

=======

# docker-mysql-gin

# 行った事

mysqlとgoをdocker Composeでコンテナを立てて連携


つぶやきを投稿してDBに保存、更新、消去、読み出しをできるようにした

ec2インスタンスを作成し、dockerなどをインストールしてインスタンス内で

$ docker-comose up -d --build

$ go run docker/go/main.go

route53とacmの設定を行った


# これからの予定

ロードバランサを使う


つぶやきを投稿するだけでなく何か追加する

クーバネティスを組み込む

circlCIとかでGitHubにプッシュすると自動でデプロイするように組み込む


AWSの設定をより細かく設定し、実際の運用に近いように構成する。

セキュリティ対策





