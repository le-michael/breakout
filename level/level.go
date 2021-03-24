package level

import (
	"fmt"
	"io/ioutil"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/le-michael/breakout/object"
	"github.com/le-michael/breakout/resmgr"
	"github.com/le-michael/breakout/sprite"
)

type GameLevel struct {
	Bricks []*object.GameObject
}

func (g *GameLevel) Draw(renderer *sprite.SpriteRenderer) {
	for _, brick := range g.Bricks {
		fmt.Println(brick)
		brick.Draw(renderer)
	}
}

func (g *GameLevel) IsCompleted() bool {
	for _, brick := range g.Bricks {
		if !brick.IsSolid && !brick.Destroyed {
			return false
		}
	}
	return true
}

func Load(file string, levelWidth int, levelHeight int) (*GameLevel, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	tileData := [][]byte{}
	tileRow := []byte{}
	for _, r := range content {
		switch r {
		case ' ':
			continue
		case '\n':
			tileData = append(tileData, tileRow)
			tileRow = []byte{}
		default:
			tileRow = append(tileRow, r-'0')
		}
	}

	gameLevel, err := newLevel(tileData, levelWidth, levelHeight)
	if err != nil {
		return nil, fmt.Errorf("unable to initalize level: %v", err)
	}

	return gameLevel, nil
}

func newLevel(tileData [][]byte, levelWidth int, levelHeight int) (*GameLevel, error) {
	height := len(tileData)
	width := len(tileData[0])
	unitWidth := float32(levelWidth) / float32(width)
	unitHeight := float32(levelHeight) / float32(height)

	colors := map[byte]mgl32.Vec3{
		0: {0.8, 0.8, 0.7},
		1: {0.2, 0.6, 1.0},
		2: {0.0, 0.7, 0.0},
		3: {0.8, 0.8, 0.4},
		4: {1.0, 0.5, 0.0},
	}

	gameLevel := &GameLevel{}

	for i, row := range tileData {
		for j, col := range row {
			pos := mgl32.Vec2{unitWidth * float32(j), unitHeight * float32(i)}
			size := mgl32.Vec2{unitWidth, unitHeight}
			vel := mgl32.Vec2{0, 0}
			switch col {
			case '0':
				continue
			case '1':
				tex, err := resmgr.GetTexture("block_solid")
				if err != nil {
					return nil, err
				}
				brick := object.New(pos, size, vel, colors[col], tex)
				brick.IsSolid = true
				gameLevel.Bricks = append(gameLevel.Bricks, brick)
			default:
				tex, err := resmgr.GetTexture("block")
				if err != nil {
					return nil, err
				}
				brick := object.New(pos, size, vel, colors[col], tex)
				gameLevel.Bricks = append(gameLevel.Bricks, brick)
			}
		}
	}

	return gameLevel, nil
}
