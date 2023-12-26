package sample

import (
	"fmt"

	"github.com/uussoop/llmmodels-go/llmmodels"
)

func SampleAgent(aiResponse *string, llm *llmmodels.Llm) (*string, error) {
	m := (*llm).GetHistoryMessages()
	fmt.Println("agent1+++++++++++++++++++++++")
	for i, v := range *m {
		fmt.Println(i, v.Role, *v.Content)
	}

	return aiResponse, nil

}
func SampleAgent2(aiResponse *string, llm *llmmodels.Llm) (*string, error) {
	m := (*llm).GetHistoryMessages()
	(*llm).ChangePrompt("some wierd prompt")
	fmt.Println("agent2+++++++++++++++++++++")
	for i, v := range *m {
		fmt.Println(i, v.Role, *v.Content)
	}
	res := "agent2"
	return &res, nil

}
func SampleAgent3(aiResponse *string, llm *llmmodels.Llm) (*string, error) {
	m := (*llm).GetHistoryMessages()
	fmt.Println("agent3+++++++++++++++++")
	for i, v := range *m {
		fmt.Println(i, v.Role, *v.Content)
	}

	return aiResponse, nil

}
