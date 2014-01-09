//This file contains some usefull vector operators

package vinamax

import (
	"math"
)

type Vector [3]float64

//Dot product between two vectors
func (x Vector) dot(y Vector) float64 {
	return (x[0]*y[0] + x[1]*y[1] + x[2]*y[2])
}

//cross product between two vectors
func (x Vector) cross(y Vector) Vector {
	return Vector{x[1]*y[2] - x[2]*y[1], y[0]*x[2] - y[2]*x[0], x[0]*y[1] - x[1]*y[0]}
}

//Set norm of a vector to one
func norm(x Vector) Vector {
	magnitude := math.Sqrt(x[0]*x[0] + x[1]*x[1] + x[2]*x[2])
	x[0] /= magnitude
	x[1] /= magnitude
	x[2] /= magnitude

	return x
}

//multiply each component of a vector by a float
func (x Vector) times(i float64) Vector {
	x[0] *= i
	x[1] *= i
	x[2] *= i
	return x
}

//Add two vectors
func (x Vector) add(i Vector) Vector {
	x[0] += i[0]
	x[1] += i[1]
	x[2] += i[2]
	return x
}

func cube(x float64) float64 {
	return x * x * x
}

func sqr(x float64) float64 {
	return x * x
}
