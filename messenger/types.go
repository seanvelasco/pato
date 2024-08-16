package messenger

type User struct {
	ID string `json:"id"`
}

type Message struct {
	Text    string `json:"text"`
	ReplyTo struct {
		MID string `json:"mid"`
	} `json:"reply_to,omitempty"`
}

type SendMessageRequest struct {
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	MessagingType string  `json:"messaging_type"`
	Message       Message `json:"message"`
}

type SendMessageResponse struct {
	RecipientID string `json:"recipient_id"`
	MessageID   string `json:"message_id"`
}
