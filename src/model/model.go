package model

type Player struct {
	Name  string
	Score int
}

type LeaderboardRecord struct {
	Game       string
	PlayerName string
	Score      int
}

type UserPassword struct {
	Username string
	Password string
}
