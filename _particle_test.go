package vinamax

import (
	"math"
	"testing"
)

func TestParticle(t *testing.T) {
	Clear()
	AddParticle(0, 0, 0)
	p := lijst[0]
	if p.alpha != 0.01 {
		t.Error("wrong default alpha")
	}
	if p.ku1 != 0.00 {
		t.Error("wrong default ku1")
	}
	if p.u[0] != 0. || p.u[1] != 0. || p.u[2] != 1. {
		t.Error("wrong default u")
	}
	if p.m[0] != 0. || p.m[1] != 0. || p.m[2] != 1. {
		t.Error("wrong default m")
	}
	if p.rc != 1e-8 {
		t.Error("wrong default rc")
	}
	if p.rh != 1e-8 {
		t.Error("wrong default rh")
	}
	if p.msat != 400e3 {
		t.Error("wrong default msat")
	}
	Alpha.Set(1.)
	Ku1.Set(1.)
	if Ku1.Get() != 1. {
		t.Error("wrong scalarvariable Get")
	}
	U_anis.Set(1., 0., 0.)
	if U_anis.Get()[0] != 1. || U_anis.Get()[1] != 0. || U_anis.Get()[2] != 0. {
		t.Error("wrong vectorvariable Get")
	}
	M.Set(1., 0., 0.)
	Rc.Set(25.e-9)
	Rh.Set(25.e-9)
	Msat.Set(100e3)
	AddParticle(3., 3., 3.)
	p = lijst[1]

	if p.x != 3. || p.y != 3. || p.z != 3. {
		t.Error("wrong set alpha")
	}
	if p.alpha != 1. {
		t.Error("wrong set alpha")
	}
	if p.ku1 != 1. {
		t.Error("wrong set ku1")
	}

	if p.u[0] != 1. || p.u[1] != 0. || p.u[2] != 0. {
		t.Error("wrong set u")
	}
	if p.m[0] != 1. || p.m[1] != 0. || p.m[2] != 0. {
		t.Error("wrong set m")
	}
	if p.rc != 25e-9 {
		t.Error("wrong set rc")
	}
	if p.rh != 25e-9 {
		t.Error("wrong set rh")
	}
	if p.msat != 100e3 {
		t.Error("wrong set msat")
	}
	if math.Abs(p.Volume()-6.544984694e-23) > 1e-29 {
		t.Error("wrong particle Volume")
	}
	p.SetM(vector{0., 2., 2.})
	if p.GetM()[0] != 0. || math.Abs(p.GetM()[1]-math.Sqrt(2.)/2.) > 1e-13 || math.Abs(p.GetM()[2]-math.Sqrt(2.)/2.) > 1e-13 {
		t.Error("wrong SetM or GetM")
	}
}
