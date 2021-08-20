FROM alpine:edge AS builder
LABEL stage=builder
RUN apk update; apk upgrade; mkdir /service; apk add --update go=1.16.5-r0 gcc=10.3.1_git20210625-r0 g++=10.3.1_git20210625-r0

WORKDIR /service

COPY . /service

RUN go mod download; \
  CGO_ENABLED=1 GOOS=linux go build -a -ldflags="-s -w" -o exchange ./main.go

FROM alpine:edge

LABEL name=exchange-service
LABEL team=kmdev

RUN apk update && apk upgrade && mkdir /exchange
WORKDIR /exchange

COPY --from=builder /service/exchange /exchange/service
COPY --from=builder /service/config /exchange/config

EXPOSE 7005

CMD ["/exchange/service"]
