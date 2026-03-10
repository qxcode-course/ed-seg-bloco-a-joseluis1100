package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func tree(p *Pen, tam float64) {
	fator := 0.75
	ang := 35.00

	if tam < 15 {
		if randInt(0, 10) < 1 {
			p.SetColor(255, 255, 255)
			p.DrawCircle(tam)
		}
		return
	}
	p.SetColor(255, 255, 255)
	p.SetStrokeWidth(int(tam / 10))
	p.Walk(tam)
	p.Rotate(-ang)
	tree(p, tam*fator)
	p.Rotate(2 * ang)
	tree(p, tam*fator)
	p.Rotate(-ang)
	p.Walk(-tam)
}

func main() {
	pen := NewPen(1000, 1000)
	pen.SetColor(255, 255, 255)
	pen.SetPosition(500, 550)
	pen.SetHeading(45)
	tree(pen, 100)

	pen.SetHeading(135)
	tree(pen, 100)

	//pen.SetHeading(180)
	//tree(pen, 100)

	//pen.SetHeading(360)
	//tree(pen, 100)
	pen.SaveToFile("tree.svg")
	fmt.Println("SVG file created successfully.")
}
