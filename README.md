## GoLang HTTP Client for ChatGPT (GPT-3.5-turbo)

A client for an official API of chat completions (known as ChatGPT) based on `gpt-3.5-turbo` model.

Here's a guide by Open AI: https://platform.openai.com/docs/guides/chat.

Reference: https://platform.openai.com/docs/api-reference/introduction

### Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	gpt35 "github.com/hleft/gpt3-5-turbo-go"
)

func main() {
	gpt35.ApiKey = ""
	if gpt35.ApiKey == "" {
		gpt35.ApiKey = os.Getenv("openai_apikey")
	}

	resp, err := gpt35.SendReq(&gpt35.RequestData{
		Model: gpt35.ModelGpt35Turbo,
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
```

output:

```
The 2020 World Series was played at Globe Life Field in Arlington, Texas.

{
	"id": "chatcmpl-6qbyGd3fmyAXmjN9AEgByNnniEvY5",
	"object": "chat.completion",
	"created": 1677997036,
	"choices": [
		{
			"index": 0,
			"message": {
				"role": "assistant",
				"content": "The 2020 World Series was played at Globe Life Field in Arlington, Texas."
			},
			"finish_reason": "stop"
		}
	],
	"usage": {
		"prompt_tokens": 56,
		"completion_tokens": 19,
		"total_tokens": 75
	}
}
```

### Thank https://github.com/AlmazDelDiablo/gpt3-5-turbo-go
