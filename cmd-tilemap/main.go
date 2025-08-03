package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

const mapPath = "../data/maps/world.tmx" // Path to your Tiled Map.
 

const (
	screenWidth  = 1024
	screenHeight = 768 
	AircraftID   = "Aircraft_1"
)

type Game struct {
	sprite   *ebiten_extended.Sprite
	rotation int
}

func NewGame() *Game {
	 	// Parse .tmx file.
	gameMap, err := tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}

	fmt.Println(gameMap)

	// You can also render the map to an in-memory image for direct
	// use with the default Renderer, or by making your own.
	renderer, err := render.NewRenderer(gameMap)
	if err != nil {
		fmt.Printf("map unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	// Render just layer 0 to the Renderer.
	err = renderer.RenderVisibleLayers()
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	    img := renderer.Result



	

	sprite := ebiten_extended.NewSprite("Aircraft_1", ebiten.NewImageFromImage(img), 0, true) 
 
 
	ebiten_extended.GameManager().World().AddNode(sprite)
	ebiten_extended.GameManager().SetIsDebug(false)

	// Clear the render result after copying the output if separation of
	// layers is desired.
	renderer.Clear()
	
	return &Game{sprite: sprite}
}

func (g *Game) Update() error {
	ebiten_extended.GameManager().Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebiten_extended.GameManager().Draw(screen, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sprite Example")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
