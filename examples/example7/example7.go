//In this example 10 particles are initialised with a random magnetisation. They have a negative anisotropy constant along the z-axis, which means that the magnetisation prefers to lie in the xy-plane. This relaxation is shown at a zero and nonzero temperature.

package main


import (
	. "github.com/JLeliaert/vinamax"
)

func main() {

	//Defines the world at location 0,0,0 and with a side of 1e-6 m	
	World(0,0,0,1e-6)

	//Adds a cube to the word with side 1e-6 m
	test := Cube{S:1e-6}

	//the particle have radius 20 nm
	Particle_radius(20e-9)

	//Adds 10 particles to the cube
	test.Addparticles(10)

	//Don't calculate the demagnetizing field
	Demag=false

	//saturation magnetisation in A/m
	Msat (860e3)

	//timestep : 0.1ps
	Dt = 1e-13
	//initialise time at zero
	T = 0.
	//temperature=300 K
	Brown=true
	Temp = 300.0
	Setrandomseed(3)
	//Gilbert damping constant=0.02
	Alpha = 0.02
	//anisotropy constant=-10 000 J/m**3
	Ku1 = -10000 

	//anisotropy axis along the z-direction
	Anisotropy_axis(0,0,1)

	//initialise the magnetisation along the y direction
	M_random()

	//write output every 4e-12s 
	Output(4e-12)
	
	//run for 100 ns
	Run(100.e-9)
}
