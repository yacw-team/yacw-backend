package models

type Prompt struct {
	Id          int    `json:"id"`
	PromptName  string `gorm:"column:promptname" json:"name"`
	Description string `json:"description"`
	Prompts     string `json:"prompts"`
	Uid         string `json:"-"`
	Designer    int64  `json:"-"`
}
type ChatConversation struct {
	Id      int
	Title   string `gorm:"column:title"`
	Uid     string `gorm:"column:uid"`
	ModelId int    `gorm:"column:modelid"`
}

type ChatMessage struct {
	Id      int
	Content string
	ChatId  int `gorm:"column:chatid"`
	Actor   string
	Show    int
}

type Personality struct {
	Id          int    `json:"id"`
	ModelName   string `gorm:"column:personalityname" json:"name"`
	Description string `json:"description"`
	Prompts     string `json:"prompts"`
	Uid         string `json:"-"`
	Designer    string `json:"-"`
}
