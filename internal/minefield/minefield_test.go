package minefield_test

import (
	"sort"
	"testing"

	"github.com/pedrohenriques/go-minesweeper/internal/minefield"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type tile struct {
	revealed      bool
	hasMine       bool
	hasFlag       bool
	adjacentMines int
}

type minefieldTestSuite struct {
	suite.Suite
	sut               minefield.IMinefield
	sutArgs           *minefield.MinefieldConfig
	expectedMinefield []tile
}

/*
validateMinefield compares each tile in the SUT against the expected minefield
and required the content of each tile to match.
*/
func (suite *minefieldTestSuite) validateMinefield() {
	for rIndex := 0; rIndex < suite.sutArgs.NumRows; rIndex++ {
		for cIndex := 0; cIndex < suite.sutArgs.NumCols; cIndex++ {
			tIndex := rIndex*suite.sutArgs.NumCols + cIndex

			tile, err := suite.sut.Tile(rIndex, cIndex)

			require.Nil(suite.T(), err)
			require.Equalf(suite.T(), suite.expectedMinefield[tIndex].adjacentMines, tile.AdjacentMines(), "row index: %v | col index: %v", rIndex, cIndex)
			require.Equalf(suite.T(), suite.expectedMinefield[tIndex].hasFlag, tile.HasFlag(), "row index: %v | col index: %v", rIndex, cIndex)
			require.Equalf(suite.T(), suite.expectedMinefield[tIndex].hasMine, tile.HasMine(), "row index: %v | col index: %v", rIndex, cIndex)
			require.Equalf(suite.T(), suite.expectedMinefield[tIndex].revealed, tile.Revealed(), "row index: %v | col index: %v", rIndex, cIndex)
		}
	}
}

func (suite *minefieldTestSuite) SetupTest() {
	suite.sutArgs = &minefield.MinefieldConfig{
		NumRows:  10,
		NumCols:  11,
		NumMines: 20,
		Seed:     "hello",
	}

	/*
		0  0  0  0  2  X  X  3  1  1  0
		1  1  0  0  2  X  X  3  X  2  1
		X  1  0  1  2  3  2  3  3  5  X
		1  1  0  1  X  1  0  1  X  X  X
		0  0  0  1  1  1  0  1  2  3  2
		0  0  0  1  1  1  0  0  0  1  1
		0  0  0  1  X  1  1  2  2  2  X
		1  1  0  1  1  1  1  X  X  2  1
		X  2  1  1  1  1  2  3  2  2  1
		X  2  1  X  1  1  X  1  0  1  X

		0  0  0  0  2  _  _  _  _  _  _
		1  1  0  0  2  _  _  _  _  _  _
		_  1  0  1  2  _  _  _  _  _  _
		1  1  0  1  _  _  _  _  _  _  _
		0  0  0  1  _  _  _  _  _  _  _
		0  0  0  1  _  _  _  _  _  _  _
		0  0  0  1  _  _  _  _  _  _  _
		1  1  0  1  _  _  _  _  _  _  _
		_  2  1  1  _  _  _  _  _  _  _
		_  _  _  _  _  _  _  _  _  _  _
	*/
	suite.expectedMinefield = []tile{
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 2, revealed: true},
		{adjacentMines: 3, hasMine: true},
		{adjacentMines: 3, hasMine: true},
		{adjacentMines: 3},
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 0},

		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 2, revealed: true},
		{adjacentMines: 3, hasMine: true},
		{adjacentMines: 3, hasMine: true},
		{adjacentMines: 3},
		{adjacentMines: 0, hasMine: true},
		{adjacentMines: 2},
		{adjacentMines: 1},

		{adjacentMines: 0, hasMine: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 2, revealed: true},
		{adjacentMines: 3},
		{adjacentMines: 2},
		{adjacentMines: 3},
		{adjacentMines: 3},
		{adjacentMines: 5},
		{adjacentMines: 2, hasMine: true},

		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, hasMine: true},
		{adjacentMines: 1},
		{adjacentMines: 0},
		{adjacentMines: 1},
		{adjacentMines: 1, hasMine: true},
		{adjacentMines: 3, hasMine: true},
		{adjacentMines: 2, hasMine: true},

		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 0},
		{adjacentMines: 1},
		{adjacentMines: 2},
		{adjacentMines: 3},
		{adjacentMines: 2},

		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 0},
		{adjacentMines: 0},
		{adjacentMines: 0},
		{adjacentMines: 1},
		{adjacentMines: 1},

		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, hasMine: true},
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 2},
		{adjacentMines: 2},
		{adjacentMines: 2},
		{adjacentMines: 0, hasMine: true},

		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 0, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 1, hasMine: true},
		{adjacentMines: 1, hasMine: true},
		{adjacentMines: 2},
		{adjacentMines: 1},

		{adjacentMines: 1, hasMine: true},
		{adjacentMines: 2, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1, revealed: true},
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 2},
		{adjacentMines: 3},
		{adjacentMines: 2},
		{adjacentMines: 2},
		{adjacentMines: 1},

		{adjacentMines: 1, hasMine: true},
		{adjacentMines: 2},
		{adjacentMines: 1},
		{adjacentMines: 0, hasMine: true},
		{adjacentMines: 1},
		{adjacentMines: 1},
		{adjacentMines: 0, hasMine: true},
		{adjacentMines: 1},
		{adjacentMines: 0},
		{adjacentMines: 1},
		{adjacentMines: 0, hasMine: true},
	}

	suite.sut = minefield.Generate(*suite.sutArgs)
}

