package gui

import (
	"github.com/pedrohenriques/go-minesweeper/internal/configs"
	"github.com/pedrohenriques/go-minesweeper/internal/game"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type statsDataBinds struct {
	minesLeft   binding.String
	livesLeft   binding.String
	timeElapsed binding.String
}

/*
Run starts the GUI.
*/
func Run(configs *configs.Configs) {
	app := app.New()

	app.SetIcon(resourceFaviconPng)

	window := app.NewWindow("Main")
	window.SetMaster()

	guiChannel := make(chan string, 1)
	go processGuiEvent(configs, &guiChannel, &window)

	guiChannel <- "setup"

	window.ShowAndRun()
}

/*
processGuiEvent listens for events on the provided channel and handles them.
*/
func processGuiEvent(config *configs.Configs, guiChannel *chan string, window *fyne.Window) {
	var gameConfig game.GameConfig
	var gameInstance game.IGame

	for event := range *guiChannel {
		if event == "setup" {
			(*window).SetContent(createSetupGui(config, func(config game.GameConfig) {
				gameConfig = config
				gameInstance = game.Generate(gameConfig)
				*guiChannel <- "game"
			}))
		} else if event == "game" {
			(*window).SetContent(createGameGui(gameInstance,
				func() {
					*guiChannel <- "setup"
				},
				func() {
					gameInstance = game.Generate(gameConfig)
					*guiChannel <- "game"
				},
				func(state int) {
					var labelText string
					switch state {
					case configs.StateWin:
						labelText = "You have won!"
					case configs.StateLoss:
						labelText = "You have lost!"
					}

					var popupWidget *widget.PopUp
					container := container.NewGridWithColumns(1)
					container.Add(widget.NewLabel(labelText))
					container.Add(widget.NewButton("Close", func() {
						popupWidget.Hide()
					}))

					popupWidget = widget.NewModalPopUp(container, (*window).Canvas())
					popupWidget.Show()
				}))
		}
	}
}
