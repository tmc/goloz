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
		index = len(assets["character0"]) - 1
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
	spriteIndices := []image.Rectangle{
		{Min: image.Point{1, 3}, Max: image.Point{17, 27}},
		{Min: image.Point{19, 3}, Max: image.Point{35, 27}},
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
