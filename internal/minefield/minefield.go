package minefield

import "fmt"

// Error: The requested tile does not exist
type tileNotFoundError struct {
	RowIndex int
	ColIndex int
}

/*
Error prints the message for this error.
*/
func (e tileNotFoundError) Error() string {
	return fmt.Sprintf(
		"Tile not found for row index '%v' and col index '%v'",
		e.RowIndex, e.ColIndex)
}

// Minefield describes the content and layout of a Minefield board
type minefield struct {
	cols  int
	rows  int
	mines int
	tiles []tile
}

// Cols returns the number of columns in the minefield
func (minefield *minefield) Cols() int {
	return minefield.cols
}

// Rows returns the number of rows in the minefield
func (minefield *minefield) Rows() int {
	return minefield.rows
}

// Mines returns the number of mines in the minefield
func (minefield *minefield) Mines() int {
	return minefield.mines
}

// Tile returns the tile in the minefield, on the provided row and col index
func (minefield *minefield) Tile(rowIndex int, colIndex int) (ITile, error) {
	tileIndex := calcTileIndex(rowIndex, colIndex, minefield.cols)

	if tileIndex < 0 || tileIndex > len(minefield.tiles)-1 {
		return &tile{}, tileNotFoundError{
			RowIndex: rowIndex,
			ColIndex: colIndex,
		}
	}

	return &minefield.tiles[tileIndex], nil
}

/*
RevealTile reveals the requested tile.
If the tile is empty the patch it belongs to will be revealed.
*/
func (minefield *minefield) RevealTile(rowIndex int, colIndex int) ([]int, error) {
	_, error := minefield.Tile(rowIndex, colIndex)
	if error != nil {
		return nil, error
	}

	tilesToReveal := findTilePatch(minefield, calcTileIndex(rowIndex, colIndex, minefield.cols))
	revealedTiles := []int{}

	for _, tileIndex := range tilesToReveal {
		if minefield.tiles[tileIndex].hasFlag {
			continue
		}
		if minefield.tiles[tileIndex].revealed {
			continue
		}

		minefield.tiles[tileIndex].revealed = true
		revealedTiles = append(revealedTiles, tileIndex)
	}

	return revealedTiles, nil
}

/*
ToggleFlag flips the flag state for the requested tile.
*/
func (minefield *minefield) ToggleFlag(rowIndex int, colIndex int) error {
	tile, error := minefield.Tile(rowIndex, colIndex)
	if error != nil {
		return error
	}

	if tile.Revealed() {
		return nil
	}

	minefield.tiles[calcTileIndex(rowIndex, colIndex, minefield.cols)].hasFlag = !tile.HasFlag()
	return nil
}

/*
ProcessAdjacentTiles applies to a revealed tile and will check the adjacent
tiles for flags.
If the number of adjacent flags is >= the number on the tile it will reveal
all adjacent tiles without a flag.
*/
func (minefield *minefield) ProcessAdjacentTiles(rowIndex int, colIndex int) ([]int, error) {
	reqTile, error := minefield.Tile(rowIndex, colIndex)

	if error != nil {
		return nil, error
	}

	if !reqTile.Revealed() {
		return nil, nil
	}

	adjacentFlags := 0
	tilesToReveal := []int{calcTileIndex(rowIndex, colIndex, minefield.cols)}
	revealedTiles := []int{}

	for rOffset := -1; rOffset <= 1; rOffset++ {
		rIndex := rowIndex + rOffset
		if rIndex < 0 || rIndex > minefield.rows-1 {
			continue
		}

		for cOffset := -1; cOffset <= 1; cOffset++ {
			cIndex := colIndex + cOffset
			if cIndex < 0 || cIndex > minefield.cols-1 {
				continue
			}

			tile, error := minefield.Tile(rIndex, cIndex)
			if error != nil {
				continue
			}
			if tile.Revealed() {
				continue
			}
			if tile.HasFlag() {
				adjacentFlags++
				continue
			}

			tilesToReveal = append(tilesToReveal, rIndex*minefield.cols+cIndex)
		}
	}

	if reqTile.AdjacentMines() > adjacentFlags {
		return nil, nil
	}

	for _, tileIndex := range tilesToReveal {
		tileRowIndex := tileIndex / minefield.cols
		tileColIndex := tileIndex % minefield.cols

		tileIndexes, _ := minefield.RevealTile(tileRowIndex, tileColIndex)
		revealedTiles = append(revealedTiles, tileIndexes...)
	}

	return uniqueTileIndexes(revealedTiles), nil
}

/*
Stats returns statistics about a minefield.
*/
func (minefield *minefield) Stats() stats {
	stats := new(stats)

	for _, tile := range minefield.tiles {
		if tile.hasFlag {
			stats.NumFlags++
		}
		if tile.revealed {
			stats.NumTilesRevealed++

			if tile.hasMine {
				stats.NumMinesRevealed++
			}
		}
	}

	return *stats
}

// Tile describes the information of a specific tile on the board
type tile struct {
	revealed      bool
	hasMine       bool
	hasFlag       bool
	adjacentMines int
}

/*
Revealed returns true if the tile is revealed and false otherwise.
*/
func (tile *tile) Revealed() bool {
	return tile.revealed
}

/*
HasMine returns true if the tile is has a mine and false otherwise.
*/
func (tile *tile) HasMine() bool {
	return tile.hasMine
}

/*
HasFlag returns true if the tile has a flag and false otherwise.
*/
func (tile *tile) HasFlag() bool {
	return tile.hasFlag
}

/*
AdjacentMines returns the number of mines in adjacent tiles.
*/
func (tile *tile) AdjacentMines() int {
	return tile.adjacentMines
}

// Stats contains statistics about a minefield.
type stats struct {
	NumTilesRevealed int
	NumMinesRevealed int
	NumFlags         int
}
