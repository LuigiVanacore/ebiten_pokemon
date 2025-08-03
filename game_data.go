package ebiten_pokemon

import (
	"image" 

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	resources "github.com/LuigiVanacore/ebiten_pokemon/graphics"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Attacks_explosion             string = "explosion"
	Attacks_fire                  string = "fire"
	Attacks_green                 string = "green"
	Attacks_ice                   string = "ice"
	Attacks_scratch               string = "scratch"
	Attacks_splash                string = "splash"
	Backgrounds_forest            string = "forest"
	Backgrounds_ice               string = "ice"
	Backgrounds_sand              string = "sand"
	Characters_blond              string = "blond"
	Characters_fire_boss          string = "fire_boss"
	Characters_grass_boss         string = "grass_boss"
	Characters_hat_girl           string = "hat_girl"
	Characters_player             string = "player"
	Characters_purple_girl        string = "purple_girl"
	Characters_straw              string = "straw"
	Characters_water_boss         string = "water_boss"
	Characters_young_girl         string = "young_girl"
	Characters_young_guy          string = "young_guy"
	Cmd___debug_bin2272826165     string = "__debug_bin2272826165"
	Cmd_script                    string = "script"
	Fonts_PixeloidSans            string = "PixeloidSans"
	Fonts_dogicapixel             string = "dogicapixel"
	Fonts_dogicapixelbold         string = "dogicapixelbold"
	Icons_Atrox                   string = "Atrox"
	Icons_Charmadillo             string = "Charmadillo"
	Icons_Cindrill                string = "Cindrill"
	Icons_Cleaf                   string = "Cleaf"
	Icons_Draem                   string = "Draem"
	Icons_Finiette                string = "Finiette"
	Icons_Finsta                  string = "Finsta"
	Icons_Friolera                string = "Friolera"
	Icons_Gulfin                  string = "Gulfin"
	Icons_Ivieron                 string = "Ivieron"
	Icons_Jacana                  string = "Jacana"
	Icons_Larvea                  string = "Larvea"
	Icons_Pluma                   string = "Pluma"
	Icons_Plumette                string = "Plumette"
	Icons_Pouch                   string = "Pouch"
	Icons_Sparchu                 string = "Sparchu"
	Monsters_Atrox                string = "Atrox"
	Monsters_Charmadillo          string = "Charmadillo"
	Monsters_Cindrill             string = "Cindrill"
	Monsters_Cleaf                string = "Cleaf"
	Monsters_Draem                string = "Draem"
	Monsters_Finiette             string = "Finiette"
	Monsters_Finsta               string = "Finsta"
	Monsters_Friolera             string = "Friolera"
	Monsters_Gulfin               string = "Gulfin"
	Monsters_Ivieron              string = "Ivieron"
	Monsters_Jacana               string = "Jacana"
	Monsters_Larvea               string = "Larvea"
	Monsters_Pluma                string = "Pluma"
	Monsters_Plumette             string = "Plumette"
	Monsters_Pouch                string = "Pouch"
	Monsters_Sparchu              string = "Sparchu"
	Objects_arean_fire            string = "arean_fire"
	Objects_arena_plant           string = "arena_plant"
	Objects_arena_water           string = "arena_water"
	Objects_gate_pillar           string = "gate_pillar"
	Objects_gate_top              string = "gate_top"
	Objects_grass                 string = "grass"
	Objects_grass_ice             string = "grass_ice"
	Objects_grassrock1            string = "grassrock1"
	Objects_grassrock2            string = "grassrock2"
	Objects_green_tree            string = "green_tree"
	Objects_green_tree_bushy      string = "green_tree_bushy"
	Objects_green_tree_small      string = "green_tree_small"
	Objects_hospital              string = "hospital"
	Objects_house_large           string = "house_large"
	Objects_house_large_alt       string = "house_large_alt"
	Objects_house_small           string = "house_small"
	Objects_house_small_alt       string = "house_small_alt"
	Objects_ice_tree              string = "ice_tree"
	Objects_icerock1              string = "icerock1"
	Objects_icerock2              string = "icerock2"
	Objects_palm                  string = "palm"
	Objects_palm_alt              string = "palm_alt"
	Objects_palm_small            string = "palm_small"
	Objects_ruin_gate             string = "ruin_gate"
	Objects_ruin_pillar           string = "ruin_pillar"
	Objects_ruin_pillar_broke     string = "ruin_pillar_broke"
	Objects_ruin_pillar_broke_alt string = "ruin_pillar_broke_alt"
	Objects_sand                  string = "sand"
	Objects_sandrock1             string = "sandrock1"
	Objects_sandrock2             string = "sandrock2"
	Objects_teal_tree             string = "teal_tree"
	Objects_teal_tree_bushy       string = "teal_tree_bushy"
	Objects_teal_tree_small       string = "teal_tree_small"
	Other_shadow                  string = "shadow"
	Other_star_animation          string = "star_animation"
	Other_star_animation1         string = "star_animation1"
	Other_star_animation2         string = "star_animation2"
	Other_star_animation3         string = "star_animation3"
	Other_star_animation4         string = "star_animation4"
	Other_star_animation5         string = "star_animation5"
	Other_star_animation6         string = "star_animation6"
	Other_star_animation7         string = "star_animation7"
	Other_star_animation8         string = "star_animation8"
	Other_star_animation9         string = "star_animation9"
	Other_star_animation10        string = "star_animation10"
	Other_star_animation11        string = "star_animation11"
	Other_star_animation12        string = "star_animation12"
	Other_star_animation13        string = "star_animation13"
	Other_star_animation14        string = "star_animation14"
	Other_star_animation15        string = "star_animation15"
	Other_star_animation16        string = "star_animation16"
	Other_star_animation17        string = "star_animation17"
	Other_star_animation18        string = "star_animation18"
	Other_star_animation19        string = "star_animation19"
	Other_star_animation20        string = "star_animation20"
	Other_star_animation21        string = "star_animation21"
	Other_star_animation22        string = "star_animation22"
	Other_star_animation23        string = "star_animation23"
	Other_star_animation24        string = "star_animation24"
	Other_star_animation25        string = "star_animation25"
	Other_star_animation26        string = "star_animation26"
	Other_star_animation27        string = "star_animation27"
	Tilesets_coast                string = "coast"
	Tilesets_indoor               string = "indoor"
	Tilesets_water0               string = "water0"
	Tilesets_water1               string = "water1"
	Tilesets_water2               string = "water2"
	Tilesets_water3               string = "water3"
	Tilesets_world                string = "world"
	Ui_arrows                     string = "arrows"
	Ui_arrows_highlight           string = "arrows_highlight"
	Ui_attack                     string = "attack"
	Ui_cross                      string = "cross"
	Ui_defense                    string = "defense"
	Ui_energy                     string = "energy"
	Ui_hand                       string = "hand"
	Ui_hand_highlight             string = "hand_highlight"
	Ui_health                     string = "health"
	Ui_notice                     string = "notice"
	Ui_recovery                   string = "recovery"
	Ui_shield                     string = "shield"
	Ui_shield_highlight           string = "shield_highlight"
	Ui_speed                      string = "speed"
	Ui_star                       string = "star"
	Ui_sword                      string = "sword"
	Ui_sword_highlight            string = "sword_highlight"
)

