package vinamax

import (
	"log"
)

var (
	//These variables can be set in the input files
	B_ext          func(t float64) (float64, float64, float64)                  // External applied field in T
	B_ext_space    func(t, x, y, z float64) (float64, float64, float64)         // External applied field in T
	Dt             float64                                              = -1    // Timestep in s
	Mindt          float64                                              = 1e-20 //smallest allowed timestep
	Maxdt          float64                                              = 1     //largest allowed timestep
	T              float64                                                      // Time in s
	Alpha          float64                                              = -1    // Gilbert damping constant
	Temp           float64                                              = -1    // Temperature in K
	Ku1            float64                                              = 0     // Uniaxial anisotropy constant in J/m**3
	Errortolerance float64                                              = 1e-5
	Thresholdbeta  float64                                              = 0.3   // The threshold value for the FMM
	universe       node                                                         // The entire universe of the simulation
	FMM            bool                                                 = false // Calculate demag with FMM method
	Demag          bool                                                 = true  // Calculate demag
	Adaptivestep   bool                                                 = false
	outdir         string                                                       // The output directory
	solver         string                                               = "rk4" // The solver used
	outputinterval float64
	maxtauwitht    float64 = 0. //maximum torque during the simulations with temperature
	//	suggest_timestep bool    = false
	order int = 4 //the order of the solver

	msatcalled          bool = false
	radiuscalled        bool = false
	uaniscalled         bool = false
	worldcalled         bool = false
	magnetisationcalled bool = false
	treecalled          bool = false
	outputcalled        bool = false
	randomseedcalled    bool = false
	tableaddcalled      bool = false
)

//initialised B_ext functions
func init() {
	B_ext = func(t float64) (float64, float64, float64) { return 0, 0, 0 }                // External applied field in T
	B_ext_space = func(t, x, y, z float64) (float64, float64, float64) { return 0, 0, 0 } // External applied field in T
}

//test the inputvalues for unnatural things
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
	if universe.number == 0 {
		log.Fatal("There are no particles in the geometry")
	}
}

//checks the inputfiles for functions that should have been called but weren't
func syntaxrun() {
	if msatcalled == false {
		log.Fatal("You have to specify msat")
	}
	if radiuscalled == false {
		log.Fatal("You have to specify the size of the particles")
	}
	if uaniscalled == false && Ku1 != 0 {
		log.Fatal("You have to specify the anisotropy-axis")
	}
	if worldcalled == false {
		log.Fatal("You have define a \"World\"")
	}
	if magnetisationcalled == false {
		log.Fatal("You have specify the initial magnetisation")
	}
	if treecalled == false && FMM == true {
		log.Fatal("You have to run Maketree() as last command in front of Run() when using the FMM method")
	}
	if Temp != 0 && randomseedcalled == false {
		log.Fatal("You have to run Setrandomseed() when using nonzero temperatures")
	}
	if tableaddcalled == true && outputcalled == false {
		log.Fatal("You have to run Output(interval) when calling tableadd")
	}
	//todo ADD bool BROWN
	//	if brown == true && Adaptivestep == true {
	//		log.Fatal("Brown Temperature can only be used with fixed timestep")
	//	}

}
