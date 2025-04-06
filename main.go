package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
}

// Update will update the information displayed on the screen
func (g *Game) Update() error {
	g.player.Update()
	return nil
}

// Draw will draw information on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

// Layout returns with width and height of the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	// cannot put the player directly into &Game{} as NewPlayer requires a game
	g := &Game{}
	// so we add the player, by referencing game first
	g.player = NewPlayer(g)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
