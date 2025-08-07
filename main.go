package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/Z-zenos/devide/internal/player"
)

type Game struct {
	player *player.Player
}

func (g *Game) Update() error {
	// TODO: Handle input, update game state, etc.
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear background
	screen.Fill(color.RGBA{30, 30, 30, 255}) // Dark gray background

	// Draw border rectangle (play area)
	borderColor := color.White
	x, y := float32(50.0), float32(50.0)
	w, h := float32(540.0), float32(380.0)
	vector.DrawFilledRect(screen, x, y, w, 1, borderColor, false)       // Top
	vector.DrawFilledRect(screen, x, y+h-1, w, 1, borderColor, false)   // Bottom
	vector.DrawFilledRect(screen, x, y, 1, h, borderColor, false)       // Left
	vector.DrawFilledRect(screen, x+w-1, y, 1, h, borderColor, false)   // Right
	// Draw player
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the size of the game area
	return 640, 480 // Fixed size for simplicity
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Devide Game")
	game := &Game{
		player: player.NewPlayer(),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}