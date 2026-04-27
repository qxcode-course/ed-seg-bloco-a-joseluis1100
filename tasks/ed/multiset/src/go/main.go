package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(cap int) *MultiSet {
	return &MultiSet{
		data:     make([]int, 0, cap),
		size:     0,
		capacity: cap,
	}
}

func (ms *MultiSet) expand() {
	if ms.capacity == 0 {
		ms.capacity = 1
	} else {
		ms.capacity *= 2
	}
	tmp := make([]int, ms.size, ms.capacity)
	copy(tmp, ms.data)
	ms.data = tmp
}

func (ms *MultiSet) search(value int) (bool, int) {
	for i := range ms.size {
		if ms.data[ms.size-i-1] == value {
			return true, ms.size - i - 1
		}
	}
	return false, -1
}

func (ms *MultiSet) Insert(value int) {
	for i := range ms.size {
		if ms.data[i] > value {
			ms.insert(value, i)
			return
		}
	}
	ms.insert(value, ms.size)
}

func (ms *MultiSet) insert(value, index int) error {
	if index < 0 || index > ms.size {
		return errors.New("index out of range")
	}
	if ms.size == ms.capacity {
		ms.expand()
	}
	ms.data = append(ms.data, 0)
	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[index] = value
	ms.size++
	return nil
}

func (ms *MultiSet) Erase(value int) error {
	existe, index := ms.search(value)
	if !existe {
		return errors.New("value not found")
	}
	return ms.erase(index)
}

func (ms *MultiSet) erase(index int) error {
	if index < 0 || index >= ms.size {
		return errors.New("index out of range")
	}
	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	return nil
}

func (ms *MultiSet) Contains(value int) bool {
	low, high := 0, ms.size-1
	for low <= high {
		mid := low + (high-low)/2
		if ms.data[mid] == value {
			return true
		}
		if ms.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func (ms *MultiSet) Count(value int) int {
	count := 0
	for i := range ms.size {
		if ms.data[i] == value {
			count++
		}
	}
	return count
}

func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}
	count := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			count++
		}
	}
	return count
}

func (ms *MultiSet) Clear() {
	ms.size = 0
}

func (ms *MultiSet) String() string {
	if ms.size == 0 {
		return "[]"
	}
	return "[" + Join(ms.data[:ms.size], ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
