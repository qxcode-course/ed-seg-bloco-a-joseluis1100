package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(cap int) *Set {
	return &Set{
		data:     make([]int, 0, cap),
		size:     0,
		capacity: cap,
	}

}

func (s *Set) reserve(newCapacity int) {
	if newCapacity <= s.capacity {
		return
	}
	tmp := make([]int, newCapacity)
	for i := range s.size {
		tmp[i] = s.data[i]
	}
	s.data = tmp
	s.capacity = newCapacity
}

func (s *Set) binarySearch(value int) int {
	low, high := 0, s.size-1
	for low <= high {
		mid := low + (high-low)/2
		if s.data[mid] == value {
			return mid
		} else if s.data[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func (s *Set) insert(value, index int) error {
	if index < 0 || index > s.size {
		return errors.New("index out of range")
	}
	if s.size == s.capacity {
		s.reserve(2 * s.capacity)
	}
	s.data = append(s.data, 0)
	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = value
	s.size++
	return nil
}

func (s *Set) erase(index int) error {
	if index < 0 || index >= s.size {
		return errors.New("index out of range")
	}
	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
	return nil
}

func (s *Set) Insert(value int) {
	if s.Contains(value) {
		return
	}
	for i := range s.size {
		if s.data[i] > value {
			s.insert(value, i)
			return
		}
	}
	s.insert(value, s.size)
}

func (s *Set) Contains(value int) bool {
	if s.binarySearch(value) == -1 {
		return false
	}
	return true
}

func (s *Set) Erase(value int) bool {
	if !s.Contains(value) {
		fmt.Println("value not found")
		return false
	}
	s.erase(s.binarySearch(value))
	return true
}

func (v *Set) String() string {
	if v.size == 0 {
		return "[]"
	}
	var res []string
	for i := 0; i < v.size; i++ {
		res = append(res, strconv.Itoa(v.data[i]))
	}
	return "[" + strings.Join(res, ", ") + "]"
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			v.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
