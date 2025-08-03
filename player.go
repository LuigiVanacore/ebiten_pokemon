package ebiten_pokemon

import (
	"github.com/LuigiVanacore/ebiten_extended"
	inputv3 "github.com/LuigiVanacore/ebiten_extended/input_v3"
	fsm "github.com/LuigiVanacore/ebiten_extended/fsm"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PLAYER_SPEED = 250
)


const (
    MOVE_UP inputv3.ActionID = iota
    MOVE_DOWN
    MOVE_LEFT
    MOVE_RIGHT
)


const (
	MOVE_STATE fsm.StateID = iota
	IDLE_STATE
)

type Player struct {
	ebiten_extended.Node2D
	movement  *MovementComponent
	animation *ebiten_extended.AnimationPlayer
}

func NewPlayer( ) *Player {

	player := &Player{
		Node2D:     *ebiten_extended.NewNode2D("player"),
	} 

	animationPlayer := ebiten_extended.NewAnimationPlayer("player_animation_player", MAIN_LAYER)

	movementComponent := NewMovementComponent(&player.Node2D, PLAYER_SPEED)

	player.movement = movementComponent
	player.animation = animationPlayer

	return player
}


func (p *Player) loadActions() {
	move_up := inputv3.NewKeyAction(ebiten.KeyW, inputv3.Hold | inputv3.PressOnce)
	move_down := inputv3.NewKeyAction(ebiten.KeyS, inputv3.Hold | inputv3.PressOnce)
	move_left := inputv3.NewKeyAction(ebiten.KeyA, inputv3.Hold | inputv3.PressOnce)
	move_right := inputv3.NewKeyAction(ebiten.KeyD, inputv3.Hold | inputv3.PressOnce)

	ebiten_extended.InputManager().RegisterAction(MOVE_UP, move_up)
	ebiten_extended.InputManager().RegisterAction(MOVE_DOWN, move_down)
	ebiten_extended.InputManager().RegisterAction(MOVE_LEFT, move_left)
	ebiten_extended.InputManager().RegisterAction(MOVE_RIGHT, move_right)
}




type PlayerMoveState struct {
	player *Player
}

func NewPlayerMoveState(player *Player) *PlayerMoveState {
	return &PlayerMoveState{
		player: player,
	}
}

func (s *PlayerMoveState) Enter() {
}

func (s *PlayerMoveState) Exit() {
}

func (s *PlayerMoveState) Update() {
 	// // horizontal movement
	// horizontal_position := math2D.AddVectors(s.player.GetPosition(), s.player.direction.MultiplyScalar(float64(s.player.speed)))

	// // vertical movement
	// vertical_position := math2D.AddVectors(s.player.GetPosition(), s.player.direction.MultiplyScalar(float64(s.player.speed)))

	// s.player.SetPosition(math2D.NewVector2D(horizontal_position.X(), vertical_position.Y()))
	// 	if s.player.direction.Y() < 0 {
	// 		s.player.orientation = UP
	// 		s.player.animationPlayer.SetCurrentAnimation(Character_Up)
	// 	} else if s.player.direction.Y() > 0 {
	// 		s.player.orientation = DOWN
	// 		s.player.animationPlayer.SetCurrentAnimation(Character_Down)
	// 	} else if s.player.direction.X() < 0 {
	// 		s.player.orientation = LEFT
	// 		s.player.animationPlayer.SetCurrentAnimation(Character_Left)
	// 	} else if s.player.direction.X() > 0 {
	// 		s.player.orientation = RIGHT
	// 		s.player.animationPlayer.SetCurrentAnimation(Character_Right)
	// 	}
}




type PlayerIdleState struct {
	player *Player
}

func NewPlayerIdleState(player *Player) *PlayerIdleState {
	return &PlayerIdleState{
		player: player,
	}
}	

func (s *PlayerIdleState) Enter() {
// switch s.player.orientation {
// 		case UP:
// 			s.player.animationPlayer.SetCurrentAnimation(Character_Up_Idle)
// 		case DOWN:
// 			s.player.animationPlayer.SetCurrentAnimation(Character_Down_Idle)
// 		case LEFT:
// 			s.player.animationPlayer.SetCurrentAnimation(Character_Left_Idle)
// 		case RIGHT:
// 			s.player.animationPlayer.SetCurrentAnimation(Character_Right_Idle)
// 		}
}

func (s *PlayerIdleState) Exit() {
}

func (s *PlayerIdleState) Update() {

}

func NewPlayerIdleTransitions(p *Player) []fsm.Transition {
	return []fsm.Transition{
		{
			Origin:  IDLE_STATE,
			Target:  MOVE_STATE,
			Trigger: func() bool { return p.movement.GetDirection().Magnitude() > 0  },
		},
	}
}

func NewPlayerMoveTransitions(p *Player) []fsm.Transition {
	return []fsm.Transition{
		{
			Origin:  MOVE_STATE,
			Target:  IDLE_STATE,
			Trigger: func() bool { return p.movement.GetDirection().Magnitude() == 0 },
		},
		 
	}
}
 


func NewPlayerStateMachine(player *Player) *fsm.StateMachine {
	stateMachine := fsm.NewStateMachine()

	moveState := NewPlayerMoveState(player)
	idleState := NewPlayerIdleState(player)

	stateMachine.AddState(MOVE_STATE, moveState) 
	stateMachine.AddState(IDLE_STATE, idleState)

	stateMachine.AddTransitions( NewPlayerMoveTransitions(player)...) 
	stateMachine.AddTransitions( NewPlayerIdleTransitions(player)...)

	stateMachine.SetState(IDLE_STATE)

	return stateMachine
}



