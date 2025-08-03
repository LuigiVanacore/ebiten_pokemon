package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LuigiVanacore/ebiten_extended"
	inputv3 "github.com/LuigiVanacore/ebiten_extended/input_v3"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_pokemon"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

const (
	screenWidth  = 1024
	screenHeight = 768
)

// func LoadTileMaps() map[string]*tiled.Map {
// 	var maps map[string]*tiled.Map = make(map[string]*tiled.Map)
// 	mapPath := "../data/maps/world.tmx"
// 	gameMap, err := tiled.LoadFile(mapPath)
// 	if err != nil {
// 		fmt.Printf("error parsing map: %s", err.Error())
// 		os.Exit(2)
// 	}
// 	maps["world"] = gameMap

// 	mapPath = "../data/maps/water.tmx"
// 	gameMap, err = tiled.LoadFile(mapPath)
// 	if err != nil {
// 		fmt.Printf("error parsing map: %s", err.Error())
// 		os.Exit(2)
// 	}
// 	maps["water"] = gameMap
// 	mapPath = "../data/maps/plant.tmx"
// 	gameMap, err = tiled.LoadFile(mapPath)
// 	if err != nil {
// 		fmt.Printf("error parsing map: %s", err.Error())
// 		os.Exit(2)
// 	}
// 	maps["plant"] = gameMap
// 	mapPath = "../data/maps/house.tmx"
// 	gameMap, err = tiled.LoadFile(mapPath)
// 	if err != nil {
// 		fmt.Printf("error parsing map: %s", err.Error())
// 		os.Exit(2)
// 	}
// 	maps["house"] = gameMap
// 	mapPath = "../data/maps/arena.tmx"
// 	gameMap, err = tiled.LoadFile(mapPath)
// 	if err != nil {
// 		fmt.Printf("error parsing map: %s", err.Error())
// 		os.Exit(2)
// 	}
// 	maps["arena"] = gameMap

// 	mapPath = "../data/maps/fire.tmx"
// 	gameMap, err = tiled.LoadFile(mapPath)
// 	if err != nil {
// 		fmt.Printf("error parsing map: %s", err.Error())
// 		os.Exit(2)
// 	}
// 	maps["fire"] = gameMap

// 	mapPath = "../data/maps/hospital.tmx"
// 	gameMap, err = tiled.LoadFile(mapPath)
// 	if err != nil {
// 		fmt.Printf("error parsing map: %s", err.Error())
// 		os.Exit(2)
// 	}
// 	maps["hospital"] = gameMap

// 	mapPath = "../data/maps/hospital2.tmx"
// 	gameMap, err = tiled.LoadFile(mapPath)
// 	if err != nil {
// 		fmt.Printf("error parsing map: %s", err.Error())
// 		os.Exit(2)
// 	}
// 	maps["hospital2"] = gameMap

// 	return maps
// }

// func Setup( start_map *tiled.Map, playerStartPos string) {

// 	var sprites []*ebiten_extended.Sprite = make([]*ebiten_extended.Sprite, 0)
// 	map_layer := start_map.Layers
// 	for _, layerName := range map_layer {
// 		if layerName.Name == "Terrain" || layerName.Name == "Terrain Top" {
// 			for _, tile := range layerName.Tiles {
// 				tileset := tile.Tileset
// 				sf, err := r.open(tile.Tileset.GetFileFullPath(tile.Tileset.Image.Source))
// 							if err != nil {
// 							return nil, err
// 						}
// 					defer sf.Close()

// 							img, _, err := image.Decode(sf)
// 							if err != nil {
// 								return nil, err
// 					}
// 				sprite := ebiten_extended.NewSprite(
// 					tileset.Name,
// 					img,
// 					ebiten_pokemon.BG_LAYER, false,
// 				)
// 				sprite.SetPosition(math2D.NewVector2D(float64(tileset.TileOffset.X*tileset.TileWidth), float64(tileset.TileOffset.Y*tileset.TileHeight)))
// 				sprites = append(sprites, sprite)
// 				ebiten_extended.GameManager().World().AddNode(sprite)
// 			}
// 		}
// 	}
// }

