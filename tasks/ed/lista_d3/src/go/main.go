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
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
}


func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}

func equals(lla, llb *LList) bool {
	currA, currB := lla.root.next, llb.root.next
	for currA != lla.root && currB != llb.root {
		if currA.Value != currB.Value {
			return false
		}
		currA = currA.next
		currB = currB.next
	}
	return currA == lla.root && currB == llb.root
}

func addsorted(lla *LList, value int) {
	curr := lla.root.next
	for curr != lla.root {
		if value <= curr.Value {
			lla.insertBefore(curr, value)
			return
		}
		curr = curr.next
	}
	lla.insertBefore(curr, value)
}

func reverse(lla *LList) {
	if lla.root.next == lla.root {
		return
	}
	curr := lla.root
	for {
		curr.prev, curr.next = curr.next, curr.prev
		curr = curr.prev
		if curr == curr.root {
			return
		}
	}
}

func merge(lla, llb *LList) *LList {
	curr := llb.root.next
	for curr != llb.root {
		addsorted(lla, curr.Value)
		curr = curr.next
	}
	return lla
}

func (ll *LList) String() string {
	str, curr := "[", ll.root.next
	for curr != ll.root {
		str += strconv.Itoa(curr.Value)
		if curr.next != ll.root {
			str += ", "
		}
		curr = curr.next
	}
	str += "]"
	return str
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
