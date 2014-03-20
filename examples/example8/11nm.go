//**********************************************************************
// VINAMAX
//
//This code calculates the magnetisation dynamics of a macrospin model
//
//November, December 2013
//Jonathan Leliaert
//Many thanks to Arne Vansteenkiste
//*********************************************************************

package main

import (
	. "github.com/JLeliaert/vinamax"
)

func main() {

	//first define geometry
	
	World(0,0,0,1e-7)
	Addsingleparticle(0,0,0)

	Particle_radius(9.0e-9)

	Demag=false
	Dt = 3e-12
	T = 0.
	Temp = 300.
	Setrandomseed(2)
	Alpha = 0.01
	Ku1 = 10000
	Msat(400e3)

	Anisotropy_axis(0, 0., 1)
	//Anisotropy_random()
	M_uniform(0,0,1)

	Output(1e-5)
	Run(1.e-1)
}
