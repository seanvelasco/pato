package main

type SendMessageRequest struct {
	MessagingType string `json:"messaging_type"`
	ThreadControl struct {
		Payload string `json:"payload"`
	} `json:"thread_control"`
	Recipient struct {
		Id string `json:"id"`
	} `json:"recipient"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}

type SendMessageResponse struct {
	RecipientId string `json:"recipient_id"`
	MessageId   string `json:"message_id"`
}
