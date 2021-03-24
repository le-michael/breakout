package sprite

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/le-michael/breakout/shader"
	"github.com/le-michael/breakout/texture"
)

type SpriteRenderer struct {
	Shader  *shader.Shader
	QuadVAO uint32
}

func (s *SpriteRenderer) init() {
	var vbo uint32
	verticies := []float32{
		// pos      // tex
		0.0, 1.0, 0.0, 1.0,
		1.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 0.0,

		0.0, 1.0, 0.0, 1.0,
		1.0, 1.0, 1.0, 1.0,
		1.0, 0.0, 1.0, 0.0,
	}

	gl.GenVertexArrays(1, &s.QuadVAO)
	gl.GenBuffers(1, &vbo)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*4, gl.Ptr(verticies), gl.STATIC_DRAW)

	gl.BindVertexArray(s.QuadVAO)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 4*4, gl.PtrOffset(0))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

func (s *SpriteRenderer) Draw(tex *texture.Texture2D, position mgl32.Vec2, size mgl32.Vec2, rotate float32, color mgl32.Vec3) {
	s.Shader.Use()
	model := mgl32.Ident4()

	model = model.Mul4(mgl32.Translate3D(position.X(), position.Y(), 0))

	model = model.Mul4(mgl32.Translate3D(0.5*size.X(), 0.5*size.Y(), 0))
	model = model.Mul4(mgl32.Rotate3DZ(rotate).Mat4())
	model = model.Mul4(mgl32.Translate3D(-0.5*size.X(), -0.5*size.Y(), 0))

	model = model.Mul4(mgl32.Scale3D(size.X(), size.Y(), 1))

	s.Shader.SetMatrix4("model\x00", model, false)

	s.Shader.SetVector3fv("spriteColor\x00", color, false)

	gl.ActiveTexture(gl.TEXTURE0)
	tex.Bind()

	gl.BindVertexArray(s.QuadVAO)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
	gl.BindVertexArray(0)
}

func New(shader *shader.Shader) *SpriteRenderer {
	renderer := &SpriteRenderer{
		Shader: shader,
	}

	renderer.init()
	return renderer
}
