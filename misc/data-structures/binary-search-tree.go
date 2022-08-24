package datastructures

// tree node
type Node struct {
	Value  int
	Left   *Node
	Right  *Node
}

func Lookup(n *Node, target int) bool {
	// base case
	if n == nil {
		return false
	} 
	// see if found here
	if target == n.Value {
		return true
	}

	if target < n.Value {
		return Lookup(n.Left, target)
	} else {
		return Lookup(n.Right, target)
	}

}


func NewNode(data int) *Node {
	n := &Node{
		Value: data,
		Left:  nil,
		Right: nil,
	}
	return n
}

func Insert(n *Node, data int) *Node {
	if n == nil {
		return NewNode(data)
	} else {
		if data <= n.Value {
			n.Left = Insert(n.Left, data)
		} else {
			n.Right = Insert(n.Right, data)
		}
		return n
	}

}