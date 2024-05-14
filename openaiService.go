package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ayush6624/go-chatgpt"
	"github.com/sashabaranov/go-openai"
	"github.com/tidwall/gjson"
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

func AskGpt(message string) string {
	openai_api_key := os.Getenv("OPENAI_API_KEY")
	client, err := chatgpt.NewClient(openai_api_key)
	if err != nil {
		fmt.Println("Error creating ChatGPT client:", err)
		return ""
	}

	ctx := context.Background()
	res, err := client.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT35Turbo,
		Messages: []chatgpt.ChatMessage{
			{
				Role:    chatgpt.ChatGPTModelRoleSystem,
				Content: message,
			},
		},
	})

	if err != nil {
		fmt.Println("Error sending message:", err)
		return ""
	}

	ans, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling response:", err)
		return ""
	}

	respuesta := gjson.Get(string(ans), "choices")
	//fmt.Println(respuesta)
	var responses []Response
	errorr := json.Unmarshal([]byte(respuesta.String()), &responses)
	if errorr != nil {
		panic(errorr)
	}

	// Access the first response and extract the "content" value
	firstResponse := responses[0]
	content := firstResponse.Message.Content
	return content
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
