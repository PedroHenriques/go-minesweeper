package game_test

import (
	"sort"
	"testing"
	"time"

	"github.com/pedrohenriques/go-minesweeper/internal/game"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type tile struct {
	revealed      bool
	hasMine       bool
	hasFlag       bool
	adjacentMines int
}

type gameTestSuite struct {
	suite.Suite
	sut               game.IGame
	sutArgs           *game.GameConfig
	expectedMinefield []tile
}

func (suite *gameTestSuite) SetupTest() {
	suite.sutArgs = &game.GameConfig{
		NumRows:      10,
		NumCols:      11,
		NumMines:     20,
		FlagsEnabled: true,
		Lives:        2,
		Seed:         "hello",
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

	suite.sut = game.Generate(*suite.sutArgs)
}

/*
validateMinefield compares each tile in the SUT against the expected minefield
and requires the content of each tile to match.
*/
func (suite *gameTestSuite) validateMinefield() {
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

/*
solveGame reveals the necessary tiles to get the game into a "win" state.
*/
func (suite *gameTestSuite) solveGame() {
	suite.sut.ToggleFlag(0, 5)
	suite.sut.ToggleFlag(0, 6)
	suite.sut.ToggleFlag(1, 5)
	suite.sut.ToggleFlag(1, 6)
	suite.sut.ToggleFlag(1, 8)
	suite.sut.ToggleFlag(2, 0)
	suite.sut.ToggleFlag(2, 10)
	suite.sut.ToggleFlag(3, 4)
	suite.sut.ToggleFlag(3, 8)
	suite.sut.ToggleFlag(3, 9)
	suite.sut.ToggleFlag(3, 10)
	suite.sut.ToggleFlag(6, 4)
	suite.sut.ToggleFlag(6, 10)
	suite.sut.ToggleFlag(7, 7)
	suite.sut.ToggleFlag(7, 8)
	suite.sut.ToggleFlag(8, 0)
	suite.sut.ToggleFlag(9, 0)
	suite.sut.ToggleFlag(9, 3)
	suite.sut.ToggleFlag(9, 6)
	suite.sut.ToggleFlag(9, 10)
	suite.sut.ProcessAdjacentTiles(2, 4)
	suite.sut.ProcessAdjacentTiles(3, 5)
	suite.sut.ProcessAdjacentTiles(6, 5)
	suite.sut.ProcessAdjacentTiles(7, 5)
	suite.sut.ProcessAdjacentTiles(8, 2)
	suite.sut.ProcessAdjacentTiles(8, 5)
	suite.sut.ProcessAdjacentTiles(8, 6)
	suite.sut.ProcessAdjacentTiles(2, 7)
	suite.sut.ProcessAdjacentTiles(1, 7)
	suite.sut.ProcessAdjacentTiles(2, 8)
	suite.sut.ProcessAdjacentTiles(1, 9)
	suite.sut.ProcessAdjacentTiles(4, 9)
	suite.sut.ProcessAdjacentTiles(9, 7)
	suite.sut.ProcessAdjacentTiles(8, 9)
}

func (suite *gameTestSuite) TestConfigReturnsTheExpectedObject() {
	type config struct {
		NumMines int
		NumRows  int
		NumCols  int
	}

	expected := config{
		NumMines: 20,
		NumRows:  10,
		NumCols:  11,
	}
	actual := suite.sut.Config()

	require.Equal(suite.T(), expected.NumCols, actual.NumCols)
	require.Equal(suite.T(), expected.NumMines, actual.NumMines)
	require.Equal(suite.T(), expected.NumRows, actual.NumRows)
}

func (suite *gameTestSuite) TestStartTimeReturnsAPlausibleTime() {
	beforeTime := time.Now()

	actual := suite.sut.StartTime().Unix()
	require.GreaterOrEqual(suite.T(), actual, beforeTime.Unix())
	require.LessOrEqual(suite.T(), actual, time.Now().Unix())
}

func (suite *gameTestSuite) TestStateReturnsZero() {
	require.Equal(suite.T(), 0, suite.sut.State())
}

func (suite *gameTestSuite) TestStateReturnsTwoIfTheNumberOfRevealedMinesIsGreaterThanOrEqualToTheNumberOfLives() {
	suite.sut.RevealTile(2, 0)
	suite.sut.RevealTile(0, 5)

	require.Equal(suite.T(), 2, suite.sut.State())
}

func (suite *gameTestSuite) TestStateReturnsOneIfTheNumberOfRevealedTilesIsEqualToTheNumberOfNonMineTilesInTheBoard() {
	suite.solveGame()

	require.Equal(suite.T(), 1, suite.sut.State())
}

func (suite *gameTestSuite) TestStateReturnsZeroIfTheNumberOfRevealedTilesIsEqualToTheNumberOfNonMineTilesInTheBoardButThereAreMinesTilesRevealed() {
	suite.sut.RevealTile(0, 7)
	suite.sut.RevealTile(0, 8)
	suite.sut.RevealTile(0, 9)
	suite.sut.RevealTile(0, 10)
	suite.sut.RevealTile(1, 7)
	suite.sut.RevealTile(1, 9)
	suite.sut.RevealTile(1, 10)
	suite.sut.RevealTile(2, 5)
	suite.sut.RevealTile(2, 6)
	suite.sut.RevealTile(2, 7)
	suite.sut.RevealTile(2, 8)
	suite.sut.RevealTile(2, 9)
	suite.sut.RevealTile(3, 5)
	suite.sut.RevealTile(3, 6) // has mine
	suite.sut.RevealTile(3, 7)
	suite.sut.RevealTile(4, 4)
	suite.sut.RevealTile(4, 5)
	suite.sut.RevealTile(4, 6)
	suite.sut.RevealTile(4, 7)
	suite.sut.RevealTile(4, 8)
	suite.sut.RevealTile(4, 9)
	suite.sut.RevealTile(4, 10)
	suite.sut.RevealTile(5, 4)
	suite.sut.RevealTile(5, 5)
	suite.sut.RevealTile(5, 6)
	suite.sut.RevealTile(5, 7)
	suite.sut.RevealTile(5, 8)
	suite.sut.RevealTile(5, 9)
	suite.sut.RevealTile(5, 10)
	suite.sut.RevealTile(6, 5)
	suite.sut.RevealTile(6, 6)
	suite.sut.RevealTile(6, 7)
	suite.sut.RevealTile(6, 8)
	suite.sut.RevealTile(6, 9)
	suite.sut.RevealTile(7, 4)
	suite.sut.RevealTile(7, 5)
	suite.sut.RevealTile(7, 6)
	suite.sut.RevealTile(7, 9)
	suite.sut.RevealTile(7, 10)
	suite.sut.RevealTile(8, 4)
	suite.sut.RevealTile(8, 5)
	suite.sut.RevealTile(8, 6)
	suite.sut.RevealTile(8, 7)
	suite.sut.RevealTile(8, 8)
	suite.sut.RevealTile(8, 9)
	suite.sut.RevealTile(8, 10)
	suite.sut.RevealTile(9, 1)
	suite.sut.RevealTile(9, 2)
	suite.sut.RevealTile(9, 4)
	suite.sut.RevealTile(9, 5)
	suite.sut.RevealTile(9, 7)
	suite.sut.RevealTile(9, 8)
	suite.sut.RevealTile(9, 9)

	require.Equal(suite.T(), 1, suite.sut.State())
}

func (suite *gameTestSuite) TestTileReturnsTheExpectedTileObject() {
	tile, error := suite.sut.Tile(2, 0)

	require.Nil(suite.T(), error)
	require.Equal(suite.T(), 0, tile.AdjacentMines())
	require.Equal(suite.T(), false, tile.HasFlag())
	require.Equal(suite.T(), true, tile.HasMine())
	require.Equal(suite.T(), false, tile.Revealed())
}

func (suite *gameTestSuite) TestTileReturnsAnErrorIfTheRequestedTileDoesNotExist() {
	_, error := suite.sut.Tile(200, 0)

	require.NotNil(suite.T(), error)
}

func (suite *gameTestSuite) TestRevealTileSetsTheRequestedTileToRevealed() {
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

func (suite *gameTestSuite) TestRevealTileReturnsTheIndexesOfTheTilesThatWereRevealed() {
	tileIndexes := []int{
		2*suite.sutArgs.NumCols + 5,
		2*suite.sutArgs.NumCols + 6,
		2*suite.sutArgs.NumCols + 7,
		3*suite.sutArgs.NumCols + 5,
		3*suite.sutArgs.NumCols + 6,
		3*suite.sutArgs.NumCols + 7,
		4*suite.sutArgs.NumCols + 5,
		4*suite.sutArgs.NumCols + 6,
		4*suite.sutArgs.NumCols + 7,
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

	revealedIndexes, error := suite.sut.RevealTile(4, 6)
	sort.Ints(revealedIndexes)

	require.EqualValues(suite.T(), tileIndexes, revealedIndexes)
	require.Nil(suite.T(), error)
}

func (suite *gameTestSuite) TestRevealTileReturnsAnErrorIfTheRequestedTileDoesNotExist() {
	_, error := suite.sut.RevealTile(40, 6)

	require.NotNil(suite.T(), error)
}

func (suite *gameTestSuite) TestRevealTileSetsTheGameEndTsFieldIfTheGameEndedWithALossAfterTheTileReveal() {
	require.Equal(suite.T(), true, suite.sut.Stats().EndTime.IsZero())

	suite.sut.RevealTile(2, 0)
	suite.sut.RevealTile(8, 0)

	require.Equal(suite.T(), false, suite.sut.Stats().EndTime.IsZero())
}

func (suite *gameTestSuite) TestRevealTileSetsTheGameEndTsFieldIfTheGameEndedWithAWinAfterTheTileReveal() {
	require.Equal(suite.T(), true, suite.sut.Stats().EndTime.IsZero())

	suite.solveGame()

	require.Equal(suite.T(), false, suite.sut.Stats().EndTime.IsZero())
}

func (suite *gameTestSuite) TestRevealTileDoesNotChangeTheRequestedTileRevealedConditionIfTheGameIsInAnEndState() {
	suite.solveGame()

	suite.sut.RevealTile(2, 0)

	tile, _ := suite.sut.Tile(2, 0)
	require.Equal(suite.T(), false, tile.Revealed())
}

func (suite *gameTestSuite) TestRevealTileDoesNotChangeTheGameEndTsIfTheGameIsInAnEndState() {
	suite.solveGame()
	expected := suite.sut.Stats().EndTime

	suite.sut.RevealTile(2, 0)

	require.Equal(suite.T(), expected, suite.sut.Stats().EndTime)
}

func (suite *gameTestSuite) TestToggleFlagSetsTheRequestedTileToRevealed() {
	suite.expectedMinefield[4*suite.sutArgs.NumCols+6].hasFlag = true

	require.Nil(suite.T(), suite.sut.ToggleFlag(4, 6))
	suite.validateMinefield()
}

func (suite *gameTestSuite) TestToggleFlagReturnsAnErrorIfTheRequestedTileDoesNotExist() {
	require.NotNil(suite.T(), suite.sut.ToggleFlag(40, 6))
}

func (suite *gameTestSuite) TestToggleFlagDoesNotChangeTheRequestedTileFlagConditionIfTheGameIsInAnEndState() {
	suite.solveGame()
	tile, _ := suite.sut.Tile(2, 0)

	suite.sut.ToggleFlag(2, 0)

	require.Equal(suite.T(), true, tile.HasFlag())
}

func (suite *gameTestSuite) TestToggleFlagDoesNotChangeTheRequestedTileFlagConditionIfTheFlagsEnabledConfigIsFalse() {
	suite.sutArgs = &game.GameConfig{
		NumRows:      10,
		NumCols:      11,
		NumMines:     20,
		FlagsEnabled: false,
		Lives:        2,
		Seed:         "hello",
	}
	suite.sut = game.Generate(*suite.sutArgs)

	suite.solveGame()
	tile, _ := suite.sut.Tile(2, 0)

	suite.sut.ToggleFlag(2, 0)

	require.Equal(suite.T(), false, tile.HasFlag())
}

func (suite *gameTestSuite) TestProcessAdjacentTilesRevealsTheAdjacentTilesWithoutAFlag() {
	suite.sut.ToggleFlag(3, 4)
	suite.expectedMinefield[3*suite.sutArgs.NumCols+4].hasFlag = true

	suite.expectedMinefield[4*suite.sutArgs.NumCols+4].revealed = true

	_, error := suite.sut.ProcessAdjacentTiles(3, 3)

	require.Nil(suite.T(), error)
	suite.validateMinefield()
}

func (suite *gameTestSuite) TestProcessAdjacentTilesReturnsTheIndexesOfTheRevealedTiles() {
	suite.sut.ToggleFlag(3, 4)

	revealedIndexes, error := suite.sut.ProcessAdjacentTiles(3, 3)
	sort.Ints(revealedIndexes)

	tileIndexes := []int{
		4*suite.sutArgs.NumCols + 4,
	}

	require.Nil(suite.T(), error)
	require.EqualValues(suite.T(), tileIndexes, revealedIndexes)
}

func (suite *gameTestSuite) TestProcessAdjacentTilesRevealsTheAdjacentTilesWithoutAFlagIncludingAPatchOfEmptyTiles() {
	suite.sutArgs = &game.GameConfig{
		NumRows:      10,
		NumCols:      10,
		NumMines:     10,
		FlagsEnabled: true,
		Lives:        1,
		Seed:         "pedrohenriques",
	}
	suite.sut = game.Generate(*suite.sutArgs)

	suite.sut.ToggleFlag(7, 8)
	suite.sut.RevealTile(6, 8)
	revealedIndexes, error := suite.sut.ProcessAdjacentTiles(6, 8)
	sort.Ints(revealedIndexes)

	tileIndexes := []int{
		4*suite.sutArgs.NumCols + 7,
		4*suite.sutArgs.NumCols + 8,
		4*suite.sutArgs.NumCols + 9,
		5*suite.sutArgs.NumCols + 7,
		5*suite.sutArgs.NumCols + 8,
		5*suite.sutArgs.NumCols + 9,
		6*suite.sutArgs.NumCols + 7,
		6*suite.sutArgs.NumCols + 9,
		7*suite.sutArgs.NumCols + 9,
	}

	require.Nil(suite.T(), error)
	require.EqualValues(suite.T(), tileIndexes, revealedIndexes)
}

func (suite *gameTestSuite) TestProcessAdjacentTilesIfTheNumberOfAdjacentFlagsIsLowerThanTheTileNumberItShouldNotRevealAnyTiles() {
	_, error := suite.sut.ProcessAdjacentTiles(3, 3)

	require.Nil(suite.T(), error)
	suite.validateMinefield()
}

func (suite *gameTestSuite) TestProcessAdjacentTilesIfTheRequestedTileIsNotRevealedItShouldNotRevealAnyTiles() {
	_, error := suite.sut.ProcessAdjacentTiles(0, 6)

	require.Nil(suite.T(), error)
	suite.validateMinefield()
}

func (suite *gameTestSuite) TestProcessAdjacentTilesReturnsAnErrorIfTheRequestedTileDoesNotExist() {
	_, error := suite.sut.ProcessAdjacentTiles(100, 0)

	require.NotNil(suite.T(), error)
}

func (suite *gameTestSuite) TestProcessAdjacentTilesSetsTheGameEndTsFieldIfTheGameEndedWithALossAfterTheTileReveal() {
	require.Equal(suite.T(), true, suite.sut.Stats().EndTime.IsZero())

	suite.sut.ToggleFlag(4, 4)
	suite.sut.ProcessAdjacentTiles(4, 3)
	suite.sut.ToggleFlag(9, 0)
	suite.sut.ToggleFlag(9, 2)
	suite.sut.ProcessAdjacentTiles(8, 1)

	require.Equal(suite.T(), false, suite.sut.Stats().EndTime.IsZero())
}

func (suite *gameTestSuite) TestProcessAdjacentTilesSetsTheGameEndTsFieldIfTheGameEndedWithAWinAfterTheTileReveal() {
	require.Equal(suite.T(), true, suite.sut.Stats().EndTime.IsZero())

	suite.solveGame()

	require.Equal(suite.T(), false, suite.sut.Stats().EndTime.IsZero())
}

func (suite *gameTestSuite) TestProcessAdjacentTilesDoesNotChangeTheRequestedTileRevealedConditionIfTheGameIsInAnEndState() {
	suite.solveGame()

	suite.sut.ProcessAdjacentTiles(1, 0)

	tile, _ := suite.sut.Tile(1, 0)
	require.Equal(suite.T(), true, tile.Revealed())
}

func (suite *gameTestSuite) TestProcessAdjacentTilesDoesNotChangeTheGameEndTsIfTheGameIsInAnEndState() {
	suite.solveGame()
	expected := suite.sut.Stats().EndTime

	suite.sut.ProcessAdjacentTiles(1, 0)

	require.Equal(suite.T(), expected, suite.sut.Stats().EndTime)
}

func (suite *gameTestSuite) TestStatsReturnsTheExpectedObject() {
	type stats struct {
		StartTime      time.Time
		EndTime        time.Time
		RemainingMines int
		RemainingLives int
	}

	expected := stats{
		StartTime:      suite.sut.StartTime(),
		RemainingMines: 20,
		RemainingLives: 2,
	}
	actual := suite.sut.Stats()

	require.Equal(suite.T(), expected.StartTime, actual.StartTime)
	require.Equal(suite.T(), expected.EndTime, actual.EndTime)
	require.Equal(suite.T(), expected.RemainingMines, actual.RemainingMines)
	require.Equal(suite.T(), expected.RemainingLives, actual.RemainingLives)
}

func (suite *gameTestSuite) TestStatsReturnsTheExpectedObjectIfFlagsHaveBeenPlaced() {
	suite.sut.ToggleFlag(3, 4)

	type stats struct {
		StartTime      time.Time
		EndTime        time.Time
		RemainingMines int
		RemainingLives int
	}

	expected := stats{
		StartTime:      suite.sut.StartTime(),
		RemainingMines: 19,
		RemainingLives: 2,
	}
	actual := suite.sut.Stats()

	require.Equal(suite.T(), expected.StartTime, actual.StartTime)
	require.Equal(suite.T(), expected.EndTime, actual.EndTime)
	require.Equal(suite.T(), expected.RemainingMines, actual.RemainingMines)
	require.Equal(suite.T(), expected.RemainingLives, actual.RemainingLives)
}

func (suite *gameTestSuite) TestStatsReturnsTheExpectedObjectIfMinesHaveBeenRevealed() {
	suite.sut.RevealTile(2, 0)

	type stats struct {
		StartTime      time.Time
		EndTime        time.Time
		RemainingMines int
		RemainingLives int
	}

	expected := stats{
		StartTime:      suite.sut.StartTime(),
		RemainingMines: 20,
		RemainingLives: 1,
	}
	actual := suite.sut.Stats()

	require.Equal(suite.T(), expected.StartTime, actual.StartTime)
	require.Equal(suite.T(), expected.EndTime, actual.EndTime)
	require.Equal(suite.T(), expected.RemainingMines, actual.RemainingMines)
	require.Equal(suite.T(), expected.RemainingLives, actual.RemainingLives)
}

func TestGameFunctionalitySuite(t *testing.T) {
	suite.Run(t, new(gameTestSuite))
}
