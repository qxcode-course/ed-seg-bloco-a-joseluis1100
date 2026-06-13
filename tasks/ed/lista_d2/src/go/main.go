package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	next *Node
	prev *Node
	root *Node
	Value int
}

func (ll *Node) Next() *Node {
	if ll.next == ll.root {
		return nil
	}
	return ll.next
}

func (ll *Node) Prev() *Node {
	if ll.prev == ll.root {
		return nil
	}
	return ll.prev
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	return &LList{root: nil, size: 0}
}

func (ll *LList) Size() int {
	count, curr := 0, ll.root
	for ; curr != nil; count++ {
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
		ll.root = &Node{Value: value, next: nil, prev: nil}
		return
	}
	curr := ll.root
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &Node{Value: value, next: nil, prev: curr}
}

func (ll *LList) PopFront() {
	if ll.root == nil || ll.root.next == nil {
		ll.Clear()
		return
	}
	ll.root = ll.root.next
	ll.root.prev = nil
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

func (ll *LList) Front() *Node {
	return ll.root
}

func (ll *LList) Back() *Node {
	if ll.root == nil {
		return nil
	}
	curr := ll.root
	for curr.next != nil {
		curr = curr.next
	}
	return curr
}

func (ll *LList) Search(value int) *Node {
	curr := ll.root
	for curr != nil {
		if curr.Value == value {
			return curr
		}
		curr = curr.next
	}
	return nil
}

func (ll *LList) Insert(node *Node, value int) {
	if node == nil {
		return
	}
	if node == ll.root {
		ll.PushFront(value)
		return
	}
	newNode := &Node{Value: value, next: node, prev: node.prev}
	node.prev.next = newNode
	node.prev = newNode
}

func (ll *LList) Remove(node *Node) {
	if node == nil {
		return
	}
	if node == ll.Front() {
		ll.PopFront()
		return
	}
	if node == ll.Back() {
		ll.PopBack()
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
