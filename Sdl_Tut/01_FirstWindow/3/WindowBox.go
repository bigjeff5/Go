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
		log.Fatal("Unable to initialize SDL! Error", err)
	}
	
	window, err = sdl.CreateWindow("SDL Window Test 3!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)
		
	if err != nil {
		log.Fatal("Unable to create Window! Error", err)
	}
	
	screenSurface, err = window.GetSurface()
	
	if err != nil {
		log.Fatal("Unable to get Surface from Window! Error", err)
	}
	
	rect := sdl.Rect{0, 0, 300, 300}
	screenSurface.FillRect(&rect, 0xFFFFF000)
	
	err = window.UpdateSurface()
	
	if err != nil {
		log.Fatal("Unable to update window Surface! Error", err)
	}
	
	sdl.Delay(3000)
	window.Destroy()
	sdl.Quit()
	
}