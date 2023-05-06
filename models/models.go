package models

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
	Actor   int
	Show    int
}

type Game struct {
	Id           int
	Uid          string
	Background   string
	Protagonist  string
	Goal         string
	SystemPrompt string `gorm:"column:systemprompt"`
}

type GameMessage struct {
	Id      int
	Content string
	Actor   int
	GameId  int `gorm:"column:gameid"`
}

type Literature struct {
	Id       int
	Prompt   string
	Name     string
	Uid      string
	Designer int
}

type Personality struct {
	Id              int
	PersonalityName string `gorm:"column:personalityname"`
	Description     string
	Prompts         string
	Uid             string
	Designer        int
}

type PsychologyConversation struct {
	Id            int
	SystemPrompt  string `gorm:"column:systemprompt"`
	Uid           string
	PersonalityId int `gorm:"column:personalityid"`
}

type PsychologyMessage struct {
	Id           int
	Content      string
	PsychologyId int `gorm:"column:psychologyid"`
	Actor        int
}

type Translation struct {
	Id         int
	Uid        string
	OriginLang string `gorm:"column:originlang"`
	GoalLang   string `gorm:"column:goallang"`
	Input      string
	Output     string
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
