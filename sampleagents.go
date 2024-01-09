package llmhelpergo

import (
	"fmt"
)

func SampleAgent(aiResponse *string, llm *Llm) (*string, error) {
	m := (*llm).GetHistoryMessages()
	fmt.Println("agent1+++++++++++++++++++++++")
	for i, v := range *m {
		fmt.Println(i, v.Role, *v.Content)
	}

	return aiResponse, nil

}
func SampleAgent2(aiResponse *string, llm *Llm) (*string, error) {
	m := (*llm).GetHistoryMessages()
	(*llm).ChangePrompt("some wierd prompt")
	fmt.Println("agent2+++++++++++++++++++++")
	for i, v := range *m {
		fmt.Println(i, v.Role, *v.Content)
	}
	res := "agent2"
	return &res, nil

}
func SampleAgent3(aiResponse *string, llm *Llm) (*string, error) {
	m := (*llm).GetHistoryMessages()
	fmt.Println("agent3+++++++++++++++++")
	for i, v := range *m {
		fmt.Println(i, v.Role, *v.Content)
	}

	return aiResponse, nil

}

func RequestSummarization(aiResponse *string, llm *Llm) (*string, error) {
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
