package minefield_test

import (
	"testing"

	"github.com/pedrohenriques/go-minesweeper/internal/minefield"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type generatorTestSuite struct {
	suite.Suite
}

func (suite *generatorTestSuite) TestItReturnsAMinefieldInstanceWithTheExpectedColsValue() {
	args := &minefield.MinefieldConfig{
		NumCols: 5,
		NumRows: 1,
	}

	require.Equal(suite.T(), 5, minefield.Generate(*args).Cols())
}

func (suite *generatorTestSuite) TestItReturnsAMinefieldInstanceWithTheExpectedRowsValue() {
	args := &minefield.MinefieldConfig{
		NumRows: 3,
		NumCols: 1,
	}

	require.Equal(suite.T(), 3, minefield.Generate(*args).Rows())
}

func (suite *generatorTestSuite) TestItReturnsAMinefieldInstanceWithTheExpectedMinesValue() {
	args := &minefield.MinefieldConfig{
		NumMines: 9,
		NumCols:  10,
		NumRows:  10,
	}

	require.Equal(suite.T(), 9, minefield.Generate(*args).Mines())
}

func (suite *generatorTestSuite) TestItReturnsAMinefieldInstanceWithTheExpectedTilesValue() {
	args := &minefield.MinefieldConfig{
		NumCols:  9,
		NumRows:  9,
		NumMines: 10,
		Seed:     "hello",
	}

	minefield := minefield.Generate(*args)

	type tile struct {
		revealed      bool
		hasMine       bool
		hasFlag       bool
		adjacentMines int
	}

	/*
		1  1  1  0  0  0  0  0  0
		1  X  1  0  0  0  1  1  1
		2  2  2  0  0  0  1  X  1
		1  X  1  0  0  1  2  3  2
		1  1  1  0  0  2  X  4  X
		0  0  0  0  0  2  X  X  3
		0  1  1  1  0  1  3  X  3
		0  1  X  1  0  0  1  2  X
		0  1  1  1  0  0  0  1  1

		_  _  1  0  0  0  0  0  0
		_  _  1  0  0  0  1  1  1
		_  _  2  0  0  0  1  _  _
		_  _  1  0  0  1  2  _  _
		1  1  1  0  0  2  _  _  _
		0  0  0  0  0  2  _  _  _
		0  1  1  1  0  1  3  _  _
		0  1  _  1  0  0  1  2  _
		0  1  _  1  0  0  0  1  _
	*/
	expected := []tile{
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 1, revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},

		{adjacentMines: 1},
		{hasMine: true},
		{adjacentMines: 1, revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},

		{adjacentMines: 2},
		{adjacentMines: 2},
		{adjacentMines: 2, revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{hasMine: true},
		{adjacentMines: 1},

		{adjacentMines: 1},
		{hasMine: true},
		{adjacentMines: 1, revealed: true},
		{revealed: true},
		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 2, revealed: true},
		{adjacentMines: 3},
		{adjacentMines: 2},

		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{revealed: true},
		{revealed: true},
		{adjacentMines: 2, revealed: true},
		{hasMine: true, adjacentMines: 2},
		{adjacentMines: 4},
		{hasMine: true, adjacentMines: 1},

		{revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},
		{adjacentMines: 2, revealed: true},
		{hasMine: true, adjacentMines: 3},
		{hasMine: true, adjacentMines: 4},
		{adjacentMines: 3},

		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 3, revealed: true},
		{hasMine: true, adjacentMines: 3},
		{adjacentMines: 3},

		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{hasMine: true},
		{adjacentMines: 1, revealed: true},
		{revealed: true},
		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 2, revealed: true},
		{hasMine: true, adjacentMines: 1},

		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1},
		{adjacentMines: 1, revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1},
	}

	for rIndex := 0; rIndex < args.NumRows; rIndex++ {
		for cIndex := 0; cIndex < args.NumCols; cIndex++ {
			tIndex := rIndex*args.NumCols + cIndex

			tile, err := minefield.Tile(rIndex, cIndex)

			require.Nil(suite.T(), err)
			require.Equalf(suite.T(), expected[tIndex].adjacentMines, tile.AdjacentMines(), "row index: %v | col index: %v", rIndex, cIndex)
			require.Equalf(suite.T(), expected[tIndex].hasFlag, tile.HasFlag(), "row index: %v | col index: %v", rIndex, cIndex)
			require.Equalf(suite.T(), expected[tIndex].hasMine, tile.HasMine(), "row index: %v | col index: %v", rIndex, cIndex)
			require.Equalf(suite.T(), expected[tIndex].revealed, tile.Revealed(), "row index: %v | col index: %v", rIndex, cIndex)
		}
	}
}

