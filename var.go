package vinamax

import (
	"log"
)

type scalarvariable struct {
	name        string
	unit        string
	description string
	called      bool
	value       float64
}

type vectorvariable struct {
	name        string
	unit        string
	description string
	called      bool
	value       vector
}

func (v *scalarvariable) Set(val float64) {
	v.value = val
}

func (v *vectorvariable) Set(valx, valy, valz float64) {
	v.value = vector{valx, valy, valz}
}

var (
	B_ext          func(t float64) (float64, float64, float64)          // External applied field in T
	B_ext_space    func(t, x, y, z float64) (float64, float64, float64) // External applied field in T
	Dt             = scalarvariable{"dt", "s", "timestep", false, 1e-15}
	MinDt          = scalarvariable{"minDt", "s", "minimum allowed timestep", false, 1e-20}
	MaxDt          = scalarvariable{"MaxDt", "s", "maximum allowed timestep", false, 1}
	T              = scalarvariable{"t", "s", "time", false, 0}
	Alpha          = scalarvariable{"alpha", "", "Gilbert damping constant", false, 0.01}
	Temp           = scalarvariable{"temp", "K", "Temperature", false, 0.}
	viscosity      = scalarvariable{"viscosity", "Pa s", "Viscosity", false, 0.001}
	solver         solvertype
	Errortolerance float64 = 1e-7
	Adaptivestep   bool    = false
	gammaoveralpha float64     //g/1+alfa^2
	Ku1            float64 = 0 // Uniaxial anisotropy constant in J/m**3
	Ku2            float64 = 0 // Uniaxial anisotropy constant in J/m**3
	global_u_anis  vector
	Kc1            float64 = 0    // Cubic anisotropy constant in J/m**3
	Universe       node           // The entire Universe of the simulation
	Demag          bool    = true // Calculate demag
	outdir         string         // The output directory
	outputinterval float64
	maxtauwitht    float64 = 0. //maximum torque during the simulations with temperature
	//	suggest_timestep bool    = false
	constradius           float64
	constradius_h         float64
	logradius_m           float64
	logradius_s           float64
	Tau0                  float64 = 1e-8
	msatcalled            bool    = false
	radiuscalled          bool    = false
	radius_hcalled        bool    = false
	constradiuscalled     bool    = false
	constradius_hcalled   bool    = false
	logradiuscalled       bool    = false
	uaniscalled           bool    = false
	c1called              bool    = false
	c2called              bool    = false
	worldcalled           bool    = false
	magnetisationcalled   bool    = false
	outputcalled          bool    = false
	randomseedcalled      bool    = false
	randomseedcalled_anis bool    = false
	tableaddcalled        bool    = false
	Jumpnoise             bool    = false
	Brown                 bool    = false
	BrownianRotation      bool    = false
	viscositycalled       bool    = false
	//noMagDyn	    bool = false //set this to true to skip calculations of magnetisation dynamics
	Condition_1 bool    = false
	Condition_2 bool    = false
	relax       bool    = false
	Test        bool    = false
	Counter     int     = 0
	Trigger     bool    = false
	Freq        float64 = 0.0
	Print1      bool    = false
	Print0      bool    = false
)

//initialised B_ext functions
func init() {
	B_ext = func(t float64) (float64, float64, float64) { return 0, 0, 0 }                // External applied field in T
	B_ext_space = func(t, x, y, z float64) (float64, float64, float64) { return 0, 0, 0 } // External applied field in T
}

//test the inputvalues for unnatural things
func testinput() {
	if Dt.value < 0 {
		log.Fatal("Dt cannot be smaller than 0, did you forget to initialise?")
	}
	if Alpha.value < 0 {
		log.Fatal("Alpha cannot be smaller than 0, did you forget to initialise?")
	}
	if Temp.value < 0 {
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
	if Temp.value != 0 && randomseedcalled == false {
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
	if Temp.value != 0 && Brown == false && Jumpnoise == false {
		log.Fatal("You have to specify which temperature you want to use: \"Jumpnoise\" or \"Brown\"")
	}
	if Brown {
		calculatetemp_prefactors(Universe.lijst)
	}
	if BrownianRotation {
		calculaterandomvprefacts(Universe.lijst)
	}
}
