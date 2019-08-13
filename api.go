//Package tamtam implements TamTam Bot API.
//Official documentation: https://dev.tamtam.chat/
package tamtam

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//Api implements main part of TamTam API
type Api struct {
	Bots          *bots
	Chats         *chats
	Messages      *messages
	Subscriptions *subscriptions
	Uploads       *uploads
	client        *client
	timeout       int
	pause         int
}

// New TamTam Api object
func New(key string) *Api {
	timeout := 30
	u, _ := url.Parse("https://botapi.tamtam.chat/")
	cl := newClient(key, "0.1.8", u, &http.Client{Timeout: time.Duration(timeout) * time.Second})
	return &Api{
		Bots:          newBots(cl),
		Chats:         newChats(cl),
		Uploads:       newUploads(cl),
		Messages:      newMessages(cl),
		Subscriptions: newSubscriptions(cl),
		client:        cl,
		timeout:       timeout,
		pause:         1,
	}
}

func (a *Api) bytesToProperUpdate(b []byte) UpdateInterface {
	u := new(Update)
	_ = json.Unmarshal(b, u)
	switch u.GetUpdateType() {
	case TypeMessageCallback:
		upd := new(MessageCallbackUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case TypeMessageCreated:
		upd := new(MessageCreatedUpdate)
		_ = json.Unmarshal(b, upd)
		for _, att := range upd.Message.Body.RawAttachments {
			upd.Message.Body.Attachments = append(upd.Message.Body.Attachments, a.bytesToProperAttachment(att))
		}
		return upd
	case TypeMessageRemoved:
		upd := new(MessageRemovedUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case TypeMessageEdited:
		upd := new(MessageEditedUpdate)
		_ = json.Unmarshal(b, upd)
		for _, att := range upd.Message.Body.RawAttachments {
			upd.Message.Body.Attachments = append(upd.Message.Body.Attachments, a.bytesToProperAttachment(att))
		}
		return upd
	case TypeBotAdded:
		upd := new(BotAddedToChatUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case TypeBotRemoved:
		upd := new(BotRemovedFromChatUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case TypeUserAdded:
		upd := new(UserAddedToChatUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case TypeUserRemoved:
		upd := new(UserRemovedFromChatUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case TypeBotStarted:
		upd := new(BotStartedUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case TypeChatTitleChanged:
		upd := new(ChatTitleChangedUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	}
	return nil
}

func (a *Api) bytesToProperAttachment(b []byte) AttachmentInterface {
	attachment := new(Attachment)
	_ = json.Unmarshal(b, attachment)
	switch attachment.GetAttachmentType() {
	case AttachmentAudio:
		res := new(AudioAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case AttachmentContact:
		res := new(ContactAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case AttachmentFile:
		res := new(FileAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case AttachmentImage:
		res := new(PhotoAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case AttachmentKeyboard:
		res := new(InlineKeyboardAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case AttachmentLocation:
		res := new(LocationAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case AttachmentShare:
		res := new(ShareAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case AttachmentSticker:
		res := new(StickerAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case AttachmentVideo:
		res := new(VideoAttachment)
		_ = json.Unmarshal(b, res)
		return res
	}
	return attachment
}

func (a *Api) getUpdates(limit int, timeout int, marker int, types []string) (*UpdateList, error) {
	result := new(UpdateList)
	values := url.Values{}
	if limit > 0 {
		values.Set("limit", strconv.Itoa(limit))
	}
	if timeout > 0 {
		values.Set("timeout", strconv.Itoa(timeout))
	}
	if marker > 0 {
		values.Set("marker", strconv.Itoa(marker))
	}
	if len(types) > 0 {
		for _, t := range types {
			values.Add("types", t)
		}
	}
	body, err := a.client.request(http.MethodGet, "updates", values, nil)
	if err != nil {
		return result, err
	}
	defer func() {
		if err := body.Close(); err != nil {
			log.Println(err)
		}
	}()
	jb, _ := ioutil.ReadAll(body)
	return result, json.Unmarshal(jb, result)
}

//GetUpdates returns updates channel
func (a *Api) GetUpdates(ctx context.Context) chan UpdateInterface {
	ch := make(chan UpdateInterface)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case <-time.After(time.Duration(a.pause) * time.Second):
				var marker int
				for {
					upds, err := a.getUpdates(50, a.timeout, marker, []string{})
					if err != nil {
						log.Println(err)
						break
					}
					if len(upds.Updates) == 0 {
						break
					}
					for _, u := range upds.Updates {
						ch <- a.bytesToProperUpdate(u)
					}
					marker = *upds.Marker
				}
			}
		}
	}()
	return ch
}

//GetHandler returns http handler for webhooks
func (a *Api) GetHandler(updates chan interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := r.Body.Close(); err != nil {
				log.Println(err)
			}
		}()
		b, _ := ioutil.ReadAll(r.Body)
		updates <- a.bytesToProperUpdate(b)
	}
}
