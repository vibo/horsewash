package main

import "time"

type Bet struct {
	ID      int    `json:"id"`
	Bet     string `json:"bet"`
	MatchID int    `json:"matchId"`
	UserID  int    `json:"userId"`
}

type EliminationRound struct {
	ID      int
	Name    string `json:"name"`
	Matches []int  `json:"matches"`
	Teams   []int  `json:"teams"`
}

type Group struct {
	ID      int
	Name    string `json:"name"`
	Matches []int  `json:"matches"`
	Teams   []int  `json:"teams"`
}

type Highscore map[int]UserHighscore

type Match struct {
	ID           int       `json:"id"`
	AwayScore    int       `json:"awayScore"`
	AwayTeamID   int       `json:"awayTeamId"`
	Date         time.Time `json:"date"`
	HomeScore    int       `json:"homeScore"`
	HomeTeamID   int       `json:"homeTeamId"`
	Outcome      string    `json:"outcome"`
	Status       string    `json:"status"`
	TournamentID int       `json:"tournamentId"`
}

type MatchDetail struct {
	ID           int         `json:"id"`
	AwayScore    int         `json:"awayScore"`
	AwayTeamID   int         `json:"awayTeamId"`
	Date         time.Time   `json:"date"`
	HomeScore    int         `json:"homeScore"`
	HomeTeamID   int         `json:"homeTeamId"`
	Outcome      string      `json:"outcome"`
	Status       string      `json:"status"`
	TournamentID int         `json:"tournamentId"`
	Bets         map[int]Bet `json:"bets"`
}

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type TournamentSummary struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TournamentDetail struct {
	ID                int                       `json:"id"`
	Bets 							map[int]Bet 							`json:"bets"`
	EliminationRounds map[int]EliminationRound  `json:"eliminationRounds"`
	Groups            map[int]Group             `json:"groups"`
	Matches           map[int]Match             `json:"matches"`
	Name              string                    `json:"name"`
	Stage             string                    `json:"stage"`
	Teams             map[int]Team              `json:"teams"`
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
