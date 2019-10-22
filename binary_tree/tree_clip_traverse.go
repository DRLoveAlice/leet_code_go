package binary_tree

// 二叉树的锯齿形层次遍历r

func TreeClipTraverse(root *TreeNode) [][]int {
	level := 1
	result := make([][]int, 0)
	if oot == nil {
		return result
	}

	tempInput := make([]*TreeNode, 0)
	tempInput = append(tempInput, root)

	for {
		if len(tempInput) == 0 {
			break
		}

		levelResult := make([]int, 0)
		if level%2 == 0 {
			for i := len(tempInput) - 1; i >= 0; i-- {
				levelResult = append(levelResult, tempInput[i].Val)
			}
		} else {
			for i := 0; i < len(tempInput); i++ {
				levelResult = append(levelResult, tempInput[i].Val)
			}
		}

		result = append(result, levelResult)
		temp := tempInput
		tempInput = make([]*TreeNode, 0)
		for _, node := range temp {
			if node.Left != nil {
				tempInput = append(tempInput, node.Left)
			}
			if node.Right != nil {
				tempInput = append(tempInput, node.Right)
			}
		}
		level += 1
	}

	return result
}
