package models

type GameMessage struct {
	Id      int
	Content string
	Actor   int
	GameId  int
}
