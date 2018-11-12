//starting from 1400 particles dipole approximation becomes faster than bruteforce demag
package main

import (
	. "github.com/JLeliaert/vinamax"
)

func main() {

	//Defines the world at location 0,0,0 and with a side of 2e-6 m	
	World(0,0,0,2e-6)

	//Adds a cube to the word with side 2e-6 m
	test := Cube{S:2e-6}

	//the particles have radius 16 nm
	Particle_radius(16e-9)


	//Adds 256 particles to the cube
	test.Addparticles(256)

	//external field along the x direction of 1mT
	//B_ext can be an arbitrary function of time

	B_ext = func(t float64) (float64, float64, float64) { return 0.001,0.,0.0}

	//calculate demag, without using the fmm method
	FMM=false
	Demag=true

	//Sets the saturation magnetisation of the particles
	Msat (860e3)

	//timestep : 0.9ps
	Dt = 5e-12
	//initialise time at zero
	T = 0.
	//temperature=0
	Temp = 0.0
	//Gilbert damping constant=0.1
	Alpha = 0.1
	//anisotropy constant=0
	Ku1 = 0 

	//anisotropy axis along the z-direction
	Anisotropy_axis(0, 0, 1)

	//initialise the magnetisation along the y direction
	M_uniform(0,1,0)
	//Adds the external field to the outputtable
	Tableadd("B_ext")

	//write output every 1.5e-10s 
	Output(1.5e-10)

	//saves the geometry of the simulation
	Save("geometry")

	//calculates the tree 
	Maketree()

	//run for 100 ns
	Run(1.e-7)
}
