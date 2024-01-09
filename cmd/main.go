package main

import (
	"fmt"
	"os"

	"github.com/uussoop/llmhelpergo"
)

func main() {

	os.Setenv("OPENAI_KEY", "")
	message := "hello world"
	llm := &llmhelpergo.GeneralLlm{
		SystemPrompt: "you are an ai",
		Messages:     nil,
		URL:          "https://api.openai.com/v1/chat/completions",
		Model:        "gpt-4",
	}
	chain := llmhelpergo.Chain(&message, llm, 450)
	chain.Use(llmhelpergo.SampleAgent)
	chain.Use(llmhelpergo.SampleAgent2)
	chain.Use(llmhelpergo.SampleAgent3)
	res, err := chain.Predict()

	fmt.Println(llm.SystemPrompt)
	if err != nil {
		panic(err)
	}
	fmt.Println(*res)

}
