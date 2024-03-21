package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/maths-lover/chatgpt/client"
	"github.com/maths-lover/chatgpt/constants"
	"github.com/maths-lover/chatgpt/utils"
)

type ChatCompletionRequest struct {
	// (Required)
	// ID of the model to use.
	Model constants.ChatGPTModel `json:"model"`

	// Required
	// The messages to generate chat completions for
	Messages []ChatMessage `json:"messages"`

	// (Optional - default: 1)
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	// We generally recommend altering this or top_p but not both.
	Temperature float64 `json:"temperature,omitempty"`

	// (Optional - default: 1)
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// We generally recommend altering this or temperature but not both.
	Top_P float64 `json:"top_p,omitempty"`

	// (Optional - default: 1)
	// How many chat completion choices to generate for each input message.
	N int `json:"n,omitempty"`

	// (Optional - default: infinite)
	// The maximum number of tokens allowed for the generated answer. By default,
	// the number of tokens the model can return will be (4096 - prompt tokens).
	MaxTokens int `json:"max_tokens,omitempty"`

	// (Optional - default: 0)
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far,
	// increasing the model's likelihood to talk about new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// (Optional - default: 0)
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// (Optional)
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
	User string `json:"user,omitempty"`
}

type ChatMessage struct {
	Role    constants.ChatGPTModelRole `json:"role"`
	Content string                     `json:"content"`
}

type ChatResponse struct {
	ID        string               `json:"id"`
	Object    string               `json:"object"`
	CreatedAt int64                `json:"created_at"`
	Choices   []ChatResponseChoice `json:"choices"`
	Usage     ChatResponseUsage    `json:"usage"`
}

type ChatResponseChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type ChatResponseUsage struct {
	Prompt_Tokens     int `json:"prompt_tokens"`
	Completion_Tokens int `json:"completion_tokens"`
	Total_Tokens      int `json:"total_tokens"`
}

func SimpleSend(ctx context.Context, c *client.Client, message string) (*ChatResponse, error) {
	req := &ChatCompletionRequest{
		Model: constants.GPT35Turbo,
		Messages: []ChatMessage{
			{
				Role:    constants.ChatGPTModelRoleUser,
				Content: message,
			},
		},
	}

	return Send(ctx, c, req)
}

func Send(ctx context.Context, c *client.Client, req *ChatCompletionRequest) (*ChatResponse, error) {
	if err := validate(req); err != nil {
		return nil, err
	}

	reqBytes, _ := json.Marshal(req)

	endpoint := "/chat/completions"
	httpReq, err := http.NewRequest("POST", c.Config.BaseURL+endpoint, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)

	res, err := c.SendRequest(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var chatResponse ChatResponse
	if err := json.NewDecoder(res.Body).Decode(&chatResponse); err != nil {
		return nil, err
	}

	return &chatResponse, nil
}

func validate(req *ChatCompletionRequest) error {
	if len(req.Messages) == 0 {
		return custom_errors.ErrNoMessages
	}

	isAllowed := false

	allowedModels := []constants.ChatGPTModel{
		constants.GPT35Turbo, constants.GPT35Turbo0613, constants.GPT35Turbo16k, constants.GPT35Turbo16k0613, constants.GPT4, constants.GPT4_0613, constants.GPT4_32k, constants.GPT4_32k_0613,
	}

	for _, model := range allowedModels {
		if req.Model == model {
			isAllowed = true
		}
	}

	if !isAllowed {
		return custom_errors.ErrInvalidModel
	}

	for _, message := range req.Messages {
		if message.Role != constants.ChatGPTModelRoleUser && message.Role != constants.ChatGPTModelRoleSystem && message.Role != constants.ChatGPTModelRoleAssistant {
			return custom_errors.ErrInvalidRole
		}
	}

	if req.Temperature < 0 || req.Temperature > 2 {
		return custom_errors.ErrInvalidTemp
	}

	if req.PresencePenalty < -2 || req.PresencePenalty > 2 {
		return custom_errors.ErrInvalidPresencePenalty
	}

	if req.FrequencyPenalty < -2 || req.FrequencyPenalty > 2 {
		return custom_errors.ErrInvalidFrequencyPenalty
	}

	return nil
}
