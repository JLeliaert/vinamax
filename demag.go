package vinamax

import (
	"math"
	//	"fmt"
)

//zie 2.51 in coey en watweuitrekenen.pdf

func calculatedemag() {

	//TODO: maketree uit main halen construct tree (eenmalig!!!)

	for i := range universe.lijst {
		universe.lijst[i].demagnetising_field = universe.lijst[i].demag()
	}
}

//demag is calculated on a position
func demag(x, y, z float64) Vector {
	prefactor := mu0 / (4 * math.Pi)
	demag := Vector{0, 0, 0}

	for i := range universe.lijst {
		if universe.lijst[i].X != x || universe.lijst[i].Y != y || universe.lijst[i].Z != z {
			volume := 4. / 3 * math.Pi * math.Pow(universe.lijst[i].r, 3)
			r_vect := Vector{x - universe.lijst[i].X, y - universe.lijst[i].Y, z - universe.lijst[i].Z}
			r := universe.lijst[i].dist(x, y, z)
			r2 := r * r
			r3 := r * r2
			r5 := r3 * r2

			demag[0] += universe.lijst[i].msat * volume * prefactor * ((3 * universe.lijst[i].m[0] * r_vect[0] * r_vect[0] / r5) - (universe.lijst[i].m[0] / r3))

			demag[1] += universe.lijst[i].msat * volume * prefactor * ((3. * universe.lijst[i].m[1] * r_vect[1] * r_vect[1] / r5) - (universe.lijst[i].m[1] / r3))

			demag[2] += universe.lijst[i].msat * volume * prefactor * ((3 * universe.lijst[i].m[2] * r_vect[2] * r_vect[2] / r5) - (universe.lijst[i].m[2] / r3))

		}

	}
	//demag = Times(demag, prefactor)
	return demag
}

//demag on a Particle
func (p Particle) demag() Vector {
	//TODO hier de keuze laten
	if FMM {
		return FMMdemag(p.X, p.Y, p.Z)
	}
	return demag(p.X, p.Y, p.Z)
}

//The distance between a Particle and a location
func (r *Particle) dist(x, y, z float64) float64 {
	return math.Sqrt(math.Pow(float64(r.X-x), 2) + math.Pow(float64(r.Y-y), 2) + math.Pow(float64(r.Z-z), 2))
}

//demag is calculated on a position
func FMMdemag(x, y, z float64) Vector {
	prefactor := mu0 / (4 * math.Pi)
	demag := Vector{0, 0, 0}

	//lijst maken met nodes
	//node universe in de box steken
	nodelist := []*node{&universe}
	//for lijst!=leeg
	for len(nodelist) > 0 {
		//	for i := range nodelist {
		i := 0
		//if aantalparticles in box==0: delete van stack
		//	if nodelist[i].number == 0 {
		//		nodelist[i] = nodelist[len(nodelist)-1]
		//		nodelist = nodelist[0 : len(nodelist)-1]
		//	}
		if nodelist[i].number == 1 {
			//if aantalparticles in box==1:
			if nodelist[i].lijst[0].X != x || nodelist[i].lijst[0].Y != y || nodelist[i].lijst[0].Z != z {
				//	if ik ben niet die ene: calculate en delete van stack
				//	CALC
				volume := 4. / 3 * math.Pi * math.Pow(nodelist[i].lijst[0].r, 3)
				r_vect := Vector{x - nodelist[i].lijst[0].X, y - nodelist[i].lijst[0].Y, z - nodelist[i].lijst[0].Z}
				r := nodelist[i].lijst[0].dist(x, y, z)
				r2 := r * r
				r3 := r * r2
				r5 := r3 * r2

				demag[0] += nodelist[i].lijst[0].msat * volume * prefactor * ((3 * nodelist[i].lijst[0].m[0] * r_vect[0] * r_vect[0] / r5) - (nodelist[i].lijst[0].m[0] / r3))

				demag[1] += nodelist[i].lijst[0].msat * volume * prefactor * ((3. * nodelist[i].lijst[0].m[1] * r_vect[1] * r_vect[1] / r5) - (nodelist[i].lijst[0].m[1] / r3))

				demag[2] += nodelist[i].lijst[0].msat * volume * prefactor * ((3 * nodelist[i].lijst[0].m[2] * r_vect[2] * r_vect[2] / r5) - (nodelist[i].lijst[0].m[2] / r3))

			}
			//	nodelist[i] = nodelist[len(nodelist)-1]
			//	nodelist = nodelist[0 : len(nodelist)-1]
		}
		if nodelist[i].number > 1 {
			//if aantalparticles in box>1:
			r_vect := Vector{x - nodelist[i].com[0], y - nodelist[i].com[1], z - nodelist[i].com[2]}
			r := math.Sqrt(r_vect[0]*r_vect[0] + r_vect[1]*r_vect[1] + r_vect[2]*r_vect[2])

			if (nodelist[i].where(Vector{x, y, z}) == -1 && (math.Sqrt(2)/2.*nodelist[i].diameter/r) < Thresholdbeta) {
				//	if voldoet aan criterium: calculate en delete van stack
				//fmt.Println(nodelist[i].number)
				m := Vector{0, 0, 0}
				volume := 0.
				//in loopje m en volume berekenen
				for j := range nodelist[i].lijst {
					volume += 4. / 3 * math.Pi * math.Pow(nodelist[i].lijst[i].r, 3)

					m[0] += nodelist[i].lijst[j].m[0] * nodelist[i].lijst[j].msat
					m[1] += nodelist[i].lijst[j].m[1] * nodelist[i].lijst[j].msat
					m[2] += nodelist[i].lijst[j].m[2] * nodelist[i].lijst[j].msat
				}
				r2 := r * r
				r3 := r * r2
				r5 := r3 * r2

				demag[0] += volume * prefactor * ((3 * m[0] * r_vect[0] * r_vect[0] / r5) - (m[0] / r3))

				demag[1] += volume * prefactor * ((3 * m[1] * r_vect[1] * r_vect[1] / r5) - (m[1] / r3))

				demag[2] += volume * prefactor * ((3 * m[2] * r_vect[2] * r_vect[2] / r5) - (m[2] / r3))

			} else {
				//	if not: add subboxen en delete van stack
				nodelist = append(nodelist, nodelist[i].tlb)
				nodelist = append(nodelist, nodelist[i].tlf)
				nodelist = append(nodelist, nodelist[i].trb)
				nodelist = append(nodelist, nodelist[i].trf)
				nodelist = append(nodelist, nodelist[i].blb)
				nodelist = append(nodelist, nodelist[i].blf)
				nodelist = append(nodelist, nodelist[i].brb)
				nodelist = append(nodelist, nodelist[i].brf)
			}
		}
		copy(nodelist[i:], nodelist[i+1:])
		nodelist[len(nodelist)-1] = nil // or the zero value of T
		nodelist = nodelist[:len(nodelist)-1]
	}
	//}
	return demag
}
