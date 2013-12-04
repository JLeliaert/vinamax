package main

import(
	"math"
	"math/rand"
	"fmt"
)

//A particle essentially constains a position, magnetisation, TODO msat, size?
type particle struct {
	x, y, z             float64
	m                   [3]float64
	demagnetising_field [3]float64
	u_anis              [3]float64 // Uniaxial anisotropy axis
}

//particles[] is a list of particles
type particles []particle

//add a particle to the list
func (l *particles) append(p particle) {
	(*l) = append(*l, p)
}

//print position and magnitisation of a particle
func (p particle) String() string {
	return fmt.Sprintf("particle@(%v, %v, %v), %v %v %v", p.x, p.y, p.z, p.m[0], p.m[1], p.m[2])
}

func anisotropy_axis(a, b, c float64) {
	norm := math.Sqrt(a*a + b*b + c*c)
	a /= norm
	b /= norm
	c /= norm
	for i := range lijst {
		lijst[i].u_anis = [3]float64{a, b, c}
	}
}

func anisotropy_random() {
	for i := range lijst {
		phi := rand.Float64() * (2 * math.Pi)
		theta := rand.Float64() * math.Pi
		lijst[i].u_anis = [3]float64{math.Sin(theta)*math.Cos(phi),math.Sin(theta)*math.Sin(phi),math.Cos(theta) }
	}
}
