/*
 * TamTam Bot API
 */
package tamtam

import "encoding/json"

type ActionRequestBody struct {
	Action SenderAction `json:"action"`
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

// Generic schema representing message attachment
type Attachment struct {
	Type AttachmentType `json:"type,omitempty"`
}

type AttachmentPayload struct {
	// Media attachment URL
	Url string `json:"url"`
}

// Request to attach some data to message
type AttachmentRequest struct {
	Type string `json:"type,omitempty"`
}

type AudioAttachment struct {
	Type    string            `json:"type,omitempty"`
	Payload AttachmentPayload `json:"payload"`
}

// Request to attach audio to message. MUST be the only attachment in message
type AudioAttachmentRequest struct {
	Type    string       `json:"type,omitempty"`
	Payload UploadedInfo `json:"payload"`
}

// You will receive this update when bot has been added to chat
type BotAddedToChatUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
	// Chat id where bot was added
	ChatId int64 `json:"chat_id"`
	// User id who added bot to chat
	UserId int64 `json:"user_id"`
}

// You will receive this update when bot has been removed from chat
type BotRemovedFromChatUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
	// Chat identifier bot removed from
	ChatId int64 `json:"chat_id"`
	// User id who removed bot from chat
	UserId int64 `json:"user_id"`
}

// Bot gets this type of update as soon as user pressed `Start` button
type BotStartedUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
	// Dialog identifier where event has occurred
	ChatId int64 `json:"chat_id"`
	// User pressed the 'Start' button
	UserId int64 `json:"user_id"`
}

type Button struct {
	Type string `json:"type,omitempty"`
	// Visible text of button
	Text string `json:"text"`
	// Intent of button. Affects clients representation.
	Intent Intent `json:"intent"`
}

// Object sent to bot when user presses button
type Callback struct {
	// Unix-time when user pressed the button
	Timestamp int64 `json:"timestamp"`
	// Current keyboard identifier
	CallbackId string `json:"callback_id"`
	// Button payload
	Payload string `json:"payload"`
	// User pressed the button
	User User `json:"user"`
}

// Send this object when your bot wants to react to when a button is pressed
type CallbackAnswer struct {
	UserId int64 `json:"user_id,omitempty"`
	// Fill this if you want to modify current message
	Message NewMessageBody `json:"message,omitempty"`
	// Fill this if you just want to send one-time notification to user
	Notification string `json:"notification,omitempty"`
}

// After pressing this type of button client sends to server payload it contains
type CallbackButton struct {
	Type string `json:"type,omitempty"`
	// Visible text of button
	Text string `json:"text"`
	// Intent of button. Affects clients representation.
	Intent Intent `json:"intent"`
	// Button payload
	Payload string `json:"payload"`
}

type Chat struct {
	// Chats identifier
	ChatId int64 `json:"chat_id"`
	// Type of chat. One of: dialog, chat, channel
	Type ChatType `json:"type"`
	// Chat status. One of:  - active: bot is active member of chat  - removed: bot was kicked  - left: bot intentionally left chat  - closed: chat was closed
	Status ChatStatus `json:"status"`
	// Visible title of chat
	Title string `json:"title"`
	// Icon of chat
	Icon Image `json:"icon"`
	// Time of last event occured in chat
	LastEventTime int64 `json:"last_event_time"`
	// Number of people in chat. Always 2 for `dialog` chat type
	ParticipantsCount int32 `json:"participants_count"`
	// Identifier of chat owner. Visible only for chat admins
	OwnerId int64 `json:"owner_id,omitempty"`
	// Participants in chat with time of last activity. Can be *null* when you request list of chats. Visible for chat admins only
	Participants map[string]int64 `json:"participants,omitempty"`
	// Is current chat publicly available. Always `false` for dialogs
	IsPublic bool `json:"is_public"`
	// Link on chat if it is public
	Link string `json:"link,omitempty"`
	// Chat description
	Description map[string]interface{} `json:"description"`
}

// ChatAdminPermission : Chat admin permissions
type ChatAdminPermission string

// List of ChatAdminPermission
const (
	READ_ALL_MESSAGES  ChatAdminPermission = "read_all_messages"
	ADD_REMOVE_MEMBERS ChatAdminPermission = "add_remove_members"
	ADD_ADMINS         ChatAdminPermission = "add_admins"
	CHANGE_CHAT_INFO   ChatAdminPermission = "change_chat_info"
	PIN_MESSAGE        ChatAdminPermission = "pin_message"
	WRITE              ChatAdminPermission = "write"
)

