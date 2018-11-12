//In this example we look at the thermal switching behaviour of a single particle 
 
package main

import (
	. "github.com/JLeliaert/vinamax"
)

func main() {

	//Define world	
	World(0,0,0,1e-7)

	//The radius of the particle is 9 nm
	Particle_radius(9.0e-9)

	//Add a single particle in the origin
	Addsingleparticle(0,0,0)

	//Don't calculate the demagnetising field
	Demag=false

	//Timestep = 3 picoseconds
	Dt = 3e-12
	//Initialize time at zero
	T = 0.
	Brown=true
	//Temperature = 0
	Temp = 300.
	//Set a randomseed for the thermal field
	Setrandomseed(2)
	//The Gilbert damping constant =0.01
	Alpha = 0.01
	//The anisotropy constant is 10000 J/m**3
	Ku1 = 10000
	//The saturation magnetisation is 400000 A/m
	Msat(400e3)

	//The particle has its anisotropy axis along the z-axis
	Anisotropy_axis(0, 0., 1)
	//Initialise the magnetisation along the z-axis
	M_uniform(0,0,1)

	//Ouput the magnetization every 1e-5 seconds
	Output(1e-5)

	//Run the simulation for 1 second
	Run(1.e-0)
}
