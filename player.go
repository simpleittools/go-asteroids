package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/simpleittools/go-asteroids/assets"
)

type Player struct {
	sprite *ebiten.Image
}

// NewPlayer is a factory method that will return a new player
func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	p := &Player{
		sprite: sprite,
	}

	return p
}

// Draw will draw the player on the screen. Ebiten requires that this function be called Draw, and it returns "screen"
func (p *Player) Draw(screen *ebiten.Image) {
	// op is the options is needed for the screen.DrawImage so you can set various options.
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(p.sprite, op)
}

// Update will update the player location
func (p *Player) Update() {}
