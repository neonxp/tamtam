# TamTam Go

[![Sourcegraph](https://sourcegraph.com/github.com/neonxp/tamtam/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/neonxp/tamtam?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/neonxp/tamtam)
[![Go Report Card](https://goreportcard.com/badge/github.com/neonxp/tamtam?style=flat-square)](https://goreportcard.com/report/github.com/neonxp/tamtam)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/neonxp/tamtam/master/LICENSE)

Простая реализация клиента к TamTam Bot API на Go. Поддерживается получение обновление как с помощью вебхуков, так и лонгполлингом.

Поддерживаемая версия API - 0.1.8

## Документация
В общем случае, методы повторяют такие из [официальной документации](https://dev.tamtam.chat/)
Так же добавлены хелпер для создания клавиатуры (`api.Messages.NewKeyboardBuilder()`) и для загрузки вложений (`api.Uploads.UploadMedia(uploadType UploadType, filename string)`). 

Пример создания клавиатуры см. ниже в примере.
 
Остальное описано тут http://godoc.org/github.com/neonxp/tamtam/ и в примерах из директории [examples](https://github.com/neonxp/tamtam/tree/master/examples)

## Пример

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/neonxp/tamtam"
)

func main() {
	api := tamtam.New(os.Getenv("TOKEN"))

	info, err := api.Bots.GetBot() // Простой метод
	log.Printf("Get me: %#v %#v", info, err)
	go api.UpdatesLoop(context.Background()) // Запуск цикла получения обновлений
	for upd := range api.GetUpdates() { // Чтение из канала с обновлениями
		log.Printf("Received: %#v", upd)
		switch upd := upd.(type) { // Определение типа пришедшего обновления
		case *tamtam.MessageCreatedUpdate:
			// Создание клавиатуры
			keyboard := api.Messages.NewKeyboardBuilder()
			keyboard.
				AddRow().
				AddGeolocation("Прислать геолокацию", true).
				AddContact("Прислать контакт")
			keyboard.
				AddRow().
				AddLink("Библиотека", tamtam.POSITIVE, "https://github.com/neonxp/tamtam").
				AddCallback("Колбек 1", tamtam.NEGATIVE, "callback_1").
				AddCallback("Колбек 2", tamtam.NEGATIVE, "callback_2")

			// Отправка сообщения с клавиатурой
			res, err := api.Messages.SendMessage(0, upd.Message.Sender.UserId, &tamtam.NewMessageBody{
				Text: fmt.Sprintf("Hello, %s! Your message: %s", upd.Message.Sender.Name, upd.Message.Body.Text),
				Attachments: []interface{}{
					tamtam.NewInlineKeyboardAttachmentRequest(keyboard.Build()),
				},
			})
			log.Printf("Answer: %#v %#v", res, err)
		case *tamtam.MessageCallbackUpdate:
			res, err := api.Messages.SendMessage(0, upd.Callback.User.UserId, &tamtam.NewMessageBody{
				Text: "Callback: " + upd.Callback.Payload,
			})
			log.Printf("Answer: %#v %#v", res, err)
		default:
			log.Printf("Unknown type: %#v", upd)
		}
	}
}
```

## Автор

Александр NeonXP Кирюхин  <a.kiryukhin@mail.ru>
