package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var err error
	img, _, err := ebitenutil.NewImageFromFile("assets/example.png")
	if err != nil {
		log.Fatal(err)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2.50, 2.50)
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100, 100
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
