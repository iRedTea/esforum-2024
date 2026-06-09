# EasyStartupForum

Небольшой Go-сервис-форум. Написан в 2024, как pet-проект, код открыт в 2026 (обновлен только Dockerfile и README! Остальной код сохранен так, как он выглядел в 2024! Никакого рефакторинга с тех пор не производилось!).

### Технологии
- Go
- Rest API
- Gin
- JWT
- SQL (MySQL)
- Swagger

### Быстрый старт

Требования: `Go >= 1.22`, Makefile (оптионально), Docker (опционально).

Сборка локально:

```bash
go build -o esforum ./cmd/release
./esforum
```


### Переменные окружения

- `ADMIN_TOKEN` - токен администратора (используется в `pkg/service/auth.go`).
- `DB_PASSWORD` - пароль для ДБ

Все остальные параметры хранятся в configs/config.yml

### API Endpoints
- `GET /api/guest/thread/:id` : Получить тред по ID
- `GET /api/guest/thread/all` : Получить список всех тредов
- `GET /api/guest/thread/:id/posts` : Получить все посты треда
- `GET /api/guest/post/:id` : Получить пост по ID
- `POST /api/thread` : Создать тред
- `DELETE /api/thread/:id` : Удалить тред по ID
- `POST /api/post` : Создать пост
- `DELETE /api/post/:id` : Удалить пост по ID
- `GET /swagger/*any` : Swagger UI и документация
