package backgrounds

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Background struct {
	Color color.RGBA
}

func (Background) Init(color color.RGBA) *Background {
	return &Background{
		Color: color,
	}
}

func (b Background) Process(delta float32) {}
func (b Background) Input()                {}
func (background Background) Draw() {
	rl.ClearBackground(background.Color)
}
func (b Background) Destroy() {}
