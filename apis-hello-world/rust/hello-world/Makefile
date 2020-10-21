# Makefile
.EXPORT_ALL_VARIABLES:	

docker:
	@echo "########## Compilando projeto ... "
	docker build --no-cache -f Dockerfile -t jeffotoni/rust-hello-world:latest .
	@echo "buiild docker completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"
run:
	@echo "########## Executando projeto ... "
	docker run --rm jeffotoni/rust-hello-world:latest 