func (suite *minefieldTestSuite) TestColsReturnsTheNumberOfColumnsInTheMinefield() {
	require.Equal(suite.T(), 11, suite.sut.Cols())
}

func (suite *minefieldTestSuite) TestRowsReturnsTheNumberOfRowsInTheMinefield() {
	require.Equal(suite.T(), 10, suite.sut.Rows())
}

func (suite *minefieldTestSuite) TestMinesReturnsTheNumberOfMinesInTheMinefield() {
	require.Equal(suite.T(), 20, suite.sut.Mines())
}

func (suite *minefieldTestSuite) TestTileReturnsTheExpectedTile() {
	tile, error := suite.sut.Tile(1, 0)

	require.Nil(suite.T(), error)
	require.Equal(suite.T(), 1, tile.AdjacentMines())
	require.Equal(suite.T(), false, tile.HasFlag())
	require.Equal(suite.T(), false, tile.HasMine())
	require.Equal(suite.T(), true, tile.Revealed())
}

func (suite *minefieldTestSuite) TestTileReturnsAnErrorIfTheRequestedTileDoesNotExist() {
	_, error := suite.sut.Tile(100, 0)

	require.NotNil(suite.T(), error)
	require.Equal(suite.T(), "Tile not found for row index '100' and col index '0'", error.Error())
}

func (suite *minefieldTestSuite) TestRevealTileUpdatesTheMinefieldToSetTheRequestedTileAsRevealed() {
	suite.expectedMinefield[3*suite.sutArgs.NumCols+4].revealed = true

	_, error := suite.sut.RevealTile(3, 4)

	require.Nil(suite.T(), error)
	suite.validateMinefield()
}

func (suite *minefieldTestSuite) TestRevealTileReturnsTheIndexesOfTheTilesThatWereRevealed() {
	tileIndexes := []int{
		3*suite.sutArgs.NumCols + 4,
	}

	revealedIndexes, error := suite.sut.RevealTile(3, 4)
	sort.Ints(revealedIndexes)

	require.EqualValues(suite.T(), tileIndexes, revealedIndexes)
	require.Nil(suite.T(), error)
}

