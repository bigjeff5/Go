package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type field struct {
	s    [][]bool
	w, h int
}

// newField(width, height) returns *field
// initializes a new field
func newField(w, h int) *field {
	s := make([][]bool, h)
	for r := range s {
		s[r] = make([]bool, w)
	}
	return &field{s: s, w: w, h: h}
}

// set(x, y, bool)
// sets the value of a given point in the field
func (f *field) set(x, y int, b bool) {
	f.s[y][x] = b
}

// alive(x, y) returns bool
// checks whether a cell is alive or dead - wraps toroidially
func (f *field) alive(x, y int) bool {
	x += f.w
	x %= f.w
	y += f.h
	y %= f.h

	return f.s[y][x]
}

// next(x, y) returns bool - determines the state of a point in the next iteration of life
func (f *field) next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i != 0 || j != 0) && f.alive(x+j, y+i) {
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

// newLife(width, height int) generates a new life game with a random initial state
func newLife(w, h int) *life {
	a := newField(w, h)
	rand.Seed(int64(time.Now().Second()))
	for i := 0; i < (w * h / 4); i++ {
		a.set(rand.Intn(w), rand.Intn(h), true)
	}
	b := newField(w, h)
	return &life{
		a: a, b: b,
		w: w, h: h,
	}
}

// step() iterates the next step of the game state
// recomputes all cells
func (l *life) step() {
	for y := range l.b.s {
		for x := range l.b.s[y] {
			l.b.set(x, y, l.a.next(x, y))
		}
	}
	l.a, l.b = l.b, l.a
}

func (l *life) render(s *sdl.Surface) {

	for y := range l.a.s {
		for x := range l.a.s[y] {
			rect := sdl.Rect{X: int32(cellPixels * x), Y: int32(cellPixels * y), W: int32(cellPixels), H: int32(cellPixels)}
			s.FillRect(&rect, 0xFFFFFFFF)
			if l.a.alive(x, y) {
				s.FillRect(&rect, 0x00000000)
			}
		}
	}
}

// String converts life state from bool to string
// must capitalize func name so fmt.Print can use it
// for console output
func (l *life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')
			if l.a.alive(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

// initSDL generates the game window and creates a hook to it
func initSDL(w, h int) (*sdl.Window, *sdl.Surface, error) {
	var err error
	var window *sdl.Window
	var surface *sdl.Surface

	sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Failed to init SDL. Error: %v", err)
		return window, surface, err
	}

	window, err = sdl.CreateWindow("Jeff's Conway's Game of Life", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN)

	if err != nil {
		fmt.Printf("Failed to set sdl.Window. Error: %v", err)
		return window, surface, err
	}

	surface, err = window.GetSurface()

	if err != nil {
		fmt.Printf("Failed to set sdl.Window. Error: %v", err)
		return window, surface, err
	}

	return window, surface, err

}

const (
	width      = 800
	height     = 600
	cellPixels = 10
)

func main() {
	w := width / cellPixels
	h := height / cellPixels

	life := newLife(w, h)
	window, surface, err := initSDL(width, height)

	if err != nil {
		fmt.Printf("Error initializing SDL. Error: %v", err)
	}

	defer sdl.Quit()

	defer window.Destroy()

	life.render(surface)
	window.UpdateSurface()
	quit := false
	var event sdl.Event

	for !quit { //r := 0; r < 300; r++ {
		now := time.Now().Second()
		life.step()
		life.render(surface)
		window.UpdateSurface()

		// Checks
		event = sdl.PollEvent()
		switch event.(type) {
		case *sdl.QuitEvent:
			quit = true
		}
		sdl.Delay(uint32(30 - (time.Now().Second() - now)))
	}

}
