package ebiten_pokemon

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

const (
	SCREEN_WIDTH         = 1280
	SCREEN_HEIGHT        = 720
	TILE_SIZE            = 64
	ANIMATION_SPEED      = 6
	BATTLE_OUTLINE_WIDTH = 4
)

// COLORS
var (
	WHITE_COLOR      = color.RGBA{0xf4, 0xfe, 0xfa, 0xff}
	PURE_WHITE_COLOR = color.RGBA{0xff, 0xff, 0xff, 0xff}
	DARK_COLOR       = color.RGBA{0x2b, 0x29, 0x2c, 0xff}
	LIGHT_COLOR      = color.RGBA{0xc8, 0xc8, 0xc8, 0xff}
	GRAY_COLOR       = color.RGBA{0x3a, 0x37, 0x3b, 0xff}
	GOLD_COLOR       = color.RGBA{0xff, 0xd7, 0x00, 0xff}
	LIGHT_GRAY_COLOR = color.RGBA{0x4b, 0x48, 0x4d, 0xff}
	FIRE_COLOR       = color.RGBA{0xf8, 0xa0, 0x60, 0xff}
	WATER_COLOR      = color.RGBA{0x50, 0xb0, 0xd8, 0xff}
	PLANT_COLOR      = color.RGBA{0x64, 0xa9, 0x90, 0xff}
	BLACK_COLOR      = color.RGBA{0x00, 0x00, 0x00, 0xff}
	RED_COLOR        = color.RGBA{0xf0, 0x31, 0x31, 0xff}
	BLUE_COLOR       = color.RGBA{0x66, 0xd7, 0xee, 0xff}
	NORMAL_COLOR     = color.RGBA{0xff, 0xff, 0xff, 0xff}
	DARK_WHITE_COLOR = color.RGBA{0xf0, 0xf0, 0xf0, 0xff}
)


// LAYERS

const (
	WATER_LAYER int = iota
	BG_LAYER
	SHADOW_LAYER
	MAIN_LAYER
	TOP_LAYER
	BATTLE_OUTLINE_LAYER
	BATTLE_NAME_LAYER
	BATTLE_MONSTER_LAYER
	BATTLE_EFFECTS_LAYER
	BATTLE_OVERLAY_LAYER
	BATTLE_CHOICES_LAYER
	BATTLE_CHOICES_ICON_LAYER
	BATTLE_CHOICES_TEXT_LAYER
	BATTLE_CHOICES_CURSOR_LAYER
	BATTLE_CHOICES_CURSOR_ICON_LAYER
	BATTLE_CHOICES_CURSOR_TEXT_LAYER
)


// BATTLE_POSITIONS 
var (
	BATTLE_LEFT_POSITION map[string]math2D.Vector2D = map[string]math2D.Vector2D{
		"top":    math2D.NewVector2D(360, 260),
		"center": math2D.NewVector2D(190, 400),
		"bottom": math2D.NewVector2D(410, 520),
	}

	BATTLE_RIGHT_POSITION map[string]math2D.Vector2D = map[string]math2D.Vector2D{
		"top":    math2D.NewVector2D(900, 260),
		"center": math2D.NewVector2D(1110, 390),
		"bottom": math2D.NewVector2D(900, 550),
	}
)
 

	

// BATTLE_CHOICES = {
// 	'full': {
// 		'fight':  {'pos' : vector(30, -60), 'icon': 'sword'},
// 		'defend': {'pos' : vector(40, -20), 'icon': 'shield'},
// 		'switch': {'pos' : vector(40, 20), 'icon': 'arrows'},
// 		'catch':  {'pos' : vector(30, 60), 'icon': 'hand'}},

// 	'limited': {
// 		'fight':  {'pos' : vector(30, -40), 'icon': 'sword'},
// 		'defend': {'pos' : vector(40, 0), 'icon': 'shield'},
// 		'switch': {'pos' : vector(30, 40), 'icon': 'arrows'}}
// }