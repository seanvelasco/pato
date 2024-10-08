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

type SuggestionResults []struct {
	Phrase string `json:"phrase"`
}

type SearchAssistArticle struct {
	Site string `json:"site"`
	Link string `json:"link"`
	Text string `json:"text"`
}

type SearchAssistSource struct {
	Article SearchAssistArticle `json:"article"`
	Section struct{}            `json:"section"`
}

type SearchAssistResults struct {
	Timestamp string               `json:"timestamp"`
	Action    string               `json:"action"`
	Answer    string               `json:"answer"`
	Sources   []SearchAssistSource `json:"sources"`
}
