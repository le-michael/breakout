package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	ID uint32
}

func (s *Shader) Use() {
	gl.UseProgram(s.ID)
}

func (s *Shader) SetFloat(name string, value float32, useShader bool) {
	if useShader {
		s.Use()
	}
	gl.Uniform1f(s.GetUniformLocation(name), value)
}

func (s *Shader) SetInteger(name string, value int32, useShader bool) {
	if useShader {
		s.Use()
	}
	gl.Uniform1i(s.GetUniformLocation(name), value)
}

func (s *Shader) SetVector2f(name string, x, y float32, useShader bool) {
	if useShader {
		s.Use()
	}
	gl.Uniform2f(s.GetUniformLocation(name), x, y)
}

func (s *Shader) SetVector2fv(name string, vec mgl32.Vec2, useShader bool) {
	if useShader {
		s.Use()
	}
	gl.Uniform2fv(s.GetUniformLocation(name), 1, &vec[0])
}

func (s *Shader) SetVector3f(name string, x, y, z float32, useShader bool) {
	if useShader {
		s.Use()
	}
	gl.Uniform3f(s.GetUniformLocation(name), x, y, z)
}

func (s *Shader) SetVector3fv(name string, vec mgl32.Vec3, useShader bool) {
	if useShader {
		s.Use()
	}
	gl.Uniform3fv(s.GetUniformLocation(name), 1, &vec[0])
}

func (s *Shader) SetVector4f(name string, x, y, z, w float32, useShader bool) {
	if useShader {
		s.Use()
	}
	gl.Uniform4f(s.GetUniformLocation(name), x, y, z, w)
}

func (s *Shader) SetVector4fv(name string, vec mgl32.Vec4, useShader bool) {
	if useShader {
		s.Use()
	}
	gl.Uniform4fv(s.GetUniformLocation(name), 1, &vec[0])
}

func (s *Shader) SetMatrix4(name string, mat mgl32.Mat4, useShader bool) {
	if useShader {
		s.Use()
	}
	gl.UniformMatrix4fv(s.GetUniformLocation(name), 1, false, &mat[0])
}

func (s *Shader) GetUniformLocation(name string) int32 {
	return gl.GetUniformLocation(s.ID, gl.Str(name+"\x00"))
}

func Compile(vertexSource, fragmentSource string) (*Shader, error) {
	sVertex, err := compile(vertexSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	sFragment, err := compile(fragmentSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, sVertex)
	gl.AttachShader(program, sFragment)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(sVertex)
	gl.DeleteShader(sFragment)

	s := &Shader{ID: program}
	return s, nil
}

func compile(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csouce, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csouce, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
