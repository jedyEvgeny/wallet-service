# Wallet Service

Этот учебный сервис предоставляет API для управления кошельками, включая операции пополнения и вывода средств, а также получение статуса кошелька. Сервис написан на языке Go и использует PostgreSQL для хранения данных.

## Стек технологий

- Golang
- PostgreSQL & pgAdmin
- Docker
- Git
- GitHub
- Linux

## Установка

1. **Клонируйте репозиторий**:

   ```bash
   git clone https://github.com/jedyEvgeny/wallet-service.git
   cd wallet-service
Запуск: go run cmd/app/main.go config.env, 
где config.env - путь к файлу конфигурации

## Для запуска в Docker-контейнере
Для запуска в Docker-контейнере в корне проекта наберите:
docker-compose up --build

Используйте Docker и Docker Compose для запуска базы данных и сервиса:
docker-compose up
