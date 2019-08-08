// +build ignore

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
	for upd := range api.GetUpdates() {      // Чтение из канала с обновлениями
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