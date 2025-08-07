package player

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y float32
	Speed float32
	Size float32
}

const (
	MapX = 50
	MapY = 50
	MapWidth = 540
	MapHeight = 380
)

func NewPlayer() *Player {
	return &Player{
		X:	 MapX,
		Y:	 MapY,
		Speed: 2,
		Size:  6,
	}
}

func (p *Player) Update() {
	log.Println(p.X, p.Y, " -- ", MapX, MapY, MapWidth, MapHeight)
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if p.Y == MapY || p.Y == MapY + MapHeight {
			if p.X < MapX + MapWidth - 1 {
				p.X += p.Speed
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if p.Y == MapY || p.Y == MapY + MapHeight {
			if p.X > MapX {
				p.X -= p.Speed
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if p.X == MapX || p.X == MapX + MapWidth {
			if p.Y < MapY+MapHeight - 1 {
				p.Y += p.Speed
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if p.X == MapX || p.X == MapX + MapWidth {
			if p.Y > MapY {
				p.Y -= p.Speed
			}
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen, 
		p.X - p.Size / 2,
		p.Y - p.Size / 2,
		p.Size,
		p.Size,
		color.RGBA{255, 0, 0, 255},
		false,
	)
}