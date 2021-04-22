package vinamax

import (
	"math"
	"math/rand"
)

var rng = rand.New(rand.NewSource(0))

//Sums the individual fields + the thermal field to the effective field working on the particle
func (p *particle) b_eff(thermField vector) vector {
	return p.demagnetising_field.add(p.anis().add(p.zeeman().add(thermField)))
}

//Set the randomseed for the temperature
func Setrandomseed(a int64) {
	rng = rand.New(rand.NewSource(a))
}

//sets the prefactor for the thermal fields
func setThermPrefac() {
	for _, p := range lijst {
		p.thermPrefac = math.Sqrt((2. * kb * p.alpha * p.temp) / (gamma0 * p.msat * Volume(p.rc)))
		p.thermRotPrefac = math.Sqrt((2. * kb * p.temp) / (xi(p.rh)))
	}
}

//Calculates the the thermal field B_therm working on a particle
func (p *particle) setThermField() {
	B_therm := vector{0., 0., 0.}
	if p.temp != 0 {
		etax := rng.NormFloat64()
		etay := rng.NormFloat64()
		etaz := rng.NormFloat64()

		B_therm = vector{etax, etay, etaz}
		B_therm = B_therm.times(p.thermPrefac / math.Sqrt(Dt.value))
	}
	p.thermField = B_therm
}

//Calculates the Zeeman field working on a particle
func (p *particle) zeeman() vector {
	x, y, z := B_ext(T.value)
	x2, y2, z2 := B_ext_space(T.value, p.x, p.y, p.z)
	return vector{x + x2, y + y2, z + z2}
}

//Calculates the anisotropy field working on a particle
//2*Ku1*(m.u)*u/p.msat
func (p *particle) anis() vector {
	mdotu := p.m.dot(p.u)
	return p.u.times(2. * p.ku1 * mdotu / p.msat)
}
