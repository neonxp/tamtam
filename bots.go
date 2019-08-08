// Package tamtam implements TamTam Bot API
// Copyright (c) 2019 Alexander Kiryukhin <a.kiryukhin@mail.ru>
package tamtam

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type bots struct {
	client *client
}

func newBots(client *client) *bots {
	return &bots{client: client}
}

func (a *bots) GetBot() (*BotInfo, error) {
	result := new(BotInfo)
	values := url.Values{}
	body, err := a.client.request(http.MethodGet, "me", values, nil)
	if err != nil {
		return result, err
	}
	defer func() {
		if err := body.Close(); err != nil {
			log.Println(err)
		}
	}()
	return result, json.NewDecoder(body).Decode(result)
}

func (a *bots) PatchBot(patch *BotPatch) (*BotInfo, error) {
	result := new(BotInfo)
	values := url.Values{}
	body, err := a.client.request(http.MethodPatch, "me", values, patch)
	if err != nil {
		return result, err
	}
	defer func() {
		if err := body.Close(); err != nil {
			log.Println(err)
		}
	}()
	return result, json.NewDecoder(body).Decode(result)
}
