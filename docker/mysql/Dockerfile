FROM mysql:5.7
#指定の場所にログを記録するファイルを作る
RUN touch /var/log/mysql/mysqld.log 

EXPOSE 3306
#COPY initdb.d/ /docker-entrypoint-initdb.d/

#######awsでデプロイする時は Composeの方でマウントができないので、これを使っておく。
COPY ./docker/mysql/conf.d/  /etc/mysql/conf.d
COPY ./docker/mysql/mysql_data/ /var/lib/mysql
ENV   MYSQL_DATABASE = golang_db  
ENV   MYSQL_USER = kouropa
ENV   MYSQL_PASSWORD = password
ENV   MYSQL_ROOT_PASSWORD = rootpassword
ENV   MYSQL_ROOT_HOST = dockerMySQL
# MySQL サーバーとも呼ばれ、MySQL インストールでのほとんどの作業を実行するメインプログラム
CMD ["mysqld"]
