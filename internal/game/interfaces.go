package game

import (
	"time"

	"github.com/pedrohenriques/go-minesweeper/internal/minefield"
)

type IGame interface {
	/*
		Config returns the configuration data for a game.
	*/
	Config() *config
	/*
		StartTime returns the Time object of when the game started.
	*/
	StartTime() time.Time
	/*
		State returns information about the game's state.
		0 = on going
		1 = ended in win
		2 = ended in loss
	*/
	State() int
	/*
		Tile searches for the tile in the requested row and column.
		The row and column coordinates are zero-indexed.
	*/
	Tile(rowIndex int, colIndex int) (minefield.ITile, error)
	/*
		RevealTile reveals the requested tile.
		If the tile is empty the patch it belongs to will be revealed.
	*/
	RevealTile(rowIndex int, colIndex int) ([]int, error)
	/*
		ToggleFlag flips the flag state for the requested tile.
	*/
	ToggleFlag(rowIndex int, colIndex int) error
	/*
		ProcessAdjacentTiles applies to a revealed tile and will check the adjacent
		tiles for flags.
		If the number of adjacent flags is >= the number on the tile it will reveal
		all adjacent tiles without a flag.
	*/
	ProcessAdjacentTiles(rowIndex int, colIndex int) ([]int, error)
	/*
		Stats returns information about the current game.
	*/
	Stats() stats
}
