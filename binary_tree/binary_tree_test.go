package binary_tree

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	node := LowestCommonAncestor(root, p, q)
	fmt.Println(node.Val)
}

func TestNormalLowestCommonAncestor(t *testing.T) {
	node := NormalLowestCommonAncestor(root, p, q)
	fmt.Println(node.Val)
}

func TestRecursionLatestCommonAncestor(t *testing.T) {
	node := RecursionLatestCommonAncestor(root, p, q)
	fmt.Println(node.Val)
}

func TestTreeClipTraverse(t *testing.T) {
	result := TreeClipTraverse(root)
	fmt.Println(result)
}
