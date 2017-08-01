package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 800
	windowHeight = 600
	cellPixels   = 10
)

type field struct {
	s    [][]bool
	w, h int
}

func newField(w, h int) *field {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &field{s: s, w: w, h: h}
}

func (f *field) set(x, y int, b bool) {
	f.s[y][x] = b
}

func (f *field) alive(x, y int) bool {
	x += f.w
	x %= f.w
	y += f.h
	y %= f.h
	return f.s[y][x]
}

func (f *field) next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.alive(x+j, y+i) {
				alive++
			}
		}
	}

	return alive == 3 || alive == 2 && f.alive(x, y)
}

type life struct {
	a, b *field
	w, h int
}

func newLife(w, h int) *life {
	a := newField(w, h)
	b := newField(w, h)
	// for i := 0; i < (w * h / 4); i++ {
	// 	a.set(rand.Intn(w), rand.Intn(h), true)
	// }
	return &life{
		a: a, b: b,
		w: w, h: h,
	}
}

func (l *life) step() {
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.set(x, y, l.a.next(x, y))
		}
	}
	// Swap next field with current field
	l.a, l.b = l.b, l.a
}

func newGlider(f *field, x, y int) {
	f.set(x+1, y, true)
	f.set(x+2, y+1, true)
	f.set(x, y+2, true)
	f.set(x+1, y+2, true)
	f.set(x+2, y+2, true)
}

func main() {
	var window *sdl.Window
	var screenSurface *sdl.Surface
	var err error

	err = sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		log.Fatal("SDL could not initialize! Error:", err)
	}

	defer sdl.Quit()

	window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)

	if err != nil {
		log.Fatal("Window could not be created! Error:", err)
	}

	defer window.Destroy()

	screenSurface, err = window.GetSurface()

	if err != nil {
		log.Fatal("Unable to get Surface from Window! Error:", err)
	}

	screenSurface.FillRect(nil, sdl.MapRGB(screenSurface.Format, 0xFF, 0xFF, 0xFF))
	err = window.UpdateSurface()

	if err != nil {
		log.Fatal("Unable to update Surface! Error:", err)
	}

	width := windowWidth / cellPixels
	height := windowHeight / cellPixels
	life := newLife(width, height)

	newGlider(life.a, 0, 40)
	newGlider(life.a, 20, 40)
	newGlider(life.a, 40, 40)
	newGlider(life.a, 60, 40)
	newGlider(life.a, 60, 0)
	newGlider(life.a, 40, 20)
	newGlider(life.a, 20, 40)

	for r := 0; r < 1000; r++ {
		// Conways GoL
		for y := range life.a.s {
			for x := range life.a.s[y] {
				rect := sdl.Rect{X: int32(x * cellPixels), Y: int32(y * cellPixels), W: cellPixels, H: cellPixels}
				screenSurface.FillRect(&rect, 0xFFFFFFFF)
				if life.a.alive(x, y) == true {
					screenSurface.FillRect(&rect, 0x00000000)
				}
			}
		}

		err = window.UpdateSurface()
		life.step()
		sdl.Delay(50)
	}
	if err != nil {
		log.Fatal("Unable to update Surface! Error:", err)
	}

	sdl.Delay(2000)

}
