package main

import "fmt"

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

// 通用类型的节点数据结构
func main() {
	root := NewNode(nil, nil)
	root.SetData("root node")

	a := NewNode(nil, nil)
	a.SetData("left node")

	b := NewNode(nil, nil)
	b.SetData("right node")

	root.le = a
	root.ri = b
	// root = &main.Node{le:(*main.Node)(0xc00004c3e0), data:"root node", ri:(*main.Node)(0xc00004c400)}, *main.Node, main.Node, *main.Node
	fmt.Printf("root = %#v, %T, %T, %T \n", root, root, Node{}, &Node{})
}
