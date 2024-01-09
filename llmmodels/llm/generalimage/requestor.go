package generalimage

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/uussoop/llmmodels-go/llmmodels/utils"

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

	Messages utils.ImageMessages `json:"messages"`
}

func (l GeneralImageLlm) Predict() (*string, error) {
	key := os.Getenv("OPENAI_KEY")
	if key == "" {

		return nil, errors.New("OPENAI_KEY is not set")
	}

	body := imageCompletionRequest{
		Model: l.Model,

		Messages: utils.ImageMessages{
			utils.ImageMessage{
				Role: "user",
				Content: []utils.Content{
					{Type: "text", Text: &l.Prompt},
					{Type: "image_url", ImageURL: &utils.ImageURL{URL: l.Image}},
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
	var responseParsed utils.CompletionResponse
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

func (l GeneralImageLlm) GetMessages() *utils.Messages {
	return nil
}

func (l GeneralImageLlm) ReplaceMessages(m *utils.Messages) {

}
func (l GeneralImageLlm) AddMessage(m utils.Message) {

}

func (l GeneralImageLlm) ReplacePrompt(prompt string) {

}
func (l GeneralImageLlm) GetHistoryMessages() *[]utils.Message {

	return nil
}
func (l GeneralImageLlm) AddHistoryMessage(m utils.Message) {

}
func (l GeneralImageLlm) ClearHistoryMessages() {

}
