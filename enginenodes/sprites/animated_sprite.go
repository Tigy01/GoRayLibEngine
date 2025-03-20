package sprites

import (
	"github.com/Tigy01/GoRayLibEngine/nodes"
	"github.com/Tigy01/GoRayLibEngine/scenes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedSprite struct {
	*Sprite
	CurrentFrame   int
	NumberOfFrames int
}

func (AnimatedSprite) Init(sprite *Sprite, currentFrame, frameCount int) *AnimatedSprite {
	return &AnimatedSprite{
		Sprite:         sprite,
		NumberOfFrames: frameCount,
		CurrentFrame:   currentFrame % frameCount,
	}
}

func (s AnimatedSprite) Input() {}

func (s AnimatedSprite) Process(delta float32) {}

func (s *AnimatedSprite) Destroy() {}

func (s AnimatedSprite) Draw() {
	if !s.Sprite.Hidden {
		s.Sprite.Hidden = true
	}
	rl.DrawTexturePro(
		s.Texture,
		s.getSourceRec(),
		s.getDestRec(),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
}

func (s AnimatedSprite) getFrameSize() (frameSize float32) {
	frameSize = float32(s.Texture.Width) / float32(s.NumberOfFrames)
	return
}

func (s AnimatedSprite) getSourceRec() rl.Rectangle {
	framePosition := s.getFrameSize() * float32(s.CurrentFrame)
	return rl.NewRectangle(framePosition, 0, s.getFrameSize(), float32(s.Texture.Height))
}

func (s AnimatedSprite) getDestRec() rl.Rectangle {
	return rl.NewRectangle(
		s.Node2d.GlobalPosition.X,
		s.Node2d.GlobalPosition.Y,
		s.getFrameSize()*s.LocalScale.X,
		float32(s.Texture.Height)*s.LocalScale.Y,
	)
}

func (s *AnimatedSprite) GetChildrenTree() scenes.Hierarchy {
	return scenes.Hierarchy{
		Scene: scenes.AsScenePtr(s),
		Children: []*scenes.Tree{
			{
				Value:    nodes.AsNodePtr(s.Sprite),
				Children: s.Sprite.GetChildrenTree().Children,
			},
		},
	}
}

func (s *AnimatedSprite) Center() {
	s.Offset(
		rl.NewVector2(
			(0 - s.getFrameSize()/2 * s.LocalScale.X),
			(0 - float32(s.Texture.Height)/2 * s.LocalScale.Y),
		),
	)
}
