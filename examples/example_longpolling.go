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
		for upd := range api.GetUpdates(ctx) {
			log.Printf("Received: %#v", upd)
			switch upd := upd.(type) {
			case *tamtam.MessageCreatedUpdate:
				err := api.Messages.Send(
					tamtam.NewMessage().
						SetUser(upd.Message.Sender.UserId).
						SetText(fmt.Sprintf("Hello, %s! Your message: %s", upd.Message.Sender.Name, upd.Message.Body.Text)),
				)
				log.Printf("Answer: %#v %#v", res, err)
			default:
				log.Printf("Unknown type: %#v", upd)
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
	<-ctx.Done()
}
