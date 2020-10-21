# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off
ENV=prd

build:
	@echo "########## Compilando nossa API ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags '-w -s' -o api.hello main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"

docker:
	@echo "########## Compilando nossa API ... "
	docker build --no-cache -f Dockerfile -t jeffotoni/api.hello:latest .
	@echo "buid docker completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"