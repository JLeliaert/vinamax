package vinamax

import (
	"fmt"
	"math"
	"math/rand"
)

//A Particle essentially constains a position, magnetisation, TODO msat, size?
type Particle struct {
	X, Y, Z             float64
	m                   Vector
	demagnetising_field Vector
	u_anis              Vector // Uniaxial anisotropy axis
}

//Particles[] is a list of Particles
type Particles []Particle

//add a Particle to the list
func (l *Particles) append(p Particle) {
	(*l) = append(*l, p)
}

//print position and magnitisation of a Particle
func (p Particle) String() string {
	return fmt.Sprintf("Particle@(%v, %v, %v), %v %v %v", p.X, p.Y, p.Z, p.m[0], p.m[1], p.m[2])
}

func Anisotropy_axis(x,y,z float64) {
	a:= norm(Vector{x,y,z})
	for i := range Lijst {
		Lijst[i].u_anis = a
	}
}

func Anisotropy_random() {
	for i := range Lijst {
		phi := rand.Float64() * (2 * math.Pi)
		theta := rand.Float64() * math.Pi
		Lijst[i].u_anis = Vector{math.Sin(theta) * math.Cos(phi), math.Sin(theta) * math.Sin(phi), math.Cos(theta)}
	}
}



func M_random() {
for i := range Lijst {
		phi := rand.Float64() * (2 * math.Pi)
		theta := rand.Float64() * math.Pi
		Lijst[i].m = Vector{math.Sin(theta) * math.Cos(phi), math.Sin(theta) * math.Sin(phi), math.Cos(theta)}
	}
}

func M_uniform(x,y,z float64) {
	a:= norm(Vector{x,y,z})
	for i := range Lijst {
		Lijst[i].m = a
	}
}
