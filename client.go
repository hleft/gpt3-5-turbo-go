package gpt35

import (
	"bytes"
	"encoding/json"
	"net/http"
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

func SendReq(rd *RequestData) (*Response, error) {
	req, err := getReq(rd, DefaultUrl)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(req)
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
