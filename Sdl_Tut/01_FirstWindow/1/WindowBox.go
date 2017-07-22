package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var window *sdl.Window
	var screenSurface *sdl.Surface
	var err error

	err = sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		log.Fatal("SDL could not initialize! Error:", err)
	}
	
	defer sdl.Quit

	window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)

	if err != nil {
		log.Fatal("Window could not be created! Error:", err)
	}

	defer window.Destroy()

	screenSurface, err = window.GetSurface()

	if err != nil {
		log.Fatal("Unable to get Surface from Window! Error:", err)
	}

	//	screenSurface.FillRect(nil, sdl.MapRGB(screenSurface.Format, 0xFF, 0xFF, 0xFF))
	//	err = window.UpdateSurface()
	//
	//	if err != nil {
	//		log.Fatal("Unable to update Surface! Error:", err)
	//	}

	rect := sdl.Rect{0, 0, 200, 200}
	screenSurface.FillRect(&rect, 0xFFFF0000)
	err = window.UpdateSurface()

	if err != nil {
		log.Fatal("Unable to update Surface! Error:", err)
	}

	sdl.Delay(2000)

	

}
