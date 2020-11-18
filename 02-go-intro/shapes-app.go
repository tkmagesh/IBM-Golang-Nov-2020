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

func main(){
	rectangle := &Rectangle{ x1 : 10, x2 : 20, y1 : 20 , y2 : 40 }
	fmt.Println(rectangle.area())

	circle := &Circle{ x : 10, y : 20, r : 30 }
	fmt.Println(circle.area())
}

/* 
to do:
	create a collection of rectangles
	create a collection of circles
	create a function 'sumOfAreas()' which will calculate the sum of areas of rectangles and circles
*/