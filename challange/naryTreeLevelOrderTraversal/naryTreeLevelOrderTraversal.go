package naryTreeLevelOrderTraversal

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func Run() {
	fmt.Println("start")

	root := Seed1()
	res := Solution(root)
	fmt.Println(res)

	fmt.Println("end")
}

func Solution(root *Node) [][]int {
	var res [][]int
	var resTemp []int
	var q []*Node
	var ptr *Node = root

	if root == nil {
		return res
	}

	q = append(q, ptr)
	q = append(q, nil)

	for {
		ptr, q = q[0], q[1:]
		resTemp = append(resTemp, ptr.Val)

		if len(ptr.Children) > 0 {
			q = append(q, ptr.Children...)
		}

		if q[0] == nil {
			res = append(res, resTemp)
			resTemp = []int{}

			q = q[1:]
			q = append(q, nil)
		}

		if len(q) == 1 && q[0] == nil {
			break
		}
	}

	return res
}
