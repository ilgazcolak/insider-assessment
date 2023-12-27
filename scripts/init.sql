drop table if exists teams;
drop table if exists players;
drop table if exists teams;
drop table if exists statistics;

-- create teams table
create table  teams(
    id serial primary key,
    name varchar not null
);

-- create players table
create table  players(
    id serial primary key,
    team_id int,
    name varchar not null,
    position varchar not null,
    FOREIGN KEY (team_id) REFERENCES teams(id)
);

-- create matches table
create table  matches(
    id serial primary key,
    home_team_id int not null,
    away_team_id int not null,
    FOREIGN KEY (home_team_id) REFERENCES teams(id),
    FOREIGN KEY (away_team_id) REFERENCES teams(id)
);

-- create statistics table
create table statistics(
    id serial primary key,
    player_id int not null,
    match_id int not null,
    score int,
    attempt_count int,
    two_points_attempt int,
    three_points_attempt int,
    two_points_success int,
    three_points_success int,
    assist int,
    FOREIGN KEY (player_id) REFERENCES players(id),
    FOREIGN KEY (match_id) REFERENCES matches(id)
);

-- populate teams table
insert into teams(id, name) values 
(1, 'Brooklyn Nets'),
(2, 'Detroit Pistons'),
(3, 'Orlando Magic'),
(4, 'Washington Wizards'),
(5, 'Atlanta Hawks'),
(6, 'Chicago Bulls'),
(7, 'Indiana Pacers'),
(8, 'Houston Rockets');

-- populate players table
insert into players(id, team_id, name, position) values
(1, 1, 'Cam Thomas', 'SG'),
(2, 1, 'Mikal Bridges', 'SF'),
(3, 1, 'Spencer Dinwiddie', 'PG'),
(4, 1, 'Cameron Johnson', 'PF'),
(5, 1, 'Nic Claxton', 'C'),
(6, 2, 'Jaden Ivey', 'SG'),
(7, 2, 'Bojan Bogdanovic', 'SF'),
(8, 2, 'Cade Cunningham', 'PG'),
(9, 2, 'Isaiah Stewart', 'PF'),
(10, 2, 'Jalen Duren', 'C'),
(11, 3, 'Jalen Suggs', 'SG'),
(12, 3, 'Franz Wagner', 'SF'),
(13, 3, 'Anthony Black', 'PG'),
(14, 3, 'Paolo Banchero', 'PF'),
(15, 3, 'Wendell Carter Jr.', 'C'),
(16, 4, 'Jordan Poole', 'SG'),
(17, 4, 'Deni Avdija', 'SF'),
(18, 4, 'Tyus Jones', 'PG'),
(19, 4, 'Kyle Kuzma', 'PF'),
(20, 4, 'Daniel Gafford', 'C'),
(21, 5, 'Dejounte Murray', 'SG'),
(22, 5, 'Saddiq Bey', 'SF'),
(23, 5, 'Trae Young', 'PG'),
(24, 5, 'Jalen Johnson', 'PF'),
(25, 5, 'Clint Capela', 'C'),
(26, 6, 'Alex Caruso', 'SG'),
(27, 6, 'DeMar DeRozan', 'SF'),
(28, 6, 'Coby White', 'PG'),
(29, 6, 'Patrick Williams', 'PF'),
(30, 6, 'Andre Drummond', 'C'),
(31, 7, 'Andrew Nembhard', 'SG'),
(32, 7, 'Aaron Nesmith', 'SF'),
(33, 7, 'Tyrese Haliburton', 'PG'),
(34, 7, 'Jalen Smith', 'PF'),
(35, 7, 'Myles Turner', 'C'),
(36, 8, 'Jalen Green', 'SG'),
(37, 8, 'Dillon Brooks', 'SF'),
(38, 8, 'Fred VanVleet', 'PG'),
(39, 8, 'Jabari Smith Jr.', 'PF'),
(40, 8, 'Alperen Sengun', 'C');

-- populate matches table
insert into matches(id, home_team_id, away_team_id) values 
(1, 1, 2),
(2, 3, 4),
(3, 5, 6),
(4, 7, 8);

-- populate statistics table
insert into statistics(player_id, match_id, score, attempt_count, two_points_attempt, three_points_attempt, two_points_success, three_points_success, assist) values
(1, 1, 0, 0, 0, 0, 0, 0, 0),
(2, 1, 0, 0, 0, 0, 0, 0, 0),
(3, 1, 0, 0, 0, 0, 0, 0, 0),
(4, 1, 0, 0, 0, 0, 0, 0, 0),
(5, 1, 0, 0, 0, 0, 0, 0, 0),
(6, 1, 0, 0, 0, 0, 0, 0, 0),
(7, 1, 0, 0, 0, 0, 0, 0, 0),
(8, 1, 0, 0, 0, 0, 0, 0, 0),
(9, 1, 0, 0, 0, 0, 0, 0, 0),
(10, 1, 0, 0, 0, 0, 0, 0, 0),
(11, 2, 0, 0, 0, 0, 0, 0, 0),
(12, 2, 0, 0, 0, 0, 0, 0, 0),
(13, 2, 0, 0, 0, 0, 0, 0, 0),
(14, 2, 0, 0, 0, 0, 0, 0, 0),
(15, 2, 0, 0, 0, 0, 0, 0, 0),
(16, 2, 0, 0, 0, 0, 0, 0, 0),
(17, 2, 0, 0, 0, 0, 0, 0, 0),
(18, 2, 0, 0, 0, 0, 0, 0, 0),
(19, 2, 0, 0, 0, 0, 0, 0, 0),
(20, 2, 0, 0, 0, 0, 0, 0, 0),
(21, 3, 0, 0, 0, 0, 0, 0, 0),
(22, 3, 0, 0, 0, 0, 0, 0, 0),
(23, 3, 0, 0, 0, 0, 0, 0, 0),
(24, 3, 0, 0, 0, 0, 0, 0, 0),
(25, 3, 0, 0, 0, 0, 0, 0, 0),
(26, 3, 0, 0, 0, 0, 0, 0, 0),
(27, 3, 0, 0, 0, 0, 0, 0, 0),
(28, 3, 0, 0, 0, 0, 0, 0, 0),
(29, 3, 0, 0, 0, 0, 0, 0, 0),
(30, 3, 0, 0, 0, 0, 0, 0, 0),
(31, 4, 0, 0, 0, 0, 0, 0, 0),
(32, 4, 0, 0, 0, 0, 0, 0, 0),
(33, 4, 0, 0, 0, 0, 0, 0, 0),
(34, 4, 0, 0, 0, 0, 0, 0, 0),
(35, 4, 0, 0, 0, 0, 0, 0, 0),
(36, 4, 0, 0, 0, 0, 0, 0, 0),
(37, 4, 0, 0, 0, 0, 0, 0, 0),
(38, 4, 0, 0, 0, 0, 0, 0, 0),
(39, 4, 0, 0, 0, 0, 0, 0, 0),
(40, 4, 0, 0, 0, 0, 0, 0, 0);
