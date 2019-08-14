package vinamax

import (
	"math"
	"testing"
)

func TestDamping(test *testing.T) {

	Rc.Set(10e-9)
	Msat.Set(400e3)
	M.Set(1, 0, 0)
	Alpha.Set(1.0)
	Dt.Set(1e-15)
	T.Set(0.)
	Temp.Set(0.)

	AddParticle(0, 0, 0)

	B_ext = func(t float64) (float64, float64, float64) { return 0., 0., 1. }

	SetSolver("dopri")
	//Relax()
	Run(1.e-9)

	if math.Abs(lijst[0].m[0]+lijst[0].m[1]) > 1e-12 || lijst[0].m[2] != 1 {
		test.Error("relax did not work as expected")
	}
}
