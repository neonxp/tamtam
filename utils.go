package tamtam

type KeyboardBuilder struct {
	rows []*KeyboardRow
}

func NewKeyboardBuilder() *KeyboardBuilder {
	return &KeyboardBuilder{
		rows: make([]*KeyboardRow, 0),
	}
}

func (k *KeyboardBuilder) AddRow() *KeyboardRow {
	kr := &KeyboardRow{}
	k.rows = append(k.rows, kr)
	return kr
}

func (k *KeyboardBuilder) Build() Keyboard {
	buttons := make([][]interface{}, 0, len(k.rows))
	for _, r := range k.rows {
		buttons = append(buttons, r.Build())
	}
	return Keyboard{Buttons: buttons}
}

type KeyboardRow struct {
	cols []interface{}
}

func (k *KeyboardRow) Build() []interface{} {
	return k.cols
}

func (k *KeyboardRow) AddLink(text string, intent Intent, url string) *KeyboardRow {
	b := LinkButton{
		Text:   text,
		Url:    url,
		Intent: intent,
		Type:   LINK,
	}
	k.cols = append(k.cols, b)
	return k
}

func (k *KeyboardRow) AddCallback(text string, intent Intent, payload string) *KeyboardRow {
	b := CallbackButton{
		Text:    text,
		Payload: payload,
		Intent:  intent,
		Type:    CALLBACK,
	}
	k.cols = append(k.cols, b)
	return k
}

func (k *KeyboardRow) AddContact(text string, intent Intent, url string) *KeyboardRow {
	b := RequestContactButton{
		Text:   text,
		Intent: intent,
		Type:   CONTACT,
	}
	k.cols = append(k.cols, b)
	return k
}

func (k *KeyboardRow) AddGeolocation(text string, intent Intent, quick bool) *KeyboardRow {
	b := RequestGeoLocationButton{
		Text:   text,
		Quick:  quick,
		Intent: intent,
		Type:   GEOLOCATION,
	}
	k.cols = append(k.cols, b)
	return k
}
