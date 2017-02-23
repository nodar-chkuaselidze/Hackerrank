package main

import "fmt"
import "math"

func vec(a, b Point) (vec Point) {
	vec.x = b.x - a.x
	vec.y = b.y - a.y
	return
}

func dotProd(a, b Point) float64 {
	return a.x*b.x + a.y*b.y
}

type Point struct {
	x float64
	y float64
}

type Canvas struct {
	w float64
	h float64
}

type Circle struct {
	c Point
	r float64
}

func (c Circle) inCircle(p Point) bool {
	len := math.Sqrt(math.Pow(float64(p.x-c.c.x), 2) + math.Pow(float64(p.y-c.c.y), 2))

	return len <= float64(c.r)
}

type Square struct {
	p1 Point
	p2 Point
	p3 Point
	p4 Point
}

func (s Square) inSquare(p Point) bool {
	AB := vec(s.p1, s.p2)
	BC := vec(s.p2, s.p3)
	AM := vec(s.p1, p)
	BM := vec(s.p2, p)

	ABxAM := dotProd(AB, AM)
	ABxAB := dotProd(AB, AB)
	BCxBM := dotProd(BC, BM)
	BCxBC := dotProd(BC, BC)

	return 0 <= ABxAM && ABxAM <= ABxAB && 0 <= BCxBM && BCxBM <= BCxBC
}

func (sq *Square) init() {
	cx := (sq.p1.x + sq.p3.x) / 2
	cy := (sq.p1.y + sq.p3.y) / 2

	dx := (sq.p1.x - sq.p3.x) / 2
	dy := (sq.p1.y - sq.p3.y) / 2

	sq.p2 = Point{
		x: cx - dy,
		y: cy + dx,
	}

	sq.p4 = Point{
		x: cx + dy,
		y: cy - dx,
	}
}

func main() {
	var canvas Canvas
	var circle Circle
	var square Square

	fmt.Scanln(&canvas.w, &canvas.h)
	fmt.Scanln(&circle.c.x, &circle.c.y, &circle.r)
	fmt.Scanln(&square.p1.x, &square.p1.y, &square.p3.x, &square.p3.y)

	square.init()

	for i := 0; i < int(canvas.h); i++ {
		for j := 0; j < int(canvas.w); j++ {
			p := Point{float64(j), float64(i)}

			if circle.inCircle(p) {
				fmt.Print("#")
				continue
			}

			if p == square.p1 || p == square.p3 {
				fmt.Print("#")
				continue
			}

			if square.inSquare(p) {
				fmt.Print("#")
				continue
			}

			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}
