package game

import (
	zone17 "github.com/dkshi/zone-17"
	"github.com/dkshi/zone-17/internal/mainscene"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	textDPI  float64 = 100
	textSize float64 = 24
)

const (
	screenWidth     = 640
	screenHeight    = 360
	playerStartPosX = 40
	playerStartPosY = 200
	playerWidth     = 20
	playerHeight    = 31
)

type Game struct {
	Sprites       *zone17.Sprites
	SceneOnScreen zone17.GameScene
}

func InitGame() (*Game, error) {
	var err error

	// Initializing sprites
	sprites, err := zone17.NewSprites()
	if err != nil {
		return &Game{}, err
	}

	// Initializing font
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		return &Game{}, err
	}

	standartFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    textDPI,
		DPI:     textSize,
		Hinting: font.HintingFull,
	})

	if err != nil {
		return &Game{}, err
	}

	// Initializing player
	player := zone17.NewGameObject(sprites.Collection["playerModel"], playerStartPosX, playerStartPosY, playerWidth, playerHeight)

	// Initializing main scene
	mainScene, err := mainscene.InitMainScene(sprites, player, &standartFont)
	if err != nil {
		return &Game{}, err
	}

	return NewGame(mainScene, sprites), nil

}

func NewGame(sceneOnScreen zone17.GameScene, sprites *zone17.Sprites) *Game {
	return &Game{
		SceneOnScreen: sceneOnScreen,
		Sprites:       sprites,
	}
}

func (g *Game) Update() error {
	g.SceneOnScreen.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneOnScreen.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
