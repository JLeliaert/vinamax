package vinamax

import (
	"math"
)

func calculatedemag() {
	N := len(lijst)
	cleandemag()
	for i := range lijst {
		demagloop(i+1, N)
	}
}

func demagloop(min, max int) {
	for j := min; j < max; j++ {
		demag2p(lijst[min-1], lijst[j])
	}
}

//resets all demag fields to zero
func cleandemag() {
	for _, p := range lijst {
		p.demagnetising_field = vector{0, 0, 0}
	}
}

//adds the demagfield of p1 to p2 and vice versa
func demag2p(p1, p2 *particle) {
	prefactor := mu0 / 3.
	ms_volume1 := cube(p1.rc) * p1.msat * prefactor
	ms_volume2 := cube(p2.rc) * p2.msat * prefactor
	r_vect := vector{p1.x - p2.x, p1.y - p2.y, p1.z - p2.z}
	r := p1.dist(p2.x, p2.y, p2.z)
	r2 := r * r
	r3 := r * r2
	r5 := r3 * r2

	dotproduct1 := p1.m.dot(r_vect)
	dotproduct2 := p2.m.dot(r_vect)

	p1.demagnetising_field = p1.demagnetising_field.add(vector{ms_volume2 * ((3 * dotproduct2 * r_vect[0] / r5) - (p2.m[0] / r3)), ms_volume2 * ((3 * dotproduct2 * r_vect[1] / r5) - (p2.m[1] / r3)), ms_volume2 * ((3 * dotproduct2 * r_vect[2] / r5) - (p2.m[2] / r3))})
	p2.demagnetising_field = p2.demagnetising_field.add(vector{ms_volume1 * ((3 * dotproduct1 * r_vect[0] / r5) - (p1.m[0] / r3)), ms_volume1 * ((3 * dotproduct1 * r_vect[1] / r5) - (p1.m[1] / r3)), ms_volume1 * ((3 * dotproduct1 * r_vect[2] / r5) - (p1.m[2] / r3))})
}

//Demag is calculated on a position
func demag(x, y, z float64) vector {
	prefactor := mu0 / (4 * math.Pi)
	demag := vector{0, 0, 0}

	for _, p := range lijst {
		if p.x != x || p.y != y || p.z != z {
			ms_volume := volume(p.rc) * p.msat * prefactor
			r_vect := vector{x - p.x, y - p.y, z - p.z}
			r := p.dist(x, y, z)
			r2 := r * r
			r3 := r * r2
			r5 := r3 * r2

			dotproduct := p.m.dot(r_vect)

			for q := 0; q < 3; q++ {
				demag[q] += ms_volume * ((3 * dotproduct * r_vect[q] / r5) - (p.m[q] / r3))
			}

		}
	}
	return demag
}

//Demag on a particle
func (p particle) demag() vector {
	return demag(p.x, p.y, p.z)
}

//The distance between a particle and a location
func (r *particle) dist(x, y, z float64) float64 {
	return math.Sqrt(sqr(float64(r.x-x)) + sqr(float64(r.y-y)) + sqr(float64(r.z-z)))
}
