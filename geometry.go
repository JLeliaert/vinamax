package vinamax

import (
	"fmt"
	"math/rand"
)

//Adds a single particle at specified coordinates
func Addsingleparticle(x, y, z float64) {
	if universe.inworld(Vector{x, y, z}) {
		a := particle{x: x, y: y, z: z}
		universe.lijst = append(universe.lijst, &a)
		universe.number += 1
		msatcalled = false
		radiuscalled = false
	} else {
		fmt.Println("error: outside of universe")
	}
}

type Cube struct {
	x, y, z float64 //position
	S       float64 //diameter
	n       int     //numberofparticles
}

//Adds a number of particles at random locations in a cubic region
func (c Cube) Addparticles(n int) {
	msatcalled = false
	radiuscalled = false

	c.n += n
	for i := 0; i < n; i++ {
		px := c.x + (-0.5+rand.Float64())*c.S
		py := c.y + (-0.5+rand.Float64())*c.S
		pz := c.z + (-0.5+rand.Float64())*c.S
		Addsingleparticle(px, py, pz)
	}
}

//Defines the universe, its center and its diameter
func World(x, y, z, r float64) {
	worldcalled = true
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
