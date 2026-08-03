# MaxBot

**MaxBot** — это легкий и гибкий Go-фреймворк для создания [MAX-ботов](https://dev.max.ru/docs). 
Он предоставляет простой и интуитивно понятный API для обработки обновлений, команд и middleware, что делает разработку ботов простой и эффективной.

## Возможности

*   **Легковесный и быстрый**: Создан с учётом опыта использования, имеет минимальный размер.
*   **Интуитивный API**: Простые для понимания методы для отправки сообщений, обработки команд и многого другого.
*   **Поддержка Middleware**: Создавайте middleware для логирования, аутентификации, ограничения частоты запросов и других задач.
*   **Обработка команд**: Простая и мощная маршрутизация для различных событий бота.
*   **Контекстная обработка**: Каждое обновление обрабатывается с контекстом, который содержит данные, специфичные для запроса.
*   **Модульная архитектура**: Легко расширяйте и настраивайте бота под свои нужды.

## Установка

Для установки MaxBot используйте `go get`:

```bash
go get github.com/max-messenger/maxbot
```

# Быстрый старт
Пример для начала работы.
```go
package main

import (
	"log"
	"os"

	"github.com/max-messenger/maxbot"
)

func main() {
	// Получите токен бота из переменных окружения
	token := os.Getenv("BOT_TOKEN")

	// Создайте новый экземпляр бота
	bot, err := maxbot.NewApi(token)
	if err != nil {
		log.Fatal(err)
	}

	// Определите обработчик для команды /start
	bot.Handle("/start", func(ctx *maxbot.Context) error {
		return ctx.Reply("Привет! Я MaxBot. Чем я могу помочь?")
	})

	// Запустите бота и начните прослушивание обновлений
	log.Println("Бот запускается...")
	bot.Start()
}
```

# Основные концепции

`Bot` — это центральный объект. Вы создаёте его с вашим [MaxBot Token](https://dev.max.ru/docs) и используете для регистрации обработчиков и запуска процесса получения обновлений ([polling](https://dev.max.ru/docs-api/methods/GET/updates) или [webhook](https://dev.max.ru/docs-api/methods/POST/subscriptions)).
```go 
bot := maxbot.New(token)
```

Вы можете настроить Webhook бота с помощью опции:

```go 
bot := maxbot.New(token, maxbot.WithWebhook("https://your-domain.com/webhook"))
```

# Обработка сообщений и команд
Используйте метод `Handle` для регистрации обработчика для конкретной команды или типа сообщения.

```go
// Обработка команды
bot.Handle("/start", startHandler)
bot.Handle("/help", helpHandler)

// Обработка всех текстовых сообщений (не начинающихся с '/')
bot.Handle(maxbot.OnText, textHandler)
```

# Объект `Context`
`*maxbot.Context` передаётся в каждый обработчик и содержит всю информацию о сообщении, а также методы для ответа.
```go
func myHandler(ctx *maxbot.Context) error {
    // Получить текст сообщения
    text := ctx.Message.Text

    // Ответить сообщением
    return ctx.Reply("Вы сказали: " + text)
}
```

# Middleware
Middleware позволяют выполнять код до или после ваших обработчиков. Это полезно для таких задач, как логирование, аутентификация или сбор метрик.
```go
// Определяем middleware для логирования
loggingMiddleware := func(next maxbot.HandlerFunc) maxbot.HandlerFunc {
    return func(ctx *maxbot.Context) error {
        log.Printf("Получено обновление от пользователя %d", ctx.Message.From.ID)
        return next(ctx)
    }
}

// Применяем middleware к конкретному обработчику
bot.Handle("/secret", secretHandler, loggingMiddleware)

// Или применяем глобально
bot.Use(loggingMiddleware)
```

# Примеры
Более подробные примеры можно найти в директории `example/` репозитория:
* [Простой пример](example/simple/README.md): Базовый бот, демонстрирующий обработку команд и ответы.

# Внесение вклада
Приветствуются любые вклады! Пожалуйста, не стесняйтесь отправлять Pull Request.
1. Сделайте форк репозитория.
2. Создайте ветку для вашей функции (`git checkout -b feature/amazing-feature`).
3. Зафиксируйте изменения (`git commit -m 'Add some amazing feature`').
4. Отправьте изменения в ветку (`git push origin feature/amazing-feature`).
5. Откройте `Pull Request`.

# Лицензия
Этот проект лицензирован под Apache License 2.0 - подробности смотрите в файле [LICENSE](LICENSE).
