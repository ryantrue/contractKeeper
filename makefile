# Makefile для сборки и запуска проекта taskManager

# Установка переменных
BINARY_NAME=contractKeeper
DB_NAME=contractKeeper.db

# Стандартная цель, которая выполняется при вызове make без аргументов
all: build

# Сборка проекта для текущей ОС
build:
	@echo "Сборка проекта для текущей ОС..."
	go build -o ${BINARY_NAME} cmd/main.go

# Сборка для macOS
build-mac:
	@echo "Сборка проекта для macOS..."
 GOOS=darwin GOARCH=amd64 go build -o ${BINARY_NAME}-mac cmd/main.go

# Сборка для Linux
build-linux:
	@echo "Сборка проекта для Linux..."
 GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}-linux cmd/main.go

# Сборка для Windows
build-windows:
	@echo "Сборка проекта для Windows..."
 GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}-windows.exe cmd/main.go

# Запуск проекта
run: build
	@echo "Запуск ${BINARY_NAME}..."
	./${BINARY_NAME}

# Чистка проекта (удаление исполняемых файлов и файла базы данных)
clean:
	@echo "Очистка..."
	go clean
	rm -f ${BINARY_NAME}*
	rm -f ${DB_NAME}

# Инициализация базы данных (опционально, если нужно создать структуру БД заранее)
init-db:
	@echo "Инициализация базы данных из schema.sql..."
	sqlite3 ${DB_NAME} < schema.sql

# Пометка целей, которые не являются файлами
.PHONY: all build build-mac build-linux build-windows run clean init-db