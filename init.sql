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

CREATE TABLE IF NOT EXISTS EliminationRound (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Stage TEXT,
    TournamentID INTEGER,
    FOREIGN KEY(TournamentID) REFERENCES Tournament(ID)
);

CREATE TABLE IF NOT EXISTS Group (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Name TEXT,
    TournamentID INTEGER,
    FOREIGN KEY(TournamentID) REFERENCES Tournament(ID)
);

CREATE TABLE IF NOT EXISTS Match (
    ID INTEGER PRIMARY KEY,
    AwayScore INTEGER,
    AwayTeamID INTEGER,
    Date TEXT,
    HomeScore INTEGER,
    HomeTeamID INTEGER,
    Outcome TEXT,
    Status TEXT,
    FOREIGN KEY(AwayTeamID) REFERENCES Team(ID),
    FOREIGN KEY(HomeTeamID) REFERENCES Team(ID),
    FOREIGN KEY(TournamentID) REFERENCES Tournament(ID),
);

CREATE TABLE IF NOT EXISTS GroupMatches (
    GroupID INTEGER,
    MatchID INTEGER,
    PRIMARY KEY(GroupID, MatchID),
    FOREIGN KEY(GroupID) REFERENCES Group(ID),
    FOREIGN KEY(MatchID) REFERENCES Match(ID)
);

CREATE TABLE IF NOT EXISTS EliminationRoundMatches (
    EliminationRoundID INTEGER,
    MatchID INTEGER,
    PRIMARY KEY(EliminationRoundID, MatchID),
    FOREIGN KEY(EliminationRoundID) REFERENCES EliminationRound(ID),
    FOREIGN KEY(MatchID) REFERENCES Match(ID)
);

CREATE TABLE IF NOT EXISTS Bet (
    ID INTEGER PRIMARY KEY,
    Bet TEXT,
    MatchID INTEGER,
    UserID INTEGER,
    FOREIGN KEY(MatchID) REFERENCES Match(ID),
    FOREIGN KEY(UserID) REFERENCES User(ID)
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

-- INSERT INTO Match (AwayScore, AwayTeamID, Date, HomeScore, HomeTeamID, Status, TournamentID, StageID) VALUES 
-- (0, 1, '2024-05-01', 0, 2, 'Scheduled', 1, 1),
-- (0, 3, '2024-05-02', 0, 4, 'Scheduled', 1, 1),
-- (0, 5, '2024-05-03', 0, 1, 'Scheduled', 1, 2),
-- (0, 2, '2024-05-04', 0, 3, 'Scheduled', 1, 2);

-- INSERT INTO Bet (MatchID, UserID, Bet) VALUES 
-- (1, 1, 'HOME'),
-- (2, 1, 'AWAY'),
-- (3, 2, 'DRAW'),
-- (4, 2, 'HOME');