func (suite *minefieldTestSuite) TestRevealTileIfTheRequestedTileIsEmptyItRevealsTheEntireTilePatch() {
	suite.expectedMinefield[2*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[2*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[2*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[3*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[3*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[3*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+8].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+9].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+8].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+9].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+8].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+9].revealed = true

	_, error := suite.sut.RevealTile(4, 6)

	require.Nil(suite.T(), error)
	suite.validateMinefield()
}

func (suite *minefieldTestSuite) TestRevealTileIfATileHasAFlagItDoesNotRevealIt() {
	suite.sut.ToggleFlag(3, 6)
	suite.expectedMinefield[3*suite.sutArgs.NumCols+6].hasFlag = true

	suite.expectedMinefield[2*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[2*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[2*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[3*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[3*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+8].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+9].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+8].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+9].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+8].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+9].revealed = true

	_, error := suite.sut.RevealTile(4, 6)

	require.Nil(suite.T(), error)
	suite.validateMinefield()
}

func (suite *minefieldTestSuite) TestRevealTileReturnsAnErrorIfTheRequestedTileDoesNotExist() {
	_, error := suite.sut.RevealTile(100, 0)

	require.NotNil(suite.T(), error)
	require.Equal(suite.T(), "Tile not found for row index '100' and col index '0'", error.Error())
}

func (suite *minefieldTestSuite) TestToggleFlagSetsTheRequestedTileAsHavingAFlag() {
	suite.expectedMinefield[3*suite.sutArgs.NumCols+6].hasFlag = true

	require.Nil(suite.T(), suite.sut.ToggleFlag(3, 6))
	suite.validateMinefield()
}

func (suite *minefieldTestSuite) TestToggleFlagIfTheTileHasAFlagItSetsTheRequestedTileAsNotHavingAFlag() {
	require.Nil(suite.T(), suite.sut.ToggleFlag(0, 0))
	suite.validateMinefield()
}

func (suite *minefieldTestSuite) TestToggleFlagIfTheTileIsRevealedItDoesNotSetTheRequestedTileAsHavingAFlag() {
	suite.expectedMinefield[3*suite.sutArgs.NumCols+6].hasFlag = true
	require.Nil(suite.T(), suite.sut.ToggleFlag(3, 6))
	suite.validateMinefield()

	suite.expectedMinefield[3*suite.sutArgs.NumCols+6].hasFlag = false
	require.Nil(suite.T(), suite.sut.ToggleFlag(3, 6))
	suite.validateMinefield()
}

func (suite *minefieldTestSuite) TestToggleFlagReturnsAnErrorIfTheRequestedTileDoesNotExist() {
	error := suite.sut.ToggleFlag(100, 0)

	require.NotNil(suite.T(), error)
	require.Equal(suite.T(), "Tile not found for row index '100' and col index '0'", error.Error())
}

func (suite *minefieldTestSuite) TestProcessAdjacentTilesRevealsTheAdjacentTilesWithoutAFlag() {
	suite.sut.ToggleFlag(3, 8)
	suite.expectedMinefield[3*suite.sutArgs.NumCols+8].hasFlag = true

	suite.sut.RevealTile(4, 7)
	suite.expectedMinefield[4*suite.sutArgs.NumCols+7].revealed = true

	suite.expectedMinefield[2*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[2*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[2*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[3*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[3*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[3*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+8].revealed = true
	suite.expectedMinefield[4*suite.sutArgs.NumCols+9].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+8].revealed = true
	suite.expectedMinefield[5*suite.sutArgs.NumCols+9].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+5].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+6].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+7].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+8].revealed = true
	suite.expectedMinefield[6*suite.sutArgs.NumCols+9].revealed = true

	_, error := suite.sut.ProcessAdjacentTiles(4, 7)

	require.Nil(suite.T(), error)
	suite.validateMinefield()
}

