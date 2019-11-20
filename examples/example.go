package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/neonxp/tamtam"
	"github.com/neonxp/tamtam/schemes"
)

func main() {
	api := tamtam.New(os.Getenv("TOKEN"))

	info, err := api.Bots.GetBot() // Простой метод
	log.Printf("Get me: %#v %#v", info, err)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		exit := make(chan os.Signal)
		signal.Notify(exit, os.Kill, os.Interrupt)
		<-exit
		cancel()
	}()
	for upd := range api.GetUpdates(ctx) { // Чтение из канала с обновлениями
		log.Printf("Received: %#v", upd)
		switch upd := upd.(type) { // Определение типа пришедшего обновления
		case *schemes.MessageCreatedUpdate:
			// Создание клавиатуры
			keyboard := api.Messages.NewKeyboardBuilder()
			keyboard.
				AddRow().
				AddGeolocation("Прислать геолокацию", true).
				AddContact("Прислать контакт")
			keyboard.
				AddRow().
				AddLink("Библиотека", schemes.POSITIVE, "https://github.com/neonxp/tamtam").
				AddCallback("Колбек 1", schemes.NEGATIVE, "callback_1").
				AddCallback("Колбек 2", schemes.NEGATIVE, "callback_2")
			keyboard.
				AddRow().
				AddCallback("Картинка", schemes.POSITIVE, "picture")

			// Отправка сообщения с клавиатурой
			err := api.Messages.Send(tamtam.NewMessage().SetUser(upd.Message.Recipient.UserId).AddKeyboard(keyboard).SetText("Привет!"))
			log.Printf("Answer: %#v", err)
		case *schemes.MessageCallbackUpdate:
			// Ответ на коллбек
			if upd.Callback.Payload == "picture" {
				photo, err := api.Uploads.UploadPhotoFromFile("./examples/example.jpg")
				if err != nil {
					log.Fatal(err)
				}
				msg := tamtam.NewMessage().SetUser(upd.Message.Recipient.UserId).AddPhoto(photo)
				if err := api.Messages.Send(msg); err != nil {
					log.Fatal(err)
				}
			}
		default:
			log.Printf("Unknown type: %#v", upd)
		}
	}
}
