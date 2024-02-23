package openai_dbchaos

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	APIkey string
	Model  string
}

func (o *OpenAI) Prompt(query string) (string, error) {
	if o.Model == "" {
		o.Model = openai.GPT3Dot5Turbo
	}

	client := openai.NewClient(o.APIkey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
