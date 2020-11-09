package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type KeyPressType string

const(
	KeyPressSurfaceDefault KeyPressType = "KEY_PRESS_SURFACE_DEFAULT"
	KeyPressSurfaceEnter                = "KEY_PRESS_SURFACE_ENTER"
	KeyPressSurfaceUp                   = "KEY_PRESS_SURFACE_UP"
	KeyPressSurfaceDown                 = "KEY_PRESS_SURFACE_DOWN"
	KeyPressSurfaceLeft                 = "KEY_PRESS_SURFACE_LEFT"
	KeyPressSurfaceRight                = "KEY_PRESS_SURFACE_RIGHT"
)

var window *sdl.Window
var gScreenSurface *sdl.Surface
var gCurrentSurface *sdl.Surface
var gKeyPressSurfaces map[KeyPressType]*sdl.Surface

func loadSurface(path string) (*sdl.Surface, error) {
	var optimizedSurface *sdl.Surface

	loadedSurface, err := sdl.LoadBMP(path)
	if err != nil {
		return nil, err
	}

	optimizedSurface, err = loadedSurface.Convert(gScreenSurface.Format, 0)
	if err != nil {
		return nil, err
	}

	loadedSurface.Free()

	return optimizedSurface, nil
}

func loadMedia() (err error) {
	gKeyPressSurfaces = make(map[KeyPressType]*sdl.Surface)

	gKeyPressSurfaces[KeyPressSurfaceDefault], err = loadSurface("assets/press.bmp")
	if err != nil {
		return err
	}

	gKeyPressSurfaces[KeyPressSurfaceEnter], err = loadSurface("assets/hello_world.bmp")
	if err != nil {
		return err
	}

	gKeyPressSurfaces[KeyPressSurfaceUp], err = loadSurface("assets/up.bmp")
	if err != nil {
		return err
	}

	gKeyPressSurfaces[KeyPressSurfaceDown], err = loadSurface("assets/down.bmp")
	if err != nil {
		return err
	}

	gKeyPressSurfaces[KeyPressSurfaceLeft], err = loadSurface("assets/left.bmp")
	if err != nil {
		return err
	}

	gKeyPressSurfaces[KeyPressSurfaceRight], err = loadSurface("assets/right.bmp")
	if err != nil {
		return err
	}

	return
}

func run() (err error) {
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow("Input", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return
	}
	defer window.Destroy()

	gScreenSurface, err = window.GetSurface()
	if err != nil {
		return
	}

	if err = loadMedia(); err != nil {
		return
	}

	gCurrentSurface = gKeyPressSurfaces[KeyPressSurfaceDefault]

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				switch e.Keysym.Sym {
					case sdl.K_RETURN:
						gCurrentSurface = gKeyPressSurfaces[KeyPressSurfaceEnter]
					case sdl.K_LEFT:
						gCurrentSurface = gKeyPressSurfaces[KeyPressSurfaceLeft]
					case sdl.K_RIGHT:
						gCurrentSurface = gKeyPressSurfaces[KeyPressSurfaceRight]
					case sdl.K_UP:
						gCurrentSurface = gKeyPressSurfaces[KeyPressSurfaceUp]
					case sdl.K_DOWN:
						gCurrentSurface = gKeyPressSurfaces[KeyPressSurfaceDown]
					case sdl.K_ESCAPE:
						gCurrentSurface = gKeyPressSurfaces[KeyPressSurfaceDefault]
				}
			}
		}

		if err = gCurrentSurface.BlitScaled(nil, gScreenSurface, nil); err != nil {
			return
		}
		if err = window.UpdateSurface(); err != nil {
			return
		}
	}

	return
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