func (suite *minefieldTestSuite) TestProcessAdjacentTilesReturnsTheIndexesOfTheRevealedTiles() {
	suite.sut.ToggleFlag(3, 8)
	suite.sut.RevealTile(4, 7)

	revealedIndexes, error := suite.sut.ProcessAdjacentTiles(4, 7)
	sort.Ints(revealedIndexes)

	tileIndexes := []int{
		2*suite.sutArgs.NumCols + 5,
		2*suite.sutArgs.NumCols + 6,
		2*suite.sutArgs.NumCols + 7,
		3*suite.sutArgs.NumCols + 5,
		3*suite.sutArgs.NumCols + 6,
		3*suite.sutArgs.NumCols + 7,
		4*suite.sutArgs.NumCols + 5,
		4*suite.sutArgs.NumCols + 6,
		4*suite.sutArgs.NumCols + 8,
		4*suite.sutArgs.NumCols + 9,
		5*suite.sutArgs.NumCols + 5,
		5*suite.sutArgs.NumCols + 6,
		5*suite.sutArgs.NumCols + 7,
		5*suite.sutArgs.NumCols + 8,
		5*suite.sutArgs.NumCols + 9,
		6*suite.sutArgs.NumCols + 5,
		6*suite.sutArgs.NumCols + 6,
		6*suite.sutArgs.NumCols + 7,
		6*suite.sutArgs.NumCols + 8,
		6*suite.sutArgs.NumCols + 9,
	}

	require.Nil(suite.T(), error)
	require.EqualValues(suite.T(), tileIndexes, revealedIndexes)
}

func (suite *minefieldTestSuite) TestProcessAdjacentTilesIfTheRequestedTileHasAPatchOnTheOtherSideOfTheBoardItDoesNotRevealIt() {
	suite.sut.ToggleFlag(2, 0)
	suite.expectedMinefield[2*suite.sutArgs.NumCols+0].hasFlag = true

	_, error := suite.sut.ProcessAdjacentTiles(1, 0)

	require.Nil(suite.T(), error)
	suite.validateMinefield()
}

func (suite *minefieldTestSuite) TestProcessAdjacentTilesIfTheNumberOfAdjacentFlagsIsLowerThanTheTileNumberItShouldNotRevealAnyTiles() {
	_, error := suite.sut.ProcessAdjacentTiles(3, 3)

	require.Nil(suite.T(), error)
	suite.validateMinefield()
}

func (suite *minefieldTestSuite) TestProcessAdjacentTilesIfTheRequestedTileIsNotRevealedItShouldNotRevealAnyTiles() {
	_, error := suite.sut.ProcessAdjacentTiles(4, 6)

	require.Nil(suite.T(), error)
	suite.validateMinefield()
}

func (suite *minefieldTestSuite) TestProcessAdjacentTilesReturnsAnErrorIfTheRequestedTileDoesNotExist() {
	_, error := suite.sut.ProcessAdjacentTiles(100, 0)

	require.NotNil(suite.T(), error)
	require.Equal(suite.T(), "Tile not found for row index '100' and col index '0'", error.Error())
}

func (suite *minefieldTestSuite) TestStatsReturnsTheExpectedObject() {
	suite.sut.RevealTile(2, 0)
	suite.sut.RevealTile(9, 0)
	suite.sut.RevealTile(9, 1)
	suite.sut.RevealTile(9, 2)
	suite.sut.RevealTile(9, 3)
	suite.sut.ToggleFlag(0, 8)
	suite.sut.ToggleFlag(0, 9)
	suite.sut.ToggleFlag(0, 10)

	type stats struct {
		NumTilesRevealed int
		NumMinesRevealed int
		NumFlags         int
	}
	expected := stats{
		NumTilesRevealed: 37 + 5,
		NumMinesRevealed: 3,
		NumFlags:         3,
	}

	actual := suite.sut.Stats()
	require.Equal(suite.T(), expected.NumFlags, actual.NumFlags)
	require.Equal(suite.T(), expected.NumMinesRevealed, actual.NumMinesRevealed)
	require.Equal(suite.T(), expected.NumTilesRevealed, actual.NumTilesRevealed)
}

func TestMinefieldFunctionalitySuite(t *testing.T) {
	suite.Run(t, new(minefieldTestSuite))
}
