package vinamax

import(
	"math"
	"math/rand"
	"fmt"
)

//A Particle essentially constains a position, magnetisation, TODO msat, size?
type Particle struct {
	X, Y, Z             float64
	M                   [3]float64
	demagnetising_field [3]float64
	u_anis              [3]float64 // Uniaxial anisotropy axis
}

//Particles[] is a list of Particles
type Particles []Particle

//add a Particle to the list
func (l *Particles) Append(p Particle) {
	(*l) = append(*l, p)
}

//print position and magnitisation of a Particle
func (p Particle) String() string {
	return fmt.Sprintf("Particle@(%v, %v, %v), %v %v %v", p.X, p.Y, p.Z, p.M[0], p.M[1], p.M[2])
}

func Anisotropy_axis(a, b, c float64) {
	norm := math.Sqrt(a*a + b*b + c*c)
	a /= norm
	b /= norm
	c /= norm
	for i := range Lijst {
		Lijst[i].u_anis = [3]float64{a, b, c}
	}
}

func Anisotropy_random() {
	for i := range Lijst {
		phi := rand.Float64() * (2 * math.Pi)
		theta := rand.Float64() * math.Pi
		Lijst[i].u_anis = [3]float64{math.Sin(theta)*math.Cos(phi),math.Sin(theta)*math.Sin(phi),math.Cos(theta) }
	}
}
