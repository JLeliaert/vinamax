//This file contains functions used for the jumpnoise

package vinamax

import (
	"math"
	//	"math/rand"
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
