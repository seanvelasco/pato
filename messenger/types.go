package messenger

type User struct {
	ID string `json:"id"`
}

type SendMessageMessage struct { // I, know :(
	Text string `json:"text"`
}

type SendMessageRequest struct {
	Recipient     User               `json:"recipient"`
	MessagingType string             `json:"messaging_type"`
	Message       SendMessageMessage `json:"message"`
}

type SendMessageResponse struct {
	RecipientID string `json:"recipient_id"`
	MessageID   string `json:"message_id"`
}

type ReplyTo struct {
	MID string `json:"mid"`
}

type Message struct {
	MID     string  `json:"mid,omitempty"`
	Text    string  `json:"text"`
	ReplyTo ReplyTo `json:"reply_to,omitempty"`
}

type Messaging struct {
	Sender    User    `json:"sender"`
	Recipient User    `json:"recipient"`
	Timestamp int     `json:"timestamp"`
	Message   Message `json:"message"`
}

type MessageEvent struct {
	Object string `json:"object"`
	Entry  []struct {
		Id        string      `json:"id"`
		Time      int         `json:"time"`
		Messaging []Messaging `json:"messaging"`
	} `json:"entry"`
}
