package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func precedencia(op string) int {
    switch op {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    default:
        return 3
    }
}

func conveter(notacao string) string {
    not := strings.Fields(notacao)
    result := make([]string, 0)
    stack := make([]string, 0)
    for i := range not {
        if not[i] == "+" || not[i] == "-" || not[i] == "*" || not[i] == "/" || not[i] == "^" {
            for len(stack) > 0 {
                if precedencia(stack[len(stack)-1]) >= precedencia(not[i]) {
                    result = append(result, stack[len(stack)-1])
                    stack = stack[:len(stack)-1]
                } else {
                    break
                }
            }
            stack = append(stack, not[i])
        } else {
            result = append(result, not[i])
        }
    }
    for len(stack) > 0 {
        result = append(result, stack[len(stack)-1])
        stack = stack[:len(stack)-1]
    }
    return strings.Join(result, " ")
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        notacao := scanner.Text()
        fmt.Println(conveter(notacao))
    }
}