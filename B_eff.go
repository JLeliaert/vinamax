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
func (p *particle) calculatetemp_prefactor() {
	p.temp_prefactor = math.Sqrt((2. * kb * Alpha * Temp) / (gamma0 * p.msat * 4. / 3. * math.Pi * cube(p.r)))
}

func calculatetemp_prefactors(lijst []*particle) {
	for i := range lijst {
		lijst[i].calculatetemp_prefactor()
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
			B_therm = B_therm.times(p.temp_prefactor / math.Sqrt(Dt))
		}
	}
	return B_therm
}

//Calculates the Zeeman field working on a particle
func (p *particle) zeeman() vector {
	x, y, z := B_ext(T)
	x2, y2, z2 := B_ext_space(T, p.x, p.y, p.z)
	return vector{x + x2, y + y2, z + z2}.add(p.biasfield)
}

//calculates the zeeman energy in a particle
func (p *particle) e_zeeman() float64 {
	return -1*p.msat*p.Volume()*p.m.dot(p.zeeman())
}

//Calculates the anisotropy field working on a particle
func (p *particle) anis() vector {
	//2*Ku1*(m.u)*u/p.msat

	mdotu := p.m.dot(p.u_anis)
	uniax := p.u_anis.times(2. * Ku1 * mdotu / p.msat)
	mdotu2 := p.m.dot(p.u2_anis)
	uniax2 := p.u2_anis.times(2. * Ku2 * mdotu2 / p.msat)

	cubic := vector{0., 0., 0.}
	if Kc1 != 0 {
		c1m := p.m.dot(p.c1_anis)
		c2m := p.m.dot(p.c2_anis)
		c3m := p.m.dot(p.c3_anis)
		firstterm := p.c1_anis.times(c1m * (c3m*c3m + c2m*c2m))
		secondterm := p.c2_anis.times(c2m * (c3m*c3m + c1m*c1m))
		thirdterm := p.c3_anis.times(c3m * (c2m*c2m + c1m*c1m))

		cubic = firstterm.add(secondterm.add(thirdterm)).times(-2. * Kc1 / p.msat)
	}
	return uniax.add(cubic).add(uniax2)
}

//calculates the anisotropy energy in a particle
func (p *particle) e_anis() float64 {
	return -0.5*p.msat*p.Volume()*p.m.dot(p.anis())
}


//calculates the thermal energy in a particle
func (p *particle) e_therm() float64 {
	return -1*p.msat*p.Volume()*p.m.dot(p.tempfield)
}


//calculates the demag energy in a particle
func (p *particle) e_demag() float64 {
	return -0.5*p.msat*p.Volume()*p.m.dot(p.demag())
}
