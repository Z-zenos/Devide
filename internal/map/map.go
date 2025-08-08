package gamemap

import (
	"image/color"

	constants "github.com/Z-zenos/devide/internal/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TileEmpty  = 0
	TileFilled = 1
	TileBorder = 2
	TileLine   = 3
	MapCols    = 640
	MapRows    = 480
)

var TileMap [MapRows][MapCols]int

func InitTileMap() {
	// Fill the tile map with empty tiles
	for y := range MapRows {
		for x := range MapCols {
			TileMap[y][x] = TileEmpty
		}
	}

	// Set the borders
	for x := 50; x <= 50+540; x++ {
		TileMap[50][x] = TileBorder
		TileMap[50+380][x] = TileBorder
	}

	for y := 50; y <= 50+380; y++ {
		TileMap[y][50] = TileBorder
		TileMap[y][50+540] = TileBorder
	}
}

func FloodFill(x, y int, from, to int) {
	if x < 0 || y < 0 || x >= MapCols || y >= MapRows {
		return
	}
	if TileMap[y][x] != from {
		return
	}
	TileMap[y][x] = to
	FloodFill(x+1, y, from, to)
	FloodFill(x-1, y, from, to)
	FloodFill(x, y+1, from, to)
	FloodFill(x, y-1, from, to)
}

func ApplyCapturedArea(path []constants.Vec2) {
	// Draw the path into the map as a wall (temporary)
	for _, pt := range path {
		x, y := int(pt.X), int(pt.Y)
		if x >= 0 && x < MapCols && y >= 0 && y < MapRows {
			TileMap[y][x] = TileLine
		}
	}
	// Assume starting from top-left is outside
	FloodFill(0, 0, TileEmpty, -1) // mark outside

	// All TileEmpty now is captured area → convert to TileFilled
	for y := range MapRows {
		for x := range MapCols {
			if TileMap[y][x] == TileEmpty {
				TileMap[y][x] = TileFilled
			}
		}
	}

	// Restore the outside to empty
	for y := range MapRows {
		for x := range MapCols {
			if TileMap[y][x] == -1 {
				TileMap[y][x] = TileEmpty
			}
		}
	}

	// Clear the temporary path
	for _, pt := range path {
		x, y := int(pt.X), int(pt.Y)
		if x >= 0 && x < MapCols && y >= 0 && y < MapRows {
			TileMap[y][x] = TileBorder
		}
	}
}

func DrawMap(screen *ebiten.Image) {
	for y := range MapRows {
		for x := range MapCols {
			switch TileMap[y][x] {
			case TileFilled:
				screen.Set(x, y, color.RGBA{50, 200, 50, 255}) // filled area
			case TileBorder:
				screen.Set(x, y, color.White) // border
			}
		}
	}
}
