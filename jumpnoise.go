//This file contains functions used for the jumpnoise

package vinamax

import (
	"math"
	"math/rand"
)

//calculates the attempt frequency of a particle
func attemptf(p particle) float64 {
	prefactor := Alpha * gamma0 * mu0 / (1 + Alpha*Alpha)
	volume := cube(p.r) * 4 / 3. * math.Pi

	//TODO replace anisotropyfield by effective field?
	anisotropyfield := size(p.anis()) / mu0
	undersqrt := cube(anisotropyfield) * p.msat * volume / (2 * math.Pi * kb * Temp)

	return prefactor * math.Sqrt(undersqrt)
}

//calculates the next switching time
func setswitchtime(p particle) {
	prob := rand.Float64()
	nextflip := -1. / attemptf(p) * math.Log(1-prob)
	p.flip = nextflip + T
}

//checks if it's time to switch and if so, switch and calculate next switchtime
func checkswitch(p particle) {
	if T > p.flip {
		switchp(p)
		setswitchtime(p)
	}
}

//switches the magnetisation of a particle
func switchp(p particle) {
	p.m = p.m.times(-1.)
}
