package nodes

type Node interface {
	Process(delta float32)
	Input()
	Draw()
	Destroy()
}

func GetNodePtr[T interface {
	*U
	Node
}, U any](n T) *Node {
    var node Node = n
	return &node
}
