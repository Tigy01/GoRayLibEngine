package scenes

import (
	"myengine/pkg/engine/nodes"
)

type Scene interface {
	nodes.Node
	GetChildrenTree() Hierarchy
}

func AsScenePtr[T interface {
	*U
	Scene
}, U any](n T) *Scene {
	var scene Scene = n
	return &scene
}

type Hierarchy struct {
	Scene    *Scene
	Children []*Tree
}

type Tree struct {
	Value    *nodes.Node
	Children []*Tree
}

func (h Hierarchy) DoOnEveryScene(function func(*Scene)) {
	function(h.Scene)
	for _, c := range h.Children {
		c.DoOnEveryScene(function)
	}

}

func (t Tree) DoOnEveryScene(function func(*Scene)) {
	if s, ok := (*t.Value).(Scene); ok {
		function(&s)
	}
	for _, c := range t.Children {
		c.DoOnEveryScene(function)
	}
}

func (h Hierarchy) DoOnEveryNode(function func(*nodes.Node)) {
	var node nodes.Node = *h.Scene
	function(&node)
	for _, c := range h.Children {
		c.DoOnEveryNode(function)
	}

}

func (t Tree) DoOnEveryNode(function func(*nodes.Node)) {
	function(t.Value)
	for _, c := range t.Children {
		c.DoOnEveryNode(function)
	}
}
