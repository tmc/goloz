package main

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
	assetsCharacter *ebiten.Image
	assetsMap       *ebiten.Image
)

func characterAsset(index int) (*ebiten.Image, error) {
	f, err := assetFS.Open("assets/character0alpha.png")
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	rgba, ok := img.(*image.NRGBA)
	if !ok {
		return nil, fmt.Errorf("characterAsset: not rgba")
	}
	return ebiten.NewImageFromImage(rgba.SubImage([]image.Rectangle{
		{
			Min: image.Point{1, 3},
			Max: image.Point{17, 27},
		},
		{
			Min: image.Point{19, 3},
			Max: image.Point{35, 27},
		},
	}[index])), err
}

func mapAsset(index int) (*ebiten.Image, error) {
	f, err := assetFS.Open("assets/map0.png")
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	rgba, ok := img.(*image.Paletted)
	if !ok {
		return nil, fmt.Errorf("mapAsset: not rgba")
	}
	return ebiten.NewImageFromImage(rgba), nil
	/*
		return ebiten.NewImageFromImage(rgba.SubImage([]image.Rectangle{
			{
				Min: image.Point{0, 0},
				Max: image.Point{3020, 3020},
			},
		}[index])), nil
	*/
}

func loadAssets() error {
	var err error
	if assetsCharacter, err = characterAsset(0); err != nil {
		return err
	}
	if assetsMap, err = mapAsset(0); err != nil {
		return err
	}
	return nil
}
