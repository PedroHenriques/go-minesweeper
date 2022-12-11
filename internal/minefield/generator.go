/*
Package minefield handles all create and mutate operations on the minefield type
*/
package minefield

import (
	"math/rand"
	"time"
)

type MinefieldConfig struct {
	NumCols  int
	NumRows  int
	NumMines int
	Seed     string
}

/*
Generate creates a minefield, using the provided configuration, and returns
a Minefield
*/
func Generate(args MinefieldConfig) IMinefield {
	minefield := &minefield{
		cols:  args.NumCols,
		rows:  args.NumRows,
		mines: args.NumMines,
		tiles: make([]tile, args.NumRows*args.NumCols),
	}

	seedRng(args.Seed)

	populateMines(minefield, &args)

	revealInitialPatch(minefield)

	return minefield
}

/*
Configures the RNG with the provided seed
If a seed is not provided then the current unix timestamp will be used as seed
*/
func seedRng(seed string) {
	var convertedSeed int64 = time.Now().Unix()
	if seed != "" {
		var sumBytes int
		for _, v := range []byte(seed) {
			sumBytes += int(v)
		}
		convertedSeed = int64(sumBytes)
	}
	rand.Seed(convertedSeed)
}

/*
Populated the minefield with mines and adds the numbers to adjacent tiles
*/
func populateMines(minefield *minefield, config *MinefieldConfig) {
	var numMineTiles int
	for numMineTiles < config.NumMines {
		tileIndex := rand.Intn(config.NumCols * config.NumRows)
		rowIndex := tileIndex / config.NumCols
		colIndex := tileIndex % config.NumCols

		if minefield.tiles[tileIndex].hasMine {
			continue
		}

		minefield.tiles[tileIndex].hasMine = true
		numMineTiles++

		for rowOffset := -1; rowOffset <= 1; rowOffset++ {
			rIndex := rowIndex + rowOffset
			if rIndex < 0 || rIndex > config.NumRows-1 {
				continue
			}

			for colOffset := -1; colOffset <= 1; colOffset++ {
				if rowOffset == 0 && colOffset == 0 {
					continue
				}

				cIndex := colIndex + colOffset
				if cIndex < 0 || cIndex > config.NumCols-1 {
					continue
				}

				minefield.tiles[rIndex*minefield.cols+cIndex].adjacentMines++
			}
		}
	}
}

/*
Reveals a patch of tiles, complying with the application configuration, to
facilitate the start of the game
*/
func revealInitialPatch(minefield *minefield) {
	iterations := 0
	for iterations <= initialPatchMaxIterations {
		focalTileIndex := rand.Intn(minefield.cols * minefield.rows)

		if minefield.tiles[focalTileIndex].adjacentMines != 0 && !minefield.tiles[focalTileIndex].hasMine {
			continue
		}
		iterations++

		tilesToReveal := findTilePatch(minefield, focalTileIndex)

		if float32(len(tilesToReveal)) < initialPatchMinCoverage*float32(minefield.rows+minefield.cols) {
			continue
		}

		for _, revealIndex := range tilesToReveal {
			minefield.tiles[revealIndex].revealed = true
		}

		break
	}
}
