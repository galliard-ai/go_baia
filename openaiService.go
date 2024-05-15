package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

var req openai.ChatCompletionRequest

func AskGpt(message string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("There was an error")
		return ""
	}

	return resp.Choices[0].Message.Content
}

func speech_to_text(filePathName string) string {
	openai_api_key := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(openai_api_key)
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: filePathName,
		Language: "es",
	}

	resp, err := client.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return ""
	}

	return string(resp.Text)
}
