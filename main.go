//**********************************************************************
// VINAMAX
//
//This code calculates the magnetisation dynamics of a macrospin model
//
//November 2013
//Jonathan Leliaert
//Many thanks to Arne Vansteenkiste
//*********************************************************************

//TODO:
//	syntax voor geometrie, mogelijkheid tot kiezen tussen bolletjes en kubusjes
//			kiezen om een kubus te definieren waarin de particles
//			op een rooster zitten of random verdeeld, concentratie als parameter 
//	code opkuisen
//	output: header aanpassen bij tableadd
//	staat u_anis random los van de temperatuur random?
//	functie demag_off
//	Check dynamics against mumax
//	automatische tests
//	deeltjes ook een distributie van sizes meegeven ipv vaste size
//	demag slimmer dan brute force maken
//	optimaliseren
//	parallelliseren

package main

import (
)

var (
	B_ext [3]float64          // External applied field in T
	dt    float64             // Timestep in s
	lijst particles           // List containing all the particles
	t     float64             // Time in s
	alpha float64             // Gilbert damping constant
	temp  float64             // Temperature in K
	Ku1   float64             // Uniaxial anisotropy constant in J/m**3
	Msat  float64    = 860000 // Saturation magnetisation in A/m
)


func main() {

	//first define geometry
	lijst.append(particle{x: -2e-9, m: [3]float64{1., 0., 0}})
	lijst.append(particle{x: 2e-9, m: [3]float64{0, 1., 0}})


	B_ext = [3]float64{0, 0, 0.01}
	dt = 1e-15
	t = 0.
	temp = 0.
	alpha = 0.01
	Ku1 = 0 //10 000

	//anisotropy_axis(0, 1, 0)
	anisotropy_random()

	Output(1e-12)
	tableadd_B_eff_at_location(0,0.01,0)

	run(10.e-9)

}
