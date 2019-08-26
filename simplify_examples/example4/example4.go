package main

import (
	. "github.com/JLeliaert/vinamax"
	"math"
	"math/rand"
)

var rng = rand.New(rand.NewSource(0))

func main() {
	//for lognormal radius OR diameter
	for i := 0; i < 10000; i++ {
		r := lognormal(20., 1.)

		Rc.Set(r)
		Rh.Set(r)
		AddParticle(float64(i)*1.e7, 0., 0.)
	}

	Save("geometry")

}

//returns a lognormally distributed number, to be used as radius
func lognormal(mu, sigma float64) float64 {
	result := 0.
	for result == 0. {
		x := rng.Float64() * 20. * mu
		f_x := 1. / (math.Sqrt(2*math.Pi) * sigma * x) * math.Exp(-1./(2.*sigma*sigma)*math.Pow(math.Log(x/mu), 2.))
		if rng.Float64() < f_x {
			result = x * 1e-9 / 2.
		}
	}
	return result
}
