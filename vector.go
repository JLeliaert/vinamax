//This file contains some usefull vector operators

package vinamax

import (
	"math"
)

//Dot product between two vectors
func Dot(x, y [3]float64) float64 {

	return (x[0]*y[0] + x[1]*y[1] + x[2]*y[2])

}

//cross product between two vectors
func Cross(x, y [3]float64) [3]float64 {
	var z [3]float64
	z[0] = x[1]*y[2] - x[2]*y[1]
	z[1] = y[0]*x[2] - y[2]*x[0]
	z[2] = x[0]*y[1] - x[1]*y[0]
	return z
}

//Set norm of a vector to one
func Norm(x [3]float64) [3]float64 {
	magnitude := math.Sqrt(x[0]*x[0] + x[1]*x[1] + x[2]*x[2])
	x[0] /= magnitude
	x[1] /= magnitude
	x[2] /= magnitude

	return x
}

//multiply each component of a vector by a float
func Times(arr [3]float64, i float64) [3]float64 {
	arr[0] *= i
	arr[1] *= i
	arr[2] *= i
	return arr
}

//Add two vectors
func Add(arr [3]float64, i [3]float64) [3]float64 {
	arr[0] += i[0]
	arr[1] += i[1]
	arr[2] += i[2]
	return arr
}