func LoadAssets() {
	LoadImages()
}



func LoadImages() {
	//Tilesets

	//Tilesets water
	ebiten_extended.ResourceManager().AddImage(Tilesets_water0, resources.Tilesets_water0)
	ebiten_extended.ResourceManager().AddImage(Tilesets_water1, resources.Tilesets_water1)
	ebiten_extended.ResourceManager().AddImage(Tilesets_water2, resources.Tilesets_water2)
	ebiten_extended.ResourceManager().AddImage(Tilesets_water3, resources.Tilesets_water3)

	//Tilesets coast
	ebiten_extended.ResourceManager().AddImage(Tilesets_coast, resources.Tilesets_coast)


	//Characters Spritesheets
	ebiten_extended.ResourceManager().AddImage(Characters_blond, resources.Characters_blond)
	ebiten_extended.ResourceManager().AddImage(Characters_fire_boss, resources.Characters_fire_boss)
	ebiten_extended.ResourceManager().AddImage(Characters_grass_boss, resources.Characters_grass_boss)
	ebiten_extended.ResourceManager().AddImage(Characters_hat_girl, resources.Characters_hat_girl)
	ebiten_extended.ResourceManager().AddImage(Characters_player, resources.Characters_player)
	ebiten_extended.ResourceManager().AddImage(Characters_purple_girl, resources.Characters_purple_girl)
	ebiten_extended.ResourceManager().AddImage(Characters_straw, resources.Characters_straw)
	ebiten_extended.ResourceManager().AddImage(Characters_water_boss, resources.Characters_water_boss)
	ebiten_extended.ResourceManager().AddImage(Characters_young_girl, resources.Characters_young_girl)
	ebiten_extended.ResourceManager().AddImage(Characters_young_guy, resources.Characters_young_guy)


	//Monsters

	//Monster icons
	ebiten_extended.ResourceManager().AddImage(Icons_Atrox, resources.Icons_Atrox)
	ebiten_extended.ResourceManager().AddImage(Icons_Charmadillo, resources.Icons_Charmadillo)
	ebiten_extended.ResourceManager().AddImage(Icons_Cindrill, resources.Icons_Cindrill)
	ebiten_extended.ResourceManager().AddImage(Icons_Cleaf, resources.Icons_Cleaf)
	ebiten_extended.ResourceManager().AddImage(Icons_Draem, resources.Icons_Draem)
	ebiten_extended.ResourceManager().AddImage(Icons_Finiette, resources.Icons_Finiette)
	ebiten_extended.ResourceManager().AddImage(Icons_Finsta, resources.Icons_Finsta)
	ebiten_extended.ResourceManager().AddImage(Icons_Friolera, resources.Icons_Friolera)
	ebiten_extended.ResourceManager().AddImage(Icons_Gulfin, resources.Icons_Gulfin)
	ebiten_extended.ResourceManager().AddImage(Icons_Ivieron, resources.Icons_Ivieron)
	ebiten_extended.ResourceManager().AddImage(Icons_Jacana, resources.Icons_Jacana)
	ebiten_extended.ResourceManager().AddImage(Icons_Larvea, resources.Icons_Larvea)
	ebiten_extended.ResourceManager().AddImage(Icons_Pluma, resources.Icons_Pluma) 
	ebiten_extended.ResourceManager().AddImage(Icons_Plumette, resources.Icons_Plumette)
	ebiten_extended.ResourceManager().AddImage(Icons_Pouch, resources.Icons_Pouch)
	ebiten_extended.ResourceManager().AddImage(Icons_Sparchu, resources.Icons_Sparchu)
	

	//Monsters Spritesheets
	ebiten_extended.ResourceManager().AddImage(Monsters_Atrox, resources.Monsters_Atrox)
	ebiten_extended.ResourceManager().AddImage(Monsters_Charmadillo, resources.Monsters_Charmadillo)
	ebiten_extended.ResourceManager().AddImage(Monsters_Cindrill, resources.Monsters_Cindrill)
	ebiten_extended.ResourceManager().AddImage(Monsters_Cleaf, resources.Monsters_Cleaf)
	ebiten_extended.ResourceManager().AddImage(Monsters_Draem, resources.Monsters_Draem)
	ebiten_extended.ResourceManager().AddImage(Monsters_Finiette, resources.Monsters_Finiette)
	ebiten_extended.ResourceManager().AddImage(Monsters_Finsta, resources.Monsters_Finsta)
	ebiten_extended.ResourceManager().AddImage(Monsters_Friolera, resources.Monsters_Friolera)
	ebiten_extended.ResourceManager().AddImage(Monsters_Gulfin, resources.Monsters_Gulfin)
	ebiten_extended.ResourceManager().AddImage(Monsters_Ivieron, resources.Monsters_Ivieron)
	ebiten_extended.ResourceManager().AddImage(Monsters_Jacana, resources.Monsters_Jacana)
	ebiten_extended.ResourceManager().AddImage(Monsters_Larvea, resources.Monsters_Larvea)
	ebiten_extended.ResourceManager().AddImage(Monsters_Pluma, resources.Monsters_Pluma)
	ebiten_extended.ResourceManager().AddImage(Monsters_Plumette, resources.Monsters_Plumette)
	ebiten_extended.ResourceManager().AddImage(Monsters_Pouch, resources.Monsters_Pouch)
	ebiten_extended.ResourceManager().AddImage(Monsters_Sparchu, resources.Monsters_Sparchu)

	//Objects
	ebiten_extended.ResourceManager().AddImage(Objects_arean_fire, resources.Objects_arean_fire)
	ebiten_extended.ResourceManager().AddImage(Objects_arena_plant, resources.Objects_arena_plant)
	ebiten_extended.ResourceManager().AddImage(Objects_arena_water, resources.Objects_arena_water)
	ebiten_extended.ResourceManager().AddImage(Objects_gate_pillar, resources.Objects_gate_pillar)
	ebiten_extended.ResourceManager().AddImage(Objects_gate_top, resources.Objects_gate_top)
	ebiten_extended.ResourceManager().AddImage(Objects_grass, resources.Objects_grass)
	ebiten_extended.ResourceManager().AddImage(Objects_grass_ice, resources.Objects_grass_ice)
	ebiten_extended.ResourceManager().AddImage(Objects_grassrock1, resources.Objects_grassrock1)
	ebiten_extended.ResourceManager().AddImage(Objects_grassrock2, resources.Objects_grassrock2)
	ebiten_extended.ResourceManager().AddImage(Objects_green_tree, resources.Objects_green_tree)
	ebiten_extended.ResourceManager().AddImage(Objects_green_tree_bushy, resources.Objects_green_tree_bushy)
	ebiten_extended.ResourceManager().AddImage(Objects_green_tree_small, resources.Objects_green_tree_small)
	ebiten_extended.ResourceManager().AddImage(Objects_hospital, resources.Objects_hospital)
	ebiten_extended.ResourceManager().AddImage(Objects_house_large, resources.Objects_house_large)
	ebiten_extended.ResourceManager().AddImage(Objects_house_large_alt, resources.Objects_house_large_alt)
	ebiten_extended.ResourceManager().AddImage(Objects_house_small, resources.Objects_house_small)
	ebiten_extended.ResourceManager().AddImage(Objects_house_small_alt, resources.Objects_house_small_alt)
	ebiten_extended.ResourceManager().AddImage(Objects_ice_tree, resources.Objects_ice_tree)
	ebiten_extended.ResourceManager().AddImage(Objects_icerock1, resources.Objects_icerock1)
	ebiten_extended.ResourceManager().AddImage(Objects_icerock2, resources.Objects_icerock2)
	ebiten_extended.ResourceManager().AddImage(Objects_palm, resources.Objects_palm)
	ebiten_extended.ResourceManager().AddImage(Objects_palm_alt, resources.Objects_palm_alt)
	ebiten_extended.ResourceManager().AddImage(Objects_palm_small, resources.Objects_palm_small)
	ebiten_extended.ResourceManager().AddImage(Objects_ruin_gate, resources.Objects_ruin_gate)
	ebiten_extended.ResourceManager().AddImage(Objects_ruin_pillar, resources.Objects_ruin_pillar)
	ebiten_extended.ResourceManager().AddImage(Objects_ruin_pillar_broke, resources.Objects_ruin_pillar_broke)
	ebiten_extended.ResourceManager().AddImage(Objects_ruin_pillar_broke_alt,	 resources.Objects_ruin_pillar_broke_alt)
	ebiten_extended.ResourceManager().AddImage(Objects_sand, resources.Objects_sand)
	ebiten_extended.ResourceManager().AddImage(Objects_sandrock1, resources.Objects_sandrock1)
	ebiten_extended.ResourceManager().AddImage(Objects_sandrock2, resources.Objects_sandrock2)
	ebiten_extended.ResourceManager().AddImage(Objects_teal_tree, resources.Objects_teal_tree)
	ebiten_extended.ResourceManager().AddImage(Objects_teal_tree_bushy, resources.Objects_teal_tree_bushy)
	ebiten_extended.ResourceManager().AddImage(Objects_teal_tree_small, resources.Objects_teal_tree_small)

	//Other
	ebiten_extended.ResourceManager().AddImage(Other_shadow, resources.Other_shadow)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation, resources.Other_star_animation_00002)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation1, resources.Other_star_animation_00003)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation2, resources.Other_star_animation_00004)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation3, resources.Other_star_animation_00005)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation4, resources.Other_star_animation_00006)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation5, resources.Other_star_animation_00007)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation6, resources.Other_star_animation_00008)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation7, resources.Other_star_animation_00009)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation8, resources.Other_star_animation_00010)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation9, resources.Other_star_animation_00011)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation10, resources.Other_star_animation_00012)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation11, resources.Other_star_animation_00013)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation12, resources.Other_star_animation_00014)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation13, resources.Other_star_animation_00015)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation14, resources.Other_star_animation_00016)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation15, resources.Other_star_animation_00017)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation16, resources.Other_star_animation_00018)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation17, resources.Other_star_animation_00019)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation18, resources.Other_star_animation_00020)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation19, resources.Other_star_animation_00021)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation20, resources.Other_star_animation_00022)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation21, resources.Other_star_animation_00023)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation22, resources.Other_star_animation_00024)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation23, resources.Other_star_animation_00025)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation24, resources.Other_star_animation_00026)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation25, resources.Other_star_animation_00027)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation26, resources.Other_star_animation_00028)
	ebiten_extended.ResourceManager().AddImage(Other_star_animation27, resources.Other_star_animation_00029)

	//Ui
	ebiten_extended.ResourceManager().AddImage(Ui_arrows, resources.Ui_arrows)
	ebiten_extended.ResourceManager().AddImage(Ui_arrows_highlight, resources.Ui_arrows_highlight)
	ebiten_extended.ResourceManager().AddImage(Ui_attack, resources.Ui_attack)
	ebiten_extended.ResourceManager().AddImage(Ui_cross, resources.Ui_cross)
	ebiten_extended.ResourceManager().AddImage(Ui_defense, resources.Ui_defense)
	ebiten_extended.ResourceManager().AddImage(Ui_energy, resources.Ui_energy)
	ebiten_extended.ResourceManager().AddImage(Ui_hand, resources.Ui_hand)
	ebiten_extended.ResourceManager().AddImage(Ui_hand_highlight, resources.Ui_hand_highlight)
	ebiten_extended.ResourceManager().AddImage(Ui_health, resources.Ui_health)
	ebiten_extended.ResourceManager().AddImage(Ui_notice, resources.Ui_notice)
	ebiten_extended.ResourceManager().AddImage(Ui_recovery, resources.Ui_recovery)
	ebiten_extended.ResourceManager().AddImage(Ui_shield, resources.Ui_shield)
	ebiten_extended.ResourceManager().AddImage(Ui_shield_highlight, resources.Ui_shield_highlight)
	ebiten_extended.ResourceManager().AddImage(Ui_speed, resources.Ui_speed)
	ebiten_extended.ResourceManager().AddImage(Ui_star, resources.Ui_star)
	ebiten_extended.ResourceManager().AddImage(Ui_sword, resources.Ui_sword)
	ebiten_extended.ResourceManager().AddImage(Ui_sword_highlight, resources.Ui_sword_highlight)

 
