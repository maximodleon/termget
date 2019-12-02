package main

import (
	"github.com/jroimartin/gocui"
)

// URL Widget
type URLWidget struct {
	*gocui.Gui
	label    string
	handlers Handlers
	*Position
	*Attributes
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

func NewURLWidget(gui *gocui.Gui, label string) *URLWidget {

	maxX, _ := gui.Size()
	return &URLWidget{
		Gui:   gui,
		label: label,
		Position: &Position{
			X: 2,
			Y: 2,
			H: 4,
			W: maxX - 3,
		},
		Attributes: &Attributes{
			textColor: gocui.ColorWhite | gocui.AttrBold,
			// textBgColor: gocui.ColorBlue,
			hlColor:   gocui.ColorBlue | gocui.AttrBold,
			hlBgColor: gocui.ColorWhite,
		},
		handlers: make(Handlers),
	}
}

func (u *URLWidget) Draw() {
	v, err := u.Gui.SetView(u.label, u.X, u.Y, u.W, u.H)

	if err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}

		v.Frame = true
		v.Editable = true
		//		v.SetCursor(len(w.body), 0)
		v.Editor = &URLEditor{}
		v.Title = "URL"

		if _, err := u.Gui.SetCurrentView(u.label); err != nil {
			panic(err)
		}
	}
}
