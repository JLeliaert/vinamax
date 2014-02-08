package vinamax

import (
	"math"
	//	"fmt"
)

//cfr. 2.51 in coey en watweuitrekenen.pdf

func calculatedemag() {

	for i := range universe.lijst {
		universe.lijst[i].demagnetising_field = universe.lijst[i].demag()
	}
}

//demag is calculated on a position
func demag(x, y, z float64) vector {
	prefactor := mu0 / (4 * math.Pi)
	demag := vector{0, 0, 0}

	for i := range universe.lijst {
		if universe.lijst[i].x != x || universe.lijst[i].y != y || universe.lijst[i].z != z {
			volume := 4. / 3 * math.Pi * cube(universe.lijst[i].r)
			r_vect := vector{x - universe.lijst[i].x, y - universe.lijst[i].y, z - universe.lijst[i].z}
			r := universe.lijst[i].dist(x, y, z)
			r2 := r * r
			r3 := r * r2
			r5 := r3 * r2

			dotproduct := universe.lijst[i].m.dot(r_vect)

			demag[0] += universe.lijst[i].msat * volume * prefactor * ((3 * dotproduct * r_vect[0] / r5) - (universe.lijst[i].m[0] / r3))

			demag[1] += universe.lijst[i].msat * volume * prefactor * ((3. * dotproduct * r_vect[1] / r5) - (universe.lijst[i].m[1] / r3))

			demag[2] += universe.lijst[i].msat * volume * prefactor * ((3 * dotproduct * r_vect[2] / r5) - (universe.lijst[i].m[2] / r3))

		}
	}
	return demag
}

//demag on a particle
func (p particle) demag() vector {
	if FMM {
		return fMMdemag(p.x, p.y, p.z)
	}
	return demag(p.x, p.y, p.z)
}

//The distance between a particle and a location
func (r *particle) dist(x, y, z float64) float64 {
	return math.Sqrt(sqr(float64(r.x-x)) + sqr(float64(r.y-y)) + sqr(float64(r.z-z)))
}

//demag is calculated on a position
func fMMdemag(x, y, z float64) vector {
	prefactor := mu0 / (4 * math.Pi)
	demag := vector{0, 0, 0}
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
			if nodelist[i].lijst[0].x != x || nodelist[i].lijst[0].y != y || nodelist[i].lijst[0].z != z {
				//	if ik ben niet die ene: calculate en delete van stack
				//	CALC
				volume := 4. / 3 * math.Pi * cube(nodelist[i].lijst[0].r)
				r_vect := vector{x - nodelist[i].lijst[0].x, y - nodelist[i].lijst[0].y, z - nodelist[i].lijst[0].z}
				r := nodelist[i].lijst[0].dist(x, y, z)
				r2 := r * r
				r3 := r * r2
				r5 := r3 * r2

				dotproduct := nodelist[i].lijst[i].m.dot(r_vect)

				demag[0] += nodelist[i].lijst[0].msat * volume * prefactor * ((3 * dotproduct * r_vect[0] / r5) - (nodelist[i].lijst[0].m[0] / r3))

				demag[1] += nodelist[i].lijst[0].msat * volume * prefactor * ((3 * dotproduct * r_vect[1] / r5) - (nodelist[i].lijst[0].m[1] / r3))

				demag[2] += nodelist[i].lijst[0].msat * volume * prefactor * ((3 * dotproduct * r_vect[2] / r5) - (nodelist[i].lijst[0].m[2] / r3))
			}
			//	nodelist[i] = nodelist[len(nodelist)-1]
			//	nodelist = nodelist[0 : len(nodelist)-1]
		}
		if nodelist[i].number > 1 {
			//if aantalparticles in box>1:
			r_vect := vector{x - nodelist[i].com[0], y - nodelist[i].com[1], z - nodelist[i].com[2]}
			r := math.Sqrt(r_vect[0]*r_vect[0] + r_vect[1]*r_vect[1] + r_vect[2]*r_vect[2])

			if (nodelist[i].where(vector{x, y, z}) == -1 && (math.Sqrt(2)/2.*nodelist[i].diameter/r) < Thresholdbeta) {
				//	if voldoet aan criterium: calculate en delete van stack
				m := vector{0, 0, 0}
				volume := 0.
				//in loopje m en volume berekenen
				for j := range nodelist[i].lijst {
					volume = 4. / 3. * math.Pi * cube(nodelist[i].lijst[i].r)
					m[0] += nodelist[i].lijst[j].m[0] * nodelist[i].lijst[j].msat * volume
					m[1] += nodelist[i].lijst[j].m[1] * nodelist[i].lijst[j].msat * volume
					m[2] += nodelist[i].lijst[j].m[2] * nodelist[i].lijst[j].msat * volume
				}
				r2 := r * r
				r3 := r * r2
				r5 := r3 * r2
				dotproduct := m.dot(r_vect)

				demag[0] += prefactor * ((3 * dotproduct * r_vect[0] / r5) - (m[0] / r3))

				demag[1] += prefactor * ((3 * dotproduct * r_vect[1] / r5) - (m[1] / r3))

				demag[2] += prefactor * ((3 * dotproduct * r_vect[2] / r5) - (m[2] / r3))

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
