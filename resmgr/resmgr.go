package resmgr

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/le-michael/breakout/shader"
	"github.com/le-michael/breakout/texture"
)

type resourceManager struct {
	Textures map[string]*texture.Texture2D
	Shaders  map[string]*shader.Shader
}

var (
	rm = &resourceManager{
		Textures: make(map[string]*texture.Texture2D),
		Shaders:  make(map[string]*shader.Shader),
	}
)

func LoadShader(vFile, fFile, name string) error {
	program, err := loadShaderFromFile(vFile, fFile)
	if err != nil {
		return fmt.Errorf("unable to load shader: %v", err)
	}
	rm.Shaders[name] = program
	return nil
}

func GetShader(name string) (*shader.Shader, error) {
	program, ok := rm.Shaders[name]
	if !ok {
		return nil, fmt.Errorf("unable to find shader program: %v", name)
	}
	return program, nil
}

func GetTexture(name string) (*texture.Texture2D, error) {
	tex, ok := rm.Textures[name]
	if !ok {
		return nil, fmt.Errorf("unable to find texture: %v", name)
	}
	return tex, nil
}

func LoadTexture(tFile string, alpha bool, name string) error {
	tex, err := loadTextureFromFile(tFile, alpha)
	if err != nil {
		return fmt.Errorf("unable to load texture %v: %v", tFile, err)
	}

	rm.Textures[name] = tex
	return nil
}

func Clear() {
	for _, program := range rm.Shaders {
		gl.DeleteProgram(program.ID)
	}
}

func loadTextureFromFile(tFile string, alpha bool) (*texture.Texture2D, error) {
	imgFile, err := os.Open(tFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open %v: %v", tFile, err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("unable to decode texture: %v", err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	tex := texture.New()
	tex.Generate(rgba)

	if alpha {
		tex.ImageFormat = gl.RGBA
		tex.InternalFormat = gl.RGBA
	}

	return tex, nil
}

func loadShaderFromFile(vFile, fFile string) (*shader.Shader, error) {
	vSource, err := ioutil.ReadFile(vFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open %v: %v", vFile, err)
	}

	fSource, err := ioutil.ReadFile(fFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open %v: %v", vFile, err)
	}

	program, err := shader.Compile(string(vSource), string(fSource))
	if err != nil {
		return nil, fmt.Errorf("unable to compile shader program: %v", err)
	}

	return program, nil
}
