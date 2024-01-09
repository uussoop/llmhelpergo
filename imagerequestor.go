package llmhelpergo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	log "github.com/sirupsen/logrus"

	"net/http"
	"os"
)

// image llm the image must be base64 string
type GeneralImageLlm struct {
	Prompt string
	Image  string
	URL    string
	Model  string
}

// Request struct
type imageCompletionRequest struct {
	Model string `json:"model"`

	Messages ImageMessages `json:"messages"`
}

func (l GeneralImageLlm) Predict() (*string, error) {
	key := os.Getenv("OPENAI_KEY")
	if key == "" {

		return nil, errors.New("OPENAI_KEY is not set")
	}

	body := imageCompletionRequest{
		Model: l.Model,

		Messages: ImageMessages{
			ImageMessage{
				Role: "user",
				Content: []Content{
					{Type: "text", Text: &l.Prompt},
					{Type: "image_url", ImageURL: &ImageURL{URL: l.Image}},
				},
			},
		},
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		log.Info(err)

		return nil, err
	}

	// Create request body reader
	bodyReader := bytes.NewReader(jsonBytes)
	req, err := http.NewRequest(
		http.MethodPost,
		l.URL,
		bodyReader,
	)
	if err != nil {
		log.Error("error creating request: %s\n", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+key)
	var responseParsed CompletionResponse
	client := &http.Client{Timeout: 0}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("error making request: %s\n", err)
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("error reading request: %s\n", err)
		return nil, err
	}

	err = json.Unmarshal(data, &responseParsed)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	if responseParsed.ID == "" {
		return nil, errors.New("model is not available")
	}
	return responseParsed.Choices[0].Message.Content, nil
}

func (l GeneralImageLlm) GetMessages() *Messages {
	return nil
}

func (l GeneralImageLlm) ReplaceMessages(m *Messages) {

}
func (l GeneralImageLlm) AddMessage(m Message) {

}

func (l GeneralImageLlm) ReplacePrompt(prompt string) {

}
func (l GeneralImageLlm) GetHistoryMessages() *[]Message {

	return nil
}
func (l GeneralImageLlm) AddHistoryMessage(m Message) {

}
func (l GeneralImageLlm) ClearHistoryMessages() {

}
