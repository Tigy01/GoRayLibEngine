package sprites

import (
	"github.com/Tigy01/GoRayLibEngine/enginenodes/nodes2d"
	"github.com/Tigy01/GoRayLibEngine/nodes"
	"github.com/Tigy01/GoRayLibEngine/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	scenes.Scene
	*nodes2d.Node2d
	texture rl.Texture2D
}

func (Sprite) Init(path string, position rl.Vector2) *Sprite {
	return &Sprite{
		texture: rl.LoadTexture(path),
		Node2d:  nodes2d.Init(position),
	}
}

func (s Sprite) Process(delta float32) {}

func (s *Sprite) Center() {
	w := s.texture.Width
	h := s.texture.Height

	s.Offset(
		rl.NewVector2(
			0-float32(w)/2,
			0-float32(h)/2,
		),
	)
}

func (s *Sprite) Offset(offset rl.Vector2) {
	s.Node2d.LocalPosition = rl.Vector2Add(s.Node2d.LocalPosition, offset)
}

func (s Sprite) Draw() {
	rl.DrawTexturePro(
		s.texture,
		s.getSourceRec(),
		s.getDestRec(),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
}

func (s *Sprite) GetChildrenTree() scenes.Hierarchy {
	return scenes.Hierarchy{
		Scene: scenes.AsScenePtr(s),
		Children: []*scenes.Tree{
			{
				Value: nodes.GetNodePtr(s.Node2d),
			},
		},
	}
}

func (s Sprite) getDestRec() rl.Rectangle {
	return rl.NewRectangle(
		s.Node2d.GlobalPosition.X, s.Node2d.GlobalPosition.Y, s.LocalScale.X, s.LocalScale.Y,
	)
}

func (s Sprite) getSourceRec() rl.Rectangle {
	return rl.NewRectangle(0, 0, float32(s.texture.Width), float32(s.texture.Height))
}

func (s *Sprite) Destroy() {
	rl.UnloadTexture(s.texture)
}

func (s Sprite) Input() {
}
