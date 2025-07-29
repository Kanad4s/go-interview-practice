package generics

import (
	"cmp"
	"errors"
	"slices"
)

// ErrEmptyCollection is returned when an operation cannot be performed on an empty collection
var ErrEmptyCollection = errors.New("collection is empty")

//
// 1. Generic Pair
//

// Pair represents a generic pair of values of potentially different types
type Pair[T, U any] struct {
	First  T
	Second U
}

// NewPair creates a new pair with the given values
func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
	// TODO: Implement this function
}

// Swap returns a new pair with the elements swapped
func (p Pair[T, U]) Swap() Pair[U, T] {
	return Pair[U, T]{First: p.Second, Second: p.First}
	// TODO: Implement this method
}

//
// 2. Generic Stack
//

// Stack is a generic Last-In-First-Out (LIFO) data structure
type Stack[T any] struct {
	// TODO: Add necessary fields
	elements []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	// TODO: Implement this function
	return &Stack[T]{elements: make([]T, 0)}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	// TODO: Implement this method
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
// Returns an error if the stack is empty
func (s *Stack[T]) Pop() (T, error) {
	// TODO: Implement this method
	var zero T
	if len(s.elements) == 0 {
		return zero, errors.New("")
	}
	elem := s.elements[len(s.elements)-1]
	s.elements = slices.Delete(s.elements, len(s.elements)-1, len(s.elements))
	return elem, nil
}

// Peek returns the top element without removing it
// Returns an error if the stack is empty
func (s *Stack[T]) Peek() (T, error) {
	// TODO: Implement this method
	var zero T
	if len(s.elements) == 0 {
		return zero, errors.New("")
	}
	return s.elements[len(s.elements)-1], nil
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	// TODO: Implement this method
	return len(s.elements)
}

// IsEmpty returns true if the stack contains no elements
func (s *Stack[T]) IsEmpty() bool {
	// TODO: Implement this method
	return len(s.elements) == 0
}

//
// 3. Generic Queue
//

// Queue is a generic First-In-First-Out (FIFO) data structure
type Queue[T any] struct {
	elements []T
	// TODO: Add necessary fields
}

// NewQueue creates a new empty queue
func NewQueue[T any]() *Queue[T] {
	// TODO: Implement this function
	return &Queue[T]{elements: make([]T, 0)}
}

// Enqueue adds an element to the end of the queue
func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
	// TODO: Implement this method
}

// Dequeue removes and returns the front element from the queue
// Returns an error if the queue is empty
func (q *Queue[T]) Dequeue() (T, error) {
	// TODO: Implement this method
	var zero T
	if len(q.elements) == 0 {
		return zero, errors.New("")
	}
	val := q.elements[0]
	q.elements = slices.Delete(q.elements, 0, 1)
	return val, nil
}

// Front returns the front element without removing it
// Returns an error if the queue is empty
func (q *Queue[T]) Front() (T, error) {
	// TODO: Implement this method
	var zero T
	if len(q.elements) == 0 {
		return zero, errors.New("")
	}
	return q.elements[0], nil
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	// TODO: Implement this method
	return len(q.elements)
}

// IsEmpty returns true if the queue contains no elements
func (q *Queue[T]) IsEmpty() bool {
	// TODO: Implement this method
	return len(q.elements) == 0
}

//
// 4. Generic Set
//

// Set is a generic collection of unique elements
type Set[T comparable] struct {
	Values []T
}

// NewSet creates a new empty set
func NewSet[T comparable]() *Set[T] {
	// TODO: Implement this function
	return &Set[T]{Values: make([]T, 0)}
}

// Add adds an element to the set if it's not already present
func (s *Set[T]) Add(value T) {
	if !s.Contains(value) {
		s.Values = append(s.Values, value)
	}
	// TODO: Implement this method
}

// Remove removes an element from the set if it exists
func (s *Set[T]) Remove(value T) {
	i := slices.Index(s.Values, value)
	if i >= 0 {
		s.Values = slices.Delete(s.Values, i, i+1)
	}
	// TODO: Implement this method
}

// Contains returns true if the set contains the given element
func (s *Set[T]) Contains(value T) bool {
	// TODO: Implement this method
	return slices.Contains(s.Values, value)
}

// Size returns the number of elements in the set
func (s *Set[T]) Size() int {
	// TODO: Implement this method
	return len(s.Values)
}

// Elements returns a slice containing all elements in the set
func (s *Set[T]) Elements() []T {
	// TODO: Implement this method
	return slices.Clone(s.Values)
}

// Union returns a new set containing all elements from both sets
func Union[T comparable](s1, s2 *Set[T]) *Set[T] {
	// TODO: Implement this function
	res := make([]T, 0, len(s1.Values)+len(s2.Values))

	res = append(res, s2.Values...)

	for _, val := range s1.Values {
		if !slices.Contains(res, val) {
			res = append(res, val)
		}
	}
	return &Set[T]{
		Values: res,
	}
}

// Intersection returns a new set containing only elements that exist in both sets
func Intersection[T comparable](s1, s2 *Set[T]) *Set[T] {
	res := []T{}
	// TODO: Implement this function
	for _, val := range s1.Values {
		if slices.Contains(s2.Values, val) {
			res = append(res, val)
		}
	}
	return &Set[T]{Values: res}
}

// Difference returns a new set with elements in s1 that are not in s2
func Difference[T comparable](s1, s2 *Set[T]) *Set[T] {
	res := []T{}
	// TODO: Implement this function
	for _, val := range s1.Values {
		if !slices.Contains(s2.Values, val) {
			res = append(res, val)
		}
	}
	return &Set[T]{Values: res}
}

//
// 5. Generic Utility Functions
//

// Filter returns a new slice containing only the elements for which the predicate returns true
func Filter[T any](slice []T, predicate func(T) bool) []T {
	res := []T{}
	for _, val := range slice {
		if predicate(val) {
			res = append(res, val)
		}
	}
	// TODO: Implement this function
	return res
}

// Map applies a function to each element in a slice and returns a new slice with the results
func Map[T, U any](slice []T, mapper func(T) U) []U {
	// TODO: Implement this function
	res := make([]U, 0, len(slice))
	for _, val := range slice {
		res = append(res, mapper(val))
	}
	return res
}

// Reduce reduces a slice to a single value by applying a function to each element
func Reduce[T, U any](slice []T, initial U, reducer func(U, T) U) U {
	// TODO: Implement this function
	for _, val := range slice {
		initial = reducer(initial, val)
	}
	return initial
}

// Contains returns true if the slice contains the given element
func Contains[T comparable](slice []T, element T) bool {
	// TODO: Implement this function
	return slices.Contains(slice, element)
}

// FindIndex returns the index of the first occurrence of the given element or -1 if not found
func FindIndex[T comparable](slice []T, element T) int {
	// TODO: Implement this function
	return slices.Index(slice, element)
}

// RemoveDuplicates returns a new slice with duplicate elements removed, preserving order
func RemoveDuplicates[T cmp.Ordered](slice []T) []T {
	// TODO: Implement this function
	res := slices.Clone(slice)
	slices.Sort(res)
	res = slices.Compact(res)
	return res
}
