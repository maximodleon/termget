package widgets
import "github.com/jroimartin/gocui"

type Key interface{}

type Handler func(g *gocui.Gui, v *gocui.View) error

type Handlers map[Key]Handler

type Position struct {
   X, Y int
   W, H int
}

type Attributes struct {
	textColor gocui.Attribute
	textBgColor gocui.Attribute
	hlColor gocui.Attribute
	hlBgColor gocui.Attribute
}

type Widget interface {
  GetName() string
  GetPosition() *Position
  Draw()
  Focus()
  UnFocus()
  Close()
}
