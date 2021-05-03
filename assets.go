package goloz

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assetFS embed.FS

var (
	assets map[string][]*ebiten.Image
)

func characterAsset(index int) *ebiten.Image {
	if index > len(assets["character0"])-1 {
		index = index % len(assets["character0"])
	}
	if index < 0 {
		index = index * -1
	}
	return assets["character0"][index]
}

func loadCharacterAssets() error {
	f, err := assetFS.Open("assets/character0alpha.png")
	if err != nil {
		return err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}
	rgba, ok := img.(*image.NRGBA)
	if !ok {
		return fmt.Errorf("loadCharacterAssets: not rgba")
	}
	characterWidth := 16
	characterHeight := 22
	points := []image.Point{
		// TODO: tweak these to match the image.
		// Row 1
		{1, 3},
		{19, 3},
		{36, 3},
		{53, 1},
		{70, 3},
		{87, 3},
		{104, 3},
		{121, 2},
		{138, 3},
		{156, 3},
		{173, 4},
		{190, 5},
		{207, 3},
		{224, 4},
		{241, 5},
		{259, 4},
		{276, 5},
		{294, 6},
		{311, 3},
		{329, 5},
		{346, 5},
		{363, 5},
		{382, 5},
		{399, 5},
		{416, 3},
		{433, 4},
		{450, 5},
		{467, 3},
		{484, 4},
		{501, 5},
		{519, 3},
		// Row 2
		{1, 32},
		{18, 33},
		{35, 34},
		{52, 37},
		{69, 34},
	}
	spriteIndices := []image.Rectangle{}
	for _, point := range points {
		spriteIndices = append(spriteIndices, image.Rectangle{
			Min: point,
			Max: image.Point{point.X + characterWidth, point.Y + characterHeight},
		})
	}
	assets["character0"] = make([]*ebiten.Image, len(spriteIndices))
	for i, rect := range spriteIndices {
		assets["character0"][i] = ebiten.NewImageFromImage(rgba.SubImage(rect))
	}
	return nil
}

func loadMapAssets() error {
	f, err := assetFS.Open("assets/map0.png")
	if err != nil {
		return err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}
	rgba, ok := img.(*image.Paletted)
	if !ok {
		return fmt.Errorf("mapAsset: not rgba")
	}
	// TODO: expand this to support more maps
	assets["map0"] = []*ebiten.Image{
		ebiten.NewImageFromImage(rgba),
	}
	return nil
}

func mapAsset(index int) *ebiten.Image {
	if index > len(assets["map0"])-1 {
		index = len(assets["map0"]) - 1
	}
	return assets["map0"][index]
}

func loadAssets() error {
	assets = make(map[string][]*ebiten.Image)
	if err := loadCharacterAssets(); err != nil {
		return err
	}
	if err := loadMapAssets(); err != nil {
		return err
	}
	return nil
}
