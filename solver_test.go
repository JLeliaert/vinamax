package vinamax

import (
	"math"
	"testing"
)

func TestSolver(t *testing.T) {
	Clear()
	Alpha.Set(1.)
	Ku1.Set(0.)
	M.Set(1., 0., 0.)
	Rc.Set(25.e-9)
	Rh.Set(35.e-9)
	Msat.Set(100e3)
	Temp.Set(300)
	Viscosity.Set(1.e-3)
	AddParticle(0, 0, 0)

	setThermPrefac()
	if math.Abs(lijst[0].thermPrefac-8.481414372111154e-08) > 1e-12 {
		t.Error("thermal prefactors not set correctly")
	}
	if math.Abs(lijst[0].thermRotPrefac-87.6789040924) > 1e-4 {
		t.Error("thermal prefactors not set correctly")
	}

	Setrandomseed(4)
	if math.Abs(rng.NormFloat64()-0.7133352143941205) > 1e-12 {
		t.Error("Random generator gives unexpected result")
	}
	Setrandomseed(5)
	if math.Abs(rng.NormFloat64()+0.7205428736258779) > 1e-12 {
		t.Error("Random generator gives unexpected result")
	}
	Setrandomseed(4)
	if math.Abs(rng.NormFloat64()-0.7133352143941205) > 1e-12 {
		t.Error("Random generator gives unexpected result")
	}

	B_ext = func(t float64) (float64, float64, float64) { return 0., 0., 1. / 28.003312237018984 / 2. / math.Pi }
	lijst[0].alpha = 0.
	SetSolver("euler")
	Dt.Set(1e-15)
	Run(1e-15)
	if math.Abs(lijst[0].m[1]-1e-6) > 1e-12 {
		t.Error("euler step gives wrong results")
	}

	lijst[0].m = vector{1, 0, 0}
	T.Set(0.)
	Adaptivestep = false
	SetSolver("dopri")
	Run(1e-15)
	if math.Abs(lijst[0].m[1]-1e-6) > 1e-12 {
		t.Error("dorpri step, fixed dt gives wrong results")
	}

	lijst[0].m = vector{1, 0, 0}
	T.Set(0.)
	Adaptivestep = true
	Errortolerance = 1e-12
	B_ext = func(t float64) (float64, float64, float64) { return 0., 0., 1. / 28.003312237018984 }
	Output(1e-15)  //
	Tableadd("Dt") //
	Run(1e-9)
	if math.Abs(lijst[0].m[1]-0) > 1e-12 {
		t.Error("dorpri step, adaptive dt gives wrong results")
	}
	if Dt.value == 1e-15 {
		t.Error("adaptive step does not adapt steps")
	}
}
