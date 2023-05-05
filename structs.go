package main

import (
	"fmt"
	"math"

	"github.com/google/go-cmp/cmp"
)

type point struct {
	x float32
	y float32
	z float32
}

type Point struct {
	X    float32
	Y    float32
	Z    float32
	Name []string
}

func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}

func (p point) length() float64 {
	return math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2) + math.Pow(float64(p.z), 2))
}

func (p *point) move(deltax, deltay, deltaz float32) {
	p.x += deltax
	p.y += deltay
	p.z += deltaz
}

func (p1 Point) Equal(p2 Point) bool {
	if p1.X == p2.X && p1.Y == p2.Y && p1.Z == p2.Z {
		return true
	}
	return false
}

func basicStructs() {
	var pt1 point
	pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 4.2
	fmt.Println(pt1.x)
	fmt.Println(pt1.y)
	fmt.Println(pt1.z)
	pt2 := point{x: 5.6, y: 3.8, z: 6.9}
	pt3 := point{x: 5.6, y: 3.8}
	fmt.Println(pt2)
	fmt.Println(pt3)
	pt4 := newPoint(7.8, 9.1, 2.3)
	fmt.Println(pt4)
	fmt.Println(pt4.x)
	pt5 := pt4
	pt5.x = 0
	fmt.Println(pt4)
	fmt.Println(pt5)
	pt6 := *pt4
	pt6.z = 0
	fmt.Println(pt4)
	fmt.Println(pt6)
	// pt7 := pt6
	// pt8 := &pt7
	fmt.Println(pt4.length())
	pt4.move(0.1, 0.1, 0.1)
	fmt.Println(pt4)
}

func main() {
	basicStructs()
	pt1 := Point{X: 5.6, Y: 3.8, Z: 6.9, Name: []string{"pt1"}}
	pt2 := Point{X: 5.6, Y: 3.8, Z: 6.9, Name: []string{"pt"}}
	pt3 := Point{X: 5.6, Y: 3.8, Z: 6.9, Name: []string{"pt"}}
	fmt.Println(cmp.Equal(pt1, pt2))
	fmt.Println(cmp.Equal(pt2, pt3))
}
