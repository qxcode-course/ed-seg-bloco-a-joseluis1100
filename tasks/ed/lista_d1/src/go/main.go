package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next *Node
	prev *Node
}

type LList struct {
	root *Node
}

func NewLList() *LList {
	return &LList{root: nil}
}

func (ll *LList) Size() int {
	count, curr := 0, ll.root
	for ;curr != nil; count++ {
		curr = curr.next
	}
	return count
}

func (ll *LList) Clear() {
	ll.root = nil
}

func (ll *LList) PushFront(value int) {
	newNode := &Node{Value: value, next: ll.root, prev: nil}
	if ll.root != nil {
		ll.root.prev = newNode
	}
	ll.root = newNode
}

func (ll *LList) PushBack(value int) {
	if ll.root == nil {
		ll.root = &Node{Value: value, next: nil, prev: ll.root}
		return
	}
	curr := ll.root
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &Node{Value: value, next: nil, prev: curr}
}

func (ll *LList) PopFront() {
	if ll.root == nil {
		return
	}
	ll.root = ll.root.next
	if ll.root != nil {
		ll.root.prev = nil
	}
}

func (ll *LList) PopBack() {
	if ll.root == nil || ll.root.next == nil {
		ll.Clear()
		return
	}
	curr := ll.root
	for curr.next != nil {
		curr = curr.next
	}
	curr.prev.next = nil
}

func (ll *LList) String() string {
	str, curr := "[", ll.root
	for curr != nil {
		str += strconv.Itoa(curr.Value)
		if curr.next != nil {
			str += ", "
		}
		curr = curr.next
	}
	str += "]"
	return str
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			 fmt.Println(ll.String())
		case "size":
			 fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			}
		case "push_front":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