func (suite *generatorTestSuite) TestIfTheFirstPatchToRevealIsTooSmallItSearchesForAnotherPatch() {
	args := &minefield.MinefieldConfig{
		NumCols:  9,
		NumRows:  9,
		NumMines: 10,
		Seed:     "hellos",
	}

	minefield := minefield.Generate(*args)

	type tile struct {
		revealed      bool
		hasMine       bool
		hasFlag       bool
		adjacentMines int
	}

	/*
		0  1  1  1  0  0  0  1  1
		1  3  X  2  0  0  0  1  X
		X  3  X  2  0  0  0  1  1
		1  2  1  1  0  0  0  1  1
		2  2  1  0  0  0  0  1  X
		X  X  1  0  0  1  1  2  1
		2  3  2  1  0  1  X  1  0
		0  1  X  1  0  1  1  2  1
		0  1  1  1  0  0  0  1  X

		_  _  _  1  0  0  0  1  _
		_  _  _  2  0  0  0  1  _
		_  _  _  2  0  0  0  1  _
		_  _  1  1  0  0  0  1  _
		_  _  1  0  0  0  0  1  _
		_  _  1  0  0  1  1  2  _
		_  _  2  1  0  1  _  _  _
		_  _  _  1  0  1  1  2  _
		_  _  _  1  0  0  0  1  _
	*/
	expected := []tile{
		{adjacentMines: 0},
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 1, revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1},

		{adjacentMines: 1},
		{adjacentMines: 3},
		{adjacentMines: 1, hasMine: true},
		{adjacentMines: 2, revealed: true},
		{revealed: true},
		{revealed: true},
		{revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, hasMine: true},

		{adjacentMines: 0, hasMine: true},
		{adjacentMines: 3},
		{adjacentMines: 1, hasMine: true},
		{adjacentMines: 2, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1},

		{adjacentMines: 1},
		{adjacentMines: 2},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1},

		{adjacentMines: 2},
		{adjacentMines: 2},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, hasMine: true},

		{adjacentMines: 1, hasMine: true},
		{adjacentMines: 1, hasMine: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 2, revealed: true},
		{adjacentMines: 1},

		{adjacentMines: 2},
		{adjacentMines: 3},
		{adjacentMines: 2, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, hasMine: true},
		{adjacentMines: 1},
		{adjacentMines: 0},

		{adjacentMines: 0},
		{adjacentMines: 1},
		{adjacentMines: 0, hasMine: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 2, revealed: true},
		{adjacentMines: 1},

		{adjacentMines: 0},
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, hasMine: true},
	}

	for rIndex := 0; rIndex < args.NumRows; rIndex++ {
		for cIndex := 0; cIndex < args.NumCols; cIndex++ {
			tIndex := rIndex*args.NumCols + cIndex

			tile, err := minefield.Tile(rIndex, cIndex)

			require.Nil(suite.T(), err)
			require.Equalf(suite.T(), expected[tIndex].adjacentMines, tile.AdjacentMines(), "row index: %v | col index: %v", rIndex, cIndex)
			require.Equalf(suite.T(), expected[tIndex].hasFlag, tile.HasFlag(), "row index: %v | col index: %v", rIndex, cIndex)
			require.Equalf(suite.T(), expected[tIndex].hasMine, tile.HasMine(), "row index: %v | col index: %v", rIndex, cIndex)
			require.Equalf(suite.T(), expected[tIndex].revealed, tile.Revealed(), "row index: %v | col index: %v", rIndex, cIndex)
		}
	}
}

func TestGeneratorSuite(t *testing.T) {
	suite.Run(t, new(generatorTestSuite))
}
