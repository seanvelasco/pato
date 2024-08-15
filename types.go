package main

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
	ID string `json:"id"`
}

type Message struct {
	MID  string `json:"mid,omitempty"`
	Text string `json:"text"`
}

type WebhookEvent struct {
	Object string `json:"object"`
	Entry  []struct {
		Id        string `json:"id"`
		Time      int    `json:"time"`
		Messaging []struct {
			Sender    User    `json:"sender"`
			Recipient User    `json:"recipient"`
			Timestamp int     `json:"timestamp"`
			Message   Message `json:"message"`
		} `json:"messaging"`
	} `json:"entry"`
}
