package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type HelpBar struct {
	*gocui.Gui
	label    string
	handlers Handlers
	*Position
	*Attributes
	shortcuts map[string]string
}

func NewHelpBar(gui *gocui.Gui, label string) *HelpBar {
	maxX, maxY := gui.Size()

	h := &HelpBar{
		Gui:   gui,
		label: label,
		Position: &Position{
			X: 2,
			Y: maxY - 3,
			W: maxX - 3,
			H: maxY - 1,
		},
		Attributes: &Attributes{
			textColor: gocui.ColorWhite | gocui.AttrBold,
			// textBgColor: gocui.ColorBlue,
			hlColor:   gocui.ColorBlue | gocui.AttrBold,
			hlBgColor: gocui.ColorWhite,
		},
		handlers: make(Handlers),
		shortcuts: map[string]string{
			"F5": "Fetch",
			"F6": "Request Head",
			"F7": "Response Head",
			"F8": "Pick Method",
		},
	}

	return h
}

func (h *HelpBar) Draw() {
	v, err := h.Gui.SetView(h.label, h.X, h.Y, h.W, h.H)

	if err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}

		v.Frame = true
		v.FgColor = h.textColor
		v.BgColor = h.textBgColor
		v.SelFgColor = h.hlColor
		v.SelBgColor = h.hlBgColor

		for key, value := range h.shortcuts {
			display := "\t\t" + key + ": " + value + "\t\t\t"
			fmt.Fprint(v, display)
		}
	}
}
