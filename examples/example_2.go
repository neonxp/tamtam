// +build ignore

/**
 * Updates loop example
 */
package main

import (
	"context"
	"fmt"
	"github.com/neonxp/tamtam"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Initialisation
	api := tamtam.New(os.Getenv("TOKEN"))

	// Some methods demo:
	info, err := api.GetMe()
	log.Printf("Get me: %#v %#v", info, err)
	chats, err := api.GetChats(0, 0)
	log.Printf("Get chats: %#v %#v", chats, err)
	chat, err := api.GetChat(chats.Chats[0].ChatId)
	log.Printf("Get chat: %#v %#v", chat, err)
	subs, _ := api.GetSubscriptions()
	for _, s := range subs.Subscriptions {
		_, _ = api.Unsubscribe(s.Url)
	}
	ch := make(chan interface{}, 1) // Channel with updates from TamTam

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case upd := <-ch:
				log.Printf("Received: %#v", upd)
				switch upd := upd.(type) {
				case tamtam.UpdateMessageCreated:
					res, err := api.SendMessage(0, upd.Message.Sender.UserId, &tamtam.NewMessageBody{
						Text: fmt.Sprintf("Hello, %s! Your message: %s", upd.Message.Sender.Name, upd.Message.Body.Text),
					})
					log.Printf("Answer: %#v %#v", res, err)
				default:
					log.Printf("Unknown type: %#v", upd)
				}
			case <-ctx.Done():
				return
			}

		}
	}()

	go func() {
		exit := make(chan os.Signal)
		signal.Notify(exit, os.Kill, os.Interrupt)
		select {
		case <-exit:
			cancel()
		case <-ctx.Done():
			return
		}
	}()

	if err := api.GetUpdatesLoop(ctx, ch); err != nil {
		log.Fatalln(err)
	}

}
