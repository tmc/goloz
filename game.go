package goloz

import (
	"bytes"
	"context"
	"fmt"
	"image/color"
	"io"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	pb "github.com/tmc/goloz/proto/goloz/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Game holds the local game state.
type Game struct {
	settings Settings

	frame   int
	bgColor struct {
		r uint8
		g uint8
		b uint8
	}

	audioContext *audio.Context
	audioPlayer  *audio.Player

	character *pb.Character
	gameMap   GameMap

	client pb.GameServerServiceClient

	// sync stream
	syncClient pb.GameServerService_SyncClient

	mu         sync.RWMutex // protects the following
	characters map[string]*pb.Character
}

// NewGame initializes a new Game.
func NewGame(ctx context.Context, settings Settings, client pb.GameServerServiceClient) (*Game, error) {
	g := &Game{
		settings:   settings,
		client:     client,
		characters: make(map[string]*pb.Character),
		character: &pb.Character{
			Pos: &pb.Position{
				X: 20,
				Y: 20,
			},
		},
		gameMap: GameMap{
			X: 2116,
			Y: 3827,
		},
	}

	// TODO: separate audio initialization.
	sampleRate := 44100
	g.audioContext = audio.NewContext(sampleRate)
	startupAudio, err := assetFS.ReadFile("assets/audio/title.mp3")
	if err != nil {
		return nil, err
	}
	d, err := mp3.Decode(g.audioContext, bytes.NewReader(startupAudio))
	if err != nil {
		return nil, err
	}
	// Create an audio.Player that has one stream.
	g.audioPlayer, err = audio.NewPlayer(g.audioContext, d)
	if err != nil {
		return nil, err
	}

	if err := loadAssets(); err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Update() error {
	// TODO: audioPlayer should be renamed to something like startupAudioPlayer.
	if g.frame == 0 && g.settings.AudioMuted == false {
		g.audioPlayer.Play()
	}

	g.frame++
	changed := false

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		changed = true
		g.character.Pos.X--
		if g.frame%5 == 0 {
			g.character.SpriteIndex--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		changed = true
		g.character.Pos.X++
		if g.frame%5 == 0 {
			g.character.SpriteIndex++
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		changed = true
		g.character.Pos.Y--
		if g.frame%5 == 0 {
			g.character.SpriteIndex--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		changed = true
		g.character.Pos.Y++
		if g.frame%5 == 0 {
			g.character.SpriteIndex++
		}
	}
	if g.character.SpriteIndex < 0 {
		g.character.SpriteIndex = g.character.SpriteIndex * -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeftBracket) {
		if g.frame%5 == 0 {
			g.character.SpriteIndex--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRightBracket) {
		if g.frame%5 == 0 {
			g.character.SpriteIndex++
		}
	}

	if changed && g.syncClient != nil {
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

func (g *Game) RunNetworkSync(ctx context.Context) {
	fmt.Println("syncing as", g.settings.UserIdentity)
	ctx = metadata.AppendToOutgoingContext(ctx,
		"id", g.settings.UserIdentity,
	)
	delay := time.Second
	for {
		// set up syncClient
		var err error
		g.syncClient, err = g.client.Sync(ctx)
		if err != nil {
			fmt.Println("issue obtaining sync stream:", err)
			time.Sleep(delay)
			continue
		}
		g.runNetworkSync(ctx, g.syncClient)
	}
}

func (g *Game) runNetworkSync(ctx context.Context, syncClient pb.GameServerService_SyncClient) {
	for {
		m, err := syncClient.Recv()
		if err == io.EOF {
			break
		}
		if status.Code(err) == codes.Internal {
			log.Println("syncClient internal error:", err)
			time.Sleep(100 * time.Millisecond)
			break
		}
		if err != nil {
			log.Println("syncClient error:", err)
			time.Sleep(100 * time.Millisecond)
			break
		}

		for key, character := range m.Characters {
			if key == g.settings.UserIdentity {
				continue
			}
			g.mu.Lock()
			g.characters[key] = character
			g.mu.Unlock()
		}
	}
}
