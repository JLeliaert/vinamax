package vinamax

//normal random number generator using marsaglias polar method
//hopefully faster than the standar golang generator
// http://en.wikipedia.org/wiki/Marsaglia_polar_method

import (
	"math"
	"math/rand"
)

var (
	newrng       = rand.New(rand.NewSource(0))
	spare        float64
	isspareready bool = false
)

func normfloat() float64 {
	if isspareready {
		isspareready = false
		return spare

	} else {
		u := (newrng.Float64() - 0.5) * 2
		v := (newrng.Float64() - 0.5) * 2
		s := u*u + v*v
		for s > 1 {
			u = (newrng.Float64() - 0.5) * 2
			v = (newrng.Float64() - 0.5) * 2
			s = u*u + v*v
		}
		lns := math.Log(s)
		num := math.Sqrt(-2 * lns / s)
		isspareready = true
		spare = num * u
		return num * v

	}
}