ebiten_extended.ResourceManager().GetImage(Ui_arrows_highlight)

	//Backgrounds
	ebiten_extended.ResourceManager().AddImage(Backgrounds_forest, resources.Backgrounds_forest)
	ebiten_extended.ResourceManager().AddImage(Backgrounds_ice, resources.Backgrounds_ice)
	ebiten_extended.ResourceManager().AddImage(Backgrounds_sand, resources.Backgrounds_sand)
}


const (
	ANIM_DOWN = iota
	ANIM_LEFT
	ANIM_RIGHT
	ANIM_UP
)



// 	// Define the mapping of key strings to their corresponding image keys
// 	animmationSets := map[string][]string{
// 		Character_Down:      {Character_Down_0, Character_Down_1, Character_Down_2, Character_Down_3},
// 		Character_Down_Axe:  {Character_Down_Axe_0, Character_Down_Axe_1},
// 		Character_Down_Hoe:  {Character_Down_Hoe_0, Character_Down_Hoe_1},
// 		Character_Down_Idle: {Character_Down_Idle_0, Character_Down_Idle_1},
// 		Character_Down_Water: {Character_Down_Water_0, Character_Down_Water_1},
// 		Character_Left:      {Character_Left_0, Character_Left_1, Character_Left_2, Character_Left_3},
// 		Character_Left_Axe:  {Character_Left_Axe_0, Character_Left_Axe_1},
// 		Character_Left_Hoe:  {Character_Left_Hoe_0, Character_Left_Hoe_1},
// 		Character_Left_Idle: {Character_Left_Idle_0, Character_Left_Idle_1},
// 		Character_Left_Water: {Character_Left_Water_0, Character_Left_Water_1},
// 		Character_Right:     {Character_Right_0, Character_Right_1, Character_Right_2, Character_Right_3},
// 		Character_Right_Axe: {Character_Right_Axe_0, Character_Right_Axe_1},
// 		Character_Right_Hoe: {Character_Right_Hoe_0, Character_Right_Hoe_1},
// 		Character_Right_Idle: {Character_Right_Idle_0, Character_Right_Idle_1},
// 		Character_Right_Water: {Character_Right_Water_0, Character_Right_Water_1},
// 		Character_Up:        {Character_Up_0, Character_Up_1, Character_Up_2, Character_Up_3},
// 		Character_Up_Axe:    {Character_Up_Axe_0, Character_Up_Axe_1},
// 		Character_Up_Hoe:    {Character_Up_Hoe_0, Character_Up_Hoe_1},
// 		Character_Up_Idle:   {Character_Up_Idle_0, Character_Up_Idle_1},
// 		Character_Up_Water:  {Character_Up_Water_0, Character_Up_Water_1},
// 	}

