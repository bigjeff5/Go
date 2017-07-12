package main

import (
	"log"
	
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth = 640
	screenHeight = 480
)

func main() {
	var window *sdl.Window
	var screenSurface *sdl.Surface
	var err error
	
	err = sdl.Init(sdl.INIT_VIDEO)
	
	if err != nil {
		log.Fatal("Unable to initialize SDL! Error:", err)
	}
	
	window, err = sdl.CreateWindow("SDL Test 2!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)
		
	if err != nil {
		log.Fatal("Unable to create window! Error:", err)
	}
	
	rect := sdl.Rect{0, 0, 200, 200}
	screenSurface, err = window.GetSurface()
	
	if err != nil {
		log.Fatal("Unable to get Surface from window! Error:", err)
	}
	
	screenSurface.FillRect(&rect, 0xFFFF0000)
	
	err = window.UpdateSurface()
	
	if err != nil {
		log.Fatal("Unable to update Surface! Error:", err)
	}
	
	sdl.Delay(2000)
	
	window.Destroy()
	sdl.Quit()
	
}