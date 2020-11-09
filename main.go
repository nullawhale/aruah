package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var window *sdl.Window
var renderer *sdl.Renderer
var texture *sdl.Texture

const ScreenWidth = 800
const ScreenHeight = 600

func loadMedia() (err error) {
	if texture, err = loadTexture("assets/texture.png"); err != nil {
		return err
	}

	return nil
}

func destroy() (err error) {
	if err = texture.Destroy(); err != nil {
		return err
	}
	texture = nil
	if err = renderer.Destroy(); err != nil {
		return err
	}
	renderer = nil
	if err = window.Destroy(); err != nil {
		return err
	}

	img.Quit()
	sdl.Quit()
	return nil
}

func loadTexture(path string) (newTexture *sdl.Texture, err error) {
	var loadedSurface *sdl.Surface
	if loadedSurface, err = img.Load(path); err != nil {
		return nil, err
	}

	if newTexture, err = renderer.CreateTextureFromSurface(loadedSurface); err != nil {
		return nil, err
	}

	loadedSurface.Free()

	return newTexture, nil
}

func run() (err error) {
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	if window, err = sdl.CreateWindow("Input",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth, ScreenHeight, sdl.WINDOW_SHOWN); err != nil {
		return err
	}
	defer window.Destroy()

	if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return err
	}

	imgFlags := img.INIT_PNG
	if err = img.Init(imgFlags); err != nil {
		return err
	}

	if err = loadMedia(); err != nil {
		return
	}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		if err = renderer.SetDrawColor(168, 235, 254, 255); err != nil {
			return err
		}
		if err = renderer.Clear(); err != nil {
			return err
		}

		var fillRect = sdl.Rect{X: ScreenWidth / 4, Y: ScreenHeight / 4, W: ScreenWidth / 2, H: ScreenHeight / 2}
		if err = renderer.SetDrawColor(87, 187, 254, 255); err != nil {
			return err
		}
		if err = renderer.FillRect(&fillRect); err != nil {
			return err
		}

		var outlineRect = sdl.Rect{X: ScreenWidth / 6, Y: ScreenHeight / 6, W: ScreenWidth * 2 / 3, H: ScreenHeight * 2 / 3}
		if err = renderer.SetDrawColor(255, 255, 255, 255); err != nil {
			return err
		}
		if err = renderer.DrawRect(&outlineRect); err != nil {
			return err
		}

		if err = renderer.SetDrawColor(255, 0, 0, 255); err != nil {
			return err
		}
		if err = renderer.DrawLine(0, ScreenHeight/2, ScreenWidth, ScreenHeight/2); err != nil {
			return err
		}

		if err = renderer.SetDrawColor(255, 255, 0, 255); err != nil {
			return err
		}
		for i := 0; i < ScreenHeight; i += 4 {
			if err = renderer.DrawPoint(ScreenWidth/2, int32(i)); err != nil {
				return err
			}
		}

		if err = renderer.Copy(texture, nil, &fillRect); err != nil {
			return err
		}
		renderer.Present()
	}

	return
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	if err := destroy(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
