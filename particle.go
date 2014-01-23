package vinamax

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

//A Particle essentially constains a position, magnetisation, TODO msat, size?
type Particle struct {
	X, Y, Z             float64
	m                   Vector
	demagnetising_field Vector
	u_anis              Vector  // Uniaxial anisotropy axis
	r                   float64 //radius
	msat                float64 // Saturation magnetisation in A/m
}

//print position and magnitisation of a Particle
func (p Particle) String() string {
	return fmt.Sprintf("Particle@(%v, %v, %v), %v %v %v", p.X, p.Y, p.Z, p.m[0], p.m[1], p.m[2])
}

//Gives all particles the same specifiekd anisotropy-axis
func Anisotropy_axis(x, y, z float64) {
	a := norm(Vector{x, y, z})
	for i := range universe.lijst {
		universe.lijst[i].u_anis = a
	}
}

//Gives all particles a random anisotropy-axis
func Anisotropy_random() {
	for i := range universe.lijst {
		phi := rand.Float64() * (2 * math.Pi)
		theta := rand.Float64() * math.Pi
		universe.lijst[i].u_anis = Vector{math.Sin(theta) * math.Cos(phi), math.Sin(theta) * math.Sin(phi), math.Cos(theta)}
	}
}

//Gives all particles with random magnetisation orientation
func M_random() {
	for i := range universe.lijst {
		phi := rand.Float64() * (2 * math.Pi)
		theta := rand.Float64() * math.Pi
		universe.lijst[i].m = Vector{math.Sin(theta) * math.Cos(phi), math.Sin(theta) * math.Sin(phi), math.Cos(theta)}
	}
}

//Gives all particles a specified magnetisation direction
func M_uniform(x, y, z float64) {
	a := norm(Vector{x, y, z})
	for i := range universe.lijst {
		universe.lijst[i].m = a
	}
}

//Sets the radius of all particles to a consant value
func Particle_radius(x float64) {
	if x < 0 {
		log.Fatal("particles can't have a negative radius")
	}
	for i := range universe.lijst {
		universe.lijst[i].r = x
	}
}

//Gives all particles a radius taken out of a lognormal distribution (mean is specified)
func Lognormal_radius(m float64) {
	mean := m
	s := 0.5
	norm := 1. / (math.Sqrt(2*math.Pi) * s * mean) * math.Exp(-sqr(math.Log(mean/mean))/(2.*s*s))

	for i := range universe.lijst {
		for {
			x := rand.Float64() * 5 * mean
			f_x := 1. / (math.Sqrt(2*math.Pi) * s * x) * math.Exp(-sqr(math.Log(x/mean))/(2.*s*s))
			if rand.Float64() > f_x/norm {
				universe.lijst[i].r = x
				break
			}
		}
	}
}

//Sets the saturation magnetisation of all particles
func Msat(x float64) {
	for i := range universe.lijst {
		universe.lijst[i].msat = x
	}
}
