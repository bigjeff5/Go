package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	err           error
	window        *sdl.Window
	screenSurface *sdl.Surface
	helloWorld    *sdl.Surface
)

func initSDL() error {
	err = sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		return err
	}

	window, err = sdl.CreateWindow("Load Image Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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
		return nil
	}

	return nil
}

func close() {
	helloWorld.Free()
	window.Destroy()
	sdl.Quit()
}

func main() {
	err = initSDL()
	if err != nil {
		log.Fatal("Error initializing SDL:", err)
	}

	err = loadMedia()
	if err != nil {
		log.Fatal("Error loading media:", err)
	}

	helloWorld.Blit(nil, screenSurface, nil)
	window.UpdateSurface()
	sdl.Delay(3000)

	close()

}
