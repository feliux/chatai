package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/subosito/gotenv"
)

func init() {
	if err := gotenv.Load(); err != nil {
		log.Fatalf("error ocurred reading .env file %s", err)
	}
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	c := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a helpful assistant",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Hello!",
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()
	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}
		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}
		fmt.Printf(response.Choices[0].Delta.Content)
	}
}
