package models

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
