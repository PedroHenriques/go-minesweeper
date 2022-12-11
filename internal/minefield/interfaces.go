package minefield

type IMinefield interface {
	/*
		Cols returns the number of columns in the minefield.
	*/
	Cols() int
	/*
		Rows returns the number of rows in the minefield.
	*/
	Rows() int
	/*
		Mines returns the number of mines in the minefield.
	*/
	Mines() int
	/*
		Tile searches for the tile in the requested row and column.
		The row and column coordinates are zero-indexed.
	*/
	Tile(rowIndex int, colIndex int) (ITile, error)
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
		Stats returns statistics about a minefield.
	*/
	Stats() stats
}

type ITile interface {
	/*
		Revealed returns true if the tile is revealed and false otherwise.
	*/
	Revealed() bool
	/*
		HasMine returns true if the tile is has a mine and false otherwise.
	*/
	HasMine() bool
	/*
		HasFlag returns true if the tile has a flag and false otherwise.
	*/
	HasFlag() bool
	/*
		AdjacentMines returns the number of mines in adjacent tiles.
	*/
	AdjacentMines() int
}
