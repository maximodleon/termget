package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

// URL Widget
type URLWidget struct {
	name string
	x, y int
	w, h int
	body string
}

type URLEditor struct {
	Insert bool
}

func (url *URLEditor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0, false)
	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0, false)
	}
}

func NewURLWidget(name string, x, y int, body string) *URLWidget {
	return &URLWidget{name: name, x: x, y: y, h: 2, w: 100, body: body}
}

func (w *URLWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true
		v.SetCursor(len(w.body), 0)
		v.Editor = &URLEditor{}
		v.Title = "URL"

		if _, err := g.SetCurrentView(w.name); err != nil {
			return err
		}

		fmt.Fprintf(v, w.body)
	}

	return nil
}

// Methods/Verbs widget
type Attributes struct {
	textColor gocui.Attribute
	textBgColor gocui.Attribute
	hlColor gocui.Attribute
	hlBgColor gocui.Attribute
}

type MethodsWidget struct {
	name          string
	x, y          int
	w, h          int
	methods       []string
	currentMethod int
	listColor *Attributes
}

var verbs = []string{"GET", "PUT", "PATCH", "POST", "DELETE"}

func (w *MethodsWidget) addAttribute (textColor, textBgColor, hlColor, hlBgColor gocui.Attribute) *MethodsWidget {
	w.listColor = &Attributes{
		textColor: textColor,
		textBgColor: textBgColor,
		hlColor: hlColor,
		hlBgColor: hlBgColor,
	}

	return w
}

func (w *MethodsWidget) cursorUp(g *gocui.Gui, v *gocui.View) error {
	maxOptions := len(w.methods)

	if maxOptions == 0 {
		return nil
	}

	v.Highlight = false
	next := w.currentMethod - 1
	if next < 0 {
		next = 0
	}

	w.currentMethod = next
	v, _ = g.SetCurrentView(w.methods[next])
	v.Highlight = true

	return nil
}

func (w *MethodsWidget) cursorDown(g *gocui.Gui, v *gocui.View) error {
	maxOptions := len(w.methods)

	if maxOptions == 0 {
		return nil
	}

	v.Highlight = false
	next := w.currentMethod + 1
	if next >= maxOptions {
		next = w.currentMethod
	}

	w.currentMethod = next

	v, _ = g.SetCurrentView(w.methods[next])

	v.Highlight = true

	return nil
}

func (w *MethodsWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.w, w.h)
	w.methods = verbs

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}


		v.Title = "Methods"
		v.Highlight = true
		y := w.y
		h := w.y + 2
		for _, method := range w.methods {

			if v, err := g.SetView(method, w.x, y, w.w, h); err != nil {
				if err != gocui.ErrUnknownView {
					return err
				}

				v.Frame = false
				v.SelFgColor = w.listColor.textColor
				v.SelBgColor = w.listColor.textBgColor
				v.FgColor = w.listColor.hlColor
				v.BgColor = w.listColor.hlBgColor
				if err := g.SetKeybinding(v.Name(), gocui.KeyArrowDown, gocui.ModNone, w.cursorDown); err != nil {
					log.Panicln(err)
				}

				if err := g.SetKeybinding(v.Name(), gocui.KeyArrowUp, gocui.ModNone, w.cursorUp); err != nil {
					log.Panicln(err)
				}

				fmt.Fprint(v, method)
			}
			y++
			h++

		}

		v, _ := g.SetCurrentView(w.methods[w.currentMethod])
		v.Highlight = true

		if err := g.SetKeybinding(w.name, gocui.KeyArrowDown, gocui.ModNone, w.cursorDown); err != nil {
			log.Panicln(err)
		}

		if err := g.SetKeybinding(w.name, gocui.KeyArrowUp, gocui.ModNone, w.cursorUp); err != nil {
			log.Panicln(err)
		}
	}

	return nil
}

func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	url := NewURLWidget("url", 20, 1, "http://localhost/v1")
	methods := &MethodsWidget{name: "methods", x: 1, y: 1, h: 7, w: 10}
	methods.addAttribute(gocui.ColorBlack, gocui.ColorWhite, gocui.ColorBlack, gocui.ColorGreen)
	g.SetManager(methods, url)
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