type ChatList struct {
	// List of requested chats
	Chats []Chat `json:"chats"`
	// Reference to the next page of requested chats
	Marker int64 `json:"marker"`
}

type ChatMember struct {
	// Users identifier
	UserId int64 `json:"user_id"`
	// Users visible name
	Name string `json:"name"`
	// Unique public user name. Can be `null` if user is not accessible or it is not set
	Username string `json:"username"`
	// URL of avatar
	AvatarUrl string `json:"avatar_url"`
	// URL of avatar of a bigger size
	FullAvatarUrl  string `json:"full_avatar_url"`
	LastAccessTime int64  `json:"last_access_time"`
	IsOwner        bool   `json:"is_owner"`
	IsAdmin        bool   `json:"is_admin"`
	JoinTime       int64  `json:"join_time"`
	// Permissions in chat if member is admin. `null` otherwise
	Permissions []ChatAdminPermission `json:"permissions"`
}

type ChatMembersList struct {
	// Participants in chat with time of last activity. Visible only for chat admins
	Members []ChatMember `json:"members"`
	// Pointer to the next data page
	Marker int64 `json:"marker"`
}

type ChatPatch struct {
	Icon  PhotoAttachmentRequestPayload `json:"icon,omitempty"`
	Title string                        `json:"title,omitempty"`
}

// ChatStatus : Chat status for current bot
type ChatStatus string

// List of ChatStatus
const (
	ACTIVE    ChatStatus = "active"
	REMOVED   ChatStatus = "removed"
	LEFT      ChatStatus = "left"
	CLOSED    ChatStatus = "closed"
	SUSPENDED ChatStatus = "suspended"
)

// Bot gets this type of update as soon as title has been changed in chat
type ChatTitleChangedUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
	// Chat identifier where event has occurred
	ChatId int64 `json:"chat_id"`
	// User who changed title
	UserId int64 `json:"user_id"`
	// New title
	Title string `json:"title"`
}

// ChatType : Type of chat. Dialog (one-on-one), chat or channel
type ChatType string

// List of ChatType
const (
	DIALOG  ChatType = "dialog"
	CHAT    ChatType = "chat"
	CHANNEL ChatType = "channel"
)

type ContactAttachment struct {
	Type    string                   `json:"type,omitempty"`
	Payload ContactAttachmentPayload `json:"payload"`
}

type ContactAttachmentPayload struct {
	// User info in VCF format
	VcfInfo string `json:"vcfInfo"`
	// User info
	TamInfo User `json:"tamInfo"`
}

// Request to attach contact card to message. MUST be the only attachment in message
type ContactAttachmentRequest struct {
	Type    string                          `json:"type,omitempty"`
	Payload ContactAttachmentRequestPayload `json:"payload"`
}

type ContactAttachmentRequestPayload struct {
	// Contact name
	Name string `json:"name"`
	// Contact identifier
	ContactId int64 `json:"contactId"`
	// Full information about contact in VCF format
	VcfInfo string `json:"vcfInfo"`
	// Contact phone in VCF format
	VcfPhone string `json:"vcfPhone"`
}

// Server returns this if there was an exception to your request
type Error struct {
	// Error
	Error string `json:"error,omitempty"`
	// Error code
	Code string `json:"code"`
	// Human-readable description
	Message string `json:"message"`
}

type FileAttachment struct {
	Type    string            `json:"type,omitempty"`
	Payload AttachmentPayload `json:"payload"`
}

// Request to attach file to message. MUST be the only attachment in message
type FileAttachmentRequest struct {
	Type    string           `json:"type,omitempty"`
	Payload UploadedFileInfo `json:"payload"`
}

// List of all WebHook subscriptions
type GetSubscriptionsResult struct {
	// Current suscriptions
	Subscriptions []Subscription `json:"subscriptions"`
}

// Generic schema describing image object
type Image struct {
	// URL of image
	Url string `json:"url"`
}

// Buttons in messages
type InlineKeyboardAttachment struct {
	Type string `json:"type,omitempty"`
	// Unique identifier of keyboard
	CallbackId string   `json:"callback_id"`
	Payload    Keyboard `json:"payload"`
}

