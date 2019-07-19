package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

// Methods/Verbs widget
type Attributes struct {
	textColor gocui.Attribute
	textBgColor gocui.Attribute
	hlColor gocui.Attribute
	hlBgColor gocui.Attribute
}

type MethodsWidget struct {
	Name          string
	X, Y          int
	W, H          int
	Methods       []string
	currentMethod int
	listColor *Attributes
}

var verbs = []string{"GET", "PUT", "PATCH", "POST", "DELETE"}

func (w *MethodsWidget) AddAttribute (textColor, textBgColor, hlColor, hlBgColor gocui.Attribute) *MethodsWidget {
	w.listColor = &Attributes{
		textColor: textColor,
		textBgColor: textBgColor,
		hlColor: hlColor,
		hlBgColor: hlBgColor,
	}

	return w
}

func (w *MethodsWidget) cursorUp(g *gocui.Gui, v *gocui.View) error {
	maxOptions := len(w.Methods)

	if maxOptions == 0 {
		return nil
	}

	v.Highlight = false
	next := w.currentMethod - 1
	if next < 0 {
		next = 0
	}

	w.currentMethod = next
	v, _ = g.SetCurrentView(w.Methods[next])
	v.Highlight = true

	return nil
}

func (w *MethodsWidget) cursorDown(g *gocui.Gui, v *gocui.View) error {
	maxOptions := len(w.Methods)

	if maxOptions == 0 {
		return nil
	}

	v.Highlight = false
	next := w.currentMethod + 1
	if next >= maxOptions {
		next = w.currentMethod
	}

	w.currentMethod = next

	v, _ = g.SetCurrentView(w.Methods[next])

	v.Highlight = true

	return nil
}

func (w *MethodsWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.Name, w.X, w.Y, w.W, w.H)
	w.Methods = verbs

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}


		v.Title = "Methods"
		v.Highlight = true
		y := w.Y
		h := w.Y + 2
		for _, method := range w.Methods {

			if v, err := g.SetView(method, w.X, y, w.W, h); err != nil {
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

		v, _ := g.SetCurrentView(w.Methods[w.currentMethod])
		v.Highlight = true

		if err := g.SetKeybinding(w.Name, gocui.KeyArrowDown, gocui.ModNone, w.cursorDown); err != nil {
			log.Panicln(err)
		}

		if err := g.SetKeybinding(w.Name, gocui.KeyArrowUp, gocui.ModNone, w.cursorUp); err != nil {
			log.Panicln(err)
		}
	}

	return nil
}
