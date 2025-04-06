package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/simpleittools/go-asteroids/assets"
	"math"
)

const (
	rotationPerSecond = math.Pi
	maxAcceleration   = 8.0
	ScreenWidth       = 1280
	ScreenHeight      = 720
)

var curAcceleration float64

type Player struct {
	game           *Game
	sprite         *ebiten.Image
	rotation       float64
	position       Vector
	playerVelocity float64
}

// NewPlayer is a factory method that will return a new player
func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	// center the player on screen
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	// position
	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	p := &Player{
		sprite:   sprite,
		game:     game,
		position: pos,
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

	op.GeoM.Translate(p.position.X, p.position.Y)

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

	p.Accelerate()

}

func (p *Player) Accelerate() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		p.keepOnScreen()

		if curAcceleration < maxAcceleration {
			curAcceleration = p.playerVelocity + 4
		}

		if curAcceleration >= 8 {
			curAcceleration = 8
		}

		p.playerVelocity = curAcceleration

		// Move in the direction we are pointing
		dx := math.Sin(p.rotation) * curAcceleration
		dy := math.Cos(p.rotation) * -curAcceleration

		// move the player on the screen
		p.position.X += dx
		p.position.Y += dy
	}
}

// keepOnScreen will see if the player has gone beyond the boundaries of the screen
func (p *Player) keepOnScreen() {
	if p.position.X >= float64(ScreenWidth) {
		p.position.X = 0
	}

	if p.position.X < 0 {
		p.position.X = ScreenWidth
	}

	if p.position.Y >= float64(ScreenHeight) {
		p.position.Y = 0
	}

	if p.position.Y < 0 {
		p.position.Y = ScreenHeight
	}
}
