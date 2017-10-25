//This example adds particles in a heterogeneous envinroment (e.g. cubes with different viscosities) and additionally calculates Brownian roation

package main

import (
	. "github.com/JLeliaert/vinamax"
	//"fmt"
)

func main() {

	//Defines the world at location 0,0,0 and with a side of 2e-6 m	
	World(0,0,0,2e-6)

	//Adds a cube to the center of the world with side 1e-8 m
	test := Cube{S:1e-8}
	
	//additionally calculate Brownian rotation
	BrownianRotation = true
	//set viscosity environement test (e.g. water 1mPas)
	test.Setviscosity(1e-3)

	//the particle have radius 16 nm
	Particle_radius(16e-9)

	//Adds 256 particles to the cube with viscosity 1 mPas
	test.Addparticles(256)

	//Adds a cube with viscosity 10 mPas next to region with 1 mPas
	test_2 := Cube{S:1e-8}
	test_2.Setorigin(5e-9,0.0,0.0)
	
	//set viscosity environement test (e.g. 10 mPas)
	test_2.Setviscosity(10e-3)

	//the particle have radius 20 nm
	Particle_radius(20e-9)

	//Adds 200 particles to the cube with viscosity 10 mPas
	test_2.Addparticles(200)
	
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
