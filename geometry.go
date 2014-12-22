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
	radius := getradius()
	//if overlap(x, y, z, radius) == true {
	//	return false
	//}

	if universe.inworld(vector{x, y, z}) {
		a := particle{x: x, y: y, z: z, r: radius}
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

func getradius() float64 {
	if constradiuscalled {
		return constradius
	}
	if logradiuscalled {
		for {
			x := rng.Float64() * 200 * logradius_m
			f_x := 1. / (math.Sqrt(2*math.Pi) * logradius_s * x) * math.Exp(-1./(2.*logradius_s*logradius_s)*sqr(math.Log(x/logradius_m)))
			if rng.Float64() < f_x {
				return x * 1e-9 / 2.
			}
		}
	}
	return 0.
}

//Sets the radius of all entries in radii to a constant value
func Particle_radius(x float64) {
	radiuscalled = true
	constradiuscalled = true

	if x < 0 {
		log.Fatal("particles can't have a negative radius")
	}
	constradius = x
}

//set the radius of all entries in radii to a diameter taken from a lognormal distribution with specified mean and stdev
func Lognormal_diameter(mean, stdev float64) {
	radiuscalled = true
	logradiuscalled = true
	logradius_m = mean * 1e9
	logradius_s = stdev * 1e9
}

//returns true if the position of a particle overlaps with another particle
//easiest implementation, assumes cubic particles instead of spheres
func overlap(x, y, z, r float64) bool {
	for i := range universe.lijst {
		x2 := universe.lijst[i].x
		y2 := universe.lijst[i].y
		z2 := universe.lijst[i].z
		r2 := universe.lijst[i].r
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