// Request to attach keyboard to message
type InlineKeyboardAttachmentRequest struct {
	Type    string                                 `json:"type,omitempty"`
	Payload InlineKeyboardAttachmentRequestPayload `json:"payload"`
}

type InlineKeyboardAttachmentRequestPayload struct {
	// Two-dimensional array of buttons
	Buttons [][]Button `json:"buttons"`
}

// Intent : Intent of button
type Intent string

// List of Intent
const (
	POSITIVE Intent = "positive"
	NEGATIVE Intent = "negative"
	DEFAULT  Intent = "default"
)

// Keyboard is two-dimension array of buttons
type Keyboard struct {
	Buttons [][]Button `json:"buttons"`
}

// After pressing this type of button user follows the link it contains
type LinkButton struct {
	Type string `json:"type,omitempty"`
	// Visible text of button
	Text string `json:"text"`
	// Intent of button. Affects clients representation.
	Intent Intent `json:"intent"`
	Url    string `json:"url"`
}

type LinkedMessage struct {
	// Type of linked message
	Type MessageLinkType `json:"type"`
	// User sent this message
	Sender User `json:"sender"`
	// Chat where message was originally posted
	ChatId  int64       `json:"chat_id"`
	Message MessageBody `json:"message"`
}

type LocationAttachment struct {
	Type      string  `json:"type,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Request to attach keyboard to message
type LocationAttachmentRequest struct {
	Type      string  `json:"type,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Message in chat
type Message struct {
	// User that sent this message
	Sender User `json:"sender"`
	// Message recipient. Could be user or chat
	Recipient Recipient `json:"recipient"`
	// Unix-time when message was created
	Timestamp int64 `json:"timestamp"`
	// Forwarder or replied message
	Link LinkedMessage `json:"link,omitempty"`
	// Body of created message. Text + attachments. Could be null if message contains only forwarded message.
	Body MessageBody `json:"body"`
}

// Schema representing body of message
type MessageBody struct {
	// Unique identifier of message
	Mid string `json:"mid"`
	// Sequence identifier of message in chat
	Seq int64 `json:"seq"`
	// Message text
	Text string `json:"text"`
	// Message attachments. Could be one of `Attachment` type. See description of this schema
	Attachments []Attachment `json:"attachments"`
	// In case this message is repled to, it is the unique identifier of the replied message
	ReplyTo string `json:"reply_to,omitempty"`
}

// You will get this `update` as soon as user presses button
type MessageCallbackUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64    `json:"timestamp"`
	Callback  Callback `json:"callback"`
	// Original message containing inline keyboard. Can be `null` in case it had been deleted by the moment a bot got this update.
	Message Message `json:"message"`
}

// You will get this `update` as soon as message is created
type MessageCreatedUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
	// Newly created message
	Message Message `json:"message"`
}

// You will get this `update` as soon as message is edited
type MessageEditedUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
	// Edited message
	Message Message `json:"message"`
}

// MessageLinkType : Type of linked message
type MessageLinkType string

// List of MessageLinkType
const (
	FORWARD MessageLinkType = "forward"
	REPLY   MessageLinkType = "reply"
)

// Paginated list of messages
type MessageList struct {
	// List of messages
	Messages []Message `json:"messages"`
}

// You will get this `update` as soon as message is removed
type MessageRemovedUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
	// Identifier of removed message
	MessageId string `json:"message_id"`
}

// You will get this `update` as soon as message is restored
type MessageRestoredUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
	// Restored message identifier
	MessageId string `json:"message_id"`
}

type NewMessageBody struct {
	// Message text
	Text string `json:"text"`
	// Single message attachment.
	Attachment interface{} `json:"attachment,omitempty"`
	// Message attachments. See `AttachmentRequest` and it's inheritors for full information.
	Attachments []interface{} `json:"attachments,omitempty"`
	// If false, chat participants wouldn't be notified
	Notify bool `json:"notify,omitempty"`
}

// Image attachment
type PhotoAttachment struct {
	Type    string                 `json:"type,omitempty"`
	Payload PhotoAttachmentPayload `json:"payload"`
}

