//TODO deze commentaar test to try demag
package main

import (
	. "github.com/JLeliaert/vinamax"
)

func main() {

	//comment TODO	
	World(0,0,0,2e-6)
	//Adds a single particle in the origin
	Addsingleparticle(-64.48e-9,0,0)
	Addsingleparticle(64.48e-9,0,0)

	//the particle has radius 16 nm
	Particle_radius(16e-9)

	//external field along the z direction of 10mT
	//B_ext can be an arbitrary function of time

	B_ext = func(t float64) (float64, float64, float64) { return 0.001,0.,0.0}

	FMM=false
	Demag=true

	//MSAT
	Msat (860e3)

	//timestep : 10fs
	Dt = 1e-12
	//initialise time at zero
	T = 0.
	//temperature=0
	Temp = 0.00
	//Gilbert damping constant=0.02
	Alpha = 0.1
	//anisotropy constant=0
	Ku1 = 0 

	//anisotropy axis along the y-direction
	Anisotropy_axis(0, 0, 1)

	//initialise the magnetisation along the x direction
	M_uniform(0,1,0)
	Tableadd("B_ext")

	//write output every 1e-13s 
	Output(1e-10)
//	SuggestTimestep()
	Save("geometry")

	//calculates the tree for the FMM demag
	Maketree()


	//run for 10 ns
	Run(100.e-9)
}
