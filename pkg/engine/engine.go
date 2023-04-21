package engine

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/notarock/dungeon/pkg/dungeon"
	"github.com/veandco/go-sdl2/sdl"
)

var game dungeon.Game

var win *sdl.Window
var ren *sdl.Renderer

func InitGame() {
	var err error

	sdl.Init(sdl.INIT_EVERYTHING)

	game, err = dungeon.NewGame()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	win, err = sdl.CreateWindow("Dungeon", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 100*8, 40*8, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Panicln(err)
	}
	defer win.Destroy()

	ren, err = sdl.CreateRenderer(win, -1, 0)
	if err != nil {
		log.Panicln(err)
	}
	defer ren.Destroy()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, moveUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, moveDown); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, moveLeft); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, moveRight); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func moveLeft(g *gocui.Gui, v *gocui.View) error {
	game.Move(dungeon.Left)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func moveRight(g *gocui.Gui, v *gocui.View) error {
	game.Move(dungeon.Right)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func moveUp(g *gocui.Gui, v *gocui.View) error {
	game.Move(dungeon.Up)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func moveDown(g *gocui.Gui, v *gocui.View) error {
	game.Move(dungeon.Down)
	dv, _ := g.View("game")
	redrawMap(dv)
	return nil
}

func layout(g *gocui.Gui) error {
	gameX, gameY := game.GetMapSize()
	if _, err := g.SetView("game", 0, 0, gameY+1, gameX+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func redrawMap(gv *gocui.View) {
	gv.Clear()
	gv.Write([]byte(game.DrawGame(ren)))
}
