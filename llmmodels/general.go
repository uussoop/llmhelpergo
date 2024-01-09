package llmmodels

import (
	"github.com/uussoop/llmmodels-go/llmmodels/llm/general"
	"github.com/uussoop/llmmodels-go/llmmodels/prompts"
	"github.com/uussoop/llmmodels-go/llmmodels/utils"

	log "github.com/sirupsen/logrus"
)

var HistoryInsert func(*[]utils.Message) (err error)
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

// returns a pointer to a chain so you can use the llm and input and response with all messages elsewhere too
func Chain(Input *string, Llm Llm, InputTokenLimit int) *chain {
	llmImpl = Llm
	if InputTokenLimit == 0 {
		InputTokenLimit = 450
	}

	return &chain{Input: Input, Llm: &llmImpl, InputTokenLimit: InputTokenLimit}
}
func (c *chain) Init() {

}
func (c *chain) Use(f func(Response *string, llm *Llm) (*string, error)) {
	c.Pipes = append(c.Pipes, f)

}

// this will initiate the chain with its middlewares and in the end all middlewares get clensed
func (c *chain) Predict() (*string, error) {

	for _, f := range c.Pipes {

		(*c.Llm).AddHistoryMessage(utils.Message{Role: "user", Content: c.Input})
		r, err := f(c.Input, c.Llm)
		if err != nil {

			return c.Response, err
		}
		if r != nil {
			messageLength, _ := utils.CountTokens(r)

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
		(*c.Llm).AddHistoryMessage(utils.Message{Role: "assistant", Content: c.Response})
	}

	c.Pipes = nil

	go c.Save()

	return c.Response, nil
}
func (c *chain) Save() {

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
	GetMessages() *utils.Messages
	ReplaceMessages(*utils.Messages)
	AddMessage(utils.Message)
	ChangePrompt(prompt string)
	ChangeModel(model string)
	AddHistoryMessage(m utils.Message)
	GetHistoryMessages() *[]utils.Message
	ClearHistoryMessages()
	ClearMessages()
}

func requestSummarization(aiResponse *string, llm *Llm) (*string, error) {
	summarizerLlm := general.GeneralLlm{
		SystemPrompt: prompts.SummarySystemPrompt,
		Messages:     &utils.Messages{{Role: "user", Content: aiResponse}},
		Model:        "gpt-3.5-turbo",
	}
	aiResponse, err := summarizerLlm.Predict()
	if err != nil {
		return nil, err
	}
	return aiResponse, nil

}
