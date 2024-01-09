package utils

type Message struct {
	Role    string  `json:"role"`
	Content *string `json:"content"`
}

type Messages []Message

type ImageMessage struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Content struct {
	Type     string    `json:"type"`
	Text     *string   `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

type ImageURL struct {
	URL string `json:"url"`
}
type ImageMessages []ImageMessage

type ChoicesResponse struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}
type CompletionResponse struct {
	ID      string            `json:"id"`
	Object  string            `json:"object"`
	Created int               `json:"created"`
	Model   string            `json:"model"`
	Choices []ChoicesResponse `json:"choices"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
