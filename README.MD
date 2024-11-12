# Wallet Service
<h3 align="left">    
    <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white">
    <img src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white">
    <img src="https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white">
    <img src="https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white">
    <img src="https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white">
    <img src="https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black">
</h3> 

## 1. Описание проекта

Этот учебный сервис предоставляет API для управления кошельками, включая операции пополнения и вывода средств, а также получение статуса кошелька. Сервис написан на языке Go и использует PostgreSQL для хранения данных.

## 2. Стек технологий

- Golang
- PostgreSQL & pgAdmin
- Docker
- Postman
- Git
- GitHub
- Linux

## 3. Установка

3.1. **Клонируйте репозиторий**:
   ```bash
   git clone https://github.com/jedyEvgeny/wallet-service.git
 ```

3.2. **Перейдите в каталог через терминал**:

   ```bash
   cd wallet-service
 ```

3.3. **Запустите проект в терминале**:

   ```bash
   go run cmd/app/main.go config.env
 ```
где config.env - путь к файлу конфигурации

## 4. Для запуска в Docker-контейнере

4.1. **Создайте контейнер**:
   ```bash
   docker-compose up --build
 ```

4.2. **Запустите контейнер**:
   ```bash
   docker-compose up
 ```

## 5. Описание работы
Приложение по REST принимает:
POST api/v1/wallet с телом запроса:
 ```bash
{
valletId: UUID,
operationType: DEPOSIT or WITHDRAW,
amount: 1000
}
 ```

Есть возможность возвращать балланс кошелька:
GET api/v1/wallets/{WALLET_UUID}


