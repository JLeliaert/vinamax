package vinamax 

import (
	"math"
	"math/rand"
)

//Sums the individual fields to the effective field working on the particle
func (p particle) B_eff() [3]float64 {
	return Add(p.demagnetising_field, Add(p.anis(), Add(p.zeeman(), p.Temp())))
}

//Calculates the the thermal field B_therm working on a particle
// TODO factor 4/3pi in "number" omdat ze bol zijn!
func (p particle) Temp() [3]float64 {
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

//Calculates the zeeman field working on a particle
func (p particle) zeeman() [3]float64 {
	return B_ext
}

//Calculates the anisotropy field working on a particle
func (p particle) anis() [3]float64 {
	//2*Ku1*(m.u)*u/Msat
	var m [3]float64
	m[0] = p.m[0]
	m[1] = p.m[1]
	m[2] = p.m[2]

	mdotu := Dot(m, p.u_anis)

	return Times(p.u_anis, (2 * Ku1 * mdotu / Msat))
}
