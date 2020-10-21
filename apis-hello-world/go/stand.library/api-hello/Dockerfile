FROM golang:1.15 as builder
WORKDIR /go/src/api.hello
COPY . .
ENV GO111MODULE=on
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=386 go build --trimpath -a -tags netgo -ldflags '-w -extldflags "-static"' -o api.hello main.go
#RUN CGO_ENABLED=0 GOOS=linux go build --trimpath -o api.hello main.go
RUN CGO_ENABLED=0 go build --trimpath -ldflags="-s -w" -o api.hello main.go
RUN cp api.hello /go/bin/api.hello
RUN ls -lh

FROM alpine:latest AS final
RUN apk update
RUN apk add --no-cache tzdata
ENV TZ="America/Sao_Paulo"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /go/bin/api.hello /

CMD ["/api.hello"]