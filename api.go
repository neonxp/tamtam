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

	"github.com/neonxp/tamtam/schemes"
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

func (a *Api) bytesToProperUpdate(b []byte) schemes.UpdateInterface {
	u := new(schemes.Update)
	_ = json.Unmarshal(b, u)
	switch u.GetUpdateType() {
	case schemes.TypeMessageCallback:
		upd := new(schemes.MessageCallbackUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case schemes.TypeMessageCreated:
		upd := new(schemes.MessageCreatedUpdate)
		_ = json.Unmarshal(b, upd)
		for _, att := range upd.Message.Body.RawAttachments {
			upd.Message.Body.Attachments = append(upd.Message.Body.Attachments, a.bytesToProperAttachment(att))
		}
		return upd
	case schemes.TypeMessageRemoved:
		upd := new(schemes.MessageRemovedUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case schemes.TypeMessageEdited:
		upd := new(schemes.MessageEditedUpdate)
		_ = json.Unmarshal(b, upd)
		for _, att := range upd.Message.Body.RawAttachments {
			upd.Message.Body.Attachments = append(upd.Message.Body.Attachments, a.bytesToProperAttachment(att))
		}
		return upd
	case schemes.TypeBotAdded:
		upd := new(schemes.BotAddedToChatUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case schemes.TypeBotRemoved:
		upd := new(schemes.BotRemovedFromChatUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case schemes.TypeUserAdded:
		upd := new(schemes.UserAddedToChatUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case schemes.TypeUserRemoved:
		upd := new(schemes.UserRemovedFromChatUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case schemes.TypeBotStarted:
		upd := new(schemes.BotStartedUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	case schemes.TypeChatTitleChanged:
		upd := new(schemes.ChatTitleChangedUpdate)
		_ = json.Unmarshal(b, upd)
		return upd
	}
	return nil
}

func (a *Api) bytesToProperAttachment(b []byte) schemes.AttachmentInterface {
	attachment := new(schemes.Attachment)
	_ = json.Unmarshal(b, attachment)
	switch attachment.GetAttachmentType() {
	case schemes.AttachmentAudio:
		res := new(schemes.AudioAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case schemes.AttachmentContact:
		res := new(schemes.ContactAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case schemes.AttachmentFile:
		res := new(schemes.FileAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case schemes.AttachmentImage:
		res := new(schemes.PhotoAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case schemes.AttachmentKeyboard:
		res := new(schemes.InlineKeyboardAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case schemes.AttachmentLocation:
		res := new(schemes.LocationAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case schemes.AttachmentShare:
		res := new(schemes.ShareAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case schemes.AttachmentSticker:
		res := new(schemes.StickerAttachment)
		_ = json.Unmarshal(b, res)
		return res
	case schemes.AttachmentVideo:
		res := new(schemes.VideoAttachment)
		_ = json.Unmarshal(b, res)
		return res
	}
	return attachment
}

func (a *Api) getUpdates(limit int, timeout int, marker int64, types []string) (*schemes.UpdateList, error) {
	result := new(schemes.UpdateList)
	values := url.Values{}
	if limit > 0 {
		values.Set("limit", strconv.Itoa(limit))
	}
	if timeout > 0 {
		values.Set("timeout", strconv.Itoa(timeout))
	}
	if marker > 0 {
		values.Set("marker", strconv.Itoa(int(marker)))
	}
	if len(types) > 0 {
		for _, t := range types {
			values.Add("types", t)
		}
	}
	body, err := a.client.request(http.MethodGet, "updates", values, nil)
	if err != nil {
		if err == errLongPollTimeout {
			return result, nil
		}
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
func (a *Api) GetUpdates(ctx context.Context) chan schemes.UpdateInterface {
	ch := make(chan schemes.UpdateInterface)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case <-time.After(time.Duration(a.pause) * time.Second):
				var marker int64
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
