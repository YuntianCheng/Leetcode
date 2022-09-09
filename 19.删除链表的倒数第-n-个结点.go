//package main
//
///*
// * @lc app=leetcode.cn id=19 lang=golang
// *
// * [19] 删除链表的倒数第 N 个结点
// */
////Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//// @lc code=start
//
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	cursor := head
//	count := 0
//	for {
//		if cursor == nil {
//			break
//		}
//		cursor = cursor.Next
//		count++
//	}
//	index := count - n
//	if index == 0 {
//		if count == 1 {
//			return nil
//		}
//		return head.Next
//	}
//	cursor = head
//	for i := 1; i < index; i++ {
//		cursor = cursor.Next
//	}
//	if index == count-1 {
//		cursor.Next = nil
//		return head
//	}
//	cursor.Next = cursor.Next.Next
//	return head
//}
//
//// @lc code=end