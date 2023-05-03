package models

type ChatConversation struct {
	Id           int
	SystemPrompt string
	Uid          string
	ModelId      int
	PromptId     int
}
