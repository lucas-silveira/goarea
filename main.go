package goarea

import "math"

// PI é uma constante matemática
const PI = 3.1416

// Circ calcula a área da circunferência
func Circ(radius float64) float64 {
	return math.Pow(radius, 2) * PI
}

// Rect calcula a área de um retângulo
func Rect(base, height float64) float64 {
	return base * height
}

// função privada
func _EqTriangle(base, height float64) float64 {
	return (base * height) / 2
}
