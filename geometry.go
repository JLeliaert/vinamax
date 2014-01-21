package vinamax

import (
	"fmt"
	"math/rand"
)

//Adds a single particle
func Addsingleparticle(x, y, z float64) {
	if universe.inworld(Vector{x, y, z}) {
		a := Particle{X: x, Y: y, Z: z, r: 20e-9, msat: 400e3} //magnetite
		universe.lijst = append(universe.lijst, &a)
		universe.number += 1
	} else {
		fmt.Println("error: outside of universe")
	}
}

// Resets the simulation to all standard values:
// Deletes all particles
// Resets B_ext to (0,0,0) T
// Dt = 1e-15 s
// T=0 s
// Alpha=0.02
// Temp=0K
// Ku1=0
// Thresholdbeta=0.7
// use FMM method

func Reset() {
	universe.lijst = nil
	B_ext = func(t float64) (float64, float64, float64) { return 0., 0., 0. }
	Dt = 1e-15
	T = 0
	Alpha = 0.02
	Temp = 0.
	Ku1 = 0
	Thresholdbeta = 0.7
	FMM = true
	//todo de 8 subnodes nil maken
}

type Cube struct {
	x, y, z float64 //position
	S       float64 //diameter
	N       int     //numberofparticles
}

//Adds a lot of particles at random locations in a cubic region
func (c Cube) Addparticles(n int) {
	c.N += n
	for i := 0; i < n; i++ {
		px := c.x + (-0.5+rand.Float64())*c.S
		py := c.y + (-0.5+rand.Float64())*c.S
		pz := c.z + (-0.5+rand.Float64())*c.S
		Addsingleparticle(px, py, pz)
	}
}

func World(x, y, z, r float64) {
	universe.origin = Vector{x, y, z}
	universe.diameter = r
}

func (w node) inworld(r Vector) bool {
	if r[0] < (w.origin[0] - w.diameter/2) {
		return false
	}
	if r[0] > (w.origin[0] + w.diameter/2) {
		return false
	}
	if r[1] < (w.origin[1] - w.diameter/2) {
		return false
	}
	if r[1] > (w.origin[1] + w.diameter/2) {
		return false
	}
	if r[2] < (w.origin[2] - w.diameter/2) {
		return false
	}
	if r[2] > (w.origin[2] + w.diameter/2) {
		return false
	}
	return true
}
