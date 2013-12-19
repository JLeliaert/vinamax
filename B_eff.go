package vinamax

import (
	"math"
	"math/rand"
)

//Sums the individual fields to the effective field working on the Particle
func (p Particle) b_eff() Vector {
	return p.demagnetising_field.add(p.anis().add(p.zeeman().add(p.temp())))
}

//Calculates the the thermal field B_therm working on a Particle
// TODO factor 4/3pi in "number" omdat ze bol zijn!
func (p Particle) temp() Vector {
	B_therm := Vector{0., 0., 0.}
	if Temp != 0. {

		etax := rand.NormFloat64()
		etay := rand.NormFloat64()
		etaz := rand.NormFloat64()
		B_therm = Vector{etax, etay, etaz}
		number := math.Sqrt((2 * kb * Alpha * Temp) / (gamma0 * Msat * 4./3. * math.Pi * math.Pow(p.r,3)*Dt))
		B_therm = B_therm.times(number)
	}
	return B_therm
}

//Calculates the zeeman field working on a Particle
func (p Particle) zeeman() Vector {
	return B_ext
}

//Calculates the anisotropy field working on a Particle
func (p Particle) anis() Vector {
	//2*Ku1*(m.u)*u/Msat
	var m Vector
	m[0] = p.m[0]
	m[1] = p.m[1]
	m[2] = p.m[2]

	mdotu := m.dot(p.u_anis)

	return p.u_anis.times(2 * Ku1 * mdotu / Msat)
}
