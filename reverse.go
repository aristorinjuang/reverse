package reverse

import (
	"errors"
	"fmt"
)

// List is a definition of linked list.
type List struct {
	Val  byte
	Next *List
}

// Init initializes a linked list.
func Init(s []byte) (*List, error) {
	if string(s) == "" {
		return nil, errors.New("list is empty")
	}

	head := &List{s[0], nil}
	curr := head

	for i := 1; i < len(s); i++ {
		node := &List{s[i], nil}
		curr.Next = node
		curr = node
	}

	return head, nil
}

// Print return a linked list as a string to print.
func Print(list *List) string {
	result := ""
	curr := list

	for curr != nil {
		result += fmt.Sprintf("%c", curr.Val)
		curr = curr.Next
	}

	return result
}

// LinkedList reverse a linked list.
func LinkedList(node *List) (*List, error) {
	if node == nil {
		return nil, errors.New("list is empty")
	}

	var curr, prev, next *List = node, nil, nil

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev, nil
}

// String reverse a string in bytes.
func String(s []byte) ([]byte, error) {
	if string(s) == "" {
		return nil, errors.New("string is empty")
	}

	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}

	return s, nil
}
