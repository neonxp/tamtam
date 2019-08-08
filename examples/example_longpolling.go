// +build ignore

/**
 * Updates loop example
 */
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/neonxp/tamtam"
)

func main() {
	// Initialisation
	api := tamtam.New(os.Getenv("TOKEN"))

	// Some methods demo:
	info, err := api.Bots.GetBot()
	log.Printf("Get me: %#v %#v", info, err)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case upd := <-api.GetUpdates():
				log.Printf("Received: %#v", upd)
				switch upd := upd.(type) {
				case *tamtam.MessageCreatedUpdate:
					res, err := api.Messages.SendMessage(0, upd.Message.Sender.UserId, &tamtam.NewMessageBody{
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

	if err := api.UpdatesLoop(ctx); err != nil {
		log.Fatalln(err)
	}

}
