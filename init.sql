CREATE TABLE IF NOT EXISTS User (
    ID INTEGER PRIMARY KEY,
    Sub TEXT,
    Username TEXT
);

CREATE TABLE IF NOT EXISTS UserHighscore (
    ID INTEGER PRIMARY KEY,
    Name TEXT,
    Score INTEGER,
    FOREIGN KEY(ID) REFERENCES User(ID)
);

CREATE TABLE IF NOT EXISTS Tournament (
    ID INTEGER PRIMARY KEY,
    Name TEXT
);


CREATE TABLE IF NOT EXISTS Team (
    ID INTEGER PRIMARY KEY,
    Name TEXT,
    Logo TEXT,
    TournamentID INTEGER,
    FOREIGN KEY(TournamentID) REFERENCES Tournament(ID)
);



CREATE TABLE IF NOT EXISTS Stage (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Name TEXT,
    Stage TEXT,
    TournamentID INTEGER,
    FOREIGN KEY(TournamentID) REFERENCES Tournament(ID)
);

CREATE TABLE IF NOT EXISTS Match (
    ID INTEGER PRIMARY KEY,
    AwayTeamID INTEGER,
    Date TEXT,
    GroupID INTEGER,
    HomeTeamID INTEGER,
    Outcome TEXT,
    Status TEXT,
    TournamentID INTEGER,
    StageID INTEGER,
    FOREIGN KEY(AwayTeamID) REFERENCES Team(ID),
    FOREIGN KEY(HomeTeamID) REFERENCES Team(ID),
    FOREIGN KEY(TournamentID) REFERENCES Tournament(ID),
    FOREIGN KEY(StageID) REFERENCES Stage(ID)
);

CREATE TABLE IF NOT EXISTS Bet (
    ID INTEGER PRIMARY KEY,
    MatchID INTEGER,
    Bet TEXT,
    FOREIGN KEY(MatchID) REFERENCES Match(ID)
);

-- INSERT INTO Tournament (ID, Name) VALUES 
-- (1, 'Champions League');

-- INSERT INTO Team (Name, Logo, TournamentID) VALUES 
-- ('AS Roma', 'ROMA', 1),
-- ('Real Madrid', 'REAL_MADRID', 1),
-- ('Bayern Münich', 'BAYERN', 1),
-- ('Barcelona', 'BARCA', 1),
-- ('Dortmund', 'DORTMUND', 1);

-- INSERT INTO Stage (Name, Stage, TournamentID) VALUES
-- ('Round of 32', "RO32", 1),
-- ('Round of 16', "RO16", 1),
-- ('Quarter finals', "QUARTER", 1),
-- ('Semi finals', "SEMI", 1),
-- ('Final', "FINAL", 1);
