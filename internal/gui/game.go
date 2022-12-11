package gui

import (
	"fmt"
	"log"
	"time"

	"github.com/pedrohenriques/go-minesweeper/internal/configs"
	"github.com/pedrohenriques/go-minesweeper/internal/game"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

/*
createGameGui generates the CanvasObject for the game screen.
*/
func createGameGui(game game.IGame, new func(), reset func(), onGameEnd func(state int)) fyne.CanvasObject {
	navContainer := buildNavContainer(new, reset)
	statsContainer, statsDataBinds := buildStatsContainer(game)
	boardContainer := buildBoardContainer(game, statsDataBinds, onGameEnd)

	return container.NewVBox(navContainer, statsContainer, boardContainer)
}

/*
clickHandler will call the Game's functionality to action the clicked tile, based on the type of click,
and will call all affected tile widgets to update their rendering.

clickType: 0 = primary | 1 = secondary | 2 = both
*/
func clickHandler(game game.IGame, tileWidgets *[]ITileWidget, clickType int, statsDataBinds *statsDataBinds, onGameEnd func(state int)) func(int, int) {
	return func(rowIndex int, colIndex int) {
		if game.State() != configs.StateOnGoing {
			return
		}

		var tileIndexes []int
		var err error

		switch clickType {
		case configs.PrimaryClick:
			tileIndexes, err = game.RevealTile(rowIndex, colIndex)
		case configs.SecondaryClick:
			gameConfig := game.Config()

			tileIndexes = []int{rowIndex*gameConfig.NumCols + colIndex}
			err = game.ToggleFlag(rowIndex, colIndex)
		case configs.BothClick:
			tileIndexes, err = game.ProcessAdjacentTiles(rowIndex, colIndex)
		}

		if err != nil {
			log.Println(err)
		}

		if game.State() != configs.StateOnGoing {
			tileIndexes = make([]int, game.Config().NumRows*game.Config().NumCols)
			for tileIndex := 0; tileIndex < game.Config().NumRows*game.Config().NumCols; tileIndex++ {
				tileIndexes = append(tileIndexes, tileIndex)
			}
			onGameEnd(game.State())
		}

		for _, tIndex := range tileIndexes {
			(*tileWidgets)[tIndex].updateWidget(game.State() != configs.StateOnGoing)
		}

		gameStats := game.Stats()

		err = statsDataBinds.livesLeft.Set(fmt.Sprint(gameStats.RemainingLives))
		if err != nil {
			log.Println(err)
		}

		err = statsDataBinds.minesLeft.Set(fmt.Sprint(gameStats.RemainingMines))
		if err != nil {
			log.Println(err)
		}
	}
}

/*
buildBoardContainer will create the container with the game statistics.
*/
func buildBoardContainer(game game.IGame, statsDataBinds *statsDataBinds, onGameEnd func(state int)) *fyne.Container {
	gameConfig := game.Config()

	boardContainer := container.NewGridWithColumns(gameConfig.NumCols)
	tileWidgets := make([]ITileWidget, gameConfig.NumRows*gameConfig.NumCols)

	primaryHandler := clickHandler(game, &tileWidgets, configs.PrimaryClick, statsDataBinds, onGameEnd)
	secondaryHandler := clickHandler(game, &tileWidgets, configs.SecondaryClick, statsDataBinds, onGameEnd)
	bothClickHandler := clickHandler(game, &tileWidgets, configs.BothClick, statsDataBinds, onGameEnd)

	for rowIndex := 0; rowIndex < gameConfig.NumRows; rowIndex++ {
		for colIndex := 0; colIndex < gameConfig.NumCols; colIndex++ {
			tile, _ := game.Tile(rowIndex, colIndex)

			widgetArgs := newTileWidgetArgs{
				rowIndex:       rowIndex,
				colIndex:       colIndex,
				tile:           tile,
				primaryClick:   primaryHandler,
				secondaryClick: secondaryHandler,
				bothClick:      bothClickHandler,
			}

			canvasObj, tileWidget := newTileWidget(widgetArgs)

			boardContainer.Add(canvasObj)
			tileWidgets[rowIndex*gameConfig.NumCols+colIndex] = tileWidget
		}
	}

	return boardContainer
}

/*
buildStatsContainer will create the container with the game statistics.
*/
func buildStatsContainer(game game.IGame) (*fyne.Container, *statsDataBinds) {
	gameStats := game.Stats()

	statsDataBinds := &statsDataBinds{
		minesLeft:   binding.NewString(),
		livesLeft:   binding.NewString(),
		timeElapsed: binding.NewString(),
	}

	var err error
	err = statsDataBinds.minesLeft.Set(fmt.Sprint(gameStats.RemainingMines))
	if err != nil {
		log.Println(err)
	}

	err = statsDataBinds.livesLeft.Set(fmt.Sprint(gameStats.RemainingLives))
	if err != nil {
		log.Println(err)
	}

	err = statsDataBinds.timeElapsed.Set("0s")
	if err != nil {
		log.Println(err)
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		var err error

		for timestamp := range ticker.C {
			if game.State() != configs.StateOnGoing {
				return
			}

			err = statsDataBinds.timeElapsed.Set(fmt.Sprint(timestamp.Sub(game.StartTime()).Truncate(time.Second)))
			if err != nil {
				log.Println(err)
			}
		}
	}()

	statsContainer := container.NewGridWithRows(1)

	minesLeftImg := canvas.NewImageFromResource(resourceMinePng)
	minesLeftImg.FillMode = canvas.ImageFillContain
	statsContainer.Add(minesLeftImg)
	statsContainer.Add(widget.NewLabelWithData(statsDataBinds.minesLeft))

	livesLeftImg := canvas.NewImageFromResource(resourceLivesPng)
	livesLeftImg.FillMode = canvas.ImageFillContain
	statsContainer.Add(livesLeftImg)
	statsContainer.Add(widget.NewLabelWithData(statsDataBinds.livesLeft))

	timerImg := canvas.NewImageFromResource(resourceTimerPng)
	timerImg.FillMode = canvas.ImageFillContain
	statsContainer.Add(timerImg)
	statsContainer.Add(widget.NewLabelWithData(statsDataBinds.timeElapsed))

	return statsContainer, statsDataBinds
}

/*
buildNavContainer will create the container with the navigation elements.
*/
func buildNavContainer(new func(), reset func()) *fyne.Container {
	navContainer := container.NewGridWithRows(1)

	navContainer.Add(widget.NewButton("New Game", new))
	navContainer.Add(widget.NewButton("Reset Game", reset))

	return navContainer
}
