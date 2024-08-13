package main

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Prompt struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type SendMessageRequestBody struct {
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

type SSE struct {
	Created int    `json:"created"`
	Id      string `json:"id"`
	Action  string `json:"action"`
	Model   string `json:"model"`
}

type StartSSE struct {
	SSE
	Role string `json:"role"`
}

type MessageSSE struct {
	SSE
	Message string `json:"message,omitempty"`
}

type EndSSE struct {
	SSE
}

type SSEType interface {
	StartSSE | MessageSSE | EndSSE
}
