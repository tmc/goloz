package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	frame   int
	bgColor struct {
		r uint8
		g uint8
		b uint8
	}

	character Character
	gameMap   GameMap
}

func (g *Game) Update() error {
	g.frame++
	g.bgColor.r++
	g.bgColor.g++
	g.bgColor.b++

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.character.X--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.character.X++
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.character.Y--
		if g.frame%10 == 0 {
			g.character.spriteIndex--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.character.Y++
		if g.frame%10 == 0 {
			g.character.spriteIndex++
		}
	}
	if g.character.spriteIndex < 0 {
		g.character.spriteIndex = g.character.spriteIndex * -1
	}
	if g.character.spriteIndex >= 2 {
		g.character.spriteIndex = 0
	}
	inc := 1
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		inc = 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		g.gameMap.X -= inc
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		g.gameMap.X += inc
	}
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.gameMap.Y -= inc
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.gameMap.Y += inc
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{g.bgColor.r, g.bgColor.g, g.bgColor.b, 0xff})

	g.drawMap(screen)
	g.drawCharacter(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"(%v,%v) (%v,%v) sprite:%v",
		g.character.X, g.character.Y,
		g.gameMap.X, g.gameMap.Y,
		g.character.spriteIndex))
}

func (g *Game) drawCharacter(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.character.X), float64(g.character.Y))
	op.GeoM.Scale(1, 1)
	img, err := characterAsset(g.character.spriteIndex)
	if err != nil {
		log.Fatal(err)
	}
	screen.DrawImage(img, op)
}

func (g *Game) drawMap(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(-1*g.gameMap.X), float64(-1*g.gameMap.Y))
	op.GeoM.Scale(1, 1)
	screen.DrawImage(assetsMap, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{
		character: Character{
			X: 105,
			Y: 74,
		},
		gameMap: GameMap{
			X: 2116,
			Y: 3827,
		},
	}
	if err := loadAssets(); err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
