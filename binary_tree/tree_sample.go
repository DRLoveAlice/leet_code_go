package binary_tree

var p = &TreeNode{
	Val:   5,
	Left:  nil,
	Right: nil,
}

var q = &TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}

// 二叉搜索树
var root = &TreeNode{
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
