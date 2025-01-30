# Nome do binário de saída
BINARY=build/main

# Comando para compilar o projeto
build:
	mkdir -p build
	go build -o $(BINARY) ./main.go

# Comando para rodar o programa após build
run: build
	./$(BINARY)

# Limpar arquivos compilados
clean:
	rm -rf build/
