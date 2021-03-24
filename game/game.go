package game

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/le-michael/breakout/level"
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

	Levels []*level.GameLevel
	level  uint32

	Renderer *sprite.SpriteRenderer
}

func (g *Game) Init() error {
	// Load Shader
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

	// Load Textures
	if err := resmgr.LoadTexture("textures/background.jpg", false, "background"); err != nil {
		return err
	}
	if err := resmgr.LoadTexture("textures/awesomeface.png", true, "face"); err != nil {
		return err
	}
	if err := resmgr.LoadTexture("textures/block.png", false, "block"); err != nil {
		return err
	}
	if err := resmgr.LoadTexture("textures/block_solid.png", false, "block_solid"); err != nil {
		return err
	}

	// Load Levels
	one, err := level.Load("levels/one.lvl", g.Width, g.Height/2)
	if err != nil {
		return err
	}
	g.Levels = append(g.Levels, one)
	return nil
}

func (g *Game) Update(dt float64) {

}

func (g *Game) ProcessInput(dt float64) {

}

func (g *Game) Render() {
	if g.State == GameActive {
		g.Levels[g.level].Draw(g.Renderer)
		fmt.Println("Rendering level")
	}
}

func New(width, height int) *Game {
	return &Game{
		State:  GameActive,
		Keys:   make([]bool, 1024),
		Width:  width,
		Height: height,
	}
}
