package vinamax

import (
	"math"
)

//zie 2.51 in coey en watweuitrekenen.pdf

func calculatedemag() {
	//FMM demag

	//construct tree (eenmalig!!!)

	//O(N**2)
	for i := range Lijst {
		Lijst[i].demagnetising_field = Lijst[i].demag()
	}
}

//demag is calculated on a position
func demag(x, y, z float64) Vector {
	//TODO dit volume beter maken en bolletjes!
	volume := math.Pow(2e-9, 3)
	prefactor := (mu0 * Msat * volume) / (4 * math.Pi)
	demag := Vector{0, 0, 0}

	for i := range Lijst {
		if Lijst[i].X != x || Lijst[i].Y != y || Lijst[i].Z != z {
			r_vect := Vector{x - Lijst[i].X, y - Lijst[i].Y, z - Lijst[i].Z}
			r := Lijst[i].dist(x, y, z)
			r2 := r * r
			r3 := r * r2
			r5 := r3 * r2

			demag[0] += prefactor * ((3 * Lijst[i].m[0] * r_vect[0] * r_vect[0] / r5) - (Lijst[i].m[0] / r3))

			demag[1] += prefactor * ((3. * Lijst[i].m[1] * r_vect[1] * r_vect[1] / r5) - (Lijst[i].m[1] / r3))

			demag[2] += prefactor * ((3 * Lijst[i].m[2] * r_vect[2] * r_vect[2] / r5) - (Lijst[i].m[2] / r3))

		}

	}
	//demag = Times(demag, prefactor)
	return demag
}

//demag on a Particle
func (p Particle) demag() Vector {
	return demag(p.X, p.Y, p.Z)
}

//The distance between a Particle and a location
func (r *Particle) dist(x, y, z float64) float64 {
	return math.Sqrt(math.Pow(float64(r.X-x), 2) + math.Pow(float64(r.Y-y), 2) + math.Pow(float64(r.Z-z), 2))
}
