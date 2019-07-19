package main

import (
	"github.com/jroimartin/gocui"
  "widgets"
	"log"
)

func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	url := widgets.NewURLWidget("url", 20, 1, "http://localhost/v1")
	methods := &widgets.MethodsWidget{Name: "methods", X: 1, Y: 1, H: 7, W: 10}
	methods.AddAttribute(gocui.ColorBlack, gocui.ColorWhite, gocui.ColorBlack, gocui.ColorGreen)
  body := widgets.NewBodyWidget("body", 20, 4, "")
	g.SetManager(methods, body, url)
	g.Cursor = true

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, toggleView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func toggleView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "methods" {
		_, err := g.SetCurrentView("url")
		return err
	}
	_, err := g.SetCurrentView("methods")
	return err
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
