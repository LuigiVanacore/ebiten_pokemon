package ebiten_pokemon





type Character struct {
  Entity
}

// func NewCharacter(pos math2D.Vector2D, frames []ebiten_extended.Sprite, groups []ebiten_extended.Node2D, facing_direction EntityState, character_data map[string]interface{}, player *Player, create_dialog func(*Character), collision_sprites []ebiten_extended.Node2D, radius float64, nurse bool, notice_sound ebiten.Sound) *Character {
// 	  character := &Character{
// 	Entity: *NewEntity(pos, frames, groups, facing_direction),
// 	character_data: character_data,
// 	player: player,
// 	create_dialog: create_dialog,
// 	collision_rects: []ebiten_extended.Node2D{},
// 	nurse: nurse,
// 	monsters: nil,
//   }

//   for _, sprite := range collision_sprites {
// 	if sprite != character {
// 	  character.collision_rects = append(character.collision_rects, sprite)
// 	}
//   }

//   if _, ok := character.character_data["monsters"]; ok {
// 	character.monsters = make(map[int]Monster)
// 	for i, (name, lvl) := range character.character_data["monsters"].(map[string]interface{}) {
// 	  character.monsters[i] = NewMonster(name.(string), lvl.(int))
// 	}
//   }

//   // movement 
//   character.has_moved = false
//   character.can_rotate = true
//   character.has_noticed = false
//   character.radius = int(radius)
//   character.view_directions = character.character_data["directions"].([]string)

//   character.timers = map[string]*Timer{
// 	"look around": NewTimer(1500, true, true, func() {character.random_view_direction()}),
// 	"notice": NewTimer(500, func() {character.start_move()}),
//   }
//   character.notice_sound = notice_sound

//   return character
// }

// func (c *Character) random_view_direction() {
//   if c.can_rotate {
// 	c.facing_direction = choice(c.view_directions)
//   }
// }

// func (c *Character) get_dialog() string {
// 	  return c.character_data["dialog"][fmt.Sprintf("%s", "defeated" if c.character_data["defeated"] else "default")]
// }

// func (c *Character) raycast() {
// 	  if check_connections(c.radius, c, c.player) && c.has_los() && !c.has_moved && !c.has_noticed {
// 	c.player.block()
// 	c.player.change_facing_direction(c.rect.center)
// 	c.timers["notice"].activate()
// 	c.can_rotate = false
// 	c.has_noticed = true
// 	c.player.noticed = true
// 	c.notice_sound.play()
//   }
// }

// func (c *Character) has_los() bool {
// 	  if vector(c.rect.center).distance_to(c.player.rect.center) < c.radius {
// 	collisions := make([]bool, len(c.collision_rects))
// 	for i, rect := range c.collision_rects {
// 	  collisions[i] = bool(rect.clipline(c.rect.center, c.player.rect.center))
// 	}
// 	return !any(collisions)
//   }
//   return false
// }

// func (c *Character) start_move() {
// 	  relation := (vector(c.player.rect.center) - vector(c.rect.center)).normalize()
// 	c.direction = vector(round(relation.x), round(relation.y))
// }

// func (c *Character) move(dt float64) {
// 	  if !c.has_moved && c.direction != nil {
// 	if !c.hitbox.inflate(10, 10).colliderect(c.player.hitbox) {
// 	  c.rect.center += c.direction * c.speed * dt
// 	  c.hitbox.center = c.rect.center
// 	} else {
// 	  c.direction = vector()
// 	  c.has_moved = true
// 	  c.create_dialog(c)
// 	  c.player.noticed = false
// 	}
//   }
// }


// func (c *Character) update(dt float64) {
// 	  for _, timer := range c.timers {
// 	timer.update()
//   }

//   c.animate(dt)
//   if c.character_data["look_around"].(bool) {
// 	c.raycast()
// 	c.move(dt)
//   }
// }





// class Character(Entity):
// 	def __init__(self, pos, frames, groups, facing_direction, character_data, player, create_dialog, collision_sprites, radius, nurse, notice_sound):
// 		super().__init__(pos, frames, groups, facing_direction)
// 		self.character_data = character_data
// 		self.player = player
// 		self.create_dialog = create_dialog
// 		self.collision_rects = [sprite.rect for sprite in collision_sprites if sprite is not self]
// 		self.nurse = nurse
// 		self.monsters = {i: Monster(name, lvl) for i, (name, lvl) in character_data['monsters'].items()} if 'monsters' in character_data else None

// 		# movement 
// 		self.has_moved = False
// 		self.can_rotate = True
// 		self.has_noticed = False
// 		self.radius = int(radius)
// 		self.view_directions = character_data['directions']

// 		self.timers = {
// 			'look around': Timer(1500, autostart = True, repeat = True, func = self.random_view_direction),
// 			'notice': Timer(500, func = self.start_move)
// 		}
// 		self.notice_sound = notice_sound

// 	def random_view_direction(self):
// 		if self.can_rotate:
// 			self.facing_direction = choice(self.view_directions)

// 	def get_dialog(self):
// 		return self.character_data['dialog'][f"{'defeated' if self.character_data['defeated'] else 'default'}"]

// 	def raycast(self):
// 		if check_connections(self.radius, self, self.player) and self.has_los() and not self.has_moved and not self.has_noticed:
// 			self.player.block()
// 			self.player.change_facing_direction(self.rect.center)
// 			self.timers['notice'].activate()
// 			self.can_rotate = False
// 			self.has_noticed = True
// 			self.player.noticed = True
// 			self.notice_sound.play()

// 	def has_los(self):
// 		if vector(self.rect.center).distance_to(self.player.rect.center) < self.radius:
// 			collisions = [bool(rect.clipline(self.rect.center, self.player.rect.center)) for rect in self.collision_rects]
// 			return not any(collisions)

// 	def start_move(self):
// 		relation = (vector(self.player.rect.center) - vector(self.rect.center)).normalize()
// 		self.direction = vector(round(relation.x), round(relation.y))

// 	def move(self, dt):
// 		if not self.has_moved and self.direction:
// 			if not self.hitbox.inflate(10,10).colliderect(self.player.hitbox):
// 				self.rect.center += self.direction * self.speed * dt
// 				self.hitbox.center = self.rect.center
// 			else:
// 				self.direction = vector()
// 				self.has_moved = True
// 				self.create_dialog(self)
// 				self.player.noticed = False

// 	def update(self, dt):
// 		for timer in self.timers.values():
// 			timer.update()

// 		self.animate(dt)
// 		if self.character_data['look_around']:
// 			self.raycast()
// 			self.move(dt)
