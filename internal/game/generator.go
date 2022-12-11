/*
Package game handles creating and managing the state of a game
*/
package game

import (
	"time"

	"github.com/pedrohenriques/go-minesweeper/internal/minefield"
)

type GameConfig struct {
	NumMines     int
	NumRows      int
	NumCols      int
	FlagsEnabled bool
	Lives        int
	Seed         string
}

/*
Generate creates a new game with the provided configuration and returns the
a game instance
*/
func Generate(args GameConfig) IGame {
	return &game{
		startTs:      time.Now(),
		numMines:     args.NumMines,
		numRows:      args.NumRows,
		numCols:      args.NumCols,
		flagsEnabled: args.FlagsEnabled,
		lives:        args.Lives,
		minefield: minefield.Generate(minefield.MinefieldConfig{
			NumCols:  args.NumCols,
			NumRows:  args.NumRows,
			NumMines: args.NumMines,
			Seed:     args.Seed,
		}),
	}
}
