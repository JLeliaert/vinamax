package vinamax 

import (
	"math"
	"math/rand"
)

//Sums the individual fields to the effective field working on the Particle
func (p Particle) B_eff() [3]float64 {
	return Add(p.demagnetising_field, Add(p.anis(), Add(p.zeeman(), p.Temp())))
}

//Calculates the the thermal field B_therm working on a Particle
// TODO factor 4/3pi in "number" omdat ze bol zijn!
func (p Particle) Temp() [3]float64 {
	B_therm := [3]float64{0., 0., 0.}
	if Temp != 0. {

		etax := rand.NormFloat64()
		etay := rand.NormFloat64()
		etaz := rand.NormFloat64()
		B_therm = [3]float64{etax, etay, etaz}
		number := math.Sqrt((2 * Kb * Alpha * Temp) / (Gamma0 * Msat * Dx * Dy * Dz * Dt))
		B_therm = Times(B_therm, number)
	}
	return B_therm
}

//Calculates the zeeman field working on a Particle
func (p Particle) zeeman() [3]float64 {
	return B_ext
}

//Calculates the anisotropy field working on a Particle
func (p Particle) anis() [3]float64 {
	//2*Ku1*(m.u)*u/Msat
	var m [3]float64
	m[0] = p.M[0]
	m[1] = p.M[1]
	m[2] = p.M[2]

	mdotu := Dot(m, p.u_anis)

	return Times(p.u_anis, (2 * Ku1 * mdotu / Msat))
}
