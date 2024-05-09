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

func (db *DB) GetTournament(id int) (*Tournament, error ) {
	tournament := &Tournament{
		Stages: make(map[int]Stage),
		Teams:  make(map[int]Team),
	}

	err := db.QueryRow("SELECT * FROM Tournament WHERE id = ?", id).Scan(&tournament.ID, &tournament.Name)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM Stage WHERE TournamentID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var stage Stage
		err = rows.Scan(&stage.ID, &stage.Name, &stage.Stage)
		if err != nil {
			return nil, err
		}
		tournament.Stages[stage.ID] = stage
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

func (db *DB) GetTournaments() (map[int]*Tournament, error) {
	tournaments := make(map[int]*Tournament)

	rows, err := db.Query(`
		SELECT 
			Tournament.ID, Tournament.Name, 
			Stage.ID, Stage.Name, Stage.Stage, 
			Team.ID, Team.Name, Team.Logo
		FROM Tournament
		LEFT JOIN Stage ON Tournament.ID = Stage.TournamentID
		LEFT JOIN Team ON Tournament.ID = Team.TournamentID
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tournamentID, stageID, teamID int
		var tournamentName, stageName, stageStage, teamName, teamLogo string
		err = rows.Scan(&tournamentID, &tournamentName, &stageID, &stageName, &stageStage, &teamID, &teamName, &teamLogo)
		if err != nil {
			return nil, err
		}

		tournament, exists := tournaments[tournamentID]
		if !exists {
			tournament = &Tournament{
				ID:     tournamentID,
				Name:   tournamentName,
				Stages: make(map[int]Stage),
				Teams:  make(map[int]Team),
			}
			tournaments[tournamentID] = tournament
		}

		tournament.Stages[stageID] = Stage{ID: stageID, Name: stageName, Stage: stageStage}
		tournament.Teams[teamID] = Team{ID: teamID, Name: teamName, Logo: teamLogo}
	}
	if rows.Err() != nil {
		return nil, err
	}

	return tournaments, nil
}