package debugUtils

import (
	"fmt"

	"github.com/pedrohenriques/go-minesweeper/internal/game"
	"github.com/pedrohenriques/go-minesweeper/internal/minefield"
)

/*
PrintMinefield prints in the console a visual representation of the minefield,
in its current state.
*/
func PrintMinefield(minefield minefield.IMinefield, revealAll bool) {
	fmt.Println("Minefield:")

	for rIndex := 0; rIndex < minefield.Rows(); rIndex++ {
		for cIndex := 0; cIndex < minefield.Cols(); cIndex++ {
			tile, _ := minefield.Tile(rIndex, cIndex)

			if tile.Revealed() || revealAll {
				if tile.HasMine() {
					fmt.Printf(" %v ", "X")
				} else {
					fmt.Printf(" %v ", tile.AdjacentMines())
				}
			} else {
				if tile.HasFlag() {
					fmt.Printf(" %v ", "F")
				} else {
					fmt.Printf(" %v ", "_")
				}
			}
		}
		fmt.Println("")
	}
}

/*
PrintMinefieldFromGame prints in the console a visual representation of the
minefield, in its current state, for the provided game.
*/
func PrintMinefieldFromGame(game game.IGame, revealAll bool) {
	fmt.Println("Minefield:")

	for rIndex := 0; rIndex < game.Config().NumRows; rIndex++ {
		for cIndex := 0; cIndex < game.Config().NumCols; cIndex++ {
			tile, _ := game.Tile(rIndex, cIndex)

			if tile.Revealed() || revealAll {
				if tile.HasMine() {
					fmt.Printf(" %v ", "X")
				} else {
					fmt.Printf(" %v ", tile.AdjacentMines())
				}
			} else {
				if tile.HasFlag() {
					fmt.Printf(" %v ", "F")
				} else {
					fmt.Printf(" %v ", "_")
				}
			}
		}
		fmt.Println("")
	}
}
