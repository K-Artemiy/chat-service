# test_qa-api  
REST API сервис чатов 

##### Используемые технологии:
1. ___Go 1.25.1___
2. ___PostgreSQL___
3. ___net/http___ для реализации HTTP API сервера
4. ___GORM___ для взаимодействия с БД
5. ___Goose___ для реализации миграций 
6. ___Docker + docker-compose___
7. ___testing + testify___ для тестов



### Реализованные модели:
1. ___Chat___ – чат: 
    - id: int 
    - title: str (длина 1-200) 
    - created_at: datetime 
2. ___Message___ – сообщение в чате: 
    - id: int 
    - chat_id: int (FK на чат) 
    - text: str (1-5000) 
    - created_at: datetime


### Методы API:
- <span style='color: green;'>POST /chats/ </span>— создать новый чат 

    *Тело запроса:*
```json: {"title": "..."}```
- <span style='color: green;'>GET /chats/{id} </span>— получить чат и limit сообщений в нём (limit задаётся от 1 до 100 в query параметрах)
- <span style='color: green;'>DELETE /chats/{id} </span>— удалить чат (вместе с сообщениями) 
- <span style='color: green;'>POST /chats/{id}/messages/ </span>— добавить сообщение в чат 

    *Тело запроса:*
```json: {"text": "..."}```

### Запуск проекта

#### 1. Установка

Склонируйте репозиторий: 
```git clone https://github.com/K-Artemiy/chat-service.git```

И перейдите в каталог проекта:
```cd chat-service```

#### 2. Копирование переменных окружения

Скопируйте переменные окружения в этот каталог:
```cp .env.example .env```

#### 3. Запуск сервиса

Запустите сервис (для этого дополнительно необходимо запустить Docker Desktop):
```docker-compose up --build```

После запуска:
- API будет доступно по адресу: http://localhost:8080
- PostgreSQL будет работать в контейнере postres

Миграции применяются автоматически при старте  приложения.

####  Запуск тестов 
Для запуска тестов:
```go test ./...```
