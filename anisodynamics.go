package vinamax

import (
	//"fmt"
	"math"
	"math/rand"
)

var rotrng = rand.New(rand.NewSource(0))

//"xi" prefactor for the mechanical rotations
func xi(radius float64) float64 {
	return 6. * Viscosity.value * volume(radius)
}

//Calculates the torque working on the uniaxial anisotropy axis of a particle
//  2KV/Xi (m.u)[(m.u)u-m]
func (p *particle) tau_u() vector {
	upart := vector{0., 0., 0.}
	mdotu := p.m.dot(p.u)
	uminm := (p.u.times(mdotu)).add(p.m.times(-1))
	upart = uminm.times((-1) * mdotu * (2 * p.ku1 * volume(p.rc)) / (xi(p.rh)))
	return upart.add(p.rotThermField)
}

//Set the randomseed for the anisotropy dynamics
func Setrandomseed_anis(a int64) {
	//randomseedcalled_anis = true
	rotrng = rand.New(rand.NewSource(a))
}

//Calculates the Brownian torques on the particles' anisotropy axis
func (p *particle) setRotThermField() {
	Rot_torque := vector{0., 0., 0.}
	if Temp.value != 0 {
		etax := rotrng.NormFloat64()
		etay := rotrng.NormFloat64()
		etaz := rotrng.NormFloat64()

		Rot_torque = vector{etax, etay, etaz}
		Rot_torque = Rot_torque.times(p.thermRotPrefac / math.Sqrt(Dt.value))
	}
	p.rotThermField = Rot_torque
}
