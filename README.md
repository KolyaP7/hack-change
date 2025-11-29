# Hack Change

## Запуск бэкенда

### Требования

- Go 1.19 или выше
- PostgreSQL 12 или выше
- Git

### Установка зависимостей

1. Перейдите в директорию бэкенда:
```bash
cd backend
```

2. Инициализируйте Go модуль (если еще не инициализирован):
```bash
go mod init hack-change-backend
```

3. Установите зависимости:
```bash
go mod tidy
```

Команда `go mod tidy` автоматически установит все необходимые зависимости, включая:
- `golang.org/x/crypto/bcrypt` - для хеширования паролей

Если нужно установить зависимость вручную:
```bash
go get golang.org/x/crypto/bcrypt
```

### Настройка базы данных

1. Создайте базу данных PostgreSQL:
```bash
createdb hack_change
```

Или через psql:
```sql
CREATE DATABASE hack_change;
```

2. Примените миграции:
```bash
psql -d hack_change -f ../migrations.sql
```

### Настройка переменных окружения

Создайте файл `.env` в директории `backend` со следующим содержимым:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=hack_change
DB_SSLMODE=disable

PORT=8080
JWT_SECRET=your_secret_key_here
```

Замените значения на ваши реальные данные для подключения к базе данных.

### Запуск сервера

1. Из директории `backend` выполните:
```bash
go run cmd/api/main.go
```

Или соберите и запустите бинарный файл:
```bash
go build -o bin/api cmd/api/main.go
./bin/api
```

2. Сервер будет доступен по адресу `http://localhost:8080` (или по порту, указанному в переменной окружения `PORT`)

### Проверка работы

## Запуск фронтенда

### Требования

- Node.js 14 или выше
- npm или yarn

### Установка зависимостей

1. Перейдите в директорию фронтенда:
```bash
cd frontend
```

2. Установите зависимости:
```bash
npm install
```

Или если используете yarn:
```bash
yarn install
```

### Запуск приложения

1. Из директории `frontend` выполните:
```bash
npm start
```

Или если используете yarn:
```bash
yarn start
```

2. Приложение автоматически откроется в браузере по адресу `http://localhost:3000`

3. Приложение будет автоматически перезагружаться при изменении файлов (hot reload)

### Сборка для продакшена

Для создания оптимизированной сборки для продакшена:
```bash
npm run build
```

Или:
```bash
yarn build
```

Собранные файлы будут находиться в директории `frontend/build`

### Проверка работы

После запуска сервера вы можете проверить его работу, отправив запрос:
```bash
curl http://localhost:8080/health
```

### Структура проекта

```
backend/
├── cmd/
│   └── api/
│       └── main.go          # Точка входа приложения
├── internal/
│   ├── handlers/            # HTTP обработчики
│   │   └── auth/           # Обработчики аутентификации
│   ├── middleware/          # Middleware (CORS, аутентификация и т.д.)
│   └── repository/          # Слой работы с базой данных
```

### Разработка

Для разработки с автоматической перезагрузкой при изменениях можно использовать [air](https://github.com/cosmtrek/air):

```bash
go install github.com/cosmtrek/air@latest
air
```

