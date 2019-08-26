//This second example is a test if the demagnetising field is implemented correctly
//To check this, we let 2 particles relax in the presence of an external field
//and check the output versus mumax. We also do the same simulation without
//calculating the demagnetising field to see if this problem is suited to
//check the implementation; i.e. to see that the demagnetising field
//makes a difference.

package main

import (
	. "github.com/JLeliaert/vinamax"
)

func main() {
	Rc.Set(16e-9)
	Rh.Set(16e-9)
	Msat.Set(860e3)
	Alpha.Set(0.1)
	Ku1.Set(0)
	M.Set(0, 1, 0)

	//Adds two particles
	AddParticle(-64.48e-9, 0, 0)
	AddParticle(64.48e-9, 0, 0)

	B_ext = func(t float64) (float64, float64, float64) { return 0.001, 0., 0.0 }
	Demag = true
	Dt.Set(1e-13)
	SetSolver("dopri")
	T.Set(0.)
	Temp.Set(0.0)
	Tableadd("B_ext")
	Output(1e-10)
	Save("geometry")
	Run(100.e-9)
}
