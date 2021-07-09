package vinamax

import (
	"math"
	"math/rand"
)

var georng = rand.New(rand.NewSource(0))

//Set the randomseed for the geometry
func Setgeorandomseed(a int64) {
	//	randomseedcalled = true
	georng = rand.New(rand.NewSource(a))
}

type Cube struct {
	x, y, z float64 //position
	S       float64 //side
	n       int     //numberofparticles
}

//Sets origin of the cube
func (c Cube) Setorigin(x1, y1, z1 float64) {
	c.x = x1
	c.y = y1
	c.z = z1

}

//Adds a number of particles at random locations in a cubic region
func (c Cube) AddParticles(n int) {
	c.n += n
	for i := 0; i < n; i++ {
		status := false
		for status == false {
			px := c.x + (-0.5+georng.Float64())*c.S
			py := c.y + (-0.5+georng.Float64())*c.S
			pz := c.z + (-0.5+georng.Float64())*c.S
			status = addParticle(px, py, pz)
		}

	}
}

type Cuboid struct {
	x, y, z             float64 //position
	Sidex, Sidey, Sidez float64 //side
	n                   int     //numberofparticles
}

//Sets origin of the cuboid
func (c Cuboid) Setorigin(x1, y1, z1 float64) {
	c.x = x1
	c.y = y1
	c.z = z1

}

//Adds a number of particles at random locations in a cubic region
func (c Cuboid) AddParticles(n int) {

	c.n += n
	for i := 0; i < n; i++ {
		status := false
		for status == false {
			px := c.x + (-0.5+georng.Float64())*c.Sidex
			py := c.y + (-0.5+georng.Float64())*c.Sidey
			pz := c.z + (-0.5+georng.Float64())*c.Sidez
			status = addParticle(px, py, pz)
		}
	}
}

//returns true if the position of a particle overlaps with another particle
//now assumes spheres instead cubic particles
func overlap(x, y, z, r_h float64) bool {
	for _, p := range lijst {
		x2 := p.x
		y2 := p.y
		z2 := p.z
		r_h2 := p.rh
		if (r_h+r_h2)-math.Sqrt(math.Pow(x-x2, 2)+math.Pow(y-y2, 2)+math.Pow(z-z2, 2)) > 1e-14 {
			return true
		}
	}
	return false
}

func Clear() {
	lijst = nil
}
