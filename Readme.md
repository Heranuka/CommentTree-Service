# CommentTree

CommentTree — сервис для работы с древовидными комментариями.  
Проект поддерживает создание, получение, удаление комментариев и построение дерева комментариев.

## Возможности

- Создание комментария.
- Получение корневых комментариев с пагинацией и поиском.
- Получение дочерних комментариев.
- Получение полного дерева комментариев.
- Удаление комментария.
- Swagger-документация API.

## Технологический стек

- Язык программирования: Go 1.25.
- Web-фреймворк: Ginext.
- База данных: PostgreSQL.
- Миграции: Migrate.
- Логирование: Zerolog.
- Документация API: Swagger.
- Кэш: Redis.
- Контейнеризация: Docker, Docker Compose.

## Структура проекта

```text
.
├── cmd
├── configs
├── internal
│   ├── app
│   ├── config
│   ├── domain
│   ├── infra
│   ├── logger
│   ├── repository
│   ├── service
│   └── transport
├── migrations
└── web
    ├── dist
    ├── node_modules
    └── src
```

## Требования

- Docker.
- Docker Compose.
- Make.
- Go 1.25.
- Node.js и npm, если используется фронтенд из папки `web`.

## Установка и запуск

### Быстрый старт

```bash
make up
```
Фронтенд-приложение доступно по адресу http://localhost:3000 (Vite dev server, проброшенный из Docker контейнера на порт 3000 хоста).
## Миграции

### Применить миграции

```bash
make migrate-up
```

### Откатить одну миграцию

```bash
make migrate-down
```

## Тестирование

### Запуск тестов

```bash
make test
```

### Запуск тестов с покрытием

```bash
go test ./... -cover
```

## Форматирование и линтинг

### Форматирование Go-кода

```bash
make fmt
```

### Проверка линтером

```bash
make lint
```

### Автоисправление линтером

```bash
make lint-fix
```

### Форматирование через golangci-lint

```bash
make fmt-ci
```

## HTTP API

### Маршруты

```go
r.GET("/home", ren.HomeHandler)
r.POST("/comments", h.CreateHandler)
r.GET("/comments", h.GetRootCommentsHandler)
r.GET("/comments/all", h.GetCommentTreeHandler)
r.GET("/comments/:parent_id/children", h.GetChildCommentsHandler)
r.DELETE("/comments/:id", h.DeleteHandler)
```

### Эндпоинты

| Method | Route | Description |
|---|---|---|
| GET | `/home` | Проверка доступности сервиса |
| POST | `/comments` | Создать комментарий |
| GET | `/comments` | Получить корневые комментарии |
| GET | `/comments/all` | Получить все комментарии деревом |
| GET | `/comments/:parent_id/children` | Получить дочерние комментарии |
| DELETE | `/comments/:id` | Удалить комментарий |

## Примеры запросов

### Создать комментарий

```bash
curl -X POST http://localhost:8080/comments \
  -H "Content-Type: application/json" \
  -d '{"content":"Hello world","parent_id":0}'
```

### Получить корневые комментарии

```bash
curl "http://localhost:8080/comments?limit=10&offset=0&search=test"
```

### Получить дерево комментариев

```bash
curl http://localhost:8080/comments/all
```

## Конфигурация

Основные параметры проекта задаются через переменные окружения и файлы конфигурации в `configs/`.  
Если используется Docker Compose, часть настроек может быть определена в `docker-compose.yml`.

## Полезные команды

```bash
make help
make up
make down
make logs
make ps
make restart
make test
make lint
make build
```

## Примечание

Папки `web/dist` и `web/node_modules` являются артефактами фронтенда и обычно не редактируются вручную.