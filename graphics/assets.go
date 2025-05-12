package resources

import (
    _ "embed"
)

var (
 
	//go:embed C:\Go_projects\ebiten_pokemon\graphics\attacks\explosion.png
	attacks_explosion []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\attacks\fire.png
	attacks_fire []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\attacks\green.png
	attacks_green []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\attacks\ice.png
	attacks_ice []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\attacks\scratch.png
	attacks_scratch []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\attacks\splash.png
	attacks_splash []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\backgrounds\forest.png
	backgrounds_forest []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\backgrounds\ice.png
	backgrounds_ice []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\backgrounds\sand.png
	backgrounds_sand []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\characters\blond.png
	characters_blond []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\characters\fire_boss.png
	characters_fire_boss []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\characters\grass_boss.png
	characters_grass_boss []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\characters\hat_girl.png
	characters_hat_girl []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\characters\player.png
	characters_player []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\characters\purple_girl.png
	characters_purple_girl []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\characters\straw.png
	characters_straw []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\characters\water_boss.png
	characters_water_boss []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\characters\young_girl.png
	characters_young_girl []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\characters\young_guy.png
	characters_young_guy []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\cmd\script.go
	cmd_script []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\fonts\PixeloidSans.ttf
	fonts_PixeloidSans []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\fonts\dogicapixel.otf
	fonts_dogicapixel []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\fonts\dogicapixelbold.otf
	fonts_dogicapixelbold []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Atrox.png
	icons_Atrox []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Charmadillo.png
	icons_Charmadillo []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Cindrill.png
	icons_Cindrill []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Cleaf.png
	icons_Cleaf []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Draem.png
	icons_Draem []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Finiette.png
	icons_Finiette []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Finsta.png
	icons_Finsta []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Friolera.png
	icons_Friolera []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Gulfin.png
	icons_Gulfin []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Ivieron.png
	icons_Ivieron []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Jacana.png
	icons_Jacana []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Larvea.png
	icons_Larvea []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Pluma.png
	icons_Pluma []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Plumette.png
	icons_Plumette []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Pouch.png
	icons_Pouch []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\icons\Sparchu.png
	icons_Sparchu []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Atrox.png
	monsters_Atrox []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Charmadillo.png
	monsters_Charmadillo []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Cindrill.png
	monsters_Cindrill []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Cleaf.png
	monsters_Cleaf []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Draem.png
	monsters_Draem []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Finiette.png
	monsters_Finiette []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Finsta.png
	monsters_Finsta []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Friolera.png
	monsters_Friolera []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Gulfin.png
	monsters_Gulfin []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Ivieron.png
	monsters_Ivieron []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Jacana.png
	monsters_Jacana []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Larvea.png
	monsters_Larvea []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Pluma.png
	monsters_Pluma []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Plumette.png
	monsters_Plumette []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Pouch.png
	monsters_Pouch []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\monsters\Sparchu.png
	monsters_Sparchu []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\arean_fire.png
	objects_arean_fire []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\arena_plant.png
	objects_arena_plant []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\arena_water.png
	objects_arena_water []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\gate_pillar.png
	objects_gate_pillar []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\gate_top.png
	objects_gate_top []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\grass.png
	objects_grass []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\grass_ice.png
	objects_grass_ice []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\grassrock1.png
	objects_grassrock1 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\grassrock2.png
	objects_grassrock2 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\green_tree.png
	objects_green_tree []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\green_tree_bushy.png
	objects_green_tree_bushy []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\green_tree_small.png
	objects_green_tree_small []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\hospital.png
	objects_hospital []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\house_large.png
	objects_house_large []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\house_large_alt.png
	objects_house_large_alt []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\house_small.png
	objects_house_small []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\house_small_alt.png
	objects_house_small_alt []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\ice_tree.png
	objects_ice_tree []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\icerock1.png
	objects_icerock1 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\icerock2.png
	objects_icerock2 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\palm.png
	objects_palm []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\palm_alt.png
	objects_palm_alt []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\palm_small.png
	objects_palm_small []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\ruin_gate.png
	objects_ruin_gate []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\ruin_pillar.png
	objects_ruin_pillar []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\ruin_pillar_broke.png
	objects_ruin_pillar_broke []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\ruin_pillar_broke_alt.png
	objects_ruin_pillar_broke_alt []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\sand.png
	objects_sand []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\sandrock1.png
	objects_sandrock1 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\sandrock2.png
	objects_sandrock2 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\teal_tree.png
	objects_teal_tree []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\teal_tree_bushy.png
	objects_teal_tree_bushy []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\objects\teal_tree_small.png
	objects_teal_tree_small []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\shadow.png
	other_shadow []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00002.png
	other_star_animation_00002 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00003.png
	other_star_animation_00003 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00004.png
	other_star_animation_00004 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00005.png
	other_star_animation_00005 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00006.png
	other_star_animation_00006 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00007.png
	other_star_animation_00007 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00008.png
	other_star_animation_00008 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00009.png
	other_star_animation_00009 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00010.png
	other_star_animation_00010 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00011.png
	other_star_animation_00011 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00012.png
	other_star_animation_00012 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00013.png
	other_star_animation_00013 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00014.png
	other_star_animation_00014 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00015.png
	other_star_animation_00015 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00016.png
	other_star_animation_00016 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00017.png
	other_star_animation_00017 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00018.png
	other_star_animation_00018 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00019.png
	other_star_animation_00019 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00020.png
	other_star_animation_00020 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00021.png
	other_star_animation_00021 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00022.png
	other_star_animation_00022 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00023.png
	other_star_animation_00023 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00024.png
	other_star_animation_00024 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00025.png
	other_star_animation_00025 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00026.png
	other_star_animation_00026 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00027.png
	other_star_animation_00027 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00028.png
	other_star_animation_00028 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\other\star_animation\00029.png
	other_star_animation_00029 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\tilesets\coast.png
	tilesets_coast []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\tilesets\indoor.png
	tilesets_indoor []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\tilesets\water\water0.png
	tilesets_water_water0 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\tilesets\water\water1.png
	tilesets_water_water1 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\tilesets\water\water2.png
	tilesets_water_water2 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\tilesets\water\water3.png
	tilesets_water_water3 []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\tilesets\world.png
	tilesets_world []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\arrows.png
	ui_arrows []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\arrows_highlight.png
	ui_arrows_highlight []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\attack.png
	ui_attack []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\cross.png
	ui_cross []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\defense.png
	ui_defense []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\energy.png
	ui_energy []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\hand.png
	ui_hand []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\hand_highlight.png
	ui_hand_highlight []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\health.png
	ui_health []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\notice.png
	ui_notice []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\recovery.png
	ui_recovery []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\shield.png
	ui_shield []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\shield_highlight.png
	ui_shield_highlight []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\speed.png
	ui_speed []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\star.png
	ui_star []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\sword.png
	ui_sword []byte

	//go:embed C:\Go_projects\ebiten_pokemon\graphics\ui\sword_highlight.png
	ui_sword_highlight []byte

)
