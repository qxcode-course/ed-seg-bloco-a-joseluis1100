package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	pen := NewPen(500, 500)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(0, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(250, 10) // move o pincel para x 250, y 500
	// Ponta 1
    pen.SetHeading(252)   
    pen.Walk(500)
    
    // Ponta 2
    pen.SetHeading(36)  
    pen.Walk(500)
    
    // Ponta 3
    pen.SetHeading(180)  
    pen.Walk(500)
    
    // Ponta 4
    pen.SetHeading(324) 
    pen.Walk(500)
    
    // Ponta 5
    pen.SetHeading(108)    
    pen.Walk(500)

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