// 	def setup(self, tmx_map, player_start_pos):
// 		# clear the map
// 		for group in (self.all_sprites, self.collision_sprites, self.transition_sprites, self.character_sprites):
// 			group.empty()

// 		# terrain
// 		for layer in ['Terrain', 'Terrain Top']:
// 			for x, y, surf in tmx_map.get_layer_by_name(layer).tiles():
// 				Sprite((x * TILE_SIZE, y * TILE_SIZE), surf, self.all_sprites, WORLD_LAYERS['bg'])

// 		# water
// 		for obj in tmx_map.get_layer_by_name('Water'):
// 			for x in range(int(obj.x), int(obj.x + obj.width), TILE_SIZE):
// 				for y in range(int(obj.y), int(obj.y + obj.height), TILE_SIZE):
// 					AnimatedSprite((x,y), self.overworld_frames['water'], self.all_sprites, WORLD_LAYERS['water'])

// 		# coast
// 		for obj in tmx_map.get_layer_by_name('Coast'):
// 			terrain = obj.properties['terrain']
// 			side = obj.properties['side']
// 			AnimatedSprite((obj.x, obj.y), self.overworld_frames['coast'][terrain][side], self.all_sprites, WORLD_LAYERS['bg'])

// 		# objects
// 		for obj in tmx_map.get_layer_by_name('Objects'):
// 			if obj.name == 'top':
// 				Sprite((obj.x, obj.y), obj.image, self.all_sprites, WORLD_LAYERS['top'])
// 			else:
// 				CollidableSprite((obj.x, obj.y), obj.image, (self.all_sprites, self.collision_sprites))

// 		# transition objects
// 		for obj in tmx_map.get_layer_by_name('Transition'):
// 			TransitionSprite((obj.x, obj.y), (obj.width, obj.height), (obj.properties['target'], obj.properties['pos']), self.transition_sprites)

// 		# collision objects
// 		for obj in tmx_map.get_layer_by_name('Collisions'):
// 			BorderSprite((obj.x, obj.y), pygame.Surface((obj.width, obj.height)), self.collision_sprites)

// 		# grass patches
// 		for obj in tmx_map.get_layer_by_name('Monsters'):
// 			MonsterPatchSprite((obj.x, obj.y), obj.image, (self.all_sprites, self.monster_sprites), obj.properties['biome'], obj.properties['monsters'], obj.properties['level'])

// 		# entities
// 		for obj in tmx_map.get_layer_by_name('Entities'):
// 			if obj.name == 'Player':
// 				if obj.properties['pos'] == player_start_pos:
// 					self.player = Player(
// 						pos = (obj.x, obj.y),
// 						frames = self.overworld_frames['characters']['player'],
// 						groups = self.all_sprites,
// 						facing_direction = obj.properties['direction'],
// 						collision_sprites = self.collision_sprites)
// 			else:
// 				Character(
// 					pos = (obj.x, obj.y),
// 					frames = self.overworld_frames['characters'][obj.properties['graphic']],
// 					groups = (self.all_sprites, self.collision_sprites, self.character_sprites),
// 					facing_direction = obj.properties['direction'],
// 					character_data = TRAINER_DATA[obj.properties['character_id']],
// 					player = self.player,
// 					create_dialog = self.create_dialog,
// 					collision_sprites = self.collision_sprites,
// 					radius = obj.properties['radius'],
// 					nurse = obj.properties['character_id'] == 'Nurse',
// 					notice_sound = self.audio['notice'])

type Game struct {
	sprite *ebiten_extended.Sprite
	maps   map[string][]*tiled.Map
}

const (
	MOVE_UP  inputv3.ActionID = iota
	MOVE_DOWN
	MOVE_LEFT
	MOVE_RIGHT
)

