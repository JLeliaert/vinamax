package vinamax

import (
	"log"
)

var (
	B_ext          func(t float64) (float64, float64, float64)         // External applied field in T
	Dt             float64                                     = 1e-15 // Timestep in s, default is 1e-15
	T              float64                                             // Time in s
	Alpha          float64                                     = 0.02  // Gilbert damping constant, default=0.02
	Temp           float64                                             // Temperature in K
	Ku1            float64                                     = 0     // Uniaxial anisotropy constant in J/m**3, default is 0
	Thresholdbeta  float64                                     = 0.7   //The threshold value for the FMM
	universe       node                                                // The entire universe of the simulation
	FMM            bool                                        = true  // Calculate demag with FMM method
	Demag          bool                                        = true  // Calculate demag
	outdir         string                                              // the output directory
	Outputinterval float64
)

func testinput() {
	if Dt < 0 {
		log.Fatal("Dt cannot be smaller than 0")
	}
	if Alpha < 0 {
		log.Fatal("Alpha cannot be smaller than 0")
	}
	if Temp < 0 {
		log.Fatal("Temp cannot be smaller than 0")
	}
}
