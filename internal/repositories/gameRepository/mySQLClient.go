package gameRepository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"minesweeper-go/internal/core/domain/game"
)

type DatastoreClient struct {
	DB    *sql.DB
}

func NewMySQLClient(db *sql.DB) DatastoreClient {
	return DatastoreClient{
		DB:    db,
	}
}

func (dc DatastoreClient) Get(ctx context.Context, id string) (game.Game, error) {
	var currentGame game.Game

	var boardSettingsJSON []byte
	var boardJSON []byte
	var playerViewJSON []byte

	query := `SELECT * FROM games WHERE id = ?`

	row := dc.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&currentGame.ID,
		&currentGame.Name,
		&currentGame.State,
		&boardSettingsJSON,
		&boardJSON,
		&playerViewJSON,
		&currentGame.CellsRemaining,
	)
	if err != nil {
		return currentGame, err
	}

	err = json.Unmarshal(boardSettingsJSON, &currentGame.BoardSettings)
	if err != nil {
		return currentGame, err
	}

	err = json.Unmarshal(boardJSON, &currentGame.Board)
	if err != nil {
		return currentGame, err
	}

	err = json.Unmarshal(playerViewJSON, &currentGame.PlayerView)
	if err != nil {
		return currentGame, err
	}


	return currentGame, nil
}

func (dc DatastoreClient) Save(ctx context.Context, currentGame game.Game) error {

	// TODO: Consider converting games into smaller data for DB

	boardSettingsJSON, err := json.Marshal(currentGame.BoardSettings)
	if err != nil {
		return err
	}

	boardJSON, err := json.Marshal(currentGame.Board)
	if err != nil {
		return err
	}

	playerViewJSON, err := json.Marshal(currentGame.PlayerView)
	if err != nil {
		return err
	}

	// Get a Tx for making transaction requests.
  tx, err := dc.DB.BeginTx(ctx, nil)
  if err != nil {
      return err
  }
  // Defer a rollback in case anything fails.
  defer tx.Rollback()

	var gameExists bool
  if err = tx.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT * FROM games WHERE id = ?)",
		currentGame.ID,
	).Scan(&gameExists); err != nil {
    if err == sql.ErrNoRows {
        return err
    }
    return err
  }

	fmt.Println(gameExists)
	log.Println(gameExists)

  if gameExists {
		query := `UPDATE games SET
			name = ?,
			state = ?,
			board_settings = ?,
			board = ?,
			player_view = ?,
			cells_remaining = ?
			WHERE id = ?`
		
		_, err = tx.ExecContext(
			ctx,
			query,
			currentGame.Name,
			currentGame.State,
			boardSettingsJSON,
			boardJSON,
			playerViewJSON,
			currentGame.CellsRemaining,
			currentGame.ID,
		)
	} else {
		query := `INSERT INTO games(
			id,
			name,
			state,
			board_settings,
			board,
			player_view,
			cells_remaining
		) VALUES (?, ?, ?, ?, ?, ?, ?)`

		_, err = tx.ExecContext(
			ctx,
			query,
			currentGame.ID,
			currentGame.Name,
			currentGame.State,
			boardSettingsJSON,
			boardJSON,
			playerViewJSON,
			currentGame.CellsRemaining,
		)
	}
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