// 	var spriteSheet []*ebiten.Image

// 	// Load images into the AnimationSet
// 	for key, animationKeys := range animmationSets {
// 		spriteSheet = nil
// 		for _, animationKey := range animationKeys {
// 			img := ebiten_extended.ResourceManager().GetImage(animationKey)
// 			if img != nil {
// 				spriteSheet = append(spriteSheet, img)
// 			}
// 		}
// 		ebiten_extended.ResourceManager().AddAnimation(key, ebiten_extended.NewAnimationSet(spriteSheet, getCenterImage(spriteSheet[0]), uint(len(spriteSheet)), float64(len(spriteSheet))  , true))
// 	}
// }

func getCenterImage(img *ebiten.Image) math2D.Vector2D {
	return math2D.NewVector2D(float64(img.Bounds().Dx()/2), float64(img.Bounds().Dy()/2))
}	


func LoadAllCharacterAnimations() {
	characters_spritesheets := []string{
		Characters_blond,
		Characters_fire_boss,
		Characters_grass_boss,
		Characters_hat_girl,
		Characters_player,
		Characters_purple_girl,
		Characters_straw,
		Characters_water_boss,
		Characters_young_girl,
		Characters_young_guy,
	}

	var characterAnimations = make(map[string]map[int]*ebiten_extended.AnimationSet)
	for _, character := range characters_spritesheets {
		spriteSheet := ebiten_extended.ResourceManager().GetImage(character)
		if spriteSheet != nil {
			characterAnimations[character] = LoadCharacterAnimation(spriteSheet)
		}
	}

}

