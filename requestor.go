package llmhelpergo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"net/http"
	"os"
)

type GeneralLlm struct {
	SystemPrompt    string
	Messages        *Messages
	HistoryMessages *[]Message
	URL             string

	Model string
}

type completionRequest struct {
	Model string `json:"model"`

	Temperature float64  `json:"temperature"`
	Messages    Messages `json:"messages"`

	Stream bool `json:"stream"`
}

func (l *GeneralLlm) Predict() (*string, error) {
	key := os.Getenv("OPENAI_KEY")
	if key == "" {

		return nil, errors.New("OPENAI_KEY is not set")
	}
	if l.Model == "" {
		l.Model = "gpt-3.5-turbo"
	}
	m := Messages{{Role: "system", Content: &l.SystemPrompt}}
	if l.Messages != nil {

		m = append(m, *l.Messages...)
	}
	body := completionRequest{
		Model: l.Model,

		Temperature: 0.7,
		Messages:    m,

		Stream: false,
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)

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
		fmt.Printf("error creating request: %s\n", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+key)
	var responseParsed CompletionResponse
	client := &http.Client{Timeout: 0}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error making request: %s\n", err)
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading request: %s\n", err)
		return nil, err
	}

	err = json.Unmarshal(data, &responseParsed)

	if err != nil {
		return nil, err
	}
	if responseParsed.ID == "" {
		return nil, errors.New("model is not available")
	}
	return responseParsed.Choices[0].Message.Content, nil
}

func (l *GeneralLlm) GetMessages() *Messages {
	return l.Messages
}

func (l *GeneralLlm) ReplaceMessages(m *Messages) {
	l.Messages = m

}
func (l *GeneralLlm) AddMessage(m Message) {
	if l.Messages == nil {
		l.Messages = &Messages{}
	}
	appendedMessages := append(*l.Messages, m)
	l.Messages = &appendedMessages
}
func (l *GeneralLlm) ClearMessages() {
	l.Messages = nil
}

func (l *GeneralLlm) ChangePrompt(prompt string) {
	l.SystemPrompt = prompt

}
func (l *GeneralLlm) ChangeModel(model string) {
	l.Model = model

}
func (l *GeneralLlm) GetHistoryMessages() *[]Message {
	if l.HistoryMessages == nil {
		l.HistoryMessages = &[]Message{}
	}

	return l.HistoryMessages
}
func (l *GeneralLlm) AddHistoryMessage(m Message) {
	if l.HistoryMessages == nil {
		l.HistoryMessages = &[]Message{}
	}
	appendedMessages := append(*l.HistoryMessages, m)
	l.HistoryMessages = &appendedMessages
}
func (l *GeneralLlm) ClearHistoryMessages() {
	l.HistoryMessages = nil
}
