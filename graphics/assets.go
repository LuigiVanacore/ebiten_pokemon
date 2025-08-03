package resources

import (
    _ "embed"
)

var (
 

	//go:embed  attacks\explosion.png
	Attacks_explosion []byte

	//go:embed  attacks\fire.png
	Attacks_fire []byte

	//go:embed  attacks\green.png
	Attacks_green []byte

	//go:embed  attacks\ice.png
	Attacks_ice []byte

	//go:embed  attacks\scratch.png
	Attacks_scratch []byte

	//go:embed  attacks\splash.png
	Attacks_splash []byte

	//go:embed  backgrounds\forest.png
	Backgrounds_forest []byte

	//go:embed  backgrounds\ice.png
	Backgrounds_ice []byte

	//go:embed  backgrounds\sand.png
	Backgrounds_sand []byte

	//go:embed  characters\blond.png
	Characters_blond []byte

	//go:embed  characters\fire_boss.png
	Characters_fire_boss []byte

	//go:embed  characters\grass_boss.png
	Characters_grass_boss []byte

	//go:embed  characters\hat_girl.png
	Characters_hat_girl []byte

	//go:embed  characters\player.png
	Characters_player []byte

	//go:embed  characters\purple_girl.png
	Characters_purple_girl []byte

	//go:embed  characters\straw.png
	Characters_straw []byte

	//go:embed  characters\water_boss.png
	Characters_water_boss []byte

	//go:embed  characters\young_girl.png
	Characters_young_girl []byte

	//go:embed  characters\young_guy.png
	Characters_young_guy []byte
  
 

	//go:embed  fonts\PixeloidSans.ttf
	Fonts_PixeloidSans []byte

	//go:embed  fonts\dogicapixel.otf
	Fonts_dogicapixel []byte

	//go:embed  fonts\dogicapixelbold.otf
	Fonts_dogicapixelbold []byte

	//go:embed  icons\Atrox.png
	Icons_Atrox []byte

	//go:embed  icons\Charmadillo.png
	Icons_Charmadillo []byte

	//go:embed  icons\Cindrill.png
	Icons_Cindrill []byte

	//go:embed  icons\Cleaf.png
	Icons_Cleaf []byte

	//go:embed  icons\Draem.png
	Icons_Draem []byte

	//go:embed  icons\Finiette.png
	Icons_Finiette []byte

	//go:embed  icons\Finsta.png
	Icons_Finsta []byte

	//go:embed  icons\Friolera.png
	Icons_Friolera []byte

	//go:embed  icons\Gulfin.png
	Icons_Gulfin []byte

	//go:embed  icons\Ivieron.png
	Icons_Ivieron []byte

	//go:embed  icons\Jacana.png
	Icons_Jacana []byte

	//go:embed  icons\Larvea.png
	Icons_Larvea []byte

	//go:embed  icons\Pluma.png
	Icons_Pluma []byte

	//go:embed  icons\Plumette.png
	Icons_Plumette []byte

	//go:embed  icons\Pouch.png
	Icons_Pouch []byte

	//go:embed  icons\Sparchu.png
	Icons_Sparchu []byte

	//go:embed  monsters\Atrox.png
	Monsters_Atrox []byte

	//go:embed  monsters\Charmadillo.png
	Monsters_Charmadillo []byte

	//go:embed  monsters\Cindrill.png
	Monsters_Cindrill []byte

	//go:embed  monsters\Cleaf.png
	Monsters_Cleaf []byte

	//go:embed  monsters\Draem.png
	Monsters_Draem []byte

	//go:embed  monsters\Finiette.png
	Monsters_Finiette []byte

	//go:embed  monsters\Finsta.png
	Monsters_Finsta []byte

	//go:embed  monsters\Friolera.png
	Monsters_Friolera []byte

	//go:embed  monsters\Gulfin.png
	Monsters_Gulfin []byte

	//go:embed  monsters\Ivieron.png
	Monsters_Ivieron []byte

	//go:embed  monsters\Jacana.png
	Monsters_Jacana []byte

	//go:embed  monsters\Larvea.png
	Monsters_Larvea []byte

	//go:embed  monsters\Pluma.png
	Monsters_Pluma []byte

	//go:embed  monsters\Plumette.png
	Monsters_Plumette []byte

	//go:embed  monsters\Pouch.png
	Monsters_Pouch []byte

	//go:embed  monsters\Sparchu.png
	Monsters_Sparchu []byte

	//go:embed  objects\arean_fire.png
	Objects_arean_fire []byte

	//go:embed  objects\arena_plant.png
	Objects_arena_plant []byte

	//go:embed  objects\arena_water.png
	Objects_arena_water []byte

	//go:embed  objects\gate_pillar.png
	Objects_gate_pillar []byte

	//go:embed  objects\gate_top.png
	Objects_gate_top []byte

	//go:embed  objects\grass.png
	Objects_grass []byte

	//go:embed  objects\grass_ice.png
	Objects_grass_ice []byte

	//go:embed  objects\grassrock1.png
	Objects_grassrock1 []byte

	//go:embed  objects\grassrock2.png
	Objects_grassrock2 []byte

	//go:embed  objects\green_tree.png
	Objects_green_tree []byte

	//go:embed  objects\green_tree_bushy.png
	Objects_green_tree_bushy []byte

	//go:embed  objects\green_tree_small.png
	Objects_green_tree_small []byte

	//go:embed  objects\hospital.png
	Objects_hospital []byte

	//go:embed  objects\house_large.png
	Objects_house_large []byte

	//go:embed  objects\house_large_alt.png
	Objects_house_large_alt []byte

	//go:embed  objects\house_small.png
	Objects_house_small []byte

	//go:embed  objects\house_small_alt.png
	Objects_house_small_alt []byte

	//go:embed  objects\ice_tree.png
	Objects_ice_tree []byte

	//go:embed  objects\icerock1.png
	Objects_icerock1 []byte

	//go:embed  objects\icerock2.png
	Objects_icerock2 []byte

	//go:embed  objects\palm.png
	Objects_palm []byte

	//go:embed  objects\palm_alt.png
	Objects_palm_alt []byte

	//go:embed  objects\palm_small.png
	Objects_palm_small []byte

	//go:embed  objects\ruin_gate.png
	Objects_ruin_gate []byte

	//go:embed  objects\ruin_pillar.png
	Objects_ruin_pillar []byte

	//go:embed  objects\ruin_pillar_broke.png
	Objects_ruin_pillar_broke []byte

	//go:embed  objects\ruin_pillar_broke_alt.png
	Objects_ruin_pillar_broke_alt []byte

	//go:embed  objects\sand.png
	Objects_sand []byte

	//go:embed  objects\sandrock1.png
	Objects_sandrock1 []byte

	//go:embed  objects\sandrock2.png
	Objects_sandrock2 []byte

	//go:embed  objects\teal_tree.png
	Objects_teal_tree []byte

	//go:embed  objects\teal_tree_bushy.png
	Objects_teal_tree_bushy []byte

	//go:embed  objects\teal_tree_small.png
	Objects_teal_tree_small []byte

	//go:embed  other\shadow.png
	Other_shadow []byte

	//go:embed  other\star_animation\00002.png
	Other_star_animation_00002 []byte

	//go:embed  other\star_animation\00003.png
	Other_star_animation_00003 []byte

	//go:embed  other\star_animation\00004.png
	Other_star_animation_00004 []byte

	//go:embed  other\star_animation\00005.png
	Other_star_animation_00005 []byte

	//go:embed  other\star_animation\00006.png
	Other_star_animation_00006 []byte

	//go:embed  other\star_animation\00007.png
	Other_star_animation_00007 []byte

	//go:embed  other\star_animation\00008.png
	Other_star_animation_00008 []byte

	//go:embed  other\star_animation\00009.png
	Other_star_animation_00009 []byte

	//go:embed  other\star_animation\00010.png
	Other_star_animation_00010 []byte

	//go:embed  other\star_animation\00011.png
	Other_star_animation_00011 []byte

	//go:embed  other\star_animation\00012.png
	Other_star_animation_00012 []byte

	//go:embed  other\star_animation\00013.png
	Other_star_animation_00013 []byte

	//go:embed  other\star_animation\00014.png
	Other_star_animation_00014 []byte

	//go:embed  other\star_animation\00015.png
	Other_star_animation_00015 []byte

	//go:embed  other\star_animation\00016.png
	Other_star_animation_00016 []byte

	//go:embed  other\star_animation\00017.png
	Other_star_animation_00017 []byte

	//go:embed  other\star_animation\00018.png
	Other_star_animation_00018 []byte

	//go:embed  other\star_animation\00019.png
	Other_star_animation_00019 []byte

	//go:embed  other\star_animation\00020.png
	Other_star_animation_00020 []byte

	//go:embed  other\star_animation\00021.png
	Other_star_animation_00021 []byte

	//go:embed  other\star_animation\00022.png
	Other_star_animation_00022 []byte

	//go:embed  other\star_animation\00023.png
	Other_star_animation_00023 []byte

	//go:embed  other\star_animation\00024.png
	Other_star_animation_00024 []byte

	//go:embed  other\star_animation\00025.png
	Other_star_animation_00025 []byte

	//go:embed  other\star_animation\00026.png
	Other_star_animation_00026 []byte

	//go:embed  other\star_animation\00027.png
	Other_star_animation_00027 []byte

	//go:embed  other\star_animation\00028.png
	Other_star_animation_00028 []byte

	//go:embed  other\star_animation\00029.png
	Other_star_animation_00029 []byte

	//go:embed  tilesets\coast.png
	Tilesets_coast []byte

	//go:embed  tilesets\indoor.png
	Tilesets_indoor []byte

	//go:embed  tilesets\water\water0.png
	Tilesets_water0 []byte

	//go:embed  tilesets\water\water1.png
	Tilesets_water1 []byte

	//go:embed  tilesets\water\water2.png
	Tilesets_water2 []byte

	//go:embed  tilesets\water\water3.png
	Tilesets_water3 []byte

	//go:embed  tilesets\world.png
	Tilesets_world []byte

	//go:embed  ui\arrows.png
	Ui_arrows []byte

	//go:embed  ui\arrows_highlight.png
	Ui_arrows_highlight []byte

	//go:embed  ui\attack.png
	Ui_attack []byte

	//go:embed  ui\cross.png
	Ui_cross []byte

	//go:embed  ui\defense.png
	Ui_defense []byte

	//go:embed  ui\energy.png
	Ui_energy []byte

	//go:embed  ui\hand.png
	Ui_hand []byte

	//go:embed  ui\hand_highlight.png
	Ui_hand_highlight []byte

	//go:embed  ui\health.png
	Ui_health []byte

	//go:embed  ui\notice.png
	Ui_notice []byte

	//go:embed  ui\recovery.png
	Ui_recovery []byte

	//go:embed  ui\shield.png
	Ui_shield []byte

	//go:embed  ui\shield_highlight.png
	Ui_shield_highlight []byte

	//go:embed  ui\speed.png
	Ui_speed []byte

	//go:embed  ui\star.png
	Ui_star []byte

	//go:embed  ui\sword.png
	Ui_sword []byte

	//go:embed  ui\sword_highlight.png
	Ui_sword_highlight []byte

)