func LoadCharacterAnimation(spriteSheet *ebiten.Image) map[int]*ebiten_extended.AnimationSet {
	animationMap := make(map[int]*ebiten_extended.AnimationSet)

	colls := 4
	rows := 4
	cell_width := spriteSheet.Bounds().Dx() / colls
	cell_height := spriteSheet.Bounds().Dy() / rows

	for row := 0; row < rows; row++ {
	var direction int
	var frameSet []*ebiten.Image = make([]*ebiten.Image, 0, colls)	
		for col := 0; col < colls; col++ {
			cutoutRect := image.Rect(
				col*cell_width,
				row*cell_height,
				(col+1)*cell_width,
				(row+1)*cell_height,
			)
			frameSet = append(frameSet, ebiten.NewImageFromImage(spriteSheet.SubImage(cutoutRect).(*ebiten.Image)))

			// Create an animation set for each direction
			
			switch row {
			case 0:
				direction = ANIM_DOWN
			case 1:
				direction = ANIM_LEFT
			case 2:
				direction = ANIM_RIGHT
			case 3:
				direction = ANIM_UP
			}
		}

			if _, exists := animationMap[direction]; !exists {
				animationMap[direction] = ebiten_extended.NewAnimationSet(frameSet, getCenterImage(frameSet[0]), uint(colls), float64(colls), true)
			}
		
	}
	// Load character animations from the sprite sheet
	return animationMap
}


