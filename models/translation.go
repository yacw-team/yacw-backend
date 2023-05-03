package models

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
