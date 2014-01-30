package vinamax

import (
	"log"
)

var (
	B_ext            func(t float64) (float64, float64, float64)         // External applied field in T
	Dt               float64                                     = -1 // Timestep in s
	T                float64                                             // Time in s
	Alpha            float64                                     = -1  // Gilbert damping constant
	Temp             float64                                     =-1        // Temperature in K
	Ku1              float64                                     = 0     // Uniaxial anisotropy constant in J/m**3
	Thresholdbeta    float64                                     = 0.7   //The threshold value for the FMM
	universe         node                                                // The entire universe of the simulation
	FMM              bool                                        = false  // Calculate demag with FMM method
	Demag            bool                                        = true  // Calculate demag
	outdir           string                                              // the output directory
	Outputinterval   float64
	maxtauwitht      float64 = 0. //maximum torque during the simulations with temperature
	suggest_timestep bool    = false



	msatcalled bool = false
	radiuscalled bool = false
	uaniscalled bool = false
	worldcalled bool = false
	magnetisationcalled bool =false
	treecalled bool = false
	outputcalled bool = false

)

func testinput() {
	if Dt < 0 {
		log.Fatal("Dt cannot be smaller than 0, did you forget to initialise?")
	}
	if Alpha < 0 {
		log.Fatal("Alpha cannot be smaller than 0, did you forget to initialise?")
	}
	if Temp < 0 {
		log.Fatal("Temp cannot be smaller than 0, did you forget to initialise?")
	}
}

func syntaxrun(){
	if msatcalled==false {
		log.Fatal("You have to specify msat")
	}
if radiuscalled==false {
		log.Fatal("You have to specify the size of the particles")
	}
if (uaniscalled==false && Ku1!=0) {
		log.Fatal("You have to specify the anisotropy-axis")
	}
if worldcalled==false {
		log.Fatal("You have define a \"World\"")
	}
if magnetisationcalled==false {
		log.Fatal("You have specify the initial magnetisation")
	}
if (treecalled==false && FMM==true) {
		log.Fatal("You have run Maketree() as last command in front of Run() when using the FMM method")
	}
}
