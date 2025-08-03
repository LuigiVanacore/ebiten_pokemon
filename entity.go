package ebiten_pokemon

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

// Entity represents a game entity with position, speed, and direction.
// It extends the Node2D class from the ebiten_extended package.
// The Entity class is used to represent any object in the game world that has a position and can move.
// It can be a player, an NPC, or any other object that needs to be represented in the game world.
// The Entity class provides methods to set and get the position, speed, and direction of the entity.
// It also provides methods to move the entity in the game world and to check if the entity is blocked or not.

type EntityState int

const (
	Idle EntityState = iota
	MovingUp
	MovingDown
	MovingLeft
	MovingRight
)

type Entity struct {
	ebiten_extended.Node2D
	speed     int
	blocked   bool
	direction math2D.Vector2D
	facing_direction EntityState
}

// func NewEntity() *Entity {
// 	return &Entity{}
// }


// func (e *Entity) SetFacingDirection(direction EntityState) {
// 	e.facing_direction = direction
// }

// func (e *Entity) GetFacingDirection() EntityState {	
// 	return e.facing_direction
// }



// func (e *Entity) get_state() EntityState {
// 	moving := e.direction.Magnitude() > 0
// 	if moving {
// 		if e.direction.X() != 0 {
// 			if e.direction.X() > 0 {
// 				return MovingRight
// 			} else {
// 				return MovingLeft
// 			}
// 		}
// 		if e.direction.Y() != 0 {
// 			if e.direction.Y() > 0 {
// 				return MovingDown
// 			} else {
// 				return MovingUp
// 			}
// 		}
// 	}
// 	return Idle
// }

// func (e *Entity) change_facing_direction(target_pos math2D.Vector2D) {
// 	relation := math2D.SubtractVectors(target_pos, e.GetPosition())
// 	if relation.Y() < 30 {
// 		if relation.X() > 0 {
// 			e.SetFacingDirection(MovingRight)
// 		} else {
// 			e.SetFacingDirection(MovingLeft)
// 		}
// 	} else {
// 		if relation.Y() > 0 {
// 			e.SetFacingDirection(MovingDown)
// 		} else {
// 			e.SetFacingDirection(MovingUp)
// 		}
// 	}
// }

// func (e *Entity) block() {
// 	e.blocked = true
// 	e.direction = math2D.Vector2D{}
// }

// func (e *Entity) unblock() {
// 	e.blocked = false
// }

// func (e *Entity) is_blocked() bool {	
// 	return e.blocked
// }

// func (e *Entity) set_speed(speed int) {
// 	e.speed = speed
// }

// func (e *Entity) get_speed() int {
// 	return e.speed
// }

// func (e *Entity) Move() {
// 	if e.direction.Magnitude() > 0 {
// 		position := math2D.AddVectors(e.direction.MultiplyScalar(float64(e.speed)), e.GetPosition())
// 		e.SetPosition(position.X(), position.Y())
// 	}
// }


// from settings import * 
// from support import check_connections
// from timer import Timer
// from random import choice
// from monster import Monster

// class Entity(pygame.sprite.Sprite):
// 	def __init__(self, pos, frames, groups, facing_direction):
// 		super().__init__(groups)
// 		self.z = WORLD_LAYERS['main']

// 		# graphics 
// 		self.frame_index, self.frames = 0, frames
// 		self.facing_direction = facing_direction

// 		# movement 
// 		self.direction = vector()
// 		self.speed = 250
// 		self.blocked = False

// 		# sprite setup
// 		self.image = self.frames[self.get_state()][self.frame_index]
// 		self.rect = self.image.get_frect(center = pos)
// 		self.hitbox = self.rect.inflate(-self.rect.width / 2, -60)

// 		self.y_sort = self.rect.centery

// 	def animate(self, dt):
// 		self.frame_index += ANIMATION_SPEED * dt
// 		self.image = self.frames[self.get_state()][int(self.frame_index % len(self.frames[self.get_state()]))]

// 	def get_state(self):
// 		moving = bool(self.direction)
// 		if moving:
// 			if self.direction.x != 0:
// 				self.facing_direction = 'right' if self.direction.x > 0 else 'left'
// 			if self.direction.y != 0:
// 				self.facing_direction = 'down' if self.direction.y > 0 else 'up'
// 		return f"{self.facing_direction}{'' if moving else '_idle'}"

// 	def change_facing_direction(self, target_pos):
// 		relation = vector(target_pos) - vector(self.rect.center)
// 		if abs(relation.y) < 30:
// 			self.facing_direction = 'right' if relation.x > 0 else 'left'
// 		else:
// 			self.facing_direction = 'down' if relation.y > 0 else 'up'

// 	def block(self):
// 		self.blocked = True
// 		self.direction = vector(0,0)

// 	def unblock(self):
// 		self.blocked = False

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
