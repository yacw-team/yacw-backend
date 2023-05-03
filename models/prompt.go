package models

type Prompt struct {
	Id            int
	ModelName     string
	Description   string
	Details       string
	ExamoleInput  string
	ExamoleOutput string
	Prompts       string
	Uid           string
	Designer      int
}
