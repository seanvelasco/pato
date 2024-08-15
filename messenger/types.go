package messenger

type User struct {
	ID string `json:"id"`
}

type SendMessageRequest struct {
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	MessagingType string `json:"messaging_type"`
	Message       struct {
		Text string `json:"text"`
	} `json:"message"`
}

type SendMessageResponse struct {
	RecipientID string `json:"recipient_id"`
	MessageID   string `json:"message_id"`
}
