package gui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pedrohenriques/go-minesweeper/internal/game"

	"github.com/pedrohenriques/go-minesweeper/internal/configs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*
createSetupGui generates the CanvasObject for the setup screen.
*/
func createSetupGui(config *configs.Configs, startGame func(config game.GameConfig)) fyne.CanvasObject {
	gameArgs := game.GameConfig{}

	container := container.NewGridWithColumns(1)

	container.Add(createDifficultySelect(config, func(option configs.SizeOption) {
		gameArgs.NumMines = option.NumMines
		gameArgs.NumRows = option.NumRows
		gameArgs.NumCols = option.NumCols
	}))

	container.Add(createFlagEnabledCheck(func(enabled bool) {
		gameArgs.FlagsEnabled = enabled
	}))

	container.Add(createNumLivesInput(func(value string) {
		if value == "" {
			return
		}

		numLives, error := strconv.Atoi(value)
		if error != nil {
			fmt.Printf("Error converting the number of lives to integer: %v\n", error)
			return
		}

		gameArgs.Lives = numLives
	}))

	container.Add(createSeedInput(func(value string) {
		gameArgs.Seed = value
	}))

	container.Add(widget.NewButton("Start Game", func() {
		startGame(gameArgs)
	}))

	return container
}

/*
createFlagEnabledCheck creates the CanvasObject with the flags enabled check.
*/
func createFlagEnabledCheck(callback func(checked bool)) fyne.CanvasObject {
	checkWidget := widget.NewCheck("Enable flags", callback)
	checkWidget.SetChecked(true)

	return checkWidget
}

/*
createDifficultySelect creates the CanvasObject with the board size options.
*/
func createDifficultySelect(config *configs.Configs, callback func(option configs.SizeOption)) fyne.CanvasObject {
	optionLabels := make([]string, len(config.SizeOptions))
	index := 0
	for key := range config.SizeOptions {
		optionLabels[index] = key
		index++
	}

	container := container.NewGridWithRows(1)

	container.Add(widget.NewLabel("Difficulty:"))
	selectWidget := widget.NewSelect(optionLabels, func(value string) {
		if option, ok := config.SizeOptions[value]; ok {
			callback(option)
		}
	})
	container.Add(selectWidget)

	selectWidget.SetSelectedIndex(0)

	return container
}

/*
createNumLivesInput creates the CanvasObject for the number of lives.
*/
func createNumLivesInput(callback func(value string)) fyne.CanvasObject {
	container := container.NewGridWithRows(1)

	container.Add(widget.NewLabel("Number of lives:"))

	inputWidget := widget.NewEntry()
	inputWidget.OnChanged = callback
	container.Add(inputWidget)

	inputWidget.SetText("1")

	return container
}

/*
createSeedInput creates the CanvasObject for the number of lives.
*/
func createSeedInput(callback func(value string)) fyne.CanvasObject {
	container := container.NewGridWithRows(1)

	container.Add(widget.NewLabel("Seed:"))

	inputWidget := widget.NewEntry()
	inputWidget.OnChanged = callback
	container.Add(inputWidget)

	inputWidget.SetText(strconv.FormatInt(time.Now().Unix(), 10))

	return container
}
