package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func (s Square) Name() string {
	return "Persegi"
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Name() string {
	return "Persegi Panjang"
}

type Triangle struct {
	Base   float64
	Height float64
	SideA  float64
	SideB  float64
	SideC  float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func (t Triangle) Name() string {
	return "Segitiga"
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Name() string {
	return "Lingkaran"
}

type Trapezoid struct {
	Base1  float64
	Base2  float64
	Height float64
	SideA  float64
	SideB  float64
}

func (t Trapezoid) Area() float64 {
	return 0.5 * (t.Base1 + t.Base2) * t.Height
}

func (t Trapezoid) Perimeter() float64 {
	return t.Base1 + t.Base2 + t.SideA + t.SideB
}

func (t Trapezoid) Name() string {
	return "Trapesium"
}

func ProcessShape(sh Shape) {
	fmt.Println("--------------------------------------")
	fmt.Println("Bangun Datar :", sh.Name())
	fmt.Println("Luas         :", sh.Area())
	fmt.Println("Keliling     :", sh.Perimeter())
	fmt.Println("--------------------------------------")
}

func main() {

	square := Square{Side: 6}
	rectangle := Rectangle{Width: 5, Height: 10}
	triangle := Triangle{Base: 4, Height: 6, SideA: 5, SideB: 6, SideC: 7}
	circle := Circle{Radius: 7}
	trapezoid := Trapezoid{Base1: 10, Base2: 6, Height: 5, SideA: 7, SideB: 8}

	ProcessShape(square)
	ProcessShape(rectangle)
	ProcessShape(triangle)
	ProcessShape(circle)
	ProcessShape(trapezoid)
}
