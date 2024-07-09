package myOpenAi

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

var Req openai.ChatCompletionRequest

func AskGpt(message string) string {

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	Req.Messages = append(Req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})

	resp, err := client.CreateChatCompletion(context.Background(), Req)
	if err != nil {
		fmt.Println("There was an error" + err.Error())
		return ""
	}

	return resp.Choices[0].Message.Content
}

func Speech_to_text(filePathName string) string {
	openai_api_key := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(openai_api_key)
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: filePathName,
		Language: "es",
	}

	responseFormat := openai.ChatCompletionResponseFormat{
		Type: openai.ChatCompletionResponseFormatTypeJSONObject,
	}

	Req.ResponseFormat = &responseFormat

	resp, err := client.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return ""
	}

	return string(resp.Text)
}
