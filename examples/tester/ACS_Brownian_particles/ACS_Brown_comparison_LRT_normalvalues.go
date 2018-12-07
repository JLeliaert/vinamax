//This examples checks if 250 particles with only Brownian relaxation behaves according to the LRT model for ACS
package main

import (
	. "github.com/JLeliaert/vinamax"
	"math"
)

func main() {

	//Defines the world at location 0,0,0 and with a side of 2 m	
	World(0,0,0,2)
	//Adds a cube to the center of the world with side 2 m
	test := Cube{S:2}

	//additionally calculate Brownian rotation
	BrownianRotation = true
	//this requires a randomn number for the anisotropy dynamics
	Setrandomseed_anis(2)
	//set viscosity environement test (e.g. water 1mPas)
	test.Setviscosity(1.e-3)

	//the particle have radius 12.5 nm and a hydrodynamic radius (core + coating together) of 15 nm 
	Particle_radius(12.5e-9)
	Particle_radius_h(15e-9)

	//Adds a single particles with radii defined above to the cube with viscosity 1 mPas
	test.Addparticles(1000)
	//fmt.Println("particles added")
	//Don't calculate the demagnetising field
	Demag=false

	//saturation magnetisation 400 000 A/m
	Msat(4e5)
	
	//apply an external field along the z direction of 1mT
	//B_ext can be an arbitrary function of time
	B_ext = func(t float64) (float64, float64, float64) { return 0, 0, 1e-3*math.Cos(2*math.Pi*t*97e3)}//field is 97 kHz which is abut 1/taub 

	Adaptivestep = true
	//timestep : 2ps
	Dt = 2e-12
	//initialise time at zero
	T = 0.
	//Temperature at 300 Kelvin
	Brown=true
	Temp = 300
	//Set a randomseed for the thermal field
	Setrandomseed(2)
	//The Gilbert damping constant =0.1
	Alpha = 0.5
	//anisotropy constant=10 000 J/m**3
	Ku1 = 10000 

	//anisotropy axis is random (so to have the Langevin thingy)
	Anisotropy_random()
	
	//initialise the magnetisation in a random way
	M_random()
	
	//curently anisodynamics only works with Euler
	Setsolver("dopri")
	
	//output u_anis of single particle (works only for single particle)
	//Tableadd("u_anis")
	//Add the external field to the outputtable
	Tableadd("B_ext")

	//write output every 0.1 µs 
	Output(0.1e-6)
	
	//fmt.Println("dt:   ", Dt)
	//run for 3*tauB 
	Run(3e-5)
}
