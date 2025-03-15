package engine

import (
	"myengine/pkg/engine/nodes"
	"myengine/pkg/engine/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var CurrentScene scenes.Scene

func EngineInit(windowSize rl.Vector2, windowTitle string) {
	rl.InitWindow(int32(windowSize.X), int32(windowSize.Y), "Test Raylib")
}

func EngineStart(scene scenes.Scene) {
	CurrentScene = scene
	rl.SetTargetFPS(60)

	EngineLoop()
}

func EngineStop() {
	(CurrentScene).GetChildrenTree().DoOnEveryNode(
		func(n *nodes.Node) {
			(*n).Destroy()
		},
	)
	rl.CloseWindow()
}

func EngineLoop() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		EngineUpdate()

		rl.EndDrawing()
	}
}

func EngineProcess() {
	rl.ClearBackground(rl.White)
	(CurrentScene).GetChildrenTree().DoOnEveryNode(
		func(n *nodes.Node) {
			(*n).Input()
			(*n).Process(rl.GetFrameTime())
			(*n).Draw()
		},
	)
}

func EngineUpdate() {
	EngineProcess()
}

func ChangeScene(scene scenes.Scene) {
	(CurrentScene).GetChildrenTree().DoOnEveryNode(
		func(n *nodes.Node) {
			(*n).Destroy()
		},
	)
	CurrentScene = scene
}
