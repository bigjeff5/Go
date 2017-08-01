package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
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

// String converts life state from bool to string
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

const (
	width      = 800
	height     = 600
	cellPixels = 10
)

func main() {
	w := width / cellPixels
	h := height / cellPixels

	life := newLife(w, h)
	//window := newSDL(x, y)
	for r := 0; r < 300; r++ {
		life.step()
		fmt.Print("\x0c", life)
		time.Sleep(time.Millisecond * 66)
		//window.UpdateSurface
	}

}
