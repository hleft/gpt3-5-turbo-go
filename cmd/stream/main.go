package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hleft/gpt3-5-turbo-go/gpt35"
)

func main() {
	gpt35.ApiKey = ""
	if gpt35.ApiKey == "" {
		gpt35.ApiKey = os.Getenv("openai_apikey")
	}

	resp, err := gpt35.GetHttpResp(context.Background(), &gpt35.RequestData{
		Messages: []*gpt35.Message{
			gpt35.NewSystemMsg("You are a helpful assistant."),
			gpt35.NewUserMsg("Who won the world series in 2020?"),
			gpt35.NewAssistantMsg("The Los Angeles Dodgers won the World Series in 2020."),
			gpt35.NewUserMsg("Where was it played?"),
		},
		Stream: true,
	})

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	messages, err := gpt35.ListenHttpResp(resp)
	if err != nil {
		fmt.Printf("Error listening to HTTP response: %v\n", err)
		return
	}

	for message := range messages {
		fmt.Print(message.Choices[0].Delta.Content)
	}
}
