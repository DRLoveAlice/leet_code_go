package binary_tree

import "fmt"

// 剑指offer讲解链接：http://www.voidcn.com/article/p-prdiwyyo-cc.html

// 二叉树最近公共父节点

// 二叉搜索树
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		return LowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return LowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

// 普通二叉树：高级算法
func NormalLowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := NormalLowestCommonAncestor(root.Left, p, q)
	right := NormalLowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

// 递归找从根节点到一个节点的链路, 然年后
func RecursionLatestCommonAncestor(root, p, q *TreeNode) *TreeNode {
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
