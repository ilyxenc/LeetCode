package tasks

// One type of problem solving:
// 1. AddTwoNumbers()
// 
// Test values are below

import (
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listNodeToSlice(n *ListNode) string {
	s := ""
	for el := n; el != nil; el = el.Next {
		s += strconv.Itoa(el.Val)
	}

	return s
}

func newListNode(val int) *ListNode {
	var n ListNode

	n.Val = val
	n.Next = nil

	return &n
}

func addListNode(n *ListNode, val int) *ListNode {
	if (n == nil) {
		n = newListNode(val)

		return n
	}

	temp := n

	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = newListNode(val)

	return n
}

func reverseListNode(n *ListNode) *ListNode  {
	if (n == nil ){
		return nil
	}

	var r *ListNode

	temp := n

	for temp != nil {
		r = &ListNode{
			Val: temp.Val,
			Next: r,
		}

		temp = temp.Next
	}

	return r
}

func sumBig(aOld, bOld string) string {
	a := new(big.Int)
	a.SetString(aOld, 10)

	b := new(big.Int)
	b.SetString(bOld, 10)

	sum := big.NewInt(0)

	sum.Add(a, b)

	return sum.String()
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1RevNew := listNodeToSlice(reverseListNode(l1))
	l2RevNew := listNodeToSlice(reverseListNode(l2))
	num := sumBig(l1RevNew, l2RevNew)

	el, _ := strconv.Atoi(string(num[0]))

	n := newListNode(el)

	for i := 1; i < len(num); i++ {
		el, _ = strconv.Atoi(string(num[i]))

		n = addListNode(n, el)
	}

	return reverseListNode(n)
}



// TEST VALUES

// l1 := &tasks.ListNode{
// 	Val: 1,
// 	Next: &tasks.ListNode{
// 		Val: 0,
// 		Next: &tasks.ListNode{
// 			Val: 0,
// 			Next: &tasks.ListNode{
// 				Val: 0,
// 				Next: &tasks.ListNode{
// 					Val: 0,
// 					Next: &tasks.ListNode{
// 						Val: 0,
// 						Next: &tasks.ListNode{
// 							Val: 0,
// 							Next: &tasks.ListNode{
// 								Val: 0,
// 								Next: &tasks.ListNode{
// 									Val: 0,
// 									Next: &tasks.ListNode{
// 										Val: 0,
// 										Next: &tasks.ListNode{
// 											Val: 0,
// 											Next: &tasks.ListNode{
// 												Val: 0,
// 												Next: &tasks.ListNode{
// 													Val: 0,
// 													Next: &tasks.ListNode{
// 														Val: 0,
// 														Next: &tasks.ListNode{
// 															Val: 0,
// 															Next: &tasks.ListNode{
// 																Val: 0,
// 																Next: &tasks.ListNode{
// 																	Val: 0,
// 																	Next: &tasks.ListNode{
// 																		Val: 0,
// 																		Next: &tasks.ListNode{
// 																			Val: 0,
// 																			Next: &tasks.ListNode{
// 																				Val: 0,
// 																				Next: &tasks.ListNode{
// 																					Val: 1,
// 																					Next: nil,
// 																				},
// 																			},
// 																		},
// 																	},
// 																},
// 															},
// 														},
// 													},
// 												},
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	},
// }

// l2 := &tasks.ListNode{
// 	Val: 5,
// 	Next: &tasks.ListNode{
// 		Val: 6,
// 		Next: &tasks.ListNode{
// 			Val: 4,
// 			Next: nil,
// 		},
// 	},
// }

// result := tasks.AddTwoNumbers(l1, l2)