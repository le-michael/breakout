package game

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
}

func (g *Game) Init() {

}

func (g *Game) Update(dt float64) {

}

func (g *Game) ProcessInput(dt float64) {

}

func (g *Game) Render() {

}

func New(width, height int) *Game {
	return &Game{
		State:  GameActive,
		Keys:   make([]bool, 1024),
		Width:  width,
		Height: height,
	}
}
