package ebiten_pokemon


type Player struct {
	Entity
}



func NewPlayer(pos math2D.Vector2D, frames []ebiten_extended.Sprite, groups []ebiten_extended.Node2D, facing_direction EntityState, collision_sprites []ebiten_extended.Node2D) *Player {
	player := &Player{
		Entity: *NewEntity(pos, frames, groups, facing_direction),
		collision_sprites: collision_sprites,
		noticed: false,
	}
	return player
}


func (p *Player) Input() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		p.direction.SetY(-1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		p.direction.SetY(1)
	} else {
		p.direction.SetY(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		p.direction.SetX(1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		p.direction.SetX(-1)
	} else {
		p.direction.SetX(0)
	}
}

func (p *Player) Move(dt float64) {
	// normalizing a vector
	if p.direction.Magnitude() > 0 {
		p.direction = p.direction.Normalize()
	}

	// horizontal movement
	horizontal_position := math2D.AddVectors(p.GetPosition(), p.direction.MultiplyScalar(float64(p.speed) * dt))

	// vertical movement
	vertical_position := math2D.AddVectors(p.GetPosition(), p.direction.MultiplyScalar(float64(p.speed) * dt))

	p.SetPosition(horizontal_position.X(), vertical_position.Y())
}


func (p *Player) Collisions(axis string) {
	for _, sprite := range p.collision_sprites {
		if sprite.GetHitbox().CollidesWith(p.GetHitbox()) {
			if axis == "horizontal" {
				if p.direction.X() > 0 {
					p.GetHitbox().SetRight(sprite.GetHitbox().Left())
				} else if p.direction.X() < 0 {
					p.GetHitbox().SetLeft(sprite.GetHitbox().Right())
				}
				p.SetPosition(p.GetHitbox().CenterX(), p.GetPosition().Y())
			} else {
				if p.direction.Y() > 0 {
					p.GetHitbox().SetBottom(sprite.GetHitbox().Top())
				} else if p.direction.Y() < 0 {
					p.GetHitbox().SetTop(sprite.GetHitbox().Bottom())
				}
				p.SetPosition(p.GetPosition().X(), p.GetHitbox().CenterY())
			}
		}
	}
}

func (p *Player) Update(dt float64) {
	p.y_sort = p.GetPosition().Y()
	if !p.blocked {
		p.Input()
		p.Move(dt)
	}
	p.Collisions("horizontal")
	p.Collisions("vertical")
	p.Animate(dt)
}





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