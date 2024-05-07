package main

type Bet struct {
	ID      int
	MatchID int
}
type Highscore struct{}
type Match struct {
	ID           int
	awayTeam     Team
	date         int
	groupID      int
	homeTeam     Team
	outcome      string
	score        []int
	status       string
	tournamentID int
}
type Team struct {
	ID           int
	name         string
	logo         string
	tournamentID int
}

type Tournament struct {
	ID    int
	name  string
	phase string
	users []int
}

type User struct {
	ID   int
	Name string
}

type Stage struct {
	name    string
	matches []Match
}

type Tournament struct {
	ID     int
	name   string
	stages []Stage
}
