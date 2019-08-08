// Package tamtam implements TamTam Bot API
// Copyright (c) 2019 Alexander Kiryukhin <a.kiryukhin@mail.ru>
package tamtam

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type chats struct {
	client *client
}

func newChats(client *client) *chats {
	return &chats{client: client}
}

func (a *chats) GetChats(count, marker int) (*ChatList, error) {
	result := new(ChatList)
	values := url.Values{}
	if count > 0 {
		values.Set("count", strconv.Itoa(int(count)))
	}
	if marker > 0 {
		values.Set("marker", strconv.Itoa(int(marker)))
	}
	body, err := a.client.request(http.MethodGet, "chats", values, nil)
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

func (a *chats) GetChat(chatID int) (*Chat, error) {
	result := new(Chat)
	values := url.Values{}
	body, err := a.client.request(http.MethodGet, fmt.Sprintf("chats/%d", chatID), values, nil)
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

func (a *chats) GetChatMembership(chatID int) (*ChatMember, error) {
	result := new(ChatMember)
	values := url.Values{}
	body, err := a.client.request(http.MethodGet, fmt.Sprintf("chats/%d/members/me", chatID), values, nil)
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

func (a *chats) GetChatMembers(chatID, count, marker int) (*ChatMembersList, error) {
	result := new(ChatMembersList)
	values := url.Values{}
	if count > 0 {
		values.Set("count", strconv.Itoa(int(count)))
	}
	if marker > 0 {
		values.Set("marker", strconv.Itoa(int(marker)))
	}
	body, err := a.client.request(http.MethodGet, fmt.Sprintf("chats/%d/members", chatID), values, nil)
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

func (a *chats) LeaveChat(chatID int) (*SimpleQueryResult, error) {
	result := new(SimpleQueryResult)
	values := url.Values{}
	body, err := a.client.request(http.MethodDelete, fmt.Sprintf("chats/%d/members/me", chatID), values, nil)
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

func (a *chats) EditChat(chatID int, update *ChatPatch) (*Chat, error) {
	result := new(Chat)
	values := url.Values{}
	body, err := a.client.request(http.MethodPatch, fmt.Sprintf("chats/%d", chatID), values, update)
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

func (a *chats) AddMember(chatID int, users UserIdsList) (*SimpleQueryResult, error) {
	result := new(SimpleQueryResult)
	values := url.Values{}
	body, err := a.client.request(http.MethodPost, fmt.Sprintf("chats/%d/members", chatID), values, users)
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

func (a *chats) RemoveMember(chatID int, userID int) (*SimpleQueryResult, error) {
	result := new(SimpleQueryResult)
	values := url.Values{}
	values.Set("user_id", strconv.Itoa(int(userID)))
	body, err := a.client.request(http.MethodDelete, fmt.Sprintf("chats/%d/members", chatID), values, nil)
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

func (a *chats) SendAction(chatID int, action SenderAction) (*SimpleQueryResult, error) {
	result := new(SimpleQueryResult)
	values := url.Values{}
	body, err := a.client.request(http.MethodPost, fmt.Sprintf("chats/%d/actions", chatID), values, ActionRequestBody{Action: action})
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
