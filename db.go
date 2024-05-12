package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "modernc.org/sqlite"
)
 

type DB struct {
	*sql.DB
}


func openDatabase(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil { 
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	initDatabase(db)

	return &DB{DB: db}, nil
}

func initDatabase(db *sql.DB) {
	initScript, err := os.ReadFile("./init.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(initScript))
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) GetUser(id string) (*User, error ) {
	user := &User{}

	err := db.QueryRow("SELECT id, username FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *DB) GetTournament(id int) (*TournamentDetail, error ) {
	tournament := &TournamentDetail{
		EliminationRounds: make(map[int]EliminationRound),
		Groups: make(map[int]Group),
		Teams: make(map[int]Team),
	}

	err := db.QueryRow("SELECT * FROM Tournament WHERE id = ?", id).Scan(&tournament.ID, &tournament.Name)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM Elimination WHERE TournamentID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var eliminationRound EliminationRound
		err = rows.Scan(&eliminationRound.ID, &eliminationRound.Name, &eliminationRound.Matches)
		if err != nil {
			return nil, err
		}
		tournament.EliminationRounds[eliminationRound.ID] = eliminationRound
	}

	rows, err = db.Query("SELECT * FROM Group WHERE TournamentID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var group Group
		err = rows.Scan(&group.ID, &group.Name, &group.Matches, &group.Teams)
		if err != nil {
			return nil, err
		}
		tournament.Groups[group.ID] = group
	}

	rows, err = db.Query("SELECT * FROM Team WHERE TournamentID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var team Team
		err = rows.Scan(&team.ID, &team.Name, &team.Logo)
		if err != nil {
			return nil, err
		}
		tournament.Teams[team.ID] = team
	}

	return tournament, nil
}

func (db *DB) GetTournaments() (map[int]TournamentSummary, error) {
	tournaments := make(map[int]TournamentSummary)

	rows, err := db.Query("SELECT * FROM Tournament")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tournament TournamentSummary
		err = rows.Scan(&tournament.ID, &tournament.Name)
		if err != nil {
			return nil, err
		}
		tournaments[tournament.ID] = tournament
	}

	return tournaments, nil
}

func (db *DB) GetTournamentBets(id int) (map[int]Bet, error) {
	bets := make(map[int]Bet)

	rows, err := db.Query("SELECT * FROM Bets WHERE TournamentID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bet Bet
		err = rows.Scan(&bet.ID, &bet.Bet, &bet.MatchID)
		if err != nil {
			return nil, err
		}
		bets[bet.ID] = bet
	}

	return bets, nil
}

func (db *DB) GetMatch(id int) (*Match, error) {
	match := &Match{}

	var dateStr string

	err := db.QueryRow("SELECT * FROM Match WHERE ID = ?", id).Scan(
		&match.ID,
		&match.AwayScore,
		&match.AwayTeamID,
		&dateStr,
		&match.HomeScore,
		&match.HomeTeamID,
		&match.Outcome,
		&match.Status,
		&match.TournamentID,
	)
	if err != nil {
		return nil, err
	}

	match.Date, err = time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (db *DB) GetMatchDetails(id int) (*MatchDetail, error) {
	match, err := db.GetMatch(id)
	if err != nil {
		return nil, err
	}

	matchDetail := &MatchDetail{
		ID:           match.ID,
		AwayScore:    match.AwayScore,
		AwayTeamID:   match.AwayTeamID,
		Date:         match.Date,
		HomeScore:    match.HomeScore,
		HomeTeamID:   match.HomeTeamID,
		Outcome:      match.Outcome,
		Status:       match.Status,
		TournamentID: match.TournamentID,
		Bets:         make(map[int]Bet),
	}

	rows, err := db.Query("SELECT * FROM Bet WHERE MatchID = ?", match.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bet Bet
		err = rows.Scan(&bet.ID, &bet.Bet, &bet.MatchID, &bet.UserID)
		if err != nil {
			return nil, err
		}
		matchDetail.Bets[bet.ID] = bet
	}

	return matchDetail, nil
}