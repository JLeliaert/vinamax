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
	gammaoveralpha float64                                                      //g/1+alfa^2
	Temp           float64                                              = -1    // Temperature in K
	Ku1            float64                                              = 0     // Uniaxial anisotropy constant in J/m**3
	Ku2            float64                                              = 0     // Uniaxial anisotropy constant in J/m**3
	global_u_anis  vector
	Kc1            float64 = 0 // Cubic anisotropy constant in J/m**3
	Errortolerance float64 = 1e-5
	Thresholdbeta  float64 = 0.3 // The threshold value for the FMM
	demagtime      float64
	Universe       node           // The entire Universe of the simulation
	FMM            bool   = false // Calculate demag with FMM method
	Demag          bool   = true  // Calculate demag
	demagevery     bool   = false // Calculate demag only after certain interval
	Adaptivestep   bool   = false
	outdir         string           // The output directory
	solver         string = "dopri" // The solver used
	outputinterval float64
	maxtauwitht    float64 = 0. //maximum torque during the simulations with temperature
	//	suggest_timestep bool    = false
	order         int = 5 //the order of the solver
	constradius   float64
	constradius_h float64
	logradius_m   float64
	logradius_s   float64
	Tau0          float64 = 1e-8
	viscosity     float64 //set viscosity value

	msatcalled        bool = false
	radiuscalled      bool = false
	radius_hcalled    bool = false
	constradiuscalled bool = false
	//constradius_hcalled bool = false
	logradiuscalled       bool = false
	uaniscalled           bool = false
	c1called              bool = false
	c2called              bool = false
	worldcalled           bool = false
	magnetisationcalled   bool = false
	treecalled            bool = false
	outputcalled          bool = false
	randomseedcalled      bool = false
	randomseedcalled_anis bool = false
	tableaddcalled        bool = false
	Jumpnoise             bool = false
	Brown                 bool = false
	BrownianRotation      bool = false
	viscositycalled       bool = false
	//noMagDyn	    bool = false //set this to true to skip calculations of magnetisation dynamics
	Condition_1    bool    = false
	Condition_2    bool    = false
	relax          bool    = false
	Test           bool    = false
	Counter        int     = 0
	Max_u_anis_x   float64 = 0.
	Max_u_anis_z   float64 = 0.
	Min_u_anis_x   float64 = 1.
	Min_u_anis_z   float64 = 1.
	Max_u_anis_x_2 float64 = 0.
	Max_u_anis_z_2 float64 = 0.
	Min_u_anis_x_2 float64 = 1.
	Min_u_anis_z_2 float64 = 1.
	Trigger        bool    = false
	Freq           float64 = 0.0
	Print1         bool    = false
	Print0         bool    = false
	Nsteps         int     = 0
)

//initialised B_ext functions
func init() {
	B_ext = func(t float64) (float64, float64, float64) { return 0, 0, 0 }                // External applied field in T
	B_ext_space = func(t, x, y, z float64) (float64, float64, float64) { return 0, 0, 0 } // External applied field in T
}

//demag every interval
func Demagevery(t float64) {
	demagevery = true
	demagtime = t
}

//test the inputvalues for unnatural things
func testinput() {
	if Demag == true && demagevery == true {
		log.Fatal("You cannot call both Demagevery and Demag, pick one")
	}
	if Dt < 0 {
		log.Fatal("Dt cannot be smaller than 0, did you forget to initialise?")
	}
	if Alpha < 0 {
		log.Fatal("Alpha cannot be smaller than 0, did you forget to initialise?")
	}
	if Temp < 0 {
		log.Fatal("Temp cannot be smaller than 0, did you forget to initialise?")
	}
	if Universe.number == 0 {
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
		log.Fatal("You have to specify the uniaxial anisotropy-axis")
	}
	if (c1called == false || c2called == false) && Kc1 != 0 {
		log.Fatal("You have to specify the cubic anisotropy-axes")
	}
	if worldcalled == false {
		log.Fatal("You have define a \"World\"")
	}
	if magnetisationcalled == false {
		log.Fatal("You have specify the initial magnetisation")
	}
	if BrownianRotation == true && viscositycalled == false {
		log.Fatal("You have to specify the viscosity of the particles' environment when taking into account Brownian relaxation")
	}
	if treecalled == false && FMM == true {
		log.Fatal("You have to run Maketree() as last command in front of Run() when using the FMM method")
	}
	if Temp != 0 && randomseedcalled == false {
		log.Fatal("You have to run Setrandomseed() when using nonzero temperatures")
	}
	if BrownianRotation == true && randomseedcalled_anis == false {
		log.Fatal("You have to run Setrandomseed_anis() when taking into account Brownian rotation (i.e. anisotropy dynamics) of the particle")
	}
	if tableaddcalled == true && outputcalled == false {
		log.Fatal("You have to run Output(interval) when calling tableadd")
	}
	//	if Brown == true && Adaptivestep == true {
	//		log.Fatal("Brown Temperature can only be used with fixed timestep")
	//	}
	if BrownianRotation == false && Condition_1 == true {
		log.Fatal("You have to calculate anisotropy dynamics for condition 1 to be true")
	}
	if BrownianRotation == false && Condition_2 == true {
		log.Fatal("You have to calculate anisotropy dynamics for condition 2 to be true")
	}
	if Condition_1 && Condition_2 == false {
		log.Fatal("This situation is not yet implemented")
	}
	//if BrownianRotation == false && noMagDyn == true {
	//	log.Fatal("You have to calculate something, e.g. anisotropy dynamics or magnetisation dynamics or both")
	//}
	//if Brown == true && Adaptivestep == true { see paper Leliaert et. al 2017
	//	log.Fatal("Brown Temperature can only be used with fixed timestep")
	//}
	if Jumpnoise == true {
		resetswitchtimes(Universe.lijst)
	}
	if Temp != 0 && Brown == false && Jumpnoise == false {
		log.Fatal("You have to specify which temperature you want to use: \"Jumpnoise\" or \"Brown\"")
	}
	if Brown {
		calculatetemp_prefactors(Universe.lijst)
	}
	if BrownianRotation {
		calculaterandomvprefacts(Universe.lijst)
	}
}