type PhotoAttachmentPayload struct {
	// Unique identifier of this image
	PhotoId int64  `json:"photo_id"`
	Token   string `json:"token"`
	// Image URL
	Url string `json:"url"`
}

type PhotoAttachmentRequest struct {
	Type    string                        `json:"type,omitempty"`
	Payload PhotoAttachmentRequestPayload `json:"payload"`
}

// Request to attach image. All fields are mutually exclusive.
type PhotoAttachmentRequestPayload struct {
	// If specified, given URL will be attached to message as image
	Url string `json:"url,omitempty"`
	// Token of any existing attachment
	Token string `json:"token,omitempty"`
	// Tokens were obtained after uploading images
	Photos map[string]PhotoToken `json:"photos,omitempty"`
}

type PhotoToken struct {
	// Encoded information of uploaded image
	Token string `json:"token"`
}

// This is information you will recieve as soon as an image uploaded
type PhotoTokens struct {
	Photos map[string]PhotoToken `json:"photos"`
}

// New message recepient. Could be user or chat
type Recipient struct {
	// Chat identifier
	ChatId int64 `json:"chat_id"`
	// Chat type
	ChatType ChatType `json:"chat_type"`
	// User identifier, if message was sent to user
	UserId int64 `json:"user_id"`
}

// After pressing this type of button client sends new message with attachment of curent user contact
type RequestContactButton struct {
	Type string `json:"type,omitempty"`
	// Visible text of button
	Text string `json:"text"`
	// Intent of button. Affects clients representation.
	Intent Intent `json:"intent"`
}

// After pressing this type of button client sends new message with attachment of current user geo location
type RequestGeoLocationButton struct {
	Type string `json:"type,omitempty"`
	// Visible text of button
	Text string `json:"text"`
	// Intent of button. Affects clients representation.
	Intent Intent `json:"intent"`
	// If *true*, sends location without asking user's confirmation
	Quick bool `json:"quick,omitempty"`
}

type SendMessageResult struct {
	Message Message `json:"message"`
}

// SenderAction : Different actions to send to chat members
type SenderAction string

// List of SenderAction
const (
	TYPING_ON     SenderAction = "typing_on"
	TYPING_OFF    SenderAction = "typing_off"
	SENDING_PHOTO SenderAction = "sending_photo"
	SENDING_VIDEO SenderAction = "sending_video"
	SENDING_AUDIO SenderAction = "sending_audio"
	MARK_SEEN     SenderAction = "mark_seen"
)

type ShareAttachment struct {
	Type    string            `json:"type,omitempty"`
	Payload AttachmentPayload `json:"payload"`
}

// Simple response to request
type SimpleQueryResult struct {
	// `true` if request was successful. `false` otherwise
	Success bool `json:"success"`
}

type StickerAttachment struct {
	Type    string            `json:"type,omitempty"`
	Payload AttachmentPayload `json:"payload"`
}

// Request to attach sticker. MUST be the only attachment request in message
type StickerAttachmentRequest struct {
	Type    string                          `json:"type,omitempty"`
	Payload StickerAttachmentRequestPayload `json:"payload"`
}

type StickerAttachmentRequestPayload struct {
	// Sticker code
	Code string `json:"code"`
}

// Schema to describe WebHook subscription
type Subscription struct {
	// WebHook URL
	Url string `json:"url"`
	// Unix-time when subscription was created
	Time int64 `json:"time"`
	// Update types bot subscribed for
	UpdateTypes []string `json:"update_types"`
	Version     string   `json:"version"`
}

// Request to set up WebHook subscription
type SubscriptionRequestBody struct {
	// URL of HTTP(S)-endpoint of your bot
	Url string `json:"url"`
	// List of update types your bot want to receive. See `Update` object for a complete list of types
	UpdateTypes []string `json:"update_types,omitempty"`
	// Version of API. Affects model representation
	Version string `json:"version,omitempty"`
}

type UpdateType string

const (
	UpdateTypeMessageCallback  UpdateType = "message_callback"
	UpdateTypeMessageCreated              = "message_created"
	UpdateTypeMessageRemoved              = "message_removed"
	UpdateTypeMessageEdited               = "message_edited"
	UpdateTypeMessageRestored             = "message_restored"
	UpdateTypeBotAdded                    = "bot_added"
	UpdateTypeBotRemoved                  = "bot_removed"
	UpdateTypeUserAdded                   = "user_added"
	UpdateTypeUserRemoved                 = "user_removed"
	UpdateTypeBotStarted                  = "bot_started"
	UpdateTypeChatTitleChanged            = "chat_title_changed"
)

