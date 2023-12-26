package summarizer

import (
	"github.com/uussoop/llmmodels-go/llmmodels"
	"github.com/uussoop/llmmodels-go/llmmodels/llm/general"
	"github.com/uussoop/llmmodels-go/llmmodels/prompts"
	"github.com/uussoop/llmmodels-go/llmmodels/utils"
)

func RequestSummarization(aiResponse *string, llm *llmmodels.Llm) (*string, error) {
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
