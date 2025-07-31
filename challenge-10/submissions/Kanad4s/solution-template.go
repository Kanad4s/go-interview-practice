// Package challenge10 contains the solution for Challenge 10.
package challenge10

import (
	"fmt"
	"math"
	"slices"
	// Add any necessary imports here
)

// Shape interface defines methods that all shapes must implement
type Shape interface {
	Area() float64
	Perimeter() float64
	fmt.Stringer // Includes String() string method
}

// Rectangle represents a four-sided shape with perpendicular sides
type Rectangle struct {
	Width  float64
	Height float64
}

// NewRectangle creates a new Rectangle with validation
func NewRectangle(width, height float64) (*Rectangle, error) {
	// TODO: Implement validation and construction
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("invalid value")
	}

	return &Rectangle{Width: width, Height: height}, nil
}

// Area calculates the area of the rectangle
func (r *Rectangle) Area() float64 {
	// TODO: Implement area calculation
	return r.Height * r.Width
}

// Perimeter calculates the perimeter of the rectangle
func (r *Rectangle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return 2 * (r.Height + r.Width)
}

// String returns a string representation of the rectangle
func (r *Rectangle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf("Rectangle(Height: %f, Width: %f)", r.Height, r.Width)
}

// Circle represents a perfectly round shape
type Circle struct {
	Radius float64
}

// NewCircle creates a new Circle with validation
func NewCircle(radius float64) (*Circle, error) {
	// TODO: Implement validation and construction
	if radius <= 0 {
		return nil, fmt.Errorf("invalid radius")
	}
	return &Circle{Radius: radius}, nil
}

// Area calculates the area of the circle
func (c *Circle) Area() float64 {
	// TODO: Implement area calculation
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the circumference of the circle
func (c *Circle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return 2 * math.Pi * c.Radius
}

// String returns a string representation of the circle
func (c *Circle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf("Circle(Radius: %f)", c.Radius)
}

// Triangle represents a three-sided polygon
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

// NewTriangle creates a new Triangle with validation
func NewTriangle(a, b, c float64) (*Triangle, error) {
	// TODO: Implement validation and construction
	if a <= 0 || b <= 0 || c <= 0 {
		return nil, fmt.Errorf("invalid value")
	} else if a+b <= c || a+c <= b || b+c <= a {
		return nil, fmt.Errorf("invalid value")
	}
	return &Triangle{SideA: a, SideB: b, SideC: c}, nil
}

// Area calculates the area of the triangle using Heron's formula
func (t *Triangle) Area() float64 {
	// TODO: Implement area calculation using Heron's formula
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

// Perimeter calculates the perimeter of the triangle
func (t *Triangle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return t.SideA + t.SideB + t.SideC
}

// String returns a string representation of the triangle
func (t *Triangle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf("Triangle - sides=%.3f, %.3f, %.3f", t.SideA, t.SideB, t.SideC)
}

// ShapeCalculator provides utility functions for shapes
type ShapeCalculator struct {
	shapes []Shape
}

// NewShapeCalculator creates a new ShapeCalculator
func NewShapeCalculator() *ShapeCalculator {
	// TODO: Implement constructor
	return &ShapeCalculator{shapes: make([]Shape, 0)}
}

// PrintProperties prints the properties of a shape
func (sc *ShapeCalculator) PrintProperties(s Shape) {
	// TODO: Implement printing shape properties
	for _, shape := range sc.shapes {
		fmt.Println(shape.String())
	}
}

// TotalArea calculates the sum of areas of all shapes
func (sc *ShapeCalculator) TotalArea(shapes []Shape) float64 {
	// TODO: Implement total area calculation
	var area float64
	for _, val := range shapes {
		area += val.Area()
	}
	return area
}

// LargestShape finds the shape with the largest area
func (sc *ShapeCalculator) LargestShape(shapes []Shape) Shape {
	// TODO: Implement finding largest shape
	var maxArea float64
	var maxShape Shape
	for _, shape := range shapes {
		curArea := shape.Area()
		if curArea > maxArea {
			maxArea = curArea
			maxShape = shape
		}
	}
	return maxShape
}

// SortByArea sorts shapes by area in ascending or descending order
func (sc *ShapeCalculator) SortByArea(shapes []Shape, ascending bool) []Shape {
	// TODO: Implement sorting shapes by area
	slices.SortFunc(shapes, func(a, b Shape) int {
		switch {
		case a.Area() > b.Area():
			if ascending {
				return 1
			}
			return -1
		case a.Area() < b.Area():
			if ascending {
				return -1
			}
			return 1
		default:
			return 0
		}

	})
	return shapes
}
