package vinamax

import (
	"math"
	"math/rand"
)

var anisrng = rand.New(rand.NewSource(0))

//Calculates the torque working on the uniaxial anisotropy axis of a particle
//using the Langevin equation
func (p *particle) tau_u(randomv vector) vector {
	mdotBext := p.m.dot(p.zeeman())
	mdotBextdotu := p.u_anis.dot(mdotBext).times((4. / 3. * math.Pi * cube(p.r) * p.msat) / (6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h)))
	randomdotu := randomv.dot(p.u_anis)
	return mdotBextdotu.add(randomdotu)
}

//Set the randomseed for the anisotropy dynamics ??Is this necessary or use temp seed? better to use independent one in case there is no temp?
func Setrandomseed_anis(a int64) {
	randomseedcalled_anis = true
	anisrng = rand.New(rand.NewSource(a))
}

func (p *particle) calculaterandomvprefact() {
	p.randomvprefact = math.Sqrt((2. * kb * Temp) / (6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h)))
}

func calculaterandomvprefacts(lijst []*particle) {
	for i := range lijst {
		lijst[i].calculaterandomvprefact()
	}
}

//Calculates the randomness working on the particles' anisotropy axis
func (p *particle) randomv() vector {
	rand_tor := vector{0., 0., 0.}
	if BrownianRotation {
		etax := anisrng.NormFloat64()
		etay := anisrng.NormFloat64()
		etaz := anisrng.NormFloat64()

		rand_tor = vector{etax, etay, etaz}
		rand_tor = rand_tor.times(p.randomvprefact/math.Sqrt(Dt))
	}
	return rand_tor
}
