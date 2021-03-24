package texture

import (
	"image"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Texture2D struct {
	ID             uint32
	Width          uint32
	Height         uint32
	InternalFormat int32
	ImageFormat    int32
	WrapS          int32
	WrapT          int32
	FilterMin      int32
	FilterMax      int32
}

func (t *Texture2D) Generate(rgba *image.RGBA) {
	t.Width = uint32(rgba.Rect.Size().X)
	t.Height = uint32(rgba.Rect.Size().Y)

	gl.BindTexture(gl.TEXTURE_2D, t.ID)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(t.Width),
		int32(t.Height),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix),
	)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, int32(t.WrapS))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, int32(t.WrapT))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(t.FilterMin))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(t.FilterMax))

	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (t *Texture2D) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.ID)
}

func New() *Texture2D {
	var id uint32
	gl.GenTextures(1, &id)
	return &Texture2D{
		ID:             id,
		Width:          0,
		Height:         0,
		InternalFormat: gl.RGB,
		ImageFormat:    gl.RGB,
		WrapS:          gl.REPEAT,
		WrapT:          gl.REPEAT,
		FilterMin:      gl.LINEAR,
		FilterMax:      gl.LINEAR,
	}
}
