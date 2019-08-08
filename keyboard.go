package tamtam

//KeyboardBuilder implements builder for inline keyboard
type KeyboardBuilder struct {
	rows []*KeyboardRow
}

//AddRow adds row to inline keyboard
func (k *KeyboardBuilder) AddRow() *KeyboardRow {
	kr := &KeyboardRow{}
	k.rows = append(k.rows, kr)
	return kr
}

//Build returns result keyboard
func (k *KeyboardBuilder) Build() Keyboard {
	buttons := make([][]ButtonInterface, 0, len(k.rows))
	for _, r := range k.rows {
		buttons = append(buttons, r.Build())
	}
	return Keyboard{Buttons: buttons}
}

//KeyboardRow represents buttons row
type KeyboardRow struct {
	cols []ButtonInterface
}

//Build returns result keyboard row
func (k *KeyboardRow) Build() []ButtonInterface {
	return k.cols
}

//AddLink button
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

//AddCallback button
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

//AddContact button
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

//AddGeolocation button
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
