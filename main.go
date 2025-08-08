package main

import (
	"log"

	gamemap "github.com/Z-zenos/devide/internal/map"
	"github.com/Z-zenos/devide/internal/player"
	"github.com/hajimehoshi/ebiten/v2"
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
	gamemap.DrawMap(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the size of the game area
	return 640, 480 // Fixed size for simplicity
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Devide Game")

	gamemap.InitTileMap()

	game := &Game{
		player: player.NewPlayer(),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
