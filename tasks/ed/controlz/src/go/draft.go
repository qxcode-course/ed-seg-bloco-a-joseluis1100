package main

import (
	"container/list"
	"fmt"
)

type Editor struct {
	texto  *list.List
	cursor *list.Element
}

func newEditor() *Editor {
	return &Editor{
		texto: list.New(),
		cursor: nil,
	}
}

func (e *Editor) clonar() *Editor {
	tmp := newEditor()
	for i := e.texto.Front(); i != nil; i = i.Next() {
		node := tmp.texto.PushBack(i.Value)
		if i == e.cursor {
			tmp.cursor = node
		}
	}
	return tmp
}

func (e *Editor) String() string {
	var str string
	for i := e.texto.Front(); i != nil; i = i.Next() {
		if i == e.cursor {
			str += "|"
		}
		str += string(i.Value.(byte))
	}
	if e.cursor == nil {
		str += "|"
	}
	return str
}

func main() {
	var texto string
	fmt.Scanln(&texto)
	e := newEditor()
    historico := []*Editor{}
	desfeitos := []*Editor{}
	for i := range texto {
		if texto[i] == 'Z' {
			if len(historico) > 0 {
				desfeitos = append(desfeitos, e.clonar())
				e = historico[len(historico)-1].clonar()
				historico = historico[:len(historico)-1]
			}
			continue
		}
		if texto[i] == 'Y' {
			if len(desfeitos) > 0 {
				e = desfeitos[len(desfeitos)-1].clonar()
				desfeitos = desfeitos[:len(desfeitos)-1]
			}
			continue
		} 
		historico = append(historico, e.clonar())
		switch texto[i] {
		case 'R':
			if e.cursor == nil {
				e.texto.PushBack(byte('\n'))
			} else {
				e.texto.InsertBefore(byte('\n'), e.cursor)
			}
		case 'B':
			if e.cursor == nil {
				if e.texto.Back() != nil {
					e.texto.Remove(e.texto.Back())
				}
			} else {
				if e.cursor.Prev() != nil {
					e.texto.Remove(e.cursor.Prev())
				}
			}
		case 'D':
			if e.cursor != nil {
				tmp := e.cursor.Next()
				e.texto.Remove(e.cursor)
				e.cursor = tmp
			}
		case '<':
			if e.cursor == nil {
				if e.texto.Back() != nil {
					e.cursor = e.texto.Back()
				}
			} else if e.cursor.Prev() != nil {
				e.cursor = e.cursor.Prev()
			}
		case '>':
			if e.cursor != nil {
				e.cursor = e.cursor.Next()
			}
		default:
			if e.cursor == nil {
				e.texto.PushBack(texto[i])
			} else {
				e.texto.InsertBefore(texto[i], e.cursor)
			}
		}
    }
	fmt.Println(e.String())
}
