package structure

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
完全二叉树

		 0
	1		  2

3		4 5		 6
[

	[0],
	[1,2],
	[3,4,5,6]
	[7,8]

]
*/
func ToTree1(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := TreeNode{Val: vals[0]}
	queue := []*TreeNode{&root}
	level := 1

	for i := 1; i < len(vals); {
		children := int(math.Pow(2, float64(level)))
		for j := children - 1; j < 2*children-1 && j < len(vals); j++ {
			node := TreeNode{Val: vals[j]}
			if j%2 == 1 {
				queue[(j-1)/2].Left = &node
			} else {
				queue[j/2-1].Right = &node
			}
			queue = append(queue, &node)
		}
		i += children
		level++
	}
	return &root
}

func ToTree2(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := TreeNode{Val: vals[0]}
	childQueue, queue := []*TreeNode{}, []*TreeNode{&root}
	level := 1

	for i := 1; i < len(vals); {
		children := int(math.Pow(2, float64(level)))
		for c := i; c < len(vals) && c < i+children; c++ {
			node := TreeNode{Val: vals[c]}
			//fmt.Printf("level:%d i:%d\n", level, c)
			if c%2 == 0 {
				queue[c/2-1-level+1].Right = &node
			} else {
				queue[(c-1)/2-level+1].Left = &node
			}
			childQueue = append(childQueue, &node)
		}

		queue = childQueue
		childQueue = []*TreeNode{}
		i += children
		level++
	}

	return &root
}
