package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/simpleittools/go-asteroids/assets"
	"math"
)

const rotationPerSecond = math.Pi

type Player struct {
	sprite   *ebiten.Image
	rotation float64
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
	//bounds gives us the boundary of the sprite. This is a rectangle.
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	// op is the options is needed for the screen.DrawImage so you can set various options.
	op := &ebiten.DrawImageOptions{}

	// this sets the sprite's center at origin point 0,0
	op.GeoM.Translate(-halfW, -halfH)
	// rotate the sprite
	op.GeoM.Rotate(p.rotation)
	// move back to where it has to go
	op.GeoM.Translate(halfW, halfH)

	screen.DrawImage(p.sprite, op)
}

// Update will update the player location
func (p *Player) Update() {
	// we want to rotate the image on the screen
	// TPS is ticks per second
	speed := rotationPerSecond / float64(ebiten.TPS())

	// TODO: add gamepad compatability
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		p.rotation += speed
	}
}
