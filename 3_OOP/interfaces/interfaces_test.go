package interfaces

import (
	"image/color"
	"testing"
)

func TestInterfaces(t *testing.T) {
	var d Door = false
	d.Open()
	d.Close()

	positionA := Position{1,2}
	positionB := Position{3,4}
	positionC := Position{5,6}
	geoObject := GeoObject{Position: positionA, Color: color.Black}

	circle := Circle{geoObject, 10}
	rectangle := Rectangle{geoObject, 2, 1}
	triangle := Triangle{geoObject, positionB, positionC}

	shapes := []Painter{circle, rectangle, triangle}

	for _, o := range shapes {
		o.Paint()
	}
}
