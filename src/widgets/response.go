package widgets

import (
	"github.com/jroimartin/gocui"
)

// Body Widget
type BodyWidget struct {
  *gocui.Gui
  label string
  handlers Handlers
  *Position
  *Attributes
	name string
	x, y int
	w, h int
	body string
}

func NewBodyWidget(gui *gocui.Gui, label string) *BodyWidget {

    maxX, _ := gui.Size()
    return &BodyWidget{
          Gui: gui,
          label: label,
          Position: &Position{
           X: 2,
           Y: 7,
           H: 20,
           W: maxX - 5,
          },
          Attributes: &Attributes{
             textColor: gocui.ColorWhite | gocui.AttrBold,
             textBgColor: gocui.ColorBlack,
             hlColor: gocui.ColorBlue | gocui.AttrBold,
             hlBgColor: gocui.ColorWhite,
          },
    }
}

func (b *BodyWidget) Draw() {
	v, err := b.Gui.SetView(b.label, b.X, b.Y, b.X+b.W, b.Y+b.H)

	if err != nil {
		if err != gocui.ErrUnknownView {
       panic(err)
		}

    v.Frame = true
		v.Editable = true
    v.FgColor = b.textColor
    v.BgColor = b.textBgColor
    v.SelFgColor = b.hlColor
    v.SelBgColor = b.hlBgColor
		//fmt.Fprintf(v, w.body)
	}
}
