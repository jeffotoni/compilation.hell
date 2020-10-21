FROM golang:1.15 as builder
WORKDIR /go/src/api.hello2
COPY . .
ENV GO111MODULE=on
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=386 go build --trimpath -a -tags netgo -ldflags '-w -extldflags "-static"' -o api.hello2 main.go
#RUN CGO_ENABLED=0 GOOS=linux go build --trimpath -o api.hello2 main.go
RUN CGO_ENABLED=0 go build --trimpath -ldflags="-s -w" -o api.hello2 main.go
RUN cp api.hello2 /go/bin/api.hello2
RUN ls -lh

FROM alpine:latest AS final
RUN apk update
RUN apk add --no-cache tzdata
ENV TZ="America/Sao_Paulo"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /go/bin/api.hello2 /

CMD ["/api.hello2"]