module github.com/Tigy01/GoRayLibEngine

go 1.23.5

replace github.com/Tigy01/GoRayLibEngine/scenes => ./scenes
replace github.com/Tigy01/GoRayLibEngine/nodes => ./nodes
replace github.com/Tigy01/GoRayLibEngine/enginenodes => ./enginenodes

require (
	github.com/ebitengine/purego v0.8.2 // indirect
	github.com/gen2brain/raylib-go/raylib v0.0.0-20250215042252-db8e47f0e5c5 // indirect
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394 // indirect
	golang.org/x/sys v0.31.0 // indirect
)
