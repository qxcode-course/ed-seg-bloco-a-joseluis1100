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
    backup := newEditor()
	for i := range texto {
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
        case 'Z':
            e = backup
		default:
			if e.cursor == nil {
				e.texto.PushBack(texto[i])
			} else {
				e.texto.InsertBefore(texto[i], e.cursor)
			}
		}
	backup = e
    }
	fmt.Println(e.String())
}
