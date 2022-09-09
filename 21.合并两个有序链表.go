package main

import "fmt"

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

// Definition for singly-linked list.

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	c1 := list1
	c2 := list2
	if c1 == nil && c2 != nil {
		return c2
	}
	if c1 != nil && c2 == nil {
		return c1
	}
	if c1 == nil && c2 == nil {
		return nil
	}
	for {
		if c2 == nil {
			break
		}
		if c2.Val >= c1.Val {
			if c1.Next != nil {
				if c1.Next.Val >= c2.Val {
					temp1 := c1.Next
					c1.Next = c2
					temp2 := c2.Next
					c2.Next = temp1
					c2 = temp2
				} else {
					c1 = c1.Next
				}
			} else {
				c1.Next = c2
				break
			}
		} else {
			temp := c2.Next
			c2.Next = c1
			list1 = c2
			c1 = list1
			c2 = temp
		}
	}
	return list1
}

// @lc code=end
func main() {
	list1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Println(mergeTwoLists(list1, list2))
}
