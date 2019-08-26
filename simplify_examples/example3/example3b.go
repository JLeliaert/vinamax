package main

import (
	. "github.com/JLeliaert/vinamax"
	"math"
)

func main() {
	Msat.Set(400e3)
	Alpha.Set(1.)
	M.Set(0.1, 1, 0)
	U_anis.Set(0, 1, 0)
	Viscosity.Set(1e-3)
	Ku1.Set(1e4)
	Temp.Set(0.)
	B_ext = func(t float64) (float64, float64, float64) {
		return 0.003 * math.Sin(2*math.Pi*t*100e3), 0., 0.
	}

	BrownianRotation = true
	MagDynamics = true
	Demag = false
	T.Set(0.)
	SetSolver("dopri")
	Dt.Set(5e-9)
	Adaptivestep = true

	Rc.Set(10e-9)
	Rh.Set(10e-9)
	AddParticle(0., 0., 0.)

	Tableadd("B_ext")
	Tableadd("U_anis")
	Tableadd("Dt")
	Output(5e-7)

	Run(5e-3)
}
