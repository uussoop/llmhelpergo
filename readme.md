# helper to make chains for llm interactions



```go
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
		SystemPrompt: "you are an ai assistant",
		Messages:     nil,
		URL:          "https://api.openai.com/v1/chat/completions",
		Model:        "gpt-3.5-turbo",
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
```

In this example, a Go program imports the `llmhelpergo` package to interact with an AI assistant model from Anthropic using their Pile API. It sets an OpenAI API key, initializes a `GeneralLlm` struct with some model parameters, creates a processing chain, adds some sample middleware agents, and makes a prediction request. The result is printed out.





## remember to set the following in your env:

    - OPENAI_KEY
> remember to input history insert function to nil if you dont have any:)

> remember to input bufferedmemmory nil if you dont have any:)

> hint: start from chat directory using Chat function:)