package vinamax

import (
	"math"
	"math/rand"
)

//Sums the individual fields to the effective field working on the Particle
func (p *Particle) b_eff() Vector {
	return p.demagnetising_field.add(p.anis().add(p.zeeman().add(p.temp())))
}

//set the randomseed for the temperatuur
func SetRandomSeed(a int64) {
	rand.Seed(a)
}

var rng = rand.New(rand.NewSource(0))

//Calculates the the thermal field B_therm working on a Particle
// TODO factor 4/3pi in "number" omdat ze bol zijn!
func (p *Particle) temp() Vector {
	B_therm := Vector{0., 0., 0.}
	if Temp != 0. {
		etax := rng.NormFloat64()
		etay := rng.NormFloat64()
		etaz := rng.NormFloat64()
		B_therm = Vector{etax, etay, etaz}
		number := math.Sqrt((2 * kb * Alpha * Temp) / (gamma0 * p.msat * 4. / 3. * math.Pi * cube(p.r) * Dt))
		B_therm = B_therm.times(number)
	}
	return B_therm
}

//Calculates the zeeman field working on a Particle
func (p *Particle) zeeman() Vector {
	return B_ext
}

//Calculates the anisotropy field working on a Particle
func (p *Particle) anis() Vector {
	//2*Ku1*(m.u)*u/p.msat

	mdotu := p.m.dot(p.u_anis)
	return p.u_anis.times(2 * Ku1 * mdotu / p.msat)
}
