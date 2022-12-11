package gui

import (
	"fmt"
	"time"

	"github.com/pedrohenriques/go-minesweeper/internal/configs"
	"github.com/pedrohenriques/go-minesweeper/internal/minefield"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Declare conformity with fyne.CanvasObject interface
var _ fyne.CanvasObject = (*tileButton)(nil)

const noButtonClick = -1
const clickDelayMS = 200

type clickTimer struct {
	timer  time.Timer
	button int
}

type ITileWidget interface {
	updateWidget(forceReveal bool)
}

// tileButton represents a revealed tile on the board.
type tileButton struct {
	widget.Button
	rowIndex       int
	colIndex       int
	tile           minefield.ITile
	primaryClick   func(rowIndex int, colIndex int)
	secondaryClick func(rowIndex int, colIndex int)
	bothClick      func(rowIndex int, colIndex int)
	clickTimer     *clickTimer
}

/*
updateWidget sets the state of the widget based on the linked tile state.
*/
func (t *tileButton) updateWidget(forceReveal bool) {
	if t.tile.Revealed() || forceReveal {
		if t.tile.HasFlag() {
			if !t.tile.HasMine() {
				t.SetIcon(resourceIncorrectFlagPng)
			}
		} else if t.tile.HasMine() {
			t.SetIcon(resourceMinePng)
		} else if t.tile.AdjacentMines() > 0 {
			t.SetIcon(nil)
			t.SetText(fmt.Sprint(fmt.Sprint(t.tile.AdjacentMines())))
		} else {
			t.SetIcon(nil)
			t.Disable()
		}
	} else if t.tile.HasFlag() {
		t.SetIcon(resourceFlagPng)
	} else {
		t.SetIcon(nil)
	}
}

/*
Tapped handles LMB clicks.
*/
func (t *tileButton) Tapped(_ *fyne.PointEvent) {
	if t.clickTimer.button == noButtonClick {
		t.clickTimer.button = configs.PrimaryClick
		t.clickTimer.timer = *time.AfterFunc(clickDelayMS*time.Millisecond, func() {
			t.clickTimer.button = noButtonClick
			t.primaryClick(t.rowIndex, t.colIndex)
		})
	} else if t.clickTimer.button != configs.PrimaryClick {
		t.clickTimer.timer.Stop()
		t.clickTimer.button = noButtonClick

		t.bothClick(t.rowIndex, t.colIndex)
	}
}

/*
TappedSecondary handles RMB clicks.
*/
func (t *tileButton) TappedSecondary(_ *fyne.PointEvent) {
	if t.clickTimer.button == noButtonClick {
		t.clickTimer.button = configs.SecondaryClick
		t.clickTimer.timer = *time.AfterFunc(clickDelayMS*time.Millisecond, func() {
			t.clickTimer.button = noButtonClick
			t.secondaryClick(t.rowIndex, t.colIndex)
		})
	} else if t.clickTimer.button != configs.SecondaryClick {
		t.clickTimer.timer.Stop()
		t.clickTimer.button = noButtonClick

		t.bothClick(t.rowIndex, t.colIndex)
	}
}

type newTileWidgetArgs struct {
	rowIndex       int
	colIndex       int
	tile           minefield.ITile
	primaryClick   func(rowIndex int, colIndex int)
	secondaryClick func(rowIndex int, colIndex int)
	bothClick      func(rowIndex int, colIndex int)
}

/*
newTileWidget creates a tileWidget instance.
*/
func newTileWidget(args newTileWidgetArgs) (fyne.CanvasObject, ITileWidget) {
	button := &tileButton{
		rowIndex:       args.rowIndex,
		colIndex:       args.colIndex,
		tile:           args.tile,
		primaryClick:   args.primaryClick,
		secondaryClick: args.secondaryClick,
		bothClick:      args.bothClick,
		clickTimer: &clickTimer{
			button: noButtonClick,
		},
	}
	button.ExtendBaseWidget(button)

	button.updateWidget(false)

	return button, button
}
