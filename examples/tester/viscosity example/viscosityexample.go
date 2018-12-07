//This example adds a particle in a certain viscosity and additionally calculates Brownian roation

package main

import (
	. "github.com/JLeliaert/vinamax"
	//"fmt"
)

func main() {

	//Defines the world at location 0,0,0 and with a side of 1e-8 m	
	World(0,0,0,1e-8)

	//the particle has radius 10 nm
	//Particle_radius(10.e-9)
	Lognormal_diameter(20e-9, 1e-9)
	Particle_radius_h(15.e-9)
	
	//additionally calculate Brownian rotation
	BrownianRotation = true
	//set viscosity environement particle (e.g. water 1mPas)
	Setviscosity(1e-3)

	//Adds a single particle in the origin
	Addsingleparticle(0,0,0)
	
	//apply an external field along the z direction of 10mT
	//B_ext can be an arbitrary function of time
	B_ext = func(t float64) (float64, float64, float64) { return 0, 0, 0.1 }
	//Don't calculate the demagnetizing field
	Demag=false
	//Don't use the dipole approximation method
	FMM=false

	//sets the saturation magnetisaton of the particles at 860e3 A/m
	Msat(860e3)	

	//timestep : 1fs
	Dt = 1e-15
	//initialise time at zero
	T = 0.
	//temperature=0
	Temp = 0.
	//Gilbert damping constant=0.02
	Alpha = 0.02
	//anisotropy constant=0
	Ku1 = 0 

	//anisotropy axis along the y-direction
	Anisotropy_axis(0, 1, 0)

	//initialise the magnetisation along the x direction
	M_uniform(1,0,0)

	//write output every 5e-12s 
	Output(5e-12)

	//fmt.Printf("particle hydrodynamic diamter is %#v \n",universe.lijst[0].r_h)


	//run for 5 ns
	Run(5.e-9)
}
