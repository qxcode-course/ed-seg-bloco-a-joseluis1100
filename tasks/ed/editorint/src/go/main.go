package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() {
		e.cursor = e.cursor.Prev()
		return
	}
	// Estamos no início da linha
	if e.line != e.lines.Front() {
		e.line = e.line.Prev()
		e.cursor = e.line.Value.End()
	}
}

func (e *Editor) KeyEnter() {
	nova := NewList[rune]()
	curr := e.cursor
	for curr != e.line.Value.End() {
		tmp := curr.Next()
		nova.PushBack(curr.Value)
		e.line.Value.Erase(curr)
		curr = tmp
	}
	e.lines.Insert(e.line.Next(), nova)
	e.line = e.line.Next()
	e.cursor = e.line.Value.Front()
}

func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.cursor.Next()
		return
	}
	// Estamos no início da linha
	if e.line.Next() != e.lines.End() {
		e.line = e.line.Next()
		e.cursor = e.line.Value.Front()
	}
}

func (e *Editor) KeyUp() {
	if e.line != e.lines.Front() {
		col := e.line.Value.IndexOf(e.cursor)
		e.line = e.line.Prev()
		e.cursor = e.line.Value.Front()
		for range col {
			e.cursor = e.cursor.Next()
		}
	}
}

func (e *Editor) KeyDown() {
	if e.line.Next() != e.lines.End() {
		col := e.line.Value.IndexOf(e.cursor)
		e.line = e.line.Next()
		e.cursor = e.line.Value.Front()
		for range col {
			e.cursor = e.cursor.Next()
		}
	}
}

func (e *Editor) KeyBackspace() {
	if e.cursor != e.line.Value.Front() {
		e.line.Value.Erase(e.cursor.Prev())
		return
	}
	if e.line != e.lines.Front() {
		col := e.line.Prev().Value.IndexOf(e.line.Prev().Value.End())
		for curr := e.line.Value.Front(); curr != e.line.Value.End(); curr = curr.Next() {
			e.line.Prev().Value.PushBack(curr.Value)
		}
		e.line = e.line.Prev()
		e.lines.Erase(e.line.Next())
		e.cursor = e.line.Value.Front()
		for range col {
			e.cursor = e.cursor.Next()
		}
	}
}

func (e *Editor) KeyDelete() {
	if e.cursor != e.line.Value.End() {
		target := e.cursor
		e.cursor = e.cursor.Next()
		e.line.Value.Erase(target)
		return
	}
	if e.line.Next() != e.lines.End() {
		col := e.line.Value.IndexOf(e.cursor)
		for curr := e.line.Next().Value.Front(); curr != e.line.Next().Value.End(); curr = curr.Next() {
			e.line.Value.PushBack(curr.Value)
		}
		e.lines.Erase(e.line.Next())
		for range col {
			e.cursor = e.cursor.Next()
		}
		e.cursor = e.cursor.Next()
	}
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
