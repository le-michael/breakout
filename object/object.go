package object

import (
	"github.com/go-gl/mathgl/mgl32"

	"github.com/le-michael/breakout/sprite"
	"github.com/le-michael/breakout/texture"
)

type Object interface {
	Draw(*sprite.SpriteRenderer)
}

type GameObject struct {
	Object
	Position  mgl32.Vec2
	Size      mgl32.Vec2
	Velocity  mgl32.Vec2
	Color     mgl32.Vec3
	Rotation  float32
	IsSolid   bool
	Destroyed bool

	Sprite *texture.Texture2D
}

func (g *GameObject) Draw(renderer *sprite.SpriteRenderer) {
	if !g.Destroyed {
		renderer.Draw(g.Sprite, g.Position, g.Size, g.Rotation, g.Color)
	}
}

func DefaultGameObject() *GameObject {
	return NewGameObject(
		mgl32.Vec2{0, 0},
		mgl32.Vec2{1, 1},
		mgl32.Vec2{0, 0},
		mgl32.Vec3{1, 1, 1},
		nil,
	)
}

func NewGameObject(position, size, velocity mgl32.Vec2, color mgl32.Vec3, sprite *texture.Texture2D) *GameObject {
	return &GameObject{
		Position:  position,
		Size:      size,
		Velocity:  velocity,
		Color:     color,
		Rotation:  0,
		IsSolid:   false,
		Destroyed: false,
		Sprite:    sprite,
	}
}
