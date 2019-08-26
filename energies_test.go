package vinamax

import (
	"math"
	"testing"
)

func TestEnergies(t *testing.T) {
	Clear()
	Ku1.Set(1.e5 / (4. / 3. * math.Pi))
	M.Set(1., 0., 0.)
	Rc.Set(1.e-8)
	Rh.Set(1.e-8)
	Msat.Set(1e5 / (4. / 3. * math.Pi))

	Temp.Set(300)
	AddParticle(0, 0, 0)
	AddParticle(1e-7, 0, 0)
	setThermPrefac()
	B_ext = func(t float64) (float64, float64, float64) { return 1., 0., 0. }
	if math.Abs(lijst[0].e_zeeman()+1e-19) > 1e-34 {
		t.Error("incorrect zeeman energy")
	}
	if math.Abs(E_zeeman()+2e-19) > 1e-34 {
		t.Error("incorrect total zeeman energy")
	}

	B_ext = func(t float64) (float64, float64, float64) { return 1., 1., 0. }
	if math.Abs(lijst[0].e_zeeman()+1e-19*math.Sin(math.Pi/4.)*math.Sqrt(2.)) > 1e-34 {
		t.Error("incorrect zeeman energy")
	}

	B_ext = func(t float64) (float64, float64, float64) { return 0., 1., 0. }
	if math.Abs(lijst[0].e_zeeman()) > 1e-34 {
		t.Error("incorrect zeeman energy")
	}
	B_ext = func(t float64) (float64, float64, float64) { return 1., 0., 0. }

	lijst[0].u = vector{1., 0., 0.}
	if math.Abs(lijst[0].e_anis()+1e-19) > 1e-34 {
		t.Error("incorrect anisotropy energy")
	}
	lijst[1].u = vector{0., 1., 0.}
	if math.Abs(lijst[1].e_anis()) > 1e-34 {
		t.Error("incorrect anisotropy energy")
	}
	if math.Abs(E_anis()+1e-19) > 1e-34 {
		t.Error("incorrect total anisotropy energy")
	}
	calculatedemag()
	if math.Abs(lijst[0].demagnetising_field[0]-2e-5) > 1e-15 {
		t.Error("particle to particle demagnetising field not correct")
	}
	if lijst[0].demagnetising_field[1] != 0. || lijst[0].demagnetising_field[2] != 0. {
		t.Error("particle to particle demagnetising field not correct")
	}

	Demag = false
	if size(lijst[0].demag()) != 0. {
		t.Error("demag is not turned off but should be")
	}
	Demag = true

	if math.Abs(lijst[0].e_demag()+1e-24) > 1e-34 {
		t.Error("incorrect demag energy")
	}
	if math.Abs(lijst[1].e_demag()+1e-24) > 1e-34 {
		t.Error("incorrect demag energy")
	}
	if math.Abs(E_demag()+2e-24) > 1e-34 {
		t.Error("incorrect total demag energy")
	}
	lijst[1].SetM(vector{0., 1., 0.})
	if math.Abs(E_demag()) > 1e-34 {
		t.Error("incorrect total demag energy")
	}
	lijst[0].SetM(vector{0., 1., 0.})
	if math.Abs(E_demag()-1e-24) > 1e-34 {
		t.Error("incorrect total demag energy")
	}
	lijst[0].SetM(vector{1., 0., 0.})
	lijst[1].SetM(vector{1., 0., 0.})

	Dt.Set(1e-15)
	Setrandomseed_anis(12)
	Setrandomseed(12)
	lijst[0].setRotThermField()
	lijst[0].setThermField()
	Setrandomseed_anis(12)
	Setrandomseed(12)
	lijst[1].setRotThermField()
	lijst[1].setThermField()
	p := lijst[0]

	if p.rotThermField[0] != -2.3477049901054058e+10 || p.rotThermField[1] != 1.986079484065659e+10 || p.rotThermField[2] != 6.731605023005455e+09 {
		t.Error("incorrect rotational torque")
	}
	if p.thermField[0] != -28.058782158285148 || p.thermField[1] != 23.736786277366 || p.thermField[2] != 8.04533106638956 {
		t.Error("incorrect thermal field ")
	}

	if math.Abs(p.e_therm()-28.05878215828e-19) > 1e-28 {
		t.Error("incorrect thermal energy")
	}
	if math.Abs(E_therm()-2.*28.05878215828e-19) > 1e-28 {
		t.Error("incorrect total thermal energy")
	}
	if math.Abs(E_total()-5.311754431657031e-18) > 1e-34 {
		t.Error("incorrect total energy")
	}
}
