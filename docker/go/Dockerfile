ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
ENV  GO111MODULE=on
#RUN go mod init gin
#RUN go get -u github.com/gin-gonic/gin 

#RUN go get -u github.com/oxequa/realize 
 #   && go get github.com/go-sql-driver/mysql
#CMD ["realize", "start"]

# RUN go get -u github.com/oxequa/realize \
#     && go get github.com/go-sql-driver/mysql

#CMD ["go", "get", "-u" ,"github.com/gin-gonic/gin" ]
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

#COPY . .
#RUN go build -o ./app ./src/main.go

#FROM alpine:latest

# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# RUN mkdir -p /api
# WORKDIR /api
# COPY --from=builder /api/app .
# COPY --from=builder /api/test.db .

EXPOSE 8080

#ENTRYPOINT ["./app"]
