package interfaces

import (
	"fmt"
	"image/color"
)

type Door bool

func (d *Door) Open() {
	*d = true
}

func (d *Door) Close() {
	*d = false
}

type Position struct {
	x, y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

type Painter interface {
	Paint()
}

type GeoObject struct {
	Painter
	Position
	Color color.Color
}

type Circle struct {
	GeoObject
	Radius int
}

func (c Circle) Paint() {
	fmt.Println(fmt.Sprintf("Circle { Radius %d, Position %s, Color %s }",
		c.Radius, c.Position, c.Color))
}

type Rectangle struct {
	GeoObject
	Width float32
	Height float32
}

func (r Rectangle) Paint() {
	fmt.Println(fmt.Sprintf("Rectangle { Width %f, Height %f, Position %s, Color %s }",
		r.Width, r.Height, r.Position, r.Color))
}

type Triangle struct {
	GeoObject
	B Position
	C Position
}

func (t Triangle) Paint() {
	fmt.Println(fmt.Sprintf("Triangle { A %s, B %s, C %s, Color %s }",
		t.Position, t.B, t.C, t.Color))
}