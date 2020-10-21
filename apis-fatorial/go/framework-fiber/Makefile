# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

build:
	@echo "########## Compilando nossa API ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o fatorial main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"

docker:
	@echo "########## Docker Build ... "
	#echo "-------------------------------------- Clean <none> images ---------------------------------------	
	docker build -f Dockerfile -t jeffotoni/fatorial .
	@echo "########## done ... "