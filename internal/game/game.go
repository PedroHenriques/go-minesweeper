package game

import (
	"time"

	"github.com/pedrohenriques/go-minesweeper/internal/configs"
	"github.com/pedrohenriques/go-minesweeper/internal/minefield"
)

// Game contains all the information about a game
type game struct {
	startTs      time.Time
	endTs        time.Time
	numMines     int
	numRows      int
	numCols      int
	flagsEnabled bool
	lives        int
	minefield    minefield.IMinefield
}

/*
Config returns the configuration data for a game.
*/
func (game *game) Config() *config {
	return &config{
		NumMines: game.numMines,
		NumRows:  game.numRows,
		NumCols:  game.numCols,
	}
}

/*
StartTime returns the Time object of when the game started.
*/
func (game *game) StartTime() time.Time {
	return game.startTs
}

/*
State returns information about the game's state.
*/
func (game *game) State() int {
	stats := game.minefield.Stats()

	if stats.NumMinesRevealed >= game.lives {
		return configs.StateLoss
	}

	numNonMineTilesRevealed := stats.NumTilesRevealed - stats.NumMinesRevealed
	if numNonMineTilesRevealed == game.numRows*game.numCols-game.numMines {
		return configs.StateWin
	}

	return configs.StateOnGoing
}

/*
Tile searches for the tile in the requested row and column.
The row and column coordinates are zero-indexed.
*/
func (game *game) Tile(rowIndex int, colIndex int) (minefield.ITile, error) {
	return game.minefield.Tile(rowIndex, colIndex)
}

/*
RevealTile reveals the requested tile.
If the tile is empty the patch it belongs to will be revealed.
*/
func (game *game) RevealTile(rowIndex int, colIndex int) ([]int, error) {
	if game.State() != configs.StateOnGoing {
		return nil, nil
	}

	tileIndexes, error := game.minefield.RevealTile(rowIndex, colIndex)

	if game.State() != configs.StateOnGoing {
		game.endTs = time.Now()
	}

	return tileIndexes, error
}

/*
ToggleFlag flips the flag state for the requested tile.
*/
func (game *game) ToggleFlag(rowIndex int, colIndex int) error {
	if !game.flagsEnabled || game.State() != configs.StateOnGoing {
		return nil
	}

	return game.minefield.ToggleFlag(rowIndex, colIndex)
}

/*
ProcessAdjacentTiles applies to a revealed tile and will check the adjacent
tiles for flags.
If the number of adjacent flags is >= the number on the tile it will reveal
all adjacent tiles without a flag.
*/
func (game *game) ProcessAdjacentTiles(rowIndex int, colIndex int) ([]int, error) {
	if game.State() != configs.StateOnGoing {
		return nil, nil
	}

	tileIndexes, error := game.minefield.ProcessAdjacentTiles(rowIndex, colIndex)

	if game.State() != configs.StateOnGoing {
		game.endTs = time.Now()
	}

	return tileIndexes, error
}

/*
Stats returns information about the current game.
*/
func (game *game) Stats() stats {
	minefieldStats := game.minefield.Stats()

	return stats{
		StartTime:      game.startTs,
		EndTime:        game.endTs,
		RemainingMines: game.numMines - minefieldStats.NumFlags,
		RemainingLives: game.lives - minefieldStats.NumMinesRevealed,
	}
}

// Config contains the setup of a game
type config struct {
	NumMines int
	NumRows  int
	NumCols  int
}

// Stats contains statistics about a game
type stats struct {
	StartTime      time.Time
	EndTime        time.Time
	RemainingMines int
	RemainingLives int
}
