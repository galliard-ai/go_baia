package main

import (
	"fmt"
	// "os"
	// "context"

	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main() {
	godotenv.Load()

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Eres un asistente virtual, todo lo que te pregunten, tienes que contestar en espanol mexicano",
			},
		},
	}

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "How much is 3 times 5?",
	})

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("There was an error")
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)

	//listenMsgs()

}
