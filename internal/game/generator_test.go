package game_test

import (
	"testing"

	"github.com/pedrohenriques/go-minesweeper/internal/game"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type generatorTestSuite struct {
	suite.Suite
}

func (suite *generatorTestSuite) TestItReturnsAGameWithTheCorrectNumberOfMines() {
	config := game.GameConfig{
		NumMines: 3,
		NumRows:  10,
		NumCols:  10,
	}

	require.Equal(suite.T(), 3, game.Generate(config).Config().NumMines)
}

func (suite *generatorTestSuite) TestItReturnsAGameWithTheCorrectNumberOfRows() {
	config := game.GameConfig{
		NumRows: 1,
		NumCols: 4,
	}

	require.Equal(suite.T(), 1, game.Generate(config).Config().NumRows)
}

func (suite *generatorTestSuite) TestItReturnsAGameWithTheCorrectNumberOfCols() {
	config := game.GameConfig{
		NumCols: 7,
		NumRows: 1,
	}

	require.Equal(suite.T(), 7, game.Generate(config).Config().NumCols)
}

func (suite *generatorTestSuite) TestItReturnsAGameWithAStateOfZero() {
	config := game.GameConfig{
		NumCols:  10,
		NumRows:  10,
		NumMines: 10,
		Lives:    1,
	}

	require.Equal(suite.T(), 0, game.Generate(config).State())
}

func TestGeneratorSuite(t *testing.T) {
	suite.Run(t, new(generatorTestSuite))
}
