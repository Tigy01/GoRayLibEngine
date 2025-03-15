package nodes2d

import (
	"github.com/Tigy01/GoRayLibEngine/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Node2d struct {
	LocalPosition  rl.Vector2
	GlobalPosition rl.Vector2
}

func Init(position rl.Vector2) *Node2d {
	return &Node2d{
		LocalPosition:  position,
		GlobalPosition: rl.NewVector2(0, 0),
	}
}

func (n Node2d) Process(delta float32) {
}
func (n Node2d) Input()   {}
func (n Node2d) Draw()    {}
func (n Node2d) Destroy() {}

func UpdateScenePositions(currentScene *scenes.Scene) {
	startPosition := rl.NewVector2(0, 0)
	node2dCount := 0

	for _, child := range (*currentScene).GetChildrenTree().Children {
		if node2dCount > 1 {
			panic("Node2d Count May Not Exceed one per scene instance")
		}

		if node2d, ok := (*child.Value).(*Node2d); ok {
			node2d.GlobalPosition = rl.Vector2Add(node2d.LocalPosition, startPosition)
			startPosition = node2d.GlobalPosition
		}
		node2dCount += 1
	}

	for _, child := range (*currentScene).GetChildrenTree().Children {
		if _, ok := (*child.Value).(*Node2d); !ok {
			UpdateTreePositions(child, startPosition)
		}
	}
}

func UpdateTreePositions(tree *scenes.Tree, startPosition rl.Vector2) {
	for _, subtree := range tree.Children {
		if node2d, ok := (*subtree.Value).(*Node2d); ok {
			node2d.GlobalPosition = rl.Vector2Add(startPosition, node2d.LocalPosition)
			startPosition = node2d.GlobalPosition
		}
	}
	for _, subtree := range tree.Children {
		if _, ok := (*subtree.Value).(*Node2d); !ok {
			UpdateTreePositions(subtree, startPosition)
		}
	}
}