func (p *Player) OnCollision(axis string) {
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

// if p.movement.direction.X() > 0 {

// } else if p.movement.direction.X() < 0 {
// }


         
//             if axis == "horizontal" {
//                 if p.Direction.X > 0 {
//                     p.Hitbox.Right = sprite.Hitbox.Left
//                 }
//                 if p.Direction.X < 0 {
//                     p.Hitbox.Left = sprite.Hitbox.Right
//                 }
//                 p.Rect.CenterX = (p.Hitbox.Left + p.Hitbox.Right) / 2
//             } else {
//                 if p.Direction.Y > 0 {
//                     p.Hitbox.Bottom = sprite.Hitbox.Top
//                 }
//                 if p.Direction.Y < 0 {
//                     p.Hitbox.Top = sprite.Hitbox.Bottom
//                 }
//                 p.Rect.CenterY = (p.Hitbox.Top + p.Hitbox.Bottom) / 2
//             }
         
    
}
// func NewPlayer(pos math2D.Vector2D, frames []ebiten_extended.Sprite, groups []ebiten_extended.Node2D, facing_direction EntityState, collision_sprites []ebiten_extended.Node2D) *Player {
// 	player := &Player{
// 		Entity: *NewEntity(pos, frames, groups, facing_direction),
// 		collision_sprites: collision_sprites,
// 		noticed: false,
// 	}
// 	return player
// }
 

 
// }

// func (p *Player) Move(dt float64) {
// 	// normalizing a vector
// 	if p.direction.Magnitude() > 0 {
// 		p.direction = p.direction.Normalize()
// 	}

// 	// horizontal movement
// 	horizontal_position := math2D.AddVectors(p.GetPosition(), p.direction.MultiplyScalar(float64(p.speed) * dt))

// 	// vertical movement
// 	vertical_position := math2D.AddVectors(p.GetPosition(), p.direction.MultiplyScalar(float64(p.speed) * dt))

// 	p.SetPosition(horizontal_position.X(), vertical_position.Y())
// }

// func (p *Player) Collisions(axis string) {
// 	for _, sprite := range p.collision_sprites {
// 		if sprite.GetHitbox().CollidesWith(p.GetHitbox()) {
// 			if axis == "horizontal" {
// 				if p.direction.X() > 0 {
// 					p.GetHitbox().SetRight(sprite.GetHitbox().Left())
// 				} else if p.direction.X() < 0 {
// 					p.GetHitbox().SetLeft(sprite.GetHitbox().Right())
// 				}
// 				p.SetPosition(p.GetHitbox().CenterX(), p.GetPosition().Y())
// 			} else {
// 				if p.direction.Y() > 0 {
// 					p.GetHitbox().SetBottom(sprite.GetHitbox().Top())
// 				} else if p.direction.Y() < 0 {
// 					p.GetHitbox().SetTop(sprite.GetHitbox().Bottom())
// 				}
// 				p.SetPosition(p.GetPosition().X(), p.GetHitbox().CenterY())
// 			}
// 		}
// 	}
// }

// func (p *Player) Update(dt float64) {
// 	p.y_sort = p.GetPosition().Y()
// 	if !p.blocked {
// 		p.Input()
// 		p.Move(dt)
// 	}
// 	p.Collisions("horizontal")
// 	p.Collisions("vertical")
// 	p.Animate(dt)
// }

// class Player(Entity):
// 	def __init__(self, pos, frames, groups, facing_direction, collision_sprites):
// 		super().__init__(pos, frames, groups, facing_direction)
// 		self.collision_sprites = collision_sprites
// 		self.noticed = False

// 	def input(self):
// 		keys = pygame.key.get_pressed()
// 		input_vector = vector()
// 		if keys[pygame.K_UP]:
// 			input_vector.y -= 1
// 		if keys[pygame.K_DOWN]:
// 			input_vector.y += 1
// 		if keys[pygame.K_LEFT]:
// 			input_vector.x -= 1
// 		if keys[pygame.K_RIGHT]:
// 			input_vector.x += 1
// 		self.direction = input_vector.normalize() if input_vector else input_vector

// 	def move(self, dt):
// 		self.rect.centerx += self.direction.x * self.speed * dt
// 		self.hitbox.centerx = self.rect.centerx
// 		self.collisions('horizontal')

// 		self.rect.centery += self.direction.y * self.speed * dt
// 		self.hitbox.centery = self.rect.centery
// 		self.collisions('vertical')

// 	def collisions(self, axis):
// 		for sprite in self.collision_sprites:
// 			if sprite.hitbox.colliderect(self.hitbox):
// 				if axis == 'horizontal':
// 					if self.direction.x > 0:
// 						self.hitbox.right = sprite.hitbox.left
// 					if self.direction.x < 0:
// 						self.hitbox.left = sprite.hitbox.right
// 					self.rect.centerx = self.hitbox.centerx
// 				else:
// 					if self.direction.y > 0:
// 						self.hitbox.bottom = sprite.hitbox.top
// 					if self.direction.y < 0:
// 						self.hitbox.top = sprite.hitbox.bottom
// 					self.rect.centery = self.hitbox.centery

// 	def update(self, dt):
// 		self.y_sort = self.rect.centery
// 		if not self.blocked:
// 			self.input()
// 			self.move(dt)
// 		self.animate(dt)