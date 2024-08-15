package telegram

type User struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username,omitempty"`
}

type Chat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	IsForum   bool   `json:"is_forum,omitempty"`
}

type Message struct {
	MessageID       int  `json:"message_id"`
	MessageThreadID int  `json:"message_thread_id,omitempty"`
	From            User `json:"from,omitempty"`
	SenderChat      Chat `json:"sender_chat,omitempty"`
	//SenderBoostContent string `json:"sender_boost_content,omitempty"`
	//SenderBusinessBot  User   `json:"sender_business_bot,omitempty"`
	Date int `json:"date"`
	//BusinessConnectionID string `json:"business_connection_id,omitempty"`
	Chat Chat   `json:"chat"`
	Text string `json:"text,omitempty"`
}

type Update struct {
	UpdatedID         int      `json:"updated_id"`
	Message           *Message `json:"message,omitempty"` // why is this a pointer? any benefits?
	EditedMessage     *Message `json:"edited_message,omitempty"`
	ChannelPost       *Message `json:"channel_post,omitempty"`
	EditedChannelPost *Message `json:"edited_channel,omitempty"`
}
