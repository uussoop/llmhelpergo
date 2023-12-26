package utils

type Message struct {
	Role    string  `json:"role"`
	Content *string `json:"content"`
}

type Messages []Message
