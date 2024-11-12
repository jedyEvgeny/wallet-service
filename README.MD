# Wallet Service
## 1. Описание проекта

Этот учебный сервис предоставляет API для управления кошельками, включая операции пополнения и вывода средств, а также получение статуса кошелька. Сервис написан на языке Go и использует PostgreSQL для хранения данных.

## 2. Стек технологий

- Golang
- PostgreSQL & pgAdmin
- Docker
- Git
- GitHub
- Linux

## 3. Установка

3.1. **Клонируйте репозиторий**:
   ```bash
   git clone https://github.com/jedyEvgeny/wallet-service.git

3.2. **Перейдите в каталог через терминал**:

   ```bash
   cd wallet-service

3.3. **Запустите проект в терминале**:

   ```bash
   go run cmd/app/main.go config.env
где config.env - путь к файлу конфигурации

## 4. Для запуска в Docker-контейнере

4.1. **Создайте контейнер**:
   ```bash
   docker-compose up --build

4.2. **Запустите контейнер**:
   ```bash
   docker-compose up

## 5. Описание работы
Приложение по REST принимает два запроса.
POST api/v1/wallet
{
valletId: UUID,
operationType: DEPOSIT or WITHDRAW,
amount: 1000
}

Есть возможность возвращать балланс кошелька:
GET api/v1/wallets/{WALLET_UUID}