func LoadTileMap() {
	
}


func LoadWaterAnimationSets()  *ebiten_extended.AnimationSet {
	spritesheet := make([]*ebiten.Image, 0, 4) 
	spriteName := []string{
		Tilesets_water0,
		Tilesets_water1,
		Tilesets_water2,
		Tilesets_water3,
	}
	// Load water animations from the sprite sheet
	for i := 0; i < 4; i++ {
		sprite := ebiten_extended.ResourceManager().GetImage(spriteName[i])
		if sprite != nil {
			spritesheet = append(spritesheet, sprite)
		}
	}

	return ebiten_extended.NewAnimationSet(spritesheet, getCenterImage(spritesheet[0]), uint(len(spritesheet)), 6, true)
}

// def import_tilemap(cols, rows, *path):
// 	frames = {}
// 	surf = import_image(*path)
// 	cell_width, cell_height = surf.get_width() / cols, surf.get_height() / rows
// 	for col in range(cols):
// 		for row in range(rows):
// 			cutout_rect = pygame.Rect(col * cell_width, row * cell_height,cell_width,cell_height)
// 			cutout_surf = pygame.Surface((cell_width, cell_height))
// 			cutout_surf.fill('green')
// 			cutout_surf.set_colorkey('green')
// 			cutout_surf.blit(surf, (0,0), cutout_rect)
// 			frames[(col, row)] = cutout_surf
// 	return frames



