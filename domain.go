package main

import "time"

type Bet struct {
	ID      int    `json:"id"`
	MatchID int    `json:"matchId"`
	Bet     string `json:"bet"` 
}

type Highscore map[int]UserHighscore

type Match struct {
	ID           int       `json:"id"`
	AwayScore    int       `json:"awayScore"`
	AwayTeamID   int       `json:"awayTeamId"`
	Date         time.Time `json:"date"`
	GroupID      int       `json:"groupId"`
	HomeScore    []int     `json:"homeScore"`
	HomeTeamID   int       `json:"homeTeamId"`
	Outcome      string    `json:"outcome"`
	Status       string    `json:"status"`
	TournamentID int       `json:"tournamentId"`
}

type Stage struct {
	ID 			int						`json:"id"`
	Matches map[int]Match `json:"matches"`
	Name    string        `json:"name"`
	Stage   string        `json:"stage"`
}

type Team struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Logo         string `json:"logo"`
}

type Tournament struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Stages map[int]Stage `json:"stages"`
	Teams  map[int]Team  `json:"teams"`
}

type User struct {
	ID       int    `json:"id"`
	Sub      string `json:"sub"`
	Username string `json:"name"`
}

type UserHighscore struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}