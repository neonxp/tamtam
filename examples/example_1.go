/**
 * Webhook example
 */
package main

import (
	"fmt"
	"github.com/neonxp/tamtam"
	"log"
	"net/http"
	"os"
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
	msgs, err := api.GetMessages(chats.Chats[0].ChatId, nil, 0, 0, 0)
	log.Printf("Get messages: %#v %#v", msgs, err)
	subs, _ := api.GetSubscriptions()
	for _, s := range subs.Subscriptions {
		_, _ = api.Unsubscribe(s.Url)
	}
	subscriptionResp, err := api.Subscribe("https://576df2ec.ngrok.io/webhook", []string{})
	log.Printf("Subscription: %#v %#v", subscriptionResp, err)

	ch := make(chan interface{}) // Channel with updates from TamTam

	http.HandleFunc("/webhook", api.GetHandler(ch))
	go func() {
		for {
			upd := <-ch
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
		}
	}()

	http.ListenAndServe(":10888", nil)
}
