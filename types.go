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

type User struct {
	Id string `json:"id"`
}

type WebhookEvent struct {
	Object string `json:"object"`
	Entry  []struct {
		Id        string `json:"id"`
		Time      int    `json:"time"`
		Messaging []struct {
			Sender    User `json:"sender"`
			Recipient User `json:"recipient"`
			Timestamp int  `json:"timestamp"`
			Message   struct {
				MID  string `json:"mid"`
				Text string `json:"text"`
			} `json:"message"`
		} `json:"messaging"`
	} `json:"entry"`
}
