package goloz

import (
	"context"
	"fmt"
	"image/color"
	"io"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	pb "github.com/tmc/goloz/proto/goloz/v1"
)

// Game holds the local game state.
type Game struct {
	frame   int
	bgColor struct {
		r uint8
		g uint8
		b uint8
	}

	character *pb.Character
	gameMap   GameMap

	syncClient pb.GameServerService_SyncClient

	mu         sync.RWMutex // protects the following
	characters map[string]*pb.Character
}

func NewGame(ctx context.Context, syncClient pb.GameServerService_SyncClient) *Game {
	g := &Game{
		syncClient: syncClient,
		characters: make(map[string]*pb.Character),
		character: &pb.Character{
			Pos: &pb.Position{
				X: 105,
				Y: 74,
			},
		},
		gameMap: GameMap{
			X: 2116,
			Y: 3827,
		},
	}
	if err := loadAssets(); err != nil {
		log.Fatal(err)
	}
	return g
}

func (g *Game) Update() error {
	defer timeit("update")()
	g.frame++
	changed := false

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		changed = true
		g.character.Pos.X--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		changed = true
		g.character.Pos.X++
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		changed = true
		g.character.Pos.Y--
		if g.frame%10 == 0 {
			g.character.SpriteIndex--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		changed = true
		g.character.Pos.Y++
		if g.frame%10 == 0 {
			g.character.SpriteIndex++
		}
	}
	if g.character.SpriteIndex < 0 {
		g.character.SpriteIndex = g.character.SpriteIndex * -1
	}
	if g.character.SpriteIndex >= 2 {
		g.character.SpriteIndex = 0
	}

	if changed {
		err := g.syncClient.Send(&pb.SyncRequest{
			Character: g.character,
		})
		if err != nil {
			fmt.Println("sync issue:", err)
		}
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

func timeit(label string) func() {
	min := 5 * time.Millisecond
	t1 := time.Now()
	return func() {

		delta := time.Now().Sub(t1)
		if delta > min {
			fmt.Printf("timeit: %v %v\n", label, delta)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	defer timeit("draw")()
	screen.Fill(color.RGBA{g.bgColor.r, g.bgColor.g, g.bgColor.b, 0xff})

	g.drawMap(screen)
	g.drawCharacter(screen)
	g.drawCharacters(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"(tps: %.1f,fps:%.1f) sprite:%v",
		ebiten.CurrentTPS(), ebiten.CurrentFPS(),
		g.character.SpriteIndex))
}

func (g *Game) drawCharacter(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.character.Pos.X), float64(g.character.Pos.Y))
	op.GeoM.Scale(1, 1)
	img := characterAsset(int(g.character.SpriteIndex))
	screen.DrawImage(img, op)
}

func (g *Game) drawCharacters(screen *ebiten.Image) {
	defer timeit("drawchars")()
	localChars := map[string]*pb.Character{}
	g.mu.RLock()
	for k, v := range g.characters {
		localChars[k] = v
	}
	g.mu.RUnlock()
	for key, character := range g.characters {
		// TODO: filter out self
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(character.GetPos().X), float64(character.GetPos().Y))
		op.GeoM.Scale(1, 1)
		img := characterAsset(int(character.SpriteIndex))
		ebitenutil.DebugPrintAt(screen, key, int(character.GetPos().X)+25, int(character.GetPos().Y))
		screen.DrawImage(img, op)
	}
}

func (g *Game) drawMap(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(-1*g.gameMap.X), float64(-1*g.gameMap.Y))
	op.GeoM.Scale(1, 1)
	screen.DrawImage(mapAsset(0), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) RunNetworkSync(ctx context.Context, identity string) {
	for {
		m, err := g.syncClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("syncClient error:", err)
			continue
		}
		// log.Println(m)

		for key, character := range m.Characters {
			if key == identity {
				continue
			}
			g.mu.Lock()
			g.characters[key] = character
			g.mu.Unlock()
		}
	}

}
