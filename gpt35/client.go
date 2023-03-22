package gpt35

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

var ApiKey = ""

const (
	DefaultUrl          = "https://api.openai.com/v1/chat/completions"
	ModelGpt35Turbo     = "gpt-3.5-turbo"
	MaxTokensGpt35Turbo = 4096
	RoleUser            = "user"
	RoleAssistant       = "assistant"
	RoleSystem          = "system"
)

func GetHttpResp(rd *RequestData) (*http.Response, error) {
	rd.Model = ModelGpt35Turbo
	req, err := getReq(rd, DefaultUrl)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return httpResp, nil
}

func ListenHttpResp(resp *http.Response) (chan SSEResponse, error) {
	reader := bufio.NewReader(resp.Body)
	messages := make(chan SSEResponse)

	go func() {
		defer close(messages)

		for {
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				panic(err)
			}

			if err == io.EOF {
				break
			}

			if strings.HasPrefix(line, "data: ") {
				if strings.TrimSpace(line[6:]) == "[DONE]" {
					break
				}
				var message SSEResponse
				// delete data:
				err := json.Unmarshal([]byte(line[6:]), &message)
				if err != nil {
					panic(err)
				}

				for range message.Choices {
					messages <- message
				}
			}
		}
	}()

	return messages, nil
}

func GetOpenAiResp(rd *RequestData) (*Response, error) {
	httpResp, err := GetHttpResp(rd)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	var resp Response
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func getReq(r *RequestData, url string) (*http.Request, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ApiKey)

	return req, err
}