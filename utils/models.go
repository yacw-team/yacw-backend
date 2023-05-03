package utils

type ChatConversation struct {
	Id           int
	SystemPrompt string
	Uid          string
	ModelId      int
	PromptId     int
}

type ChatMessage struct {
	Id      int
	Content string
	ChatId  int
	Actor   int
}

type Game struct {
	Id           int
	Uid          string
	Background   string
	Protagonist  string
	Goal         string
	SystemPrompt string
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
	Id        int
	ModelName string
	Details   string
	Prompt    string
	Uid       string
	Designer  int
}

type Prompt struct {
	Id            int
	ModelName     string
	Description   string
	Details       string
	ExampleInput  string
	ExampleOutput string
	Prompts       string
	Uid           string
	Designer      int
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
