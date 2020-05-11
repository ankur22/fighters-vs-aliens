package game

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type entity interface {
	Update(screen *ebiten.Image, tick float64) error
	Draw(screen *ebiten.Image)
}

type Game struct {
	name         string
	screenWidth  int
	screenHeight int
	entities     []entity
}

func (g *Game) Update(screen *ebiten.Image) error {
	tick := 1000 / ebiten.CurrentTPS()
	for _, e := range g.entities {
		if err := e.Update(screen, tick); err != nil {
			return nil
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.name)
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS())
	ebitenutil.DebugPrintAt(screen, msg, 0, 15)

	for _, e := range g.entities {
		e.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}

func Run() error {
	f22Img, _, err := ebitenutil.NewImageFromFile("f22-1.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	gameName := "Fighter vs Aliens"
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle(gameName)

	f22Ent := f22{
		img:         f22Img,
		frameHeight: 128,
		frameWidth:  128,
	}

	entities := []entity{&f22Ent}

	g := &Game{
		name:         gameName,
		screenHeight: 480,
		screenWidth:  640,
		entities:     entities,
	}
	if err := ebiten.RunGame(g); err != nil {
		return fmt.Errorf("On startup - %w", err)
	}
	return nil
}
