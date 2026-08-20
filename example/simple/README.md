# Пример использования MaxBot: simple

Этот пример демонстрирует базовые возможности фреймворка для создания MaxBot-ботов. Он показывает, как настраивать бота,
обрабатывать команды, callback-запросы и различные типы событий.

## Обзор

Файл [`main.go`](main.go) представляет собой работающий
пример бота, который умеет:

* Отвечать на команды (`/help`, `/command`, `/reply`).
* Обрабатывать нажатия на кнопки (callback).
* Реагировать на события (например, изменение названия чата).
* Отправлять сообщения с клавиатурой.
* Логировать входящие текстовые сообщения.

## Код и пояснения

### 1. Инициализация бота

```go
opts := []maxbot.Opt{
  maxbot.WithHTTPClient(&http.Client{Timeout: 25 * time.Second}),
  // maxbot.WithWebhook("http://my-bot.cloud.hooli.local/webhook", "secret", []string{
  //     maxbot.OnBotAdded,
  //     maxbot.OnMessageCreated,
  //     maxbot.OnMessageCallback,
  // }),
  }
  
  token := os.Getenv("BOT_TOKEN")
  bot, err := maxbot.NewApi(token, opts...)
  if err != nil {
  log.Fatal(err)
}
```

* Настройка: Создаётся список опций (opts). В примере устанавливается таймаут для HTTP-клиента.
* Webhook (закомментирован): Показано, как настроить бота на получение обновлений через webhook вместо long polling.
  Можно указать URL, секрет и список типов событий для подписки.
* Создание: Бот создаётся с помощью maxbot.NewApi(), используя токен из переменной окружения BOT_TOKEN.

### 2. Обработка команд

Команда `/help`

```go
bot.Handle("/help", func (c maxbot.Context) error {
  kb := model.NewKeyboard()
  kb.AddRow().
  AddLink("Документация", "https://dev.max.ru/docs").
  AddCallBack("нажми на меня", "pushBtn")
  
  return c.Send("Основная информация:", maxbot.WithKeyboard(kb))
})
```

* **Действие**: При получении команды `/help` бот создаёт клавиатуру с двумя кнопками: ссылкой и callback-кнопкой.
* **Ответ**: Отправляет сообщение с этой клавиатурой.

### Команда `/command`

```go
bot.Handle("/command", func (c maxbot.Context) error {
  command := c.Update().GetCommand()
  msg := fmt.Sprintf(
  "command: %s\nbot name: %s\n params: \n%s\n text: %s\n",
  command.Command, command.BotName,
  strings.Join(command.Params, "\n"),
  command.RemainingText,
  )
  
  return c.Send(msg)
})
```

* Действие: Эта команда предназначена для отладки. Она извлекает из контекста структуру команды (`GetCommand()`) и
  отправляет обратно всю информацию о ней: имя команды, имя бота, параметры и оставшийся текст.

### Команда /reply

```go
bot.Handle("/reply", func (c maxbot.Context) error {
  kb := model.NewKeyboard()
  kb.AddRow().
  AddLink("docs", "https://dev.max.ru/docs")
  
  return c.Reply("reply", maxbot.WithKeyboard(kb))
})
```

* **Действие**: Отвечает на сообщение, вызвавшее команду. Использует `c.Reply()`, чтобы ответить "reply" с клавиатурой,
  содержащей ссылку.

### 3. Обработка Callback-запросов

```go
bot.HandleCallback("pushBtn", func (c maxbot.Context) error {
kb := model.NewKeyboard()
kb.AddRow().
AddLink("Документация", "https://dev.max.ru/docs")

return c.Answer("Изменено", maxbot.WithKeyboard(kb))
})
```

* **Регистрация**: `bot.HandleCallback("pushBtn", ...)` связывает callback-данные "pushBtn" с обработчиком.
* **Действие**: Когда пользователь нажимает кнопку "push me baby" из команды `/help`, бот отвечает на callback (
  `c.Answer()`) сообщением "Изменено" с новой клавиатурой.

### 4. Обработка событий

```go
bot.Handle(maxbot.OnChatTitleChangedEvent, func (c maxbot.Context) error {
return c.Send("Заголовок чата изменен")
})
```

* **Событие**: Обработчик регистрируется на событие `maxbot.OnChatTitleChangedEvent`.
* **Действие**: При любом изменении названия чата, в котором находится бот, он отправляет уведомление "Заголовок
  изменен".

### 5. Обработка текстовых сообщений

```go
bot.Handle(maxbot.OnMessageCreated, func (c maxbot.Context) error {
//err = c.Send(fmt.Sprintf("%s - принято", c.Update().GetMessage().Body.Text))
//if err != nil {
//	return err
//}
fmt.Println("-->", c.Update().GetMessage().Body.Text)

return nil
})
```

* **Событие**: Обработчик для `maxbot.OnMessageCreated` срабатывает на все текстовые сообщения, которые не являются командами.
* **Действие**: В текущей версии код отправки ответа закомментирован, и бот просто выводит полученный текст в
  стандартный
  вывод (`fmt.Println`). Это полезно для отладки или логирования.

### 6. Запуск бота

```go
bot.Start()
```

* **Запуск**: Последняя строка запускает бота. Он начинает получать и обрабатывать обновления, используя настроенный
  метод (
  long polling или webhook).

### Заключение

Этот пример охватывает основные сценарии использования MaxBot:

* Настройка клиента.
* Регистрация обработчиков для команд, callback-ов и событий.
* Отправка сообщений и создание клавиатур.
* Работа с контекстом запроса.

Вы можете использовать этот код как основу для вашего собственного бота, модифицируя обработчики под ваши задачи и
раскомментируя/настраивая webhook, если требуется.
