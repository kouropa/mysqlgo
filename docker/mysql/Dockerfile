FROM mysql:5.7
#指定の場所にログを記録するファイルを作る
RUN touch /var/log/mysql/mysqld.log 

EXPOSE 3306
COPY initdb.d/ /docker-entrypoint-initdb.d/


# MySQL サーバーとも呼ばれ、MySQL インストールでのほとんどの作業を実行するメインプログラム
CMD ["mysqld"]
