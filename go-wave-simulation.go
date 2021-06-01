package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const WINDOW_WIDTH int = 640
const WINDOW_HEIGHT int = 480

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var newPixels [WINDOW_HEIGHT * WINDOW_WIDTH * 4]byte

	for y := 0; y < WINDOW_HEIGHT; y++ {
		for x := 0; x < WINDOW_WIDTH; x++ {
			for c := 0; c < 4; c++ {
				newPixels[y*WINDOW_WIDTH*4+x*4+c] = byte(x * y)
			}
		}
	}

	screen.ReplacePixels(newPixels[:])
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOW_WIDTH, WINDOW_HEIGHT
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Hello World")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
