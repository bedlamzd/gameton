package app

import (
	"github.com/rivo/tview"

	"strconv"

	"github.com/bedlamzd/gameton/internal/api/client"
)

var app = tview.NewApplication()
var main = tview.NewGrid().
	SetRows(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1).
	SetColumns(2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2).
	SetBorders(true)

var grid = tview.NewGrid()

func Start() {
	DrawShit()

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func DrawShit() {
	var abob = client.GetUniverse()
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	menu := tview.NewList().
		AddItem("Capacity X", strconv.Itoa(abob.Ship.CapacityX), 1, nil).
		AddItem("Capacity Y", strconv.Itoa(abob.Ship.CapacityY), 1, nil).
		AddItem("Fuel Used", strconv.Itoa(abob.Ship.FuelUsed), 1, nil).
		AddItem("Planet", abob.Ship.Planet.Name, 1, nil)

	button := tview.NewButton("Update").SetSelectedFunc(func() { DrawShit() })
	button.SetBorder(true).SetRect(0, 0, 5, 5)
	sideBar := button

	grid.SetRows(1, 0).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive(abob.Name), 0, 0, 1, 3, 0, 0, false)

	grid.AddItem(menu, 1, 0, 1, 1, 0, 0, false).
		AddItem(main, 1, 1, 1, 1, 0, 0, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 0, false)
}
