package tree

type Element interface{}

type TreeNode struct {
	Value    Element
	parent   *TreeNode
	children []*TreeNode
}

func (node *TreeNode) AddChild(newNode *TreeNode) {
	node.children = append(node.children, newNode)
	newNode.parent = node
}

func (node *TreeNode) Search(value Element) *TreeNode {
	if value == node.Value {
		return node
	}
	for _, child := range node.children {
		if found := child.Search(value); found != nil {
			return found
		}
	}
	return nil
}

func New(value Element) *TreeNode {
	return &TreeNode{value, nil, nil}
}
