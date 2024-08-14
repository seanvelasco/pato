package main

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Prompt struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type User struct {
	Id string `json:"id"`
}

type EventNotificationEntryMessagingMessage struct {
	Message string `json:"message"`
}

type EventNotificationEntryMessaging struct {
	Sender    User                                     `json:"sender"`
	Recipient User                                     `json:"recipient"`
	Messaging []EventNotificationEntryMessagingMessage `json:"messaging,omitempty"`
}

type EventNotificationEntry struct {
	Id        string                            `json:"id"`
	Time      int                               `json:"time"`
	Messaging []EventNotificationEntryMessaging `json:"messaging"`
}

type EventNotification struct {
	Object string                            `json:"object"`
	Entry  []EventNotificationEntryMessaging `json:"entry"`
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

type WebhookEvent struct {
	Field string `json:"field"`
	Value struct {
		Sender    User   `json:"sender"`
		Recipient User   `json:"recipient"`
		Timestamp string `json:"timestamp"`
		Message   struct {
			MID      string `json:"mid"`
			Text     string `json:"text"`
			Commands []struct {
				Name string `json:"name"`
			} `json:"commands"`
		} `json:"message"`
	} `json:"value"`
}
