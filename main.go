package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"httputils"
	"log"
	"widgets"
  "io/ioutil"
)

var methods widgets.MethodsWidget = widgets.MethodsWidget{Name: "methods", X: 1, Y: 1, H: 7, W: 10}

// TODO: Add shortcut to copy body results
// to clipboard
func main() {

	g, err := gocui.NewGui(gocui.Output256)

	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

//	url := widgets.NewURLWidget("url", 20, 1, "https://jsonplaceholder.typicode.com/todos")
	methods.AddAttribute(gocui.ColorBlack, gocui.ColorWhite, 0, 0)
//	body := widgets.NewBodyWidget("body", 20, 4, "")
  helper := widgets.NewHelpBar(g, "help")
  helper.Draw()
//	g.SetManager(&methods, body, url)
	g.Cursor = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	// TODO: add other key binding to quit functionality
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, toggleView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyF5, gocui.ModNone, displayRequestResults); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// TODO: add other views to toggle between when using
// the keyboard
func toggleView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "body" {
		_, err := g.SetCurrentView("url")
		g.Cursor = true
		return err
	}
  _, err := g.SetCurrentView("body")
	g.Cursor = false
	return err
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// TODO: Return correct value
func displayRequestResults(g *gocui.Gui, v *gocui.View) error {
	// TODO: handle error returnded by View function
	g.SetCurrentView("body")
  // TODO: create helper to get view and not repeate these lines?
	bodyView, err := g.View("body")
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	urlView, err := g.View("url")
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	method := methods.GetSelected()

	bodyView.Clear()
	// TODO: send body for POST, PATCH and PUT requests
	err, response := httputils.GetResponse(method, urlView.Buffer(), "")

	if err != nil {
		fmt.Fprint(bodyView, err)
		return nil
	}
  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)

  if err != nil {
   return err
  }

	fmt.Fprint(bodyView, string(body))
	//TODO: return correct value
	return nil
}
