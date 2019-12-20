// Difficult: Medium

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

package main

import "fmt"

// ListNode is a struct for linked List
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	item1 := &ListNode{2, nil}
	item2 := &ListNode{4, nil}
	item3 := &ListNode{3, nil}

	item4 := &ListNode{5, nil}
	item5 := &ListNode{6, nil}
	item6 := &ListNode{4, nil}

	items := item1  // list1
	items2 := item4 // list2

	items = addNodeToEnd(item2, items)
	items = addNodeToEnd(item3, items)

	items2 = addNodeToEnd(item5, items)
	items2 = addNodeToEnd(item6, items)

	items = reverseRecurrsive(items)
	items2 = reverseRecurrsive(items2)

	for list := items; list != nil; list = list.Next {
		fmt.Println(list)
	}
	for list := items2; list != nil; list = list.Next {
		fmt.Println(list)
	}

	// addTwoNumbers(items, items2)

}

func addNodeToEnd(li, item *ListNode) *ListNode {
	if item == nil {
		return item
	}
	for list := item; list != nil; list = list.Next {
		if list.Next == nil {
			list.Next = li
			return item
		}
	}
	return item
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	for list := l1; list != nil; list = list.Next {
// 		fmt.Println(list)
// 	}
// }

func reverseRecurrsive(list *ListNode) *ListNode {
	if list == nil {
		return list
	}
	l := list

	if l.Next == nil {
		return l
	} else {
		newHead := reverseRecurrsive(l.Next)
		l.Next.Next = l
		l.Next = nil
		return newHead
	}
}
