package main
import "fmt"

type Bomba struct {
    gas, dist int
}

type Queue struct {
    data []Bomba
    totalgas int
}

func newQueue(n int) *Queue {
    return &Queue{data: make([]Bomba, 0, n)}
}

func (q *Queue) Push(v Bomba) {
    q.data = append(q.data, v)
}

func (q *Queue) Pop() {
    q.data = q.data[1:]
}

func (q *Queue) isEmpty() bool {
    return len(q.data) == 0
}

func (q *Queue) pode() bool {
    if q.totalgas + q.data[0].gas - q.data[0].dist < 0 {
        return false
    }
    q.totalgas += q.data[0].gas - q.data[0].dist
    return true
}

func main() {
    var n int
    fmt.Scan(&n)
    bombas := make([]Bomba, n)
    for i := range n {
        var gas, dist int
        fmt.Scan(&gas, &dist)
        bombas[i] = Bomba{gas: gas, dist: dist}
    }
    for i := range n {
        queue := newQueue(n)
        for j := range bombas {
            queue.Push(bombas[(j+i) % len(bombas)])
        }
        for range n {
            if queue.pode() {
                queue.Pop()
                } else {
                    break
                }
            }
            if queue.isEmpty() {
                fmt.Println(i)
                return
            }
    }
}