package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hleft/gpt3-5-turbo-go/gpt35"
)

func main() {
	gpt35.ApiKey = ""
	if gpt35.ApiKey == "" {
		gpt35.ApiKey = os.Getenv("openai_apikey")
	}

	resp, err := gpt35.GetOpenAiResp(context.Background(), &gpt35.RequestData{
		Messages: []*gpt35.Message{
			gpt35.NewSystemMsg("You are a helpful assistant."),
			gpt35.NewUserMsg("Who won the world series in 2020?"),
			gpt35.NewAssistantMsg("The Los Angeles Dodgers won the World Series in 2020."),
			gpt35.NewUserMsg("Where was it played?"),
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%s\n\n", resp.Choices[0].Message.Content)

	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
