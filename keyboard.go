// Package tamtam implements TamTam Bot API
// Copyright (c) 2019 Alexander Kiryukhin <a.kiryukhin@mail.ru>
package tamtam

type KeyboardBuilder struct {
	rows []*KeyboardRow
}

func (k *KeyboardBuilder) AddRow() *KeyboardRow {
	kr := &KeyboardRow{}
	k.rows = append(k.rows, kr)
	return kr
}

func (k *KeyboardBuilder) Build() Keyboard {
	buttons := make([][]ButtonInterface, 0, len(k.rows))
	for _, r := range k.rows {
		buttons = append(buttons, r.Build())
	}
	return Keyboard{Buttons: buttons}
}

type KeyboardRow struct {
	cols []ButtonInterface
}

func (k *KeyboardRow) Build() []ButtonInterface {
	return k.cols
}

func (k *KeyboardRow) AddLink(text string, intent Intent, url string) *KeyboardRow {
	b := LinkButton{
		Url: url,
		Button: Button{
			Text: text,
			Type: LINK,
		},
	}
	k.cols = append(k.cols, b)
	return k
}

func (k *KeyboardRow) AddCallback(text string, intent Intent, payload string) *KeyboardRow {
	b := CallbackButton{
		Payload: payload,
		Intent:  intent,
		Button: Button{
			Text: text,
			Type: CALLBACK,
		},
	}
	k.cols = append(k.cols, b)
	return k
}

func (k *KeyboardRow) AddContact(text string) *KeyboardRow {
	b := RequestContactButton{
		Button: Button{
			Text: text,
			Type: CONTACT,
		},
	}
	k.cols = append(k.cols, b)
	return k
}

func (k *KeyboardRow) AddGeolocation(text string, quick bool) *KeyboardRow {
	b := RequestGeoLocationButton{
		Quick: quick,
		Button: Button{
			Text: text,
			Type: GEOLOCATION,
		},
	}
	k.cols = append(k.cols, b)
	return k
}
