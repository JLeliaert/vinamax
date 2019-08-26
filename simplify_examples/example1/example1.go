//This example adds a single particle without anisotropy,
//gives it a well defined magnetisation direction,
//applies an external field and sees how the magnetisation
//rotates around and damps towards this field.
//The gyration frequency should be 28 GHz/T

package main

import (
	. "github.com/JLeliaert/vinamax"
)

func main() {
	Msat.Set(860e3)
	Alpha.Set(0.02)
	M.Set(1, 0, 0)
	Rc.Set(10.e-9)
	Rh.Set(10.e-9)

	AddParticle(0., 0., 0.)

	B_ext = func(t float64) (float64, float64, float64) { return 0, 0, 0.1 }
	Demag = false
	Dt.Set(1e-15)
	T.Set(0.)
	Temp.Set(0.)
	Tableadd("Dt")
	Output(5e-12)
	SetSolver("dopri")
	Adaptivestep = true
	Run(5.e-9)
}
