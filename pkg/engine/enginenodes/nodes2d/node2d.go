package nodes2d

import (
	"mygame/pkg/engine/nodes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Node2d struct {
	nodes.Node
	LocalPosition  rl.Vector2
	GlobalPosition rl.Vector2
}

func Init(position rl.Vector2) *Node2d {
	return &Node2d{
		LocalPosition:  position,
		GlobalPosition: position,
	}
}

func (n *Node2d) Process(delta float32) {}
func (n *Node2d) Input()                {}
func (n *Node2d) Draw()                 {}
func (n *Node2d) Destroy()              {}
