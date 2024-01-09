package main

import (
	"fmt"
	"os"

	"github.com/uussoop/llmmodels-go/llmmodels"

	"github.com/uussoop/llmmodels-go/llmmodels/agents/sample"
	"github.com/uussoop/llmmodels-go/llmmodels/llm/general"
)

func main() {

	os.Setenv("OPENAI_KEY", "")
	message := "hello world"
	llm := &general.GeneralLlm{
		SystemPrompt: "you are an ai",
		Messages:     nil,
		URL:          "https://api.openai.com/v1/chat/completions",
		Model:        "gpt-4",
	}
	chain := llmmodels.Chain(&message, llm, 450)
	chain.Use(sample.SampleAgent)
	chain.Use(sample.SampleAgent2)
	chain.Use(sample.SampleAgent3)
	res, err := chain.Predict()

	fmt.Println(llm.SystemPrompt)
	if err != nil {
		panic(err)
	}
	fmt.Println(*res)

}
