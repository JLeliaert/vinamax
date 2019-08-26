package vinamax

import (
	"math"
	"testing"
)

func TestRelax(test *testing.T) {
	Clear()
	Rc.Set(10e-9)
	Msat.Set(400e3)
	M.Set(1, 0, 0)
	Alpha.Set(1.0)
	Ku1.Set(2.e4 * math.Pi)
	Dt.Set(1e-15)
	T.Set(0.)
	Temp.Set(0.)

	AddParticle(0, 0, 0)

	SetSolver("dopri")
	Relax()
	B_ext = func(t float64) (float64, float64, float64) { return 0., 0., 1. }
	Relax()
	if math.Abs(lijst[0].m[0]+lijst[0].m[1]) > 1e-5 || 1-lijst[0].m[2] >= 1e-5 {
		test.Error("relax did not work as expected")
	}
}
