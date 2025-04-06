package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

// Update will update the information displayed on the screen
func (g *Game) Update() error {
	return nil
}

// Draw will draw information on the screen
func (g *Game) Draw(screen *ebiten.Image) {
}

// Layout returns with width and height of the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
