//This file contains functions used for the jumpnoise

package vinamax

import (
//	"fmt"
	"math"
	"math/rand"
)

//calculates the attempt frequency of a particle
func attemptf(p *particle) float64 {
	prefactor := Alpha * gamma0 * mu0 / (1 + Alpha*Alpha)
	volume := cube(p.r) * 4 / 3. * math.Pi

	anisotropyfield := 2. * Ku1 / (p.msat * mu0)
	undersqrt := cube(anisotropyfield) * p.msat * volume / (2 * math.Pi * kb * Temp)

	barrier := volume * Ku1
	// exp( -Delta E/kt)
	postfactor := math.Exp(-barrier / (kb * Temp))
	bx, by, bz := B_ext(T)
	bextvector := vector{bx,by,bz}
	hoverhk := math.Abs(bextvector.dot(p.u_anis)) / (anisotropyfield * mu0 )
	if math.Signbit(bextvector.dot(p.m)) {
		hoverhk *= -1
	}
	postpostfactor := (1 - hoverhk) * (1 - hoverhk*hoverhk)

	return prefactor * math.Sqrt(undersqrt) * postfactor * postpostfactor
}

//calculates the next switching time
func setswitchtime(p *particle) {
	prob := rand.Float64()
	nextflip := -1. / attemptf(p) * math.Log(1-prob)
	p.flip = nextflip + T
}

//checks if it's time to switch and if so, switch and calculate next switchtime
func checkswitch(p *particle) {
	if T > p.flip {
		switchp(p)
		setswitchtime(p)
	}
}

//switches the magnetisation of a particle
func switchp(p *particle) {
	p.m = p.m.times(-1.)
}

//resets all switchtimes
func resetswitchtimes(Lijst []*particle) {
	for _, p := range Lijst {
		setswitchtime(p)
	}
}

//checks switch for all particles
func checkallswitch(Lijst []*particle) {
	for _, p := range Lijst {
		checkswitch(p)
	}
}
