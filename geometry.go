package vinamax

import (
	"fmt"
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
func (c Cube) Addparticles(n int) {
	c.n += n
	for i := 0; i < n; i++ {
		fmt.Println(i, "th particle to be added")
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
func (c Cuboid) Addparticles(n int) {

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

//set the radius of all entries in radii to a diameter taken from a lognormal distribution with specified mean and stdev
func Lognormal_diameter(mean, stdev float64) {
	logradius_m = mean * 1e9
	logradius_s = stdev * 1e9
}

//returns true if the position of a particle overlaps with another particle
//easiest implementation, assumes cubic particles instead of spheres
func overlap(x, y, z, r_h float64) bool {
	for _, p := range lijst {
		x2 := p.x
		y2 := p.y
		z2 := p.z
		r_h2 := p.rh
		if math.Abs(x-x2) < (r_h + r_h2) {
			if math.Abs(y-y2) < (r_h + r_h2) {
				if math.Abs(z-z2) < (r_h + r_h2) {
					return true
				}
			}
		}
	}
	return false
}
