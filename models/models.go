package models

type ErrCode struct {
	ErrCode string `json:"errCode"`
}

type Prompt struct {
	Id            int    `json:"id"`
	ModelName     string `gorm:"column:modelname" json:"name"`
	Description   string `json:"description"`
	Details       string `json:"details"`
	ExampleInput  string `gorm:"column:exampleinput" json:"exampleInput"`
	ExampleOutput string `gorm:"column:exampleoutput" json:"exampleOutput"`
	Prompts       string `json:"content"`
	Uid           string `json:"-"`
	Designer      string `json:"-"`
}

type ChatConversation struct {
	Id      int    `json:"id"`
	Uid     string `json:"uid"`
	ModelId int    `gorm:"column:modelid"`
	Title   string `json:"title"`
}

type ChatMessage struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	ChatId  int    `gorm:"column:chatid"`
	Actor   string `json:"actor"`
	Show    int    `json:"show"`
}

type Game struct {
	Id           int    `json:"id"`
	Uid          string `json:"uid"`
	Background   string `json:"background"`
	Protagonist  string `json:"protagonist"`
	Goal         string `json:"goal"`
	SystemPrompt string `gorm:"column:systemprompt" json:"systemPrompt"`
}

type GameMessage struct {
	Id      int
	Content string
	Actor   int
	GameId  int
}

type Literature struct {
	Id       int
	Prompt   string
	Name     string
	Uid      string
	Designer int
}

type Personality struct {
	Id          int    `json:"id"`
	ModelName   string `gorm:"column:personalityname" json:"name"`
	Description string `json:"description"`
	Prompts     string `json:"prompts"`
	Uid         string `json:"-"`
	Designer    string `json:"-"`
}

type PsychologyConversation struct {
	Id            int
	SystemPrompt  string
	Uid           string
	PersonalityId int
}

type PsychologyMessage struct {
	Id           int
	Content      string
	PsychologyId int
	Actor        int
}

type Translation struct {
	Id           int
	Uid          string
	OriginLang   string
	GoalLang     string
	EmotionId    int
	LiteratureId int
	Input        string
	Output       string
}

type TranslationEmotion struct {
	Id       int
	Prompt   string
	Name     string
	Uid      string
	Designer int
}

type User struct {
	Uid string
}
