package models

type Prompt struct {
	Id          int    `json:"id"`
	PromptName  string `gorm:"column:promptname" json:"name"`
	Description string `json:"description"`
	Prompts     string `json:"prompts"`
	Uid         string `json:"-"`
	Designer    int64  `json:"-"`
}
