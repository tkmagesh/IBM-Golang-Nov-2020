package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (rect *Rectangle) area() float64 {
	x := rect.x2 - rect.x1
	y := rect.y2 - rect.y1
	return math.Sqrt(x*x + y*y)
}

type Circle struct{
	x, y, r float64
}

func (circle *Circle) area() float64 {
	return math.Pi * circle.r * circle.r
}

type Triangle struct {
	x,y,z float64
}

func (triangle *Triangle) area() float64 {
	return float64(30)
}

type Shape interface {
	area() float64
}

func main(){
	
	/* shapes = append(shapes, &Rectangle{ x1 : 10, x2 : 20, y1 : 20 , y2 : 40 }) */

	/* rectangles := []*Rectangle{
		&Rectangle{ x1 : 10, x2 : 20, y1 : 20 , y2 : 40 },
		&Rectangle{ x1 : 10, x2 : 20, y1 : 20 , y2 : 40 },
	} */
	
	/*
	circles := []*Circle{ 
		&Circle{x : 10, y : 20, r : 10 },
		&Circle{x : 10, y : 20, r : 20 },
		&Circle{x : 10, y : 20, r : 30 },
	}

	triangles := []*Triangle{
		&Triangle{x : 10, y : 20, z : 30},
		&Triangle{x : 10, y : 20, z : 30},
		&Triangle{x : 10, y : 20, z : 30},
	} */

	shapes := []Shape{
		&Rectangle{ x1 : 10, x2 : 20, y1 : 20 , y2 : 40 },
		&Rectangle{ x1 : 10, x2 : 20, y1 : 20 , y2 : 40 },
		&Circle{x : 10, y : 20, r : 10 },
		&Circle{x : 10, y : 20, r : 20 },
		&Circle{x : 10, y : 20, r : 30 },
		&Triangle{x : 10, y : 20, z : 30},
		&Triangle{x : 10, y : 20, z : 30},
		&Triangle{x : 10, y : 20, z : 30},
	}
	fmt.Println(sumOfAreas(shapes))
	
}

func sumOfAreas(shapes []Shape) float64 {
	sumOfArea := float64(0);
	for  _, shape := range shapes{
		sumOfArea += shape.area()
	}
	return sumOfArea
}

/* 
to do:
	create a collection of rectangles
	create a collection of circles
	create a function 'sumOfAreas()' which will calculate the sum of areas of rectangles and circles
*/