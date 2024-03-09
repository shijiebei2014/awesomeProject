package utils

import (
	"awesomeProject/pkg/structure"
	"fmt"
)

/*
0表示空节点
*/
func ToTree(vals []int) *structure.TreeNode {
	var root *structure.TreeNode
	var offset int

	for i := 0; i < len(vals); {
		if root == nil {
			root, offset = TreeNodeFromArray(vals)
			i += offset
		} else {
			if root.Left == nil {
				i += TreeNodeFromNode(root.Right, vals[i:])
			} else {
				i += TreeNodeFromNode(root.Left, vals[i:])
				i += TreeNodeFromNode(root.Right, vals[i:])
			}
		}
	}

	return root
}

func PrintTree(root *structure.TreeNode) {
	if root == nil {
		fmt.Println()
		return
	}
	printDfs([]*structure.TreeNode{root})
}

func printDfs(children []*structure.TreeNode) {
	nodes := make([]*structure.TreeNode, 0)
	leafs := 0
	for _, child := range children {
		if child != nil {
			if child.Left == nil && child.Right == nil {
				leafs++
			}
			nodes = append(nodes, child.Left, child.Right)
			fmt.Printf("%d ", child.Val)
		} else {
			leafs++
			//nodes = append(nodes, nil, nil)
			fmt.Printf("nil ")
		}
	}

	if leafs == len(children) {
		return
	}
	printDfs(nodes)
}

func TreeNodeFromNode(root *structure.TreeNode, vals []int) int {
	switch len(vals) {
	case 0:
		return 0
	case 1:
		root.Left = &structure.TreeNode{Val: vals[0]}
		return 1
	default:
		root.Left = &structure.TreeNode{Val: vals[0]}
		root.Right = &structure.TreeNode{Val: vals[1]}
		return 2
	}
}

func TreeNodeFromArray(vals []int) (*structure.TreeNode, int) {
	switch len(vals) {
	case 0:
		return nil, 0
	case 1:
		return treeNodeFromArray(vals[0], 0, 0), 1
	case 2:
		return treeNodeFromArray(vals[0], vals[1], 0), 2
	default:
		return treeNodeFromArray(vals[0], vals[1], vals[2]), 3
	}
}

func treeNodeFromArray(rootVal, leftVal, rightVal int) *structure.TreeNode {
	if rootVal == 0 {
		return nil
	}
	root := structure.TreeNode{Val: rootVal}
	if leftVal != 0 {
		root.Left = &structure.TreeNode{Val: leftVal}
	}
	if rightVal != 0 {
		root.Right = &structure.TreeNode{Val: rightVal}
	}
	return &root
}