// `Update` object repsesents different types of events that happened in chat. See its inheritors
type Update struct {
	UpdateType UpdateType `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
}

type UpdateMessageCallback struct {
	Update
	Callback Callback `json:"callback"`
	Message  Message  `json:"message"`
}

type UpdateMessageCreated struct {
	Update
	Message Message `json:"message"`
}

type UpdateMessageEdited struct {
	Update
	Message Message `json:"message"`
}

type UpdateMessageRestored struct {
	Update
	MessageID string `json:"message_id"`
}

type UpdateMessageRemoved struct {
	Update
	MessageID string `json:"message_id"`
}

type UpdateBotAdded struct {
	Update
	ChatID int64 `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type UpdateBotRemoved struct {
	Update
	ChatID int64 `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type UpdateBotStarted struct {
	Update
	ChatID int64 `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type UpdateChatTitleChanged struct {
	Update
	ChatID int64  `json:"chat_id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
}

type UpdateUserAdded struct {
	Update
	ChatID    int64 `json:"chat_id"`
	UserID    int64 `json:"user_id"`
	InviterID int64 `json:"inviter_id"`
}

type UpdateUserRemoved struct {
	Update
	ChatID  int64 `json:"chat_id"`
	UserID  int64 `json:"user_id"`
	AdminID int64 `json:"admin_id"`
}

// List of all updates in chats your bot participated in
type UpdateList struct {
	// Page of updates
	Updates []json.RawMessage `json:"updates"`
	// Pointer to the next data page
	Marker int64 `json:"marker"`
}

// Endpoint you should upload to your binaries
type UploadEndpoint struct {
	// URL to upload
	Url string `json:"url"`
}

// UploadType : Type of file uploading
type UploadType string

// List of UploadType
const (
	PHOTO UploadType = "photo"
	VIDEO UploadType = "video"
	AUDIO UploadType = "audio"
	FILE  UploadType = "file"
)

// This is information you will recieve as soon as a file is uploaded
type UploadedFileInfo struct {
	// Unique file identifier
	FileId int64 `json:"fileId"`
}

// This is information you will recieve as soon as audio/video is uploaded
type UploadedInfo struct {
	Id int64 `json:"id"`
}

type User struct {
	// Users identifier
	UserId int64 `json:"user_id"`
	// Users visible name
	Name string `json:"name"`
	// Unique public user name. Can be `null` if user is not accessible or it is not set
	Username string `json:"username"`
}

// You will receive this update when user has been added to chat where bot is administrator
type UserAddedToChatUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
	// Chat identifier where event has occured
	ChatId int64 `json:"chat_id"`
	// User added to chat
	UserId int64 `json:"user_id"`
	// User who added user to chat
	InviterId int64 `json:"inviter_id"`
}

type UserIdsList struct {
	UserIds []int64 `json:"user_ids"`
}

// You will receive this update when user has been removed from chat where bot is administrator
type UserRemovedFromChatUpdate struct {
	UpdateType string `json:"update_type,omitempty"`
	// Unix-time when event has occured
	Timestamp int64 `json:"timestamp"`
	// Chat identifier where event has occured
	ChatId int64 `json:"chat_id"`
	// User removed from chat
	UserId int64 `json:"user_id"`
	// Administrator who removed user from chat
	AdminId int64 `json:"admin_id"`
}

type UserWithPhoto struct {
	// Users identifier
	UserId int64 `json:"user_id"`
	// Users visible name
	Name string `json:"name"`
	// Unique public user name. Can be `null` if user is not accessible or it is not set
	Username string `json:"username"`
	// URL of avatar
	AvatarUrl string `json:"avatar_url"`
	// URL of avatar of a bigger size
	FullAvatarUrl string `json:"full_avatar_url"`
}

type VideoAttachment struct {
	Type    string            `json:"type,omitempty"`
	Payload AttachmentPayload `json:"payload"`
}

// Request to attach video to message
type VideoAttachmentRequest struct {
	Type    string       `json:"type,omitempty"`
	Payload UploadedInfo `json:"payload"`
}
