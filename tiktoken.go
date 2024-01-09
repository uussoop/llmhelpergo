package llmhelpergo

import (
	"github.com/pkoukk/tiktoken-go"
)

func CountTokens(text *string) (int, error) {

	encoding := "gpt-3.5-turbo"
	tke, err := tiktoken.EncodingForModel(encoding)
	if err != nil {
		return 0, err
	}

	// encode
	token := tke.Encode(*text, nil, nil)

	return len(token), nil
}
