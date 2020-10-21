# tart by building the authlication.
# Build em fatorial com distroless
FROM golang:1.15 as builder
WORKDIR /go/src/fatorial
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o fatorial main.go
RUN cp fatorial /go/bin/fatorial

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/fatorial /
CMD ["/fatorial"]