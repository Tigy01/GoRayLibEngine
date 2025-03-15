package nodes2d

import (
	"github.com/Tigy01/GoRayLibEngine/nodes"
	"github.com/Tigy01/GoRayLibEngine/scenes"

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

func UpdateScenePositions(currentScene scenes.Scene) {
	startPosition := rl.NewVector2(0, 0)
	node2dCount := 0
	for _, child := range currentScene.GetChildrenTree().Children {
		if node2dCount > 1 {
			panic("Node2d Count May Not Exceed one per scene instance")
		}
		if node2d, ok := (*child.Value).(*Node2d); ok {
			node2d.GlobalPosition = rl.Vector2Add(node2d.LocalPosition, startPosition)
			startPosition = node2d.GlobalPosition
		}
		UpdateTreePositions(child.Children, startPosition)
		node2dCount += 1
	}
}

func UpdateTreePositions(trees []*scenes.Tree, startPosition rl.Vector2) {
	for _, tree := range trees {
		if node2d, ok := (*tree.Value).(*Node2d); ok {
			node2d.GlobalPosition = rl.Vector2Add(startPosition, node2d.LocalPosition)
		}
	}
}
