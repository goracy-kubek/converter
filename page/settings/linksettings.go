package settings

import (
	"converter/storage"
	"converter/view/cwidget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewLinkSettings() fyne.CanvasObject {
	return container.NewBorder(
		widget.NewLabel("Links settings"),
		nil, nil, nil,
		cwidget.NewUriList(storage.GetLinkStorage()),
	)
}
