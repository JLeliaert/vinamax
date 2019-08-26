package vinamax

import (
	"log"
	"math"
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
	v.called = true
	v.value = val
}

func (v *scalarvariable) Get() float64 {
	return v.value
}

func (v *vectorvariable) Set(valx, valy, valz float64) {
	v.called = true
	v.value = norm(vector{valx, valy, valz})
}

func (v *vectorvariable) Get() vector {
	return v.value
}

var (
	// Global variables
	B_ext       func(t float64) (float64, float64, float64)          // External applied field in T
	B_ext_space func(t, x, y, z float64) (float64, float64, float64) // External applied field in T

	Dt                    = scalarvariable{"dt", "s", "timestep", false, 1e-15}
	MinDt                 = scalarvariable{"minDt", "s", "minimum allowed timestep", false, 1e-20}
	MaxDt                 = scalarvariable{"MaxDt", "s", "maximum allowed timestep", false, 1}
	T                     = scalarvariable{"t", "s", "time", false, 0}
	Temp                  = scalarvariable{"temp", "K", "Temperature", false, 0.}
	Viscosity             = scalarvariable{"viscosity", "Pa s", "Viscosity", false, 0.001}
	Demag          bool   = true // Calculate demag
	outdir         string        // The output directory
	outputinterval float64
	solver         solvertype
	Errortolerance float64     = 1e-5
	Adaptivestep   bool        = false
	lijst          []*particle //lijst met alle particles

	//"default values for particle-specific variables"
	Alpha  = scalarvariable{"alpha", "", "Gilbert damping constant", false, 0.01}
	Ku1    = scalarvariable{"Ku1", "J/m**3", "uniaxial anistropy constant", false, 0.}
	U_anis = vectorvariable{"U_ani", "", "uniaxial anistropy axis", false, vector{0., 0., 1.}}
	M      = vectorvariable{"M", "", "normalized magnetization direction", false, vector{0., 0., 1.}}
	Rc     = scalarvariable{"Rc", "m", "core radius", false, 1.e-8}
	Rh     = scalarvariable{"Rh", "m", "hydrodynamic radius", false, 1.e-8}
	Msat   = scalarvariable{"Msat", "A/m", "Saturation magnetizatoin", false, 400e3}

	magTorque float64 = 0. //maximum torque during the simulations with temperature
	rotTorque float64 = 0. //maximum torque during the simulations with temperature
	totalErr  float64 = 0. //maximum error during the simulations with temperature

	//	suggest_timestep bool    = false
	constradius   float64
	constradius_h float64
	logradius_m   float64
	logradius_s   float64
	Tau0          float64 = 1e-8

	worldcalled           bool = false
	outputcalled          bool = false
	randomseedcalled      bool = false
	randomseedcalled_anis bool = false
	tableaddcalled        bool = false
	BrownianRotation      bool = false
	viscositycalled       bool = false
	//noMagDyn	    bool = false //set this to true to skip calculations of magnetisation dynamics
	relax   bool    = false
	Test    bool    = false
	Counter int     = 0
	Trigger bool    = false
	Freq    float64 = 0.0
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
	if Temp.value < 0 {
		log.Fatal("Temp cannot be smaller than 0, did you forget to initialise?")
	}
}

//checks the inputfiles for functions that should have been called but weren't
func syntaxrun() {
	if tableaddcalled == true && outputcalled == false {
		log.Fatal("You have to run Output(interval) when calling tableadd")
	}
}

func Volume(radius float64) float64 {
	return 4. / 3. * math.Pi * cube(radius)
}
