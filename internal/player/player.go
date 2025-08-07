package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Vec2 struct {
	X, Y float32
}

type Player struct {
	X, Y float32
	Speed float32
	Size float32
	IsDrawing bool
	Path []Vec2
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
		IsDrawing: false,
		Path:  []Vec2{},
	}
}

func (p *Player) Update() {
	// Check if the player is trying to start drawing
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !p.IsDrawing && isOnEdge(p.X, p.Y) {
		p.IsDrawing = true
		p.Path = []Vec2{{X: p.X, Y: p.Y}} // Save the starting position
	}

	oldX, oldY := p.X, p.Y

	// Handle player movement based on input
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if (p.Y == MapY || p.Y == MapY + MapHeight && !p.IsDrawing) || (p.Y >= MapY && p.Y <= MapY + MapHeight && p.IsDrawing) {
			if p.X < MapX + MapWidth - 1 {
				p.X += p.Speed
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if (p.Y == MapY || p.Y == MapY + MapHeight && !p.IsDrawing) || (p.Y >= MapY && p.Y <= MapY + MapHeight && p.IsDrawing) {
			if p.X > MapX {
				p.X -= p.Speed
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if (p.X == MapX || p.X == MapX + MapWidth && !p.IsDrawing) || (p.X >= MapX && p.X <= MapX + MapWidth && p.IsDrawing) {
			if p.Y < MapY+MapHeight - 1 {
				p.Y += p.Speed
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if (p.X == MapX || p.X == MapX + MapWidth && !p.IsDrawing) || (p.X >= MapX && p.X <= MapX + MapWidth && p.IsDrawing) {
			if p.Y > MapY {
				p.Y -= p.Speed
			}
		}
	}

	// If the player is drawing and has moved, add the new position to the path
	if p.IsDrawing && (p.X != oldX || p.Y != oldY) {
		p.Path = append(p.Path, Vec2{X: p.X, Y: p.Y})
	}

	// If the player is drawing and is on the edge, stop drawing if the path has more than one point
	if p.IsDrawing && isOnEdge(p.X, p.Y) && len(p.Path) > 1 {
		p.IsDrawing = false
		// TODO: Handle the drawn path (e.g., save it, render it, etc.)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Draw player
	vector.DrawFilledRect(
		screen, 
		p.X - p.Size / 2,
		p.Y - p.Size / 2,
		p.Size,
		p.Size,
		color.RGBA{255, 0, 0, 255},
		false,
	)

	// Draw the path if the player is drawing
	if p.IsDrawing && len(p.Path) > 1 {
		for i := 0; i < len(p.Path) - 1; i++ {
			a := p.Path[i]
			b := p.Path[i+1]
			vector.StrokeLine(
				screen,
				a.X, a.Y,
				b.X, b.Y,
				1.5, color.RGBA{0, 200, 255, 255},
				false,
			)
		}
	}
}

func isOnEdge(x, y float32) bool {
	return (x == MapX || x == MapX + MapWidth) && (y >= MapY && y <= MapY + MapHeight) ||
		(y == MapY || y == MapY + MapHeight) && (x >= MapX && x <= MapX + MapWidth)
}