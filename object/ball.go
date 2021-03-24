package object

import (
	"github.com/go-gl/mathgl/mgl32"

	"github.com/le-michael/breakout/texture"
)

type Ball struct {
	GameObject
	Radius float32
	Stuck  bool
}

func (b *Ball) Move(dt float32, windowWidth int) mgl32.Vec2 {
	if !b.Stuck {
		b.Position = b.Position.Add(b.Velocity.Mul(dt))
		if b.Position.X() <= 0 {
			b.Velocity = mgl32.Vec2{-b.Velocity.X(), b.Velocity.Y()}
			b.Position = mgl32.Vec2{0, b.Position.Y()}
		} else if b.Position.X()+b.Size.X() >= float32(windowWidth) {
			b.Velocity = mgl32.Vec2{-b.Velocity.X(), b.Velocity.Y()}
			b.Position = mgl32.Vec2{float32(windowWidth) - b.Size.X(), b.Position.Y()}
		}

		if b.Position.Y() <= 0 {
			b.Velocity = mgl32.Vec2{b.Velocity.X(), -b.Velocity.Y()}
			b.Position = mgl32.Vec2{b.Position.X(), 0}
		}
	}
	return b.Position
}

func (b *Ball) Reset(pos, vel mgl32.Vec2) {
	b.Position = pos
	b.Velocity = vel
	b.Stuck = true
}

func NewBall(pos mgl32.Vec2, radius float32, vel mgl32.Vec2, sprite *texture.Texture2D) *Ball {
	b := &Ball{}
	b.Position = pos
	b.Radius = radius
	b.Velocity = vel
	b.Sprite = sprite
	b.Stuck = true
	b.Size = mgl32.Vec2{radius * 2, radius * 2}
	b.Color = mgl32.Vec3{1, 1, 1}
	return b
}
