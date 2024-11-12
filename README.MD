# Wallet Service

Этот учебный сервис предоставляет API для управления кошельками, включая операции пополнения и вывода средств, а также получение статуса кошелька. Сервис написан на языке Go и использует PostgreSQL для хранения данных.

## Стек технологий

- Golang
- PostgreSQL
- Docker

## Установка

1. **Клонируйте репозиторий**:

   ```bash
   git clone https://github.com/jedyEvgeny/wallet-service.git
   cd wallet-service
Создайте файл конфигурации:

Создайте файл config.env в корневой директории проекта, содержащий необходимые переменные окружения, например:

env

Copy
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=wallet_service
SERVER_HOST=0.0.0.0
SERVER_PORT=8080
Запустите Docker:

Используйте Docker и Docker Compose для запуска базы данных и сервиса:

bash

Copy
docker-compose up
Это создаст и запустит контейнеры PostgreSQL и вашего сервиса.

API
Создание/Обновление кошелька
POST /api/v1/wallet

Запрос
json

Copy
{
  "valletId": "UUID",
  "operationType": "DEPOSIT", // или "WITHDRAW"
  "amount": 1000
}
Ответ
201 Created: При успешном выполнении операции.
400 Bad Request: В случае ошибок валидации.
500 Internal Server Error: В случае ошибок сервиса.
Получение статуса кошелька
GET /api/v1/wallets/{WALLET_UUID}

Запрос
WALLET_UUID: идентификатор кошелька.
Ответ
200 OK: Возвращает текущий баланс кошелька.
400 Bad Request: В случае некорректного UUID.
404 Not Found: Если кошелек не найден.
500 Internal Server Error: В случае ошибок сервиса.
Запуск тестов
Вы можете запускать тесты с помощью команды:

bash

Copy
go test -v ./...
Для очистки кэша тестов перед запуском можно использовать:

bash

Copy
go clean -testcache
Лицензия
Данный проект лицензирован под MIT License. Вы можете свободно использовать и изменять его в соответствии с условиями лицензии.