func NewGame() *Game {


	move_up := inputv3.NewKeyAction(ebiten.KeyW, inputv3.Hold | inputv3.PressOnce)
	move_down := inputv3.NewKeyAction(ebiten.KeyS, inputv3.Hold | inputv3.PressOnce)
	move_left := inputv3.NewKeyAction(ebiten.KeyA, inputv3.Hold | inputv3.PressOnce)
	move_right := inputv3.NewKeyAction(ebiten.KeyD, inputv3.Hold | inputv3.PressOnce)
	 

	ebiten_extended.InputManager().RegisterAction(MOVE_UP, move_up)
	ebiten_extended.InputManager().RegisterAction(MOVE_DOWN, move_down)
	ebiten_extended.InputManager().RegisterAction(MOVE_LEFT, move_left)
	ebiten_extended.InputManager().RegisterAction(MOVE_RIGHT, move_right)
 
	// Parse .tmx file.
	// gameMap, err := tiled.LoadFile(mapPath)
	// if err != nil {
	// 	fmt.Printf("error parsing map: %s", err.Error())
	// 	os.Exit(2)
	// }

	// // You can also render the map to an in-memory image for direct
	// // use with the default Renderer, or by making your own.
	// renderer, err := render.NewRenderer(gameMap)
	// if err != nil {
	// 	fmt.Printf("map unsupported for rendering: %s", err.Error())
	// 	os.Exit(2)
	// }

	// // Render just layer 0 to the Renderer.
	// err = renderer.RenderVisibleLayers()
	// if err != nil {
	// 	fmt.Printf("layer unsupported for rendering: %s", err.Error())
	// 	os.Exit(2)
	// }

	//     img := renderer.Result

	// tileMaps := LoadTileMaps()

	// Setup(tileMaps["world"], "world")

	// Clear the render result after copying the output if separation of
	// layers is desired.
	// renderer.Clear()

	return &Game{}
}

func (g *Game) Update() error {
	g.Input()
	ebiten_extended.GameManager().Update()
	return nil
}

func (g *Game) Input() {
	if ebiten_extended.InputManager().IsActionActive(MOVE_UP) {
		ebiten_extended.GameManager().World().Camera().Translate(math2D.NewVector2D(0, -10))
	}
	if ebiten_extended.InputManager().IsActionActive(MOVE_DOWN) {
		ebiten_extended.GameManager().World().Camera().Translate(math2D.NewVector2D(0, 10))
	}
	if ebiten_extended.InputManager().IsActionActive(MOVE_LEFT) {
		ebiten_extended.GameManager().World().Camera().Translate(math2D.NewVector2D(-10, 0))
	}
	if ebiten_extended.InputManager().IsActionActive(MOVE_RIGHT) {
		ebiten_extended.GameManager().World().Camera().Translate(math2D.NewVector2D(10, 0))
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebiten_extended.GameManager().Draw(screen, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ebiten_pokemon.SCREEN_WIDTH, ebiten_pokemon.SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(ebiten_pokemon.SCREEN_WIDTH, ebiten_pokemon.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Ebiten Pokemon")

	filepath := "../data/maps/world.tmx"

	// tileMap := ebiten_pokemon.TileMap{}
	// if err := tileMap.Import(filepath); err != nil {
	// 	log.Fatal(err)
	// }

	// ebiten_extended.GameManager().World().AddNode(&tileMap)

	gameMap, err := tiled.LoadFile(filepath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}

	fmt.Println(gameMap)

	// You can also render the map to an in-memory image for direct
	// use with the default Renderer, or by making your own.
	// renderer, err := render.NewRenderer(gameMap)
	// if err != nil {
	// 	fmt.Printf("map unsupported for rendering: %s", err.Error())
	// 	os.Exit(2)
	// }

	// // Render just layer 0 to the Renderer.
	// err = renderer.RenderVisibleLayers()
	// if err != nil {
	// 	fmt.Printf("layer unsupported for rendering: %s", err.Error())
	// 	os.Exit(2)
	// }

	// img := renderer.Result

	tileMap := ebiten_pokemon.NewTileMap(gameMap)


	err = tileMap.RenderVisibleLayers()
	if err != nil {
		fmt.Printf("error rendering tile map: %s", err.Error())
		os.Exit(2)
	}

img := tileMap.Result

 

	sprite := ebiten_extended.NewSprite("Aircraft_1", ebiten.NewImageFromImage(img), ebiten_pokemon.BG_LAYER, false)
	//sprite.Node2D.SetPosition(math2D.NewVector2D(2720, 2720))
	
 
	ebiten_pokemon.LoadAssets()
	waterAnimationSet := ebiten_pokemon.LoadWaterAnimationSets()
	

	for _, objectGroup := range gameMap.ObjectGroups {
		if objectGroup.Name == "Water" {
			for _, object := range objectGroup.Objects {
				 
					for x := int(object.X); x < int(object.X+object.Width); x += gameMap.TileWidth {
						for y := int(object.Y); y < int(object.Y+object.Height); y += gameMap.TileHeight {
							animatedSprite := ebiten_extended.NewAnimationSprite("water", waterAnimationSet, ebiten_pokemon.WATER_LAYER, true)
							animatedSprite.SetPosition(math2D.NewVector2D(float64(x), float64(y)))
							ebiten_extended.GameManager().World().AddNode(animatedSprite)
						}
					}
				 
			}
		}
	}

// coast': coast_importer(24, 12, '..', 'graphics', 'tilesets', 'coast'),
// def coast_importer(cols, rows, *path):
// 	frame_dict = import_tilemap(cols, rows, *path)
// 	new_dict = {}
// 	terrains = ['grass', 'grass_i', 'sand_i', 'sand', 'rock', 'rock_i', 'ice', 'ice_i']
// 	sides = {
// 		'topleft': (0,0), 'top': (1,0), 'topright': (2,0), 
// 		'left': (0,1), 'right': (2,1), 'bottomleft': (0,2), 
// 		'bottom': (1,2), 'bottomright': (2,2)}
// 	for index, terrain in enumerate(terrains):
// 		new_dict[terrain] = {}
// 		for key, pos in sides.items():
// 			new_dict[terrain][key] = [frame_dict[(pos[0] + index * 3, pos[1] + row)] for row in range(0,rows, 3)]
// 	return new_dict


	// frames = {}
	// surf = import_image(*path)
	// cell_width, cell_height = surf.get_width() / cols, surf.get_height() / rows
	// for col in range(cols):
	// 	for row in range(rows):
	// 		cutout_rect = pygame.Rect(col * cell_width, row * cell_height,cell_width,cell_height)
	// 		cutout_surf = pygame.Surface((cell_width, cell_height))
	// 		cutout_surf.fill('green')
	// 		cutout_surf.set_colorkey('green')
	// 		cutout_surf.blit(surf, (0,0), cutout_rect)
	// 		frames[(col, row)] = cutout_surf
	// return frames

	

	terrains := []string{"grass", "grass_i", "sand_i", "sand", "rock", "rock_i", "ice", "ice_i"}
	sides := map[string][2]int{
		"topleft":    {0, 0},
		"top":        {1, 0},
		"topright":   {2, 0},
		"left":       {0, 1},
		"right":      {2, 1},
		"bottomleft": {0, 2},
		"bottom":     {1, 2},
		"bottomright": {2, 2},
	}

	frames_dict := make(map[[2]int]*ebiten.Image)

//	new_dict[terrain][key] = [frame_dict[(pos[0] + index * 3, pos[1] + row)] for row in range(0,rows, 3)]

	var new_dict = make(map[string]map[string][]*ebiten.Image)
	for index, terrain := range terrains {
		for key, pos := range sides {
			for row := 0; row < 12; row++ {
				new_dict[terrain][key] = append(new_dict[terrain][key], frames_dict[[2]int{pos[0] + index*3, pos[1] + row}])
			}
		}
	}

 			// animatedSprite := ebiten_extended.NewAnimationSprite("water", waterAnimationSet, ebiten_pokemon.WATER_LAYER, true)
			// 				animatedSprite.SetPosition(math2D.NewVector2D(0, 0))
			// 				ebiten_extended.GameManager().World().AddNode(animatedSprite)

		// for x in range(int(obj.x), int(obj.x + obj.width), TILE_SIZE):
		// 		for y in range(int(obj.y), int(obj.y + obj.height), TILE_SIZE):
		// 			AnimatedSprite((x,y), self.overworld_frames['water'], self.all_sprites, WORLD_LAYERS['water'])
 
	 

	ebiten_extended.GameManager().World().AddNode(sprite)
	ebiten_extended.GameManager().SetIsDebug(false)

	// Clear the render result after copying the output if separation of
	// layers is desired.
	tileMap.Clear()

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
 


// import (
//     "github.com/hajimehoshi/ebiten/v2"
//     "github.com/hajimehoshi/ebiten/v2/ebitenutil"
//     "image/color"
//     "log"
//     "math"
// )

// type Vector2 struct {
//     X, Y float64
// }

// func (v Vector2) Add(o Vector2) Vector2 {
//     return Vector2{v.X + o.X, v.Y + o.Y}
// }
// func (v Vector2) Mul(s float64) Vector2 {
//     return Vector2{v.X * s, v.Y * s}
// }
// func (v Vector2) Normalize() Vector2 {
//     l := math.Hypot(v.X, v.Y)
//     if l == 0 {
//         return Vector2{0, 0}
//     }
//     return Vector2{v.X / l, v.Y / l}
// }

// type Rect struct {
//     X, Y, W, H float64
// }

// func (r *Rect) Center() Vector2 {
//     return Vector2{r.X + r.W/2, r.Y + r.H/2}
// }
// func (r *Rect) SetCenter(v Vector2) {
//     r.X = v.X - r.W/2
//     r.Y = v.Y - r.H/2
// }
// func (r *Rect) Colliderect(o *Rect) bool {
//     return r.X < o.X+o.W && r.X+r.W > o.X && r.Y < o.Y+o.H && r.Y+r.H > o.Y
// }

// type Entity struct {
//     Rect     Rect
//     Dir      Vector2
//     Speed    float64
//     Color    color.Color
//     Blocked  bool
// }

// func (e *Entity) Move(dt float64, collisionRects []*Rect) {
//     // Horizontal
//     e.Rect.X += e.Dir.X * e.Speed * dt
//     for _, c := range collisionRects {
//         if e.Rect.Colliderect(c) {
//             if e.Dir.X > 0 {
//                 e.Rect.X = c.X - e.Rect.W
//             } else if e.Dir.X < 0 {
//                 e.Rect.X = c.X + c.W
//             }
//         }
//     }
//     // Vertical
//     e.Rect.Y += e.Dir.Y * e.Speed * dt
//     for _, c := range collisionRects {
//         if e.Rect.Colliderect(c) {
//             if e.Dir.Y > 0 {
//                 e.Rect.Y = c.Y - e.Rect.H
//             } else if e.Dir.Y < 0 {
//                 e.Rect.Y = c.Y + c.H
//             }
//         }
//     }
// }

// type Game struct {
//     Player   *Entity
//     Obstacles []*Rect
// }

// func (g *Game) Update() error {
//     // Input
//     dir := Vector2{}
//     if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
//         dir.Y -= 1
//     }
//     if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
//         dir.Y += 1
//     }
//     if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
//         dir.X -= 1
//     }
//     if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
//         dir.X += 1
//     }
//     g.Player.Dir = dir.Normalize()
//     g.Player.Move(1, g.Obstacles) // dt=1 for simplicity
//     return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
//     // Draw player
//     ebitenutil.DrawRect(screen, g.Player.Rect.X, g.Player.Rect.Y, g.Player.Rect.W, g.Player.Rect.H, g.Player.Color)
//     // Draw obstacles
//     for _, o := range g.Obstacles {
//         ebitenutil.DrawRect(screen, o.X, o.Y, o.W, o.H, color.RGBA{100, 100, 100, 255})
//     }
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
//     return 640, 480
// }

// func main() {
//     player := &Entity{
//         Rect:  Rect{X: 100, Y: 100, W: 32, H: 32},
//         Speed: 2.5,
//         Color: color.RGBA{0, 200, 0, 255},
//     }
//     obstacles := []*Rect{
//         {X: 200, Y: 200, W: 64, H: 64},
//         {X: 400, Y: 100, W: 32, H: 128},
//     }
//     game := &Game{
//         Player:   player,
//         Obstacles: obstacles,
//     }
//     ebiten.SetWindowSize(640, 480)
//     ebiten.SetWindowTitle("Go Game Structure Example")
//     if err := ebiten.RunGame(game); err != nil {
//         log.Fatal(err)
//     }
// }
