package vinamax

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

//A Particle essentially constains a position, magnetisation
type Particle struct {
	x, y, z             float64
	m                   Vector
	demagnetising_field Vector
	u_anis              Vector  // Uniaxial anisotropy axis
	r                   float64 //radius
	msat                float64 // Saturation magnetisation in A/m
}

//print position and magnitisation of a Particle
func (p Particle) string() string {
	return fmt.Sprintf("Particle@(%v, %v, %v), %v %v %v", p.x, p.y, p.z, p.m[0], p.m[1], p.m[2])
}

//Gives all particles the same specified anisotropy-axis
func Anisotropy_axis(x, y, z float64) {
	uaniscalled = true
	a := norm(Vector{x, y, z})
	for i := range universe.lijst {
		universe.lijst[i].u_anis = a
	}
}

//Gives all particles a random anisotropy-axis
func Anisotropy_random() {
	uaniscalled = true
	for i := range universe.lijst {
		phi := rand.Float64() * (2 * math.Pi)
		theta := rand.Float64() * math.Pi
		universe.lijst[i].u_anis = Vector{math.Sin(theta) * math.Cos(phi), math.Sin(theta) * math.Sin(phi), math.Cos(theta)}
	}
}

//Gives all particles with random magnetisation orientation
func M_random() {
	magnetisationcalled = true
	for i := range universe.lijst {
		phi := rand.Float64() * (2 * math.Pi)
		theta := rand.Float64() * math.Pi
		universe.lijst[i].m = Vector{math.Sin(theta) * math.Cos(phi), math.Sin(theta) * math.Sin(phi), math.Cos(theta)}
	}
}

//Gives all particles a specified magnetisation direction
func M_uniform(x, y, z float64) {
	magnetisationcalled = true
	a := norm(Vector{x, y, z})
	for i := range universe.lijst {
		universe.lijst[i].m = a
	}
}

//Sets the radius of all particles to a consant value
func Particle_radius(x float64) {
	radiuscalled = true
	if x < 0 {
		log.Fatal("particles can't have a negative radius")
	}
	for i := range universe.lijst {
		universe.lijst[i].r = x
	}
}

//Gives all particles a diameter taken from a lognormal distribution with specified mean and std
func Lognormal_diameter(me, st float64) {
	m := me * 1e9
	s := st * 1e9
	radiuscalled = true
	for i := range universe.lijst {
		for {
			x := rand.Float64() * 200 * m
			f_x := 1. / (math.Sqrt(2*math.Pi) * s * x) * math.Exp(-1./(2.*s*s)*sqr(math.Log(x/m)))
			if rand.Float64() < f_x {
				universe.lijst[i].r = x * 1e-9 / 2.
				break
			}
		}
	}
}

//Sets the saturation magnetisation of all particles
func Msat(x float64) {
	msatcalled = true
	for i := range universe.lijst {
		universe.lijst[i].msat = x
	}
}
