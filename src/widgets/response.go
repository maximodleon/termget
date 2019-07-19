package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

// Body Widget
type BodyWidget struct {
	name string
	x, y int
	w, h int
	body string
}

func NewBodyWidget(name string, x, y int, body string) *BodyWidget {
    return &BodyWidget{name: name, x: x, y: y, h: 20, w: 100, body: body }
}


func (w *BodyWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true
		v.SetCursor(len(w.body), 0)
		v.Title = "Body"

		if _, err := g.SetCurrentView(w.name); err != nil {
			return err
		}

		fmt.Fprintf(v, w.body)
	}

	return nil
}
