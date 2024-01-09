package llmhelpergo

import (
	log "github.com/sirupsen/logrus"
)

var HistoryInsert func(*[]Message) (err error)
var QaInsert func(question, answer *string) error
var QaRead func(question *string) (*string, *string, error)
var llmImpl Llm
var InputTokenLimit = 450

type chain struct {
	Input           *string
	Llm             *Llm
	Response        *string
	InputTokenLimit int
	Pipes           []func(Response *string, llm *Llm) (*string, error)
}
type ChainType chain

// returns a pointer to a chain so you can use the llm and input and response with all messages elsewhere too
func Chain(Input *string, Llm Llm, InputTokenLimit int) *ChainType {
	llmImpl = Llm
	if InputTokenLimit == 0 {
		InputTokenLimit = 450
	}
	var ct ChainType
	ct.Input = Input
	ct.Llm = &Llm
	ct.InputTokenLimit = InputTokenLimit
	return &ct
}

// func (c *ChainType) Init() {

// }
func (c *ChainType) Use(f func(Response *string, llm *Llm) (*string, error)) {
	c.Pipes = append(c.Pipes, f)

}

// this will initiate the chain with its middlewares and in the end all middlewares get clensed
func (c *ChainType) Predict() (*string, error) {

	for _, f := range c.Pipes {

		(*c.Llm).AddHistoryMessage(Message{Role: "user", Content: c.Input})
		r, err := f(c.Input, c.Llm)
		if err != nil {

			return c.Response, err
		}
		if r != nil {
			messageLength, _ := CountTokens(r)

			if messageLength > c.InputTokenLimit {
				summarized, err := requestSummarization(r, nil)
				if err != nil {
					c.Response = nil
					go c.Save()
					return nil, err
				}
				r = summarized
			}
		}
		c.Input = r
		c.Response = r
		(*c.Llm).AddHistoryMessage(Message{Role: "assistant", Content: c.Response})
	}

	c.Pipes = nil

	go c.Save()

	return c.Response, nil
}
func (c *ChainType) Save() {

	if HistoryInsert == nil {
		log.Debug("error saving history function not specified.")
		return

	}
	err := HistoryInsert((*c.Llm).GetHistoryMessages())
	if err != nil {
		log.Debug("error saving history: ", err)
	}
	(*c.Llm).ClearHistoryMessages()
}

type Llm interface {
	Predict() (*string, error)
	GetMessages() *Messages
	ReplaceMessages(*Messages)
	AddMessage(Message)
	ChangePrompt(prompt string)
	ChangeModel(model string)
	AddHistoryMessage(m Message)
	GetHistoryMessages() *[]Message
	ClearHistoryMessages()
	ClearMessages()
}

func requestSummarization(aiResponse *string, llm *Llm) (*string, error) {
	summarizerLlm := GeneralLlm{
		SystemPrompt: SummarySystemPrompt,
		Messages:     &Messages{{Role: "user", Content: aiResponse}},
		Model:        "gpt-3.5-turbo",
	}
	aiResponse, err := summarizerLlm.Predict()
	if err != nil {
		return nil, err
	}
	return aiResponse, nil

}
