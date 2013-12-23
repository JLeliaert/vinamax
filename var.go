package vinamax

import(
	"log"
)

var (
	B_ext         Vector             // External applied field in T
	Dt            float64   =1e-15         // Timestep in s, default is 1e-15
	Lijst         Particles          // List containing all the Particles
	T             float64            // Time in s
	Alpha         float64  =0.02          // Gilbert damping constant, default=0.02
	Temp          float64            // Temperature in K
	Ku1           float64    =0        // Uniaxial anisotropy constant in J/m**3, default is 0
	Msat          float64   = 860000 // Saturation magnetisation in A/m, default = 860e3
	Thresholdbeta float64   =0.7            //The threshold value for the FMM
	Universe      node               // The entire universe of the simulation
	FMM	      bool 	=true	 // Calculate demag with FMM method
	outdir	      string		 // the output directory
)

func testinput(){
if Dt <0{
log.Fatal("Dt cannot be smaller than 0")
}
if Alpha <0{
log.Fatal("Alpha cannot be smaller than 0")
}
if Temp <0{
log.Fatal("Temp cannot be smaller than 0")
}
if Msat <0{
log.Fatal("Msat cannot be smaller than 0")
}
}
