package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	low, high := 0, len(slice)-1
	for low <= high {
		mid := low + (high-low)/2
		if slice[mid] == value {
			return true, mid
		} else if slice[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false, low
}


func MagicSearch(slice []int, value int) int {
	ok, index := BetterSearch(slice, value)
	if ok {
		for i := range slice {
			if slice[len(slice)-i-1] == value {
				return len(slice)-i-1
			}
		}
	}
	return index
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
