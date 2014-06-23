package vinamax

import (
	"log"
	"math"
	"math/rand"
)

var georng = rand.New(rand.NewSource(0))

//Set the randomseed for the geometry
func Setgeorandomseed(a int64) {
	randomseedcalled = true
	georng = rand.New(rand.NewSource(a))
}

//Adds a single particle at specified coordinates, returns false if unsuccesfull
func addparticle(x, y, z float64) bool {
	if radiuscalled == false {
		log.Fatal("You have to specify the size of the particles before creating new particles")
	}
	r := radii[radiusindex]
	if overlap(x, y, z, r) == true {
		return false
	}

	if universe.inworld(vector{x, y, z}) {
		a := particle{x: x, y: y, z: z, rindex: radiusindex}
		universe.lijst = append(universe.lijst, &a)
		universe.number += 1
		msatcalled = false
	} else {
		log.Fatal("Trying to add particle at location (", x, ",", y, ",", z, ") which lies outside of universe")
	}
	return true
}

func Addsingleparticle(x, y, z float64) {
	if addparticle(x, y, z) == false {
		log.Fatal("Trying to add particle at overlapping locations")
	}
	radiusindex += 1
	if radiusindex == len(radii) {
		radiusindex = 0
	}
}

type Cube struct {
	x, y, z float64 //position
	S       float64 //side
	n       int     //numberofparticles
}

//Adds a number of particles at random locations in a cubic region
func (c Cube) Addparticles(n int) {
	c.n += n
	for i := 0; i < n; i++ {
		status := false
		for status == false {
			px := c.x + (-0.5+georng.Float64())*c.S
			py := c.y + (-0.5+georng.Float64())*c.S
			pz := c.z + (-0.5+georng.Float64())*c.S
			status = addparticle(px, py, pz)
		}
		radiusindex += 1
		if radiusindex == len(radii) {
			radiusindex = 0
		}
	}
}

type Cuboid struct {
	x, y, z             float64 //position
	Sidex, Sidey, Sidez float64 //side
	n                   int     //numberofparticles
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
			status = addparticle(px, py, pz)
		}
		radiusindex += 1
		if radiusindex == len(radii) {
			radiusindex = 0
		}
	}
}

//Defines the universe, its center and its diameter
func World(x, y, z, r float64) {
	worldcalled = true
	universe.origin = vector{x, y, z}
	universe.diameter = r
}

func (w node) inworld(r vector) bool {
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

//Sets the radius of all entries in radii to a consant value
func Particle_radius(x float64) {
	radiuscalled = true
	if x < 0 {
		log.Fatal("particles can't have a negative radius")
	}
	size := len(radii)
	for i := 0; i < size; i++ {
		radii[i] = x
	}
}

//set the radius of all entries in radii to a diameter taken from a lognormal distribution with specified mean and stdev
func Lognormal_diameter(mean, stdev float64) {
	m := mean * 1e9
	s := stdev * 1e9
	radiuscalled = true
	size := len(radii)
	for i := 0; i < size; i++ {
		for {
			x := rng.Float64() * 200 * m
			f_x := 1. / (math.Sqrt(2*math.Pi) * s * x) * math.Exp(-1./(2.*s*s)*sqr(math.Log(x/m)))
			if rng.Float64() < f_x {
				radii[i] = x * 1e-9 / 2.
				break
			}
		}
	}
}

//returns true if the position of a particle overlaps with another particle
//easiest implementation, assumes cubic particles instead of spheres
func overlap(x, y, z, r float64) bool {
	for i := range universe.lijst {
		x2 := universe.lijst[i].z
		y2 := universe.lijst[i].y
		z2 := universe.lijst[i].z
		r2 := radii[universe.lijst[i].rindex]
		if math.Abs(x-x2) < (r + r2) {
			if math.Abs(y-y2) < (r + r2) {
				if math.Abs(z-z2) < (r + r2) {
					return true
				}
			}
		}
	}
	return false
}
