package models

type ErrCode struct {
	ErrCode string `json:"errCode"`
}

type Prompt struct {
	Id          int    `json:"-"`
	PromptName  string `gorm:"column:promptname" json:"name"`
	Description string `json:"description"`
	Prompts     string `json:"prompts"`
	Uid         string `json:"-"`
	Designer    int64  `json:"-"`
}
type ChatConversation struct {
	Id      string
	Title   string `gorm:"column:title"`
	Uid     string `gorm:"column:uid"`
	ModelId int    `gorm:"column:modelid"`
}

type ChatMessage struct {
	Id      int
	Content string
	ChatId  string `gorm:"column:chatid"`
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

type Game struct {
	GameId      string `gorm:"column:gameId"`
	Name        string
	Description string
}

type GameMessage struct {
	Uid    string `gorm:"uid" json:"-"`
	Story  string `json:"story"`
	Chocie string `json:"choice"`
	Round  int    `json:"round"`
	GameId string `gorm:"column:gameId" json:"-"`
}
