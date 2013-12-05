package vinamax

import (
	"math/rand"
)

//Adds a single particle
func Addsingleparticle(x, y, z float64) {
	Lijst.append(Particle{X: x, Y: y, Z: z})
}

//Deletes all particles
func Reset() {
	Lijst = nil
}

type Cube struct {
	x, y, z float64 //position
	S       float64 //diameter
	N       int     //numberofparticles
}

//Adds a lot of particles at random locations in a cubic region
func (c Cube) Addparticles(n int) {
	c.N+=n
	for i := 0; i < n; i++ {
		px := c.x + (-0.5+rand.Float64())*c.S
		py := c.y + (-0.5+rand.Float64())*c.S
		pz := c.z + (-0.5+rand.Float64())*c.S
		Addsingleparticle(px, py, pz)
	}
}
