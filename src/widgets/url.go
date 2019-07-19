package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
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
