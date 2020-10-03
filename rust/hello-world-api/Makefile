# Makefile
.EXPORT_ALL_VARIABLES:	

docker:
	@echo "########## Compilando projeto ... "
	docker build --no-cache -f Dockerfile -t jeffotoni/rust-hello-world-api:latest .
	@echo "buiild docker completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"
run:
	@echo "########## Executando projeto ... "
	docker run --init --rm -p 8080:8080 jeffotoni/rust-hello-world-api:latest
