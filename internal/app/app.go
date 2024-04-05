package app

import (
	"github.com/rivo/tview"

	"strconv"

	"github.com/bedlamzd/gameton/internal/api/client"
)

var app = tview.NewApplication()
var main = tview.NewGrid()

func Start() {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	var abob = client.GetUniverse()
	menu := tview.NewList().
		AddItem("Capacity X", strconv.Itoa(abob.Ship.CapacityX), 1, nil).
		AddItem("Capacity Y", strconv.Itoa(abob.Ship.CapacityY), 1, nil).
		AddItem("Fuel Used", strconv.Itoa(abob.Ship.FuelUsed), 1, nil).
		AddItem("Planet", abob.Ship.Planet.Name, 1, nil)

	sideBar := tview.NewBox()

	grid := tview.NewGrid().
		SetRows(1, 0).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive(client.GetUniverse().Name), 0, 0, 1, 3, 0, 0, false)

	grid.AddItem(menu, 1, 0, 1, 1, 0, 0, false).
		AddItem(main, 1, 1, 1, 1, 0, 0, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 0, false)

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
