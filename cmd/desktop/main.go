package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("erp-opity")
	w := a.NewWindow("ERP Opity")

	menuItems := []string{"Clientes", "Produtos", "Vendas", "Relatórios", "Configurações"}

	sidebar := widget.NewList(
		func() int { return len(menuItems) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(menuItems[i])
		},
	)

	content := widget.NewLabel("Bem-vindo ao ERP Opity Desktop!")
	mainArea := container.NewMax(content)

	sidebar.OnSelected = func(id widget.ListItemID) {
		switch id {
		case 0:
			mainArea.Objects = []fyne.CanvasObject{widget.NewLabel("Clientes - Em breve CRUD completo!")}
		case 1:
			mainArea.Objects = []fyne.CanvasObject{widget.NewLabel("Produtos - Em breve CRUD completo!")}
		case 2:
			mainArea.Objects = []fyne.CanvasObject{widget.NewLabel("Vendas - Em breve PDV!")}
		case 3:
			mainArea.Objects = []fyne.CanvasObject{widget.NewLabel("Relatórios - Em breve dashboards!")}
		case 4:
			mainArea.Objects = []fyne.CanvasObject{widget.NewLabel("Configurações do sistema")}
		}
		mainArea.Refresh()
	}

	topbar := container.NewHBox(
		widget.NewLabelWithStyle("ERP Opity", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	)

	w.SetContent(
		container.NewBorder(
			topbar, nil, sidebar, nil, mainArea,
		),
	)
	w.Resize(fyne.NewSize(1200, 800))
	w.ShowAndRun()
}
