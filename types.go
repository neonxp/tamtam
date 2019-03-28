package tamtam

type BotInfo struct {
	UserID        int64  `json:"user_id"`
	Name          string `json:"name"`
	Username      string `json:"username,omitempty"`
	AvatarURL     string `json:"avatar_url"`
	FullAvatarURL string `json:"full_avatar_url"`
}

type ChatType string

const (
	TypeDialog  ChatType = "dialog"
	TypeChat             = "chat"
	TypeChannel          = "channel"
)

type StatusType string

const (
	StatusActive    StatusType = "active"
	StatusRemoved              = "removed"
	StatusLeft                 = "left"
	StatusClosed               = "closed"
	StatusSuspended            = "suspended"
)

type Chat struct {
	ChatID int64      `json:"chat_id"`
	Type   ChatType   `json:"type"`
	Status StatusType `json:"status"`
	Title  string     `json:"title"`
	Icon   struct {
		URL string `json:"url"`
	} `json:"icon"`
	LastEventTime     int64       `json:"last_event_time"`
	ParticipantsCount int32       `json:"participants_count"`
	OwnerID           int64       `json:"owner_id"`
	Participants      interface{} `json:"participants,omitempty"`
	IsPublic          bool        `json:"is_public"`
	Link              string      `json:"link,omitempty"`
	Description       string      `json:"description,omitempty"`
}

type Chats struct {
	Chats  []Chat `json:"chats"`
	Marker int    `json:"marker"`
}

type Participant struct {
	UserID   int64  `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username,omitempty"`
}

type Recipient struct {
	ChatID   int64    `json:"chat_id"`
	ChatType ChatType `json:"chat_type"`
	UserID   int64    `json:"user_id,omitempty"`
}

type LinkType string

const (
	LinkForward = "forward"
	LinkReply   = "reply"
)

type Message struct {
	Sender    Participant `json:"sender"`
	Recipient Recipient   `json:"recipient"`
	Timestamp int64       `json:"timestamp"`
	Link      struct {
		Type    LinkType    `json:"type"`
		Sender  Participant `json:"sender"`
		ChatID  int64       `json:"chat_id"`
		Message MessageBody `json:"message"`
	} `json:"link"`
	Body MessageBody `json:"body"`
}

type AttachmentType string

const (
	AttachmentImage    AttachmentType = "image"
	AttachmentVideo                   = "video"
	AttachmentAudio                   = "audio"
	AttachmentFile                    = "file"
	AttachmentContact                 = "contact"
	AttachmentSticker                 = "sticker"
	AttachmentShare                   = "share"
	AttachmentLocation                = "location"
	AttachmentKeyboard                = "inline_keyboard"
)

type Attachment struct {
	Type    AttachmentType `json:"type"`
	Payload interface{}    `json:"payload"`
}

type MessageBody struct {
	MID         string       `json:"mid"`
	Seq         int64        `json:"seq"`
	Text        string       `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments"`
}
