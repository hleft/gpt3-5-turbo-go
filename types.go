package gpt35

type RequestData struct {
	Model            string      `json:"model"`
	Messages         []*Message  `json:"messages"`
	Temperature      float64     `json:"temperature,omitempty"`
	TopP             float64     `json:"top_p,omitempty"`
	N                int         `json:"n,omitempty"`
	Stream           bool        `json:"stream,omitempty"`
	Stop             interface{} `json:"stop,omitempty"`
	MaxTokens        int         `json:"max_tokens,omitempty"`
	PresencePenalty  float64     `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64     `json:"frequency_penalty,omitempty"`
	LogitBias        interface{} `json:"logit_bias,omitempty"`
	User             string      `json:"user,omitempty"`
}

type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []*struct {
		Index        int      `json:"index"`
		Message      *Message `json:"message"`
		FinishReason string   `json:"finish_reason"`
	} `json:"choices"`
	Usage *struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Param   string `json:"param"`
		Code    string `json:"code"`
	} `json:"error,omitempty"`
}

type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content"`
}

func NewSystemMsg(content string) *Message {
	return &Message{Role: RoleSystem, Content: content}
}
func NewUserMsg(content string) *Message {
	return &Message{Role: RoleUser, Content: content}
}
func NewAssistantMsg(content string) *Message {
	return &Message{Role: RoleAssistant, Content: content}
}
