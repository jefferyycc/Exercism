package pov

type Tree struct {
	value
	nodes
}

type value string
type nodes []*Tree

// New creates and returns a new Tree with the given root value and children.
func New(root string, children ...*Tree) *Tree {
	return &Tree{value: value(root), nodes: children}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return string(tr.value)
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.nodes
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	if tr == nil {
		return nil
	}
	var path []*Tree
	if !findNodePath(tr, from, &path) {
		return nil
	}
	root := clone(path[len(path)-1])
	var ancestor *Tree
	for i := 0; i < len(path)-1; i++ {
		parent := path[i]
		children := make([]*Tree, 0, len(parent.Children()))
		for _, child := range parent.Children() {
			if child.Value() != path[i+1].Value() {
				children = append(children, clone(child))
			}
		}
		if ancestor != nil {
			children = append(children, ancestor)
		}
		ancestor = New(parent.Value(), children...)
	}
	if ancestor != nil {
		root = appendChild(root, ancestor)
	}
	return root
}

func findNodePath(tr *Tree, target string, path *[]*Tree) bool {
	*path = append(*path, tr)
	if tr.Value() == target {
		return true
	}
	for _, child := range tr.Children() {
		if findNodePath(child, target, path) {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	if tr == nil {
		return nil
	}
	// Find each node's path from the root, then join the paths at their
	// lowest common ancestor.
	var fromPath, toPath []string
	path := make([]string, 0)
	findPath := func(target string) []string {
		path = path[:0]
		if findTarget(tr, target, &path) {
			return append([]string(nil), path...)
		}
		return nil
	}
	fromPath, toPath = findPath(from), findPath(to)
	if fromPath == nil || toPath == nil {
		return nil
	}
	i := 0
	for i < len(fromPath) && i < len(toPath) && fromPath[i] == toPath[i] {
		i++
	}
	result := make([]string, 0, len(fromPath)+len(toPath)-2*(i-1))
	for j := len(fromPath) - 1; j >= i-1; j-- {
		result = append(result, fromPath[j])
	}
	result = append(result, toPath[i:]...)
	return result
}

func clone(tr *Tree) *Tree {
	if tr == nil {
		return nil
	}
	children := make([]*Tree, len(tr.Children()))
	for i, child := range tr.Children() {
		children[i] = clone(child)
	}
	return New(tr.Value(), children...)
}

func appendChild(tr *Tree, child *Tree) *Tree {
	children := append([]*Tree(nil), tr.Children()...)
	children = append(children, child)
	return New(tr.Value(), children...)
}

func findTarget(tr *Tree, target string, path *[]string) bool {
	*path = append(*path, tr.Value())
	if tr.Value() == target {
		return true
	}
	for _, child := range tr.Children() {
		if findTarget(child, target, path) {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}
