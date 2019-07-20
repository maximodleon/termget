package main

import (
	"github.com/jroimartin/gocui"
  "widgets"
	"log"
  "httputils"
  "fmt"
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

  if err := g.SetKeybinding("", gocui.KeyF5, gocui.ModNone, request); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func toggleView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "methods" {
		_, err := g.SetCurrentView("url")
    g.Cursor = true
		return err
	}
	_, err := g.SetCurrentView("methods")
  g.Cursor = false
	return err
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// TODO: Return correct value
func request (g *gocui.Gui, v *gocui.View) error {
        // TODO: handle error returnded by View function
        g.SetCurrentView("body")
        bodyView, _  := g.View("body")
        urlView, _ := g.View("url")

        bodyView.Clear()
        err, body := httputils.MakeRequest(urlView.Buffer())

        if err != nil {
          fmt.Fprint(bodyView, err)
          return nil
        }

        fmt.Fprint(bodyView, body)
        //TODO: return correct value
        return nil
}
