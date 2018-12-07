package vinamax

import (
	"fmt"
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
	

	var radius_h float64	
	if radius_hcalled == false { //when no hydrodynamic radius is specified, consider it equal to core radius
		radius_h = radius 
	} 
	if logradiuscalled { //when distribution of core sizes use a fixed coating size
		radius_h = getradius_h() + radius
	}
	if constradiuscalled {
		radius_h = getradius_h() 
	}
	if overlap(x, y, z, radius_h) == true {
		return false
	}
	
	if BrownianRotation == true && viscositycalled == false {
		log.Fatal("You have to specify the viscosity of the particles' surroundings before adding new particles")
	}	

	if universe.inworld(vector{x, y, z}) {
		a := particle{x: x, y: y, z: z, r: radius, r_h: radius_h}
		if BrownianRotation {
			a.eta = viscosity
		}
		universe.lijst = append(universe.lijst, &a)
		universe.number += 1
		msatcalled = false
	} else {
		log.Fatal("Trying to add particle at location (", x, ",", y, ",", z, ") which lies outside of universe")
	}
	//fmt.Printf("particle core diamter is %#v \n",radius)
	//fmt.Printf("particle hydrodynamic diamter is %#v \n",radius_h)
	//fmt.Printf("Viscosity of particle is %#v \n", viscosity)
	//fmt.Printf("Viscosity called %#v \n", viscositycalled)
	//fmt.Printf("Brownian rotation? %#v \n", BrownianRotation)

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

//Sets origin of the cube 
func (c Cube) Setorigin(x1,y1,z1 float64) {
	c.x = x1
	c.y = y1
	c.z = z1
	
}


//Sets viscosity of particles in the cube (e.g. different viscosity regions possible)
func (c Cube) Setviscosity(visc float64) {
	if BrownianRotation == false {
		log.Fatal("You need to set BrownianRotation in order to calculate the impact of the particles' viscosity")
	}
	viscositycalled = true
	viscosity = visc
	
}

//Adds a number of particles at random locations in a cubic region
func (c Cube) Addparticles(n int) {
	c.n += n
	for i := 0; i < n; i++ {
	fmt.Println(i,"th particle to be added")
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

//Sets origin of the cuboid 
func (c Cuboid) Setorigin(x1,y1,z1 float64) {
	c.x = x1
	c.y = y1
	c.z = z1
	
}


//Sets viscosity of particles in the cuboid (e.g. different viscosity regions possible)
func (c Cuboid) Setviscosity(visc float64) {
	if BrownianRotation == false {
		log.Fatal("You need to set BrownianRotation in order to calculate the impact of the particles' viscosity")
	}
	viscositycalled = true
	viscosity = visc
	
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

//Sets viscosity of particles to be added directly to the universe
func Setviscosity(visc float64) {
	//check if Brownian rotation calculations are on
	if BrownianRotation == false {
		log.Fatal("You need to set BrownianRotation in order to calculate the impact of the particles' viscosity")
	}
	viscositycalled = true
	viscosity = visc
	
}

func getradius() float64 {
	if constradiuscalled {
		return constradius
	}
	if logradiuscalled {
		for {
			x := rng.Float64() * 2. * logradius_m
			f_x := 1. / (math.Sqrt(2*math.Pi) * logradius_s * x) * math.Exp(-1./(2.*logradius_s*logradius_s)*sqr(math.Log(x/logradius_m)))
			if rng.Float64() < f_x {
				return x * 1e-9 / 2.
			}
		}
	}
	return 0.
}

func getradius_h() float64 {
	return constradius_h 		
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

//Sets the hydrodynamic radius of all entries in radii to a constant value or constant coating in case core distribution
func Particle_radius_h(x float64) {
	radius_hcalled = true
	
	if radiuscalled == false {
		log.Fatal("You have to specify the core size of the particles before setting the hydrodynamic radius")
	}

	if constradiuscalled {
		if x < constradius {
			log.Fatal("particles can't have a hydrodynamic radius (core and coating together) smaller than the core radius")
		}	
	}
	constradius_h = x
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
func overlap(x, y, z, r_h float64) bool {
	for i := range universe.lijst {
		x2 := universe.lijst[i].x
		y2 := universe.lijst[i].y
		z2 := universe.lijst[i].z
		r_h2 := universe.lijst[i].r_h
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
