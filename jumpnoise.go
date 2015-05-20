//This file contains functions used for the jumpnoise

package vinamax

import (
	"math"
	//	"fmt"
)

//calculates the attempt frequency of a particle
func attemptf1(p *particle) float64 {
	volume := cube(p.r) * 4. / 3. * math.Pi
	hk := 2. * Ku1 / (p.msat * mu0)
	gprime := Alpha * gamma0 * mu0 / (1. + (Alpha * Alpha))
	delta := Ku1 * volume / (kb * Temp)
	bx, by, bz := B_ext(T)
	bextvector := vector{bx / mu0, by / mu0, bz / mu0}
	hoverhk := math.Abs(bextvector.dot(p.u_anis)) / (hk)
	if math.Signbit(bextvector.dot(p.m)) == false {
		hoverhk *= -1.
	}
	//works only with aligned particles with B_ext at the moment
	fieldfactor := (1. - 1*hoverhk)
	//fmt.Println(gprime*hk*math.Sqrt(delta/math.Pi)*math.Exp(-delta*fieldfactor))
	return gprime * hk * math.Sqrt(delta/math.Pi) * math.Exp(-delta*fieldfactor)
}

//calculates the attempt frequency of a particle
func attemptf2(p *particle) float64 {
	volume := cube(p.r) * 4 / 3. * math.Pi
	hk := 2. * Ku1 / (p.msat * mu0)

	gprime := Alpha * gamma0 * mu0 / (1. + (Alpha * Alpha))

	delta := Ku1 * volume / (kb * Temp)

	bx, by, bz := B_ext(T)
	bextvector := vector{bx / mu0, by / mu0, bz / mu0}
	hoverhk := math.Abs(bextvector.dot(p.u_anis)) / hk
	if math.Signbit(bextvector.dot(p.m)) {
		hoverhk *= -1.
	}

	postpostfactor := 1. / (math.Erf(math.Sqrt(delta) * (1 - hoverhk)))

	return gprime * hk / math.Sqrt(delta*math.Pi) * postpostfactor * math.Exp(-delta)

}

//calculates the attempt frequency of a particle
func attemptf3(p *particle) float64 {
	volume := cube(p.r) * 4 / 3. * math.Pi
	hk := 2. * Ku1 / (p.msat * mu0)

	gprime := Alpha * gamma0 * mu0 / (1. + (Alpha * Alpha))

	delta := Ku1 * volume / (kb * Temp)

	bx, by, bz := B_ext(T)
	bextvector := vector{bx / mu0, by / mu0, bz / mu0}
	hoverhk := math.Abs(bextvector.dot(p.u_anis)) / hk
	if math.Signbit(bextvector.dot(p.m)) {
		hoverhk *= -1.
	}

	postpostfactor := 1. / (math.Log(2. / (1 + hoverhk)))

	return gprime * hk / (4 * delta) * postpostfactor

}

//calculates the attempt frequency of a particle
func attemptf0(p *particle) float64 {
	volume := cube(p.r) * 4. / 3. * math.Pi
	delta := Ku1 * volume / (kb * Temp)

	return 1./Tau0*math.Exp(-delta)

}

//calculates the next switching time
func setswitchtime(p *particle) {
	prob := rng.Float64()

	//TODO choose based on the barrier??? see which one corresponds when with brownian noise

	nextflip := (-1. / attemptf0(p)) * math.Log((1. - prob))
	//nextflip := (-1. / attemptf1(p)) * math.Log((1. - prob))
	//nextflip := (-1. / attemptf2(p)) * math.Log(1.-prob)
	//nextflip := (-1. / attemptf3(p)) * math.Log(1.-prob)

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
