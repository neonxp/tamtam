//Package tamtam implements TamTam Bot API
//Copyright (c) 2019 Alexander Kiryukhin <a.kiryukhin@mail.ru>
package tamtam

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type messages struct {
	client *client
}

func newMessages(client *client) *messages {
	return &messages{client: client}
}

//GetMessages returns messages in chat: result page and marker referencing to the next page. Messages traversed in reverse direction so the latest message in chat will be first in result array. Therefore if you use from and to parameters, to must be less than from
func (a *messages) GetMessages(chatID int, messageIDs []string, from int, to int, count int) (*MessageList, error) {
	result := new(MessageList)
	values := url.Values{}
	if chatID != 0 {
		values.Set("chat_id", strconv.Itoa(int(chatID)))
	}
	if len(messageIDs) > 0 {
		for _, mid := range messageIDs {
			values.Add("message_ids", mid)
		}
	}
	if from != 0 {
		values.Set("from", strconv.Itoa(int(from)))
	}
	if to != 0 {
		values.Set("to", strconv.Itoa(int(to)))
	}
	if count > 0 {
		values.Set("count", strconv.Itoa(int(count)))
	}
	body, err := a.client.request(http.MethodGet, "messages", values, nil)
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

//SendMessage sends a message to a chat. As a result for this method new message identifier returns.
func (a *messages) SendMessage(chatID int, userID int, message *NewMessageBody) (*Message, error) {
	result := new(Message)
	values := url.Values{}
	if chatID != 0 {
		values.Set("chat_id", strconv.Itoa(int(chatID)))
	}
	if userID != 0 {
		values.Set("user_id", strconv.Itoa(int(userID)))
	}
	body, err := a.client.request(http.MethodPost, "messages", values, message)
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

//EditMessage updates message by id
func (a *messages) EditMessage(messageID int, message *NewMessageBody) (*SimpleQueryResult, error) {
	result := new(SimpleQueryResult)
	values := url.Values{}
	values.Set("message_id", strconv.Itoa(int(messageID)))
	body, err := a.client.request(http.MethodPut, "messages", values, message)
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

//DeleteMessage deletes message by id
func (a *messages) DeleteMessage(messageID int) (*SimpleQueryResult, error) {
	result := new(SimpleQueryResult)
	values := url.Values{}
	values.Set("message_id", strconv.Itoa(int(messageID)))
	body, err := a.client.request(http.MethodDelete, "messages", values, nil)
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

//AnswerOnCallback should be called to send an answer after a user has clicked the button. The answer may be an updated message or/and a one-time user notification.
func (a *messages) AnswerOnCallback(callbackID int, callback *CallbackAnswer) (*SimpleQueryResult, error) {
	result := new(SimpleQueryResult)
	values := url.Values{}
	values.Set("callback_id", strconv.Itoa(int(callbackID)))
	body, err := a.client.request(http.MethodPost, "answers", values, callback)
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

//NewKeyboardBuilder returns new keyboard builder helper
func (a *messages) NewKeyboardBuilder() *KeyboardBuilder {
	return &KeyboardBuilder{
		rows: make([]*KeyboardRow, 0),
	}
}
