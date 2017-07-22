package main

import (
	"log"
	
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth = 800
	screenHeight = 600
)

var (
	err				error
	window			*sdl.Window
	screenSurface	*sdl.Surface
	helloWorld		*sdl.Surface
)

func sdlInit() error {
	err = sdl.Init(sdl.INIT_EVERYTHING)
	
	if err != nil {
		return err
	}
	
	window, err = sdl.CreateWindow("Load Image 2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			screenWidth, screenHeight, sdl.WINDOW_SHOWN)
			
	if err != nil {
		return err
	}
	
	screenSurface, err = window.GetSurface()
	if err != nil {
		return err
	}
	
	return nil
}

func loadMedia() error {
	helloWorld, err = sdl.LoadBMP("hello_world.bmp")
	if err != nil {
		return err
	}
	
	return nil
}

func close() {
//	helloWorld.Free()
//	window.Destroy()
	sdl.Quit()
}

func main() {
	err = sdlInit()
	if err != nil {
		log.Fatal("Unable to initialize SDL! Error:", err)
	}
	
	err = loadMedia()
	if err != nil {
		log.Fatal("Unable to load media! Error:", err)
	}
	
	helloWorld.Blit(nil, screenSurface, nil)
	window.UpdateSurface()
	sdl.Delay(3000)
	
	close()
}