// def character_importer(cols, rows, *path):
// 	frame_dict = import_tilemap(cols, rows, *path)
// 	new_dict = {}
// 	for row, direction in enumerate(('down', 'left', 'right', 'up')):
// 		new_dict[direction] = [frame_dict[(col, row)] for col in range(cols)]
// 		new_dict[f'{direction}_idle'] = [frame_dict[(0, row)]]
// 	return new_dict

// def all_character_import(*path):
// 	new_dict = {}
// 	for _, __, image_names in walk(join(*path)):
// 		for image in image_names:
// 			image_name = image.split('.')[0]
// 			new_dict[image_name] = character_importer(4,4,*path, image_name)
// 	return new_dict


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




// 	def import_assets(self):
// 		self.tmx_maps = tmx_importer('..', 'data', 'maps')

// 		self.overworld_frames = {
// 			'water': import_folder('..', 'graphics', 'tilesets', 'water'),
// 			'coast': coast_importer(24, 12, '..', 'graphics', 'tilesets', 'coast'),
// 			'characters': all_character_import('..', 'graphics', 'characters')
// 		}

// 		self.monster_frames = {
// 			'icons': import_folder_dict('..', 'graphics', 'icons'),
// 			'monsters': monster_importer(4,2,'..', 'graphics', 'monsters'),
// 			'ui': import_folder_dict('..', 'graphics', 'ui'),
// 			'attacks': attack_importer('..', 'graphics', 'attacks')
// 		}
// 		self.monster_frames['outlines'] = outline_creator(self.monster_frames['monsters'], 4)

// 		self.fonts = {
// 			'dialog': pygame.font.Font(join('..', 'graphics', 'fonts', 'PixeloidSans.ttf'), 30),
// 			'regular': pygame.font.Font(join('..', 'graphics', 'fonts', 'PixeloidSans.ttf'), 18),
// 			'small': pygame.font.Font(join('..', 'graphics', 'fonts', 'PixeloidSans.ttf'), 14),
// 			'bold': pygame.font.Font(join('..', 'graphics', 'fonts', 'dogicapixelbold.otf'), 20),
// 		}
// 		self.bg_frames = import_folder_dict('..', 'graphics', 'backgrounds')
// 		self.start_animation_frames = import_folder('..', 'graphics', 'other', 'star animation')

// 		self.audio = audio_importer('..', 'audio')