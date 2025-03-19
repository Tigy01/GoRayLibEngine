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

func (s AnimatedSprite) getSourceRec() rl.Rectangle {
	frameSize := float32(s.Texture.Width) / float32(s.NumberOfFrames)
	framePosition := frameSize * float32(s.CurrentFrame)

	return rl.NewRectangle(framePosition, 0, frameSize, float32(s.Texture.Height))
}

func (s AnimatedSprite) Input() {}

func (s AnimatedSprite) Process(delta float32) {}

func (s *AnimatedSprite) GetChildrenTree() scenes.Hierarchy {
	return scenes.Hierarchy{
		Scene: scenes.AsScenePtr(s),
		Children: []*scenes.Tree{
			{
				Value:    nodes.GetNodePtr(s.Sprite),
				Children: s.Sprite.GetChildrenTree().Children,
			},
		},
	}
}
