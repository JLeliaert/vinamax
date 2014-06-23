package vinamax

import (
	"math"
	"math/rand"
)

var rng = rand.New(rand.NewSource(0))

//Sums the individual fields to the effective field working on the particle
func (p *particle) b_eff(temp vector) vector {
	return p.demagnetising_field.add(p.anis().add(p.zeeman().add(temp)))
}

//Set the randomseed for the temperature
func Setrandomseed(a int64) {
	randomseedcalled = true
	rng = rand.New(rand.NewSource(a))
}

// factor 4/3pi in "number" because they are spherical
func (p *particle) calculatetempnumber() {
	radius := radii[p.rindex]
	p.tempnumber = math.Sqrt((2. * kb * Alpha * Temp) / (gamma0 * p.msat * 4. / 3. * math.Pi * cube(radius) * Dt))
}

func calculatetempnumbers(lijst []*particle) {
	for i := range lijst {
		lijst[i].calculatetempnumber()
	}
}

//Calculates the the thermal field B_therm working on a particle
func (p *particle) temp() vector {
	B_therm := vector{0., 0., 0.}
	if Brown {
		if Temp != 0. {
			etax := rng.NormFloat64()
			etay := rng.NormFloat64()
			etaz := rng.NormFloat64()

			B_therm = vector{etax, etay, etaz}
			B_therm = B_therm.times(p.tempnumber)
		}
	}
	return B_therm
}

//Calculates the Zeeman field working on a particle
func (p *particle) zeeman() vector {
	x, y, z := B_ext(T)
	x2, y2, z2 := B_ext_space(T, p.x, p.y, p.z)
	return vector{x + x2, y + y2, z + z2}
}

//Calculates the anisotropy field working on a particle
func (p *particle) anis() vector {
	//2*Ku1*(m.u)*u/p.msat

	mdotu := p.m.dot(p.u_anis)
	return p.u_anis.times(2. * Ku1 * mdotu / p.msat)
}
