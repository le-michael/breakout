package game

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/le-michael/breakout/level"
	"github.com/le-michael/breakout/object"
	"github.com/le-michael/breakout/resmgr"
	"github.com/le-michael/breakout/sprite"
)

type GameState int

const (
	GameActive GameState = iota
	GameMenu
	GameWin
)

var (
	playerSize     = mgl32.Vec2{200, 20}
	playerVelocity = float32(500.0)

	ballRadius   = float32(12.5)
	ballVelocity = mgl32.Vec2{-100, -350}
)

type Game struct {
	State  GameState
	Keys   []bool
	Width  int
	Height int

	Levels []*level.GameLevel
	level  uint32

	Player *object.GameObject
	Ball   *object.Ball

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
	if err := resmgr.LoadTexture("textures/paddle.png", false, "paddle"); err != nil {
		return err
	}

	// Load Levels
	one, err := level.Load("levels/one.lvl", g.Width, g.Height/2)
	if err != nil {
		return err
	}
	g.Levels = append(g.Levels, one)

	// Player
	paddleSpr, err := resmgr.GetTexture("paddle")
	if err != nil {
		return err
	}

	playerPos := mgl32.Vec2{float32(g.Width)/2 - playerSize.X()/2, float32(g.Height) - playerSize.Y()}
	g.Player = object.NewGameObject(playerPos, playerSize, mgl32.Vec2{}, mgl32.Vec3{1, 1, 1}, paddleSpr)

	// Ball
	ballSpr, err := resmgr.GetTexture("face")
	if err != nil {
		return err
	}
	ballPos := playerPos.Add(mgl32.Vec2{playerSize.X()/2 - ballRadius, -ballRadius * 2})
	g.Ball = object.NewBall(ballPos, ballRadius, ballVelocity, ballSpr)

	return nil
}

func (g *Game) Update(dt float32) {
	g.Ball.Move(dt, g.Width)
}

func (g *Game) ProcessInput(dt float32) {
	if g.State == GameActive {
		velocity := playerVelocity * dt
		if g.Keys[glfw.KeyA] {
			if g.Player.Position.X() >= 0 {
				g.Player.Position = g.Player.Position.Add(mgl32.Vec2{-velocity, 0})
				if g.Ball.Stuck {
					g.Ball.Position = g.Ball.Position.Add(mgl32.Vec2{-velocity, 0})
				}
			}
		}
		if g.Keys[glfw.KeyD] {
			if g.Player.Position.X() <= float32(g.Width)-playerSize.X() {
				g.Player.Position = g.Player.Position.Add(mgl32.Vec2{velocity, 0})
				if g.Ball.Stuck {
					g.Ball.Position = g.Ball.Position.Add(mgl32.Vec2{velocity, 0})
				}
			}
		}
		if g.Keys[glfw.KeySpace] {
			g.Ball.Stuck = false
		}
	}
}

func (g *Game) Render() {
	if g.State == GameActive {
		g.Levels[g.level].Draw(g.Renderer)
		g.Player.Draw(g.Renderer)
		g.Ball.Draw(g.Renderer)
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
