package ebitenpokemon

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

type Entity struct {
	ebiten_extended.Node2D
	speed     int
	blocked   bool
	direction math2D.Vector2D
	facing_direction string
}

func (e *Entity) SetFacingDirection(direction string) {
	e.facing_direction = direction
}

func (e *Entity) GetFacingDirection() string {	
	return e.facing_direction
}

func NewEntity() *Entity {
	return &Entity{}
}


func (e *Entity) get_state() string {
	moving := e.direction.Magnitude() > 0
	if moving {
		if e.direction.X() != 0 {
			if e.direction.X() > 0 {
				return "right"
			} else {
				return "left"
			}
		}
		if e.direction.Y() != 0 {
			if e.direction.Y() > 0 {
				return "down"
			} else {
				return "up"
			}
		}
	}
	return "idle"
}

func (e *Entity) change_facing_direction(target_pos math2D.Vector2D) {
	relation := math2D.SubtractVectors(target_pos, e.GetPosition())
	if relation.Y < 30 {
		if relation.X > 0 {
			e.SetFacingDirection("right")
		} else {
			e.SetFacingDirection("left")
		}
	} else {
		if relation.Y > 0 {
			e.SetFacingDirection("down")
		} else {
			e.SetFacingDirection("up")
		}
	}
}

func (e *Entity) block() {
	e.blocked = true
	e.direction = math2D.Vector2D{}
}

func (e *Entity) unblock() {
	e.blocked = false
}

func (e *Entity) is_blocked() bool {	
	return e.blocked
}

func (e *Entity) set_speed(speed int) {
	e.speed = speed
}

func (e *Entity) get_speed() int {
	return e.speed
}

func (e *Entity) move() {
	if e.direction.Magnitude() > 0 {
		position := math2D.AddVectors(e.direction.MultiplyScalar(float64(e.speed)), e.GetPosition())
		e.SetPosition(position.X(), position.Y())
	}
}
