//+build ignore

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
	. "."
	"flag"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	//first define geometry

	World(0, 0, 0, 1e-6)
	test := Cube{S: 1e-8}
	test.Addparticles(500)
	Maketree()

	Particle_radius(7.0e-9)
	//	Lognormal_radius(1.6e-9)

	//An example of an external field (in Tesla, t in seconds)
	B_ext = func(t float64) (float64, float64, float64) { return 0.01, 0.001*math.Sin(t), t / 2 }

	Demag = false
	FMM = false
	Dt = 1e-11
	T = 0.
	Temp = 300.
	SetRandomSeed(3)
	Alpha = 0.01
	Ku1 = 10000

	//Anisotropy_axis(0, 0, 1)
	Anisotropy_random()
	M_uniform(0, 0, 1)

	//Tableadd_B_eff_at_location(0,0.0,0)
	Tableadd("B_ext")
	Output(5e-10)
	Save("geometry")

	Run(1e-8)
}
