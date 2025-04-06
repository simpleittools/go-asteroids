package assets

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

//go:embed *

// assets will give go access to all of these items as the file system
var assets embed.FS

var PlayerSprite = mustLoadImage("images/player.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
