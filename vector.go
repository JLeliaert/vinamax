//This file contains some useful vector operators

package vinamax

import "math"

type vector [3]float64

//Dot product between two vectors
func (x vector) dot(y vector) float64 {
	return (x[0]*y[0] + x[1]*y[1] + x[2]*y[2])
}

//cross product between two vectors
func (x *vector) cross(y vector) vector {
	return vector{x[1]*y[2] - x[2]*y[1], y[0]*x[2] - y[2]*x[0], x[0]*y[1] - x[1]*y[0]}
}

//Set norm of a vector to one
// uses the taylor expansion of sqrt because it's close to 1 anyway, and adds lots of speed
func norm(x vector) vector {
	magnitude := math.Sqrt(x[0]*x[0] + x[1]*x[1] + x[2]*x[2])
	//	magnitude := ((x[0]*x[0]+x[1]*x[1]+x[2]*x[2])-1)/2. + 1
	return x.times(1. / magnitude)
}

//returns the norm of a vector
func size(x vector) float64 {
	return math.Sqrt(x[0]*x[0] + x[1]*x[1] + x[2]*x[2])
}

//multiply each component of a vector by a float
func (x vector) times(i float64) vector {
	return vector{x[0] * i, x[1] * i, x[2] * i}
}

//Add two vectors
func (x vector) add(i vector) vector {
	x[0] += i[0]
	x[1] += i[1]
	x[2] += i[2]
	return x
}

func (x *vector) directadd(i vector) {
	x[0] += i[0]
	x[1] += i[1]
	x[2] += i[2]
}

//cubes a number
func cube(x float64) float64 {
	return x * x * x
}

//squares a number
func sqr(x float64) float64 {
	return x * x
}
