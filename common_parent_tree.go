package leet_code_go

import "fmt"

// 剑指offer讲解链接：http://www.voidcn.com/article/p-prdiwyyo-cc.html

// 二叉树最近公共父节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

// 普通二叉树：高级算法
func lowestCommonAncestorNew(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestorNew(root.Left, p, q)
	right := lowestCommonAncestorNew(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

// 递归找从根节点到一个节点的链路, 然年后
func getLatestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	pathP := make([]*TreeNode, 0)
	pathP = append(pathP, root)
	pathQ := make([]*TreeNode, 0)
	pathQ = append(pathQ, root)

	_, pathP = getPath(root, p, pathP)
	_, pathQ = getPath(root, q, pathQ)
	printPath(pathQ, "q")
	printPath(pathP, "p")
	var Node *TreeNode
	for i := 0; i < len(pathP) && i < len(pathQ); i++ {
		if pathP[i] == pathQ[i] {
			Node = pathP[i]
			continue
		}
		break
	}

	return Node
}

func getPath(root, node *TreeNode, path []*TreeNode) (bool, []*TreeNode) {
	if root == node {
		return true, path
	}

	if root.Left != nil {
		path = append(path, root.Left)
		if found, newPath := getPath(root.Left, node, path); found {
			return true, newPath
		}
		path = path[0 : len(path)-1]
	}
	if root.Right != nil {
		path = append(path, root.Right)
		if found, newPath := getPath(root.Right, node, path); found {

			return true, newPath
		}
		path = path[0 : len(path)-1]
	}
	return false, path

}

func printPath(path []*TreeNode, tag string) {
	fmt.Print(tag)
	for _, node := range path {
		fmt.Print(node.Val)
	}

	fmt.Println()
}

func main() {
	p := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	q := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: q,
		},
		Right: &TreeNode{
			Val:  6,
			Left: p,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	node := getLatestCommonAncestor(root, p, q)
	println(node.Val)
	node = lowestCommonAncestorNew(root, p, q)
	println(node.Val)
	node = lowestCommonAncestor(root, p, q)
	println(node.Val)
}
