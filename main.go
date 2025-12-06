// package main

// import (
// 	"os"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/theme"
// 	"fyne.io/fyne/v2/widget"
// 	"fyne.io/fyne/v2/layout"
// )

// func init() {
// 	os.Setenv("FYNE_THEME", "dark")
// }

// func main() {
//     a := app.New()

//     w := a.NewWindow("Converter")
//     w.SetMaster()

// 	title :=  container.NewHBox(layout.NewSpacer(), widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {}))
// 	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
//     w.SetContent(container.NewVBox(title, widget.NewSeparator(), intro))

//     w.Resize(fyne.NewSize(640, 460))
//     w.ShowAndRun()
// }

// func main() {
//     a := app.New()

//     w := a.NewWindow("Converter")

//     w.SetMaster()

//     content := container.NewStack()
//     title := widget.NewLabel("Component name")
//     intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
//     intro.Wrapping = fyne.TextWrapWord

//     tutorial := container.NewBorder( container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)

//     split := container.NewHSplit(makeNav(), tutorial)
//     split.Offset = 0

//     w.SetContent(split)

//     w.Resize(fyne.NewSize(640, 460))
//     w.ShowAndRun()
// }

// func makeNav() fyne.CanvasObject {
// 	var menuItems = map[string][]string{
// 		"": {"welcome", "collections", "advanced"},
// 		"collections": {"list", "table"},
// 	}
//     tree := widget.NewTreeWithStrings(menuItems)

// 	tree.OnSelected =func(uid widget.TreeNodeID) {
// 		print(uid)
// 	}

//     return container.NewBorder(nil, nil, nil, nil, tree)
// }

package main

import (
	"converter/view/cwidget"
	// "image/color"

	// "image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"

	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func init() {
	os.Setenv("FYNE_THEME", "dark")
}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("TabContainer Widget")

	tabs := container.NewAppTabs(
		// container.NewTabItem("Converter", converter(window)),
		container.NewTabItem("Autoconverting", autoconverting()),
		// container.NewTabItem("Settings", widget.NewLabel("World!")),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	window.Resize(fyne.NewSize(640, 480))

	window.SetContent(tabs)
	window.ShowAndRun()
}

func autoconverting() fyne.CanvasObject {
	return container.NewVBox(
		container.NewGridWrap(
			fyne.NewSize(150, 50),
			widget.NewCheck("Scan after every new file", func(value bool) {println("Check set to", value)}),
			widget.NewCheck("Delete original", func(value bool) {println("Check set to", value)}),
			widget.NewCheck("Transfer original", func(value bool) {println("Check set to", value)}),
		),
	)

	// autoconverting periods
	// checking folder for updates
	// delete converts file
	// transfer files

	// adding folder
	// list of folders
}

func converter(w fyne.Window) fyne.CanvasObject {
	return container.New(
		layout.NewCustomPaddedLayout(20, 0, 0, 0),
		container.NewVBox(
			container.NewCenter(
				cwidget.NewDragAndDrop(w, func(u []fyne.URI) { print(u[0].Path()) }),
			),
		),
	)
}
