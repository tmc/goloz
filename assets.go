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
		{1, 3},
		{19, 3},
		{36, 3},
		{53, 3},
		{70, 3},
		{87, 3},
		{104, 3},
		{121, 3},
		{138, 3},
		{155, 3},
		{173, 3},
		{190, 3},
		{207, 3},
		{224, 3},
		{241, 3},
		{259, 3},
		{276, 3},
		{295, 3},
		{312, 3},
		{329, 3},
		{346, 3},
		{363, 3},
		{382, 3},
		{399, 3},
		{416, 3},
		{433, 3},
		{450, 3},
		{467, 3},
		{484, 3},
		{501, 3},
		{515, 3},
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
