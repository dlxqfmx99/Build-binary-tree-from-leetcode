package BinaryTree

import (
	"fmt"
	"math"
)

const Null = math.MaxInt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据力扣给的遍历结果创建二叉树
func CreateBT(a []int) *TreeNode {
	list := []*TreeNode{}
	root := new(TreeNode)
	i := 0
	root.Val = a[i]
	i++
	list = append(list, root)
	for i < len(a) {
		size := len(list)
		for _, node := range list {
			if i == len(a) {
				break
			}
			if a[i] != Null {
				node.Left = new(TreeNode)
				node.Left.Val = a[i]
				list = append(list, node.Left)
			}
			i++
			if i == len(a) {
				break
			}
			if a[i] != Null {
				node.Right = new(TreeNode)
				node.Right.Val = a[i]
				list = append(list, node.Right)
			}
			i++
		}
		list = list[size:]
	}
	return root
}

// 前序遍历
func Preorder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	Preorder(node.Left)
	Preorder(node.Right)
}

func INTMAX(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func INTMIN(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func INTsMAX(nums... int) int {
	ans := nums[0]
	for _, num := range nums {
		if ans < num {
			ans = num
		}
	}
	return ans
}

func INTsMIN(nums... int) int {
	ans := nums[0]
	for _, num := range nums {
		if ans > num {
			ans = num
		}
	}
	return ans
}