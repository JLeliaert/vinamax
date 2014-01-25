//This example adds a single particle without anisotropy
//gives it a well defined magnetisation direction
//applies an external field and sees how the magnetisation
//rotates around and damps towards this field
//The gyration frequency should be 28 GHz/T


package main

import (
	. "github.com/JLeliaert/vinamax"
)

func main() {

	//comment TODO	
	World(0,0,0,1e-8)
	//Adds a single particle in the origin
	Addsingleparticle(0,0,0)

	//calculates the tree for the FMM demag
	Maketree()

	//the particle has radius 10 nm
	Particle_radius(10.e-9)

	//external field along the z direction of 10mT
	//B_ext can be an arbitrary function of time
	B_ext = func(t float64) (float64, float64, float64) { return 0, 0, 0.1 }
	//Don't use the fast multipole method
	FMM=false
	//Don't calculate demag at all
	Demag=false

	//timestep : 1fs
	Dt = 1e-15
	//initialise time at zero
	T = 0.
	//temperature=0
	Temp = 0.
	//damping constant=0.01
	Alpha = 0.01
	//anisotropy constant=0
	Ku1 = 0 

	//anisotropy axis along the y-direction
	Anisotropy_axis(0, 1, 0)

	//initialise the magnetisation along the x direction
	M_uniform(1,0,0)

	//write output every 1e-13s 
	Output(5e-12)

	//run for 1 ns
	Run(5.e-9)
}
