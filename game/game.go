package game

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/le-michael/breakout/resmgr"
	"github.com/le-michael/breakout/sprite"
)

type GameState int

const (
	GameActive GameState = iota
	GameMenu
	GameWin
)

type Game struct {
	State  GameState
	Keys   []bool
	Width  int
	Height int

	Renderer *sprite.SpriteRenderer
}

func (g *Game) Init() error {
	if err := resmgr.LoadShader("shaders/sprite.vert", "shaders/sprites.frg", "sprite"); err != nil {
		return err
	}

	projection := mgl32.Ortho(0, float32(g.Width), float32(g.Height), 0, -1, 1)
	spriteShader, err := resmgr.GetShader("sprite")
	if err != nil {
		return err
	}

	spriteShader.SetInteger("image", 0, true)
	spriteShader.SetMatrix4("projection", projection, false)

	g.Renderer = sprite.New(spriteShader)

	if err := resmgr.LoadTexture("textures/awesomeface.png", true, "face"); err != nil {
		return err
	}
	return nil
}

func (g *Game) Update(dt float64) {

}

func (g *Game) ProcessInput(dt float64) {

}

func (g *Game) Render() {

	tex, err := resmgr.GetTexture("face")
	if err != nil {
		log.Fatalln(err)
	}
	g.Renderer.Draw(
		tex,
		mgl32.Vec2{200, 200},
		mgl32.Vec2{200, 300},
		45,
		mgl32.Vec3{0, 1, 0},
	)
}

func New(width, height int) *Game {
	return &Game{
		State:  GameActive,
		Keys:   make([]bool, 1024),
		Width:  width,
		Height: height,
	}
}
