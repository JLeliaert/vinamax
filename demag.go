package main

import (
	"math"
)

//zie 2.51 in coey en watweuitrekenen.pdf



func calculatedemag(){
	for i := range lijst {
		lijst[i].demagnetising_field=lijst[i].demag()
	}
}


//demag is calculated on a position
func demag(x, y, z float64) [3]float64 {
	//TODO dit volume beter maken en bolletjes!
	volume := math.Pow(2e-9, 3)
	prefactor := (Mu0 * Msat*volume) / (4 * math.Pi)
	demag := [3]float64{0, 0, 0}

	for i := range lijst {
		if lijst[i].x != x || lijst[i].y != y || lijst[i].z != z {
			r_vect := [3]float64{x - lijst[i].x, y - lijst[i].y, z - lijst[i].z}
			r := lijst[i].dist(x, y, z)
			r2 := r * r
			r3 := r * r2
			r5 := r3 * r2

			demag[0] += prefactor*((3 * lijst[i].m[0] * r_vect[0] * r_vect[0] / r5) - (lijst[i].m[0] / r3))

			demag[1] += prefactor*((3. * lijst[i].m[1] * r_vect[1] * r_vect[1] / r5) - (lijst[i].m[1] / r3))

			demag[2] += prefactor*((3 * lijst[i].m[2] * r_vect[2] * r_vect[2] / r5) - (lijst[i].m[2] / r3))

			
		}

	}
	//demag = Times(demag, prefactor)
	return demag
}

//demag on a particle
func (p particle) demag() [3]float64 {
	return demag(p.x, p.y, p.z)
}

//The distance between a particle and a location
func (r *particle) dist(x, y, z float64) float64 {
	return math.Sqrt(math.Pow(float64(r.x-x), 2) + math.Pow(float64(r.y-y), 2) + math.Pow(float64(r.z-z), 2))
}
