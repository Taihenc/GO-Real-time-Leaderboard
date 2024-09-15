package model

type Player struct {
	Name  string
	Score int
}

type AddScoreRequest struct {
	Game       string
	PlayerName string
	Score      int
}
