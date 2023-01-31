DROP TABLE IF EXISTS `games`;

CREATE TABLE games (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    state varchar(255) NOT NULL,
    board_settings JSON NOT NULL,
    board JSON NOT NULL,
    player_view JSON NOT NULL,
    cells_remaining int NOT NULL
);
