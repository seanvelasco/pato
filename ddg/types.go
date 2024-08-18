package ddg

// Question: can I parse string URLs into Go URLs?

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Prompt struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type TextResults struct {
	Results []TextResult `json:"results"`
}

type TextResult struct {
	Body  string `json:"a"`
	Title string `json:"t"`
	URL   string `json:"u"`
}

type ImageResults struct {
	Results []ImageResult `json:"results"`
}

type ImageResult struct {
	Title     string `json:"title"`
	Image     string `json:"image"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	Source    string `json:"source"`
}
