package main
import "fmt"

type Stack struct {
    data []byte
}

func newStack() *Stack {
    return &Stack{data: make([]byte, 0)}
}

func (s *Stack) Top() byte {
    return s.data[len(s.data)-1]
}

func (s *Stack) Push(v byte) {
    s.data = append(s.data, v)
}

func (s *Stack) Pop() {
    s.data = s.data[:len(s.data)-1]
}

func (s *Stack) isEmpty() bool {
    return len(s.data) == 0
}

func ehBalanceado(text string) bool {
    abre := newStack()
    for i := range text {
        if text[i] == '(' || text[i] == '[' {
            abre.Push(text[i])
            continue
        }
        if (text[i] == ')' || text[i] == ']') && abre.isEmpty() {
            return false
        }
        if text[i] == ')' && abre.Top() != '(' || text[i] == ']' && abre.Top() != '[' {
            return false
        }
        if text[i] == ')' && abre.Top() == '(' || text[i] == ']' && abre.Top() == '[' {
            abre.Pop()
        }
    }
    if !abre.isEmpty() {
        return false
    }
    return true
}
func main() {
    var text string
    fmt.Scan(&text)
    if ehBalanceado(text) {
        fmt.Println("balanceado")
    } else {

        fmt.Println("nao balanceado")
    }
}