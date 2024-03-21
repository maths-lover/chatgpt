package constants

type ChatGPTModel string

const (
	GPT35Turbo ChatGPTModel = "gpt-3.5-turbo"

	// Deprecated: Use gpt-3.5-turbo-0613 instead, model discontinued on 09/13/2023
	GPT35Turbo0301 ChatGPTModel = "gpt-3.5-turbo-0301"

	GPT35Turbo0613    ChatGPTModel = "gpt-3.5-turbo-0613"
	GPT35Turbo16k     ChatGPTModel = "gpt-3.5-turbo-16k"
	GPT35Turbo16k0613 ChatGPTModel = "gpt-3.5-turbo-16k-0613"
	GPT4              ChatGPTModel = "gpt-4"

	// Deprecated: Use gpt-4-0613 instead, model discontinued on 09/13/2023
	GPT4_0314 ChatGPTModel = "gpt-4-0314"

	GPT4_0613 ChatGPTModel = "gpt-4-0613"
	GPT4_32k  ChatGPTModel = "gpt-4-32k"

	// Deprecated: Use gpt-4-32k-0613 instead, model discontinued on 09/13/2023
	GPT4_32k_0314 ChatGPTModel = "gpt-4-32k-0314"

	GPT4_32k_0613 ChatGPTModel = "gpt-4-32k-0613"
)

type ChatGPTModelRole string

const (
	ChatGPTModelRoleUser      ChatGPTModelRole = "user"
	ChatGPTModelRoleSystem    ChatGPTModelRole = "system"
	ChatGPTModelRoleAssistant ChatGPTModelRole = "assistant"
)
