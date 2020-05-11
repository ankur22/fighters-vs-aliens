package game

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type f22 struct {
	img         *ebiten.Image
	frameWidth  int
	frameHeight int
	left, right bool
}

func (f *f22) Update(screen *ebiten.Image, tick float64) error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && ebiten.IsKeyPressed(ebiten.KeyRight) {
		f.left = false
		f.right = false
	} else {
		f.left = false
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			f.left = true
		}

		f.right = false
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			f.right = true
		}
	}
	return nil
}

func (f *f22) Draw(screen *ebiten.Image) {
	sx, sy := 0, 2*f.frameHeight
	if f.left {
		sy = f.frameHeight
	} else if f.right {
		sy = 0
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(f.frameWidth)/2, -float64(f.frameHeight)/2)
	op.GeoM.Translate(float64(screen.Bounds().Size().X/2), float64(screen.Bounds().Size().Y/2))

	screen.DrawImage(f.img.SubImage(image.Rect(sx, sy, sx+f.frameWidth, sy+f.frameHeight)).(*ebiten.Image), op)
}
