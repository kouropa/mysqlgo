
# ヴァージョンを挙げろと言われたのであげる。1.11-1.12 
#ARG GO_VERSION=1.12.4
#alpineの時はbin/sh 
#eFROM golang:${GO_VERSION}-alpine AS builder
#FROM golang:1.14.4 as builder
#FROM golang:1.12.0-alpine3.9
FROM golang:latest as builder

#アルパインようなので、通常バーの時は外す
#RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/* 
# for time 
#RUN apk add mysql-client

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app



ENV  GO111MODULE=on

# # waitライブラリの追加
# ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.6.0/wait /go/src/app/wait 
# # ライブラリの権限変更
#RUN chmod 750 /go/src/app/start.sh


# COPY --from=build /work/target/dtest-0.0.1-SNAPSHOT.jar /go/src/app/wait

# CMD /wait 
#CMD ["go","run","go/src/app/docker/go/main.go"]

#　 Compose　upしただけでアクセスできないといけないからコンテナに入ってコマンドをうつのをここでやっておく必要がある。
# RUN go run /go/src/app/docker/go/main.go docker execでは行けるがこれだとgo　runできない・
#RUN go mod init gin
RUN go get -u github.com/gin-gonic/gin 

##$#RUN go get -u github.com/oxequa/realize 
 #   && go get github.com/go-sql-driver/mysql
#CMD ["realize", "start"]

# RUN go get -u github.com/oxequa/realize \
RUN go get  -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm

#CMD ["go", "get", "-u" ,"github.com/gin-gonic/gin" ]
# COPY go.mod .
# COPY go.sum .
# RUN go mod download


EXPOSE 80

#EXPOSE 8080
#ENV GIN_MODE=release
#ENTRYPOINT ["./app"]
#RUN go run /go/src/app/docker/go/main.go

##本当はこれでいきたいが、コンテナに入ってから売った方が安全？
#CMD ["go", "run", "/go/src/app/docker/go/main.go"]
