package vinamax

import (
	"math"
)

func calculatedemag() {
	if FMM {
		universe.calculatem()
	}

	for i := range universe.lijst {
		universe.lijst[i].demagnetising_field = universe.lijst[i].demag()
	}
}

//Demag is calculated on a position
func demag(x, y, z float64) vector {
	prefactor := mu0 / (4 * math.Pi)
	demag := vector{0, 0, 0}

	for i := range universe.lijst {
		if universe.lijst[i].x != x || universe.lijst[i].y != y || universe.lijst[i].z != z {
			radius := *universe.lijst[i].rindex
			volume := 4. / 3 * math.Pi * cube(radius)
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

//Demag on a particle
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

//Dipole approximation Demag is calculated on a position
func fMMdemag(x, y, z float64) vector {

	prefactor := mu0 / (4 * math.Pi)
	demag := vector{0, 0, 0}
	//make list with nodes
	//put node universe in box
	nodelist := []*node{&universe}
	//for lijst!=empty
	for len(nodelist) > 0 {
		i := 0
		if nodelist[i].number == 1 {
			//if numberofparticles in box==1:
			if nodelist[i].lijst[0].x != x || nodelist[i].lijst[0].y != y || nodelist[i].lijst[0].z != z {
				//	if i'm not the one: calculate and delete from stack

				volume := nodelist[i].volume

				r_vect := vector{x - nodelist[i].lijst[0].x, y - nodelist[i].lijst[0].y, z - nodelist[i].lijst[0].z}
				r := nodelist[i].lijst[0].dist(x, y, z)
				r2 := r * r
				r3 := r * r2
				r5 := r3 * r2

				dotproduct := nodelist[i].lijst[0].m.dot(r_vect)

				demag[0] += nodelist[i].lijst[0].msat * volume * prefactor * ((3 * dotproduct * r_vect[0] / r5) - (nodelist[i].lijst[0].m[0] / r3))

				demag[1] += nodelist[i].lijst[0].msat * volume * prefactor * ((3 * dotproduct * r_vect[1] / r5) - (nodelist[i].lijst[0].m[1] / r3))

				demag[2] += nodelist[i].lijst[0].msat * volume * prefactor * ((3 * dotproduct * r_vect[2] / r5) - (nodelist[i].lijst[0].m[2] / r3))
			}
		}
		if nodelist[i].number > 1 {
			//if number of particles in box>1:
			r_vect := vector{x - nodelist[i].com[0], y - nodelist[i].com[1], z - nodelist[i].com[2]}
			r := math.Sqrt(r_vect[0]*r_vect[0] + r_vect[1]*r_vect[1] + r_vect[2]*r_vect[2])

			if (nodelist[i].where(vector{x, y, z}) == -1 && (math.Sqrt(2)/2.*nodelist[i].diameter/r) < Thresholdbeta) {
				//	if criterium is ok: calculate and delete from stack

				m := nodelist[i].m

				r2 := r * r
				r3 := r * r2
				r5 := r3 * r2
				dotproduct := m.dot(r_vect)

				demag[0] += prefactor * ((3 * dotproduct * r_vect[0] / r5) - (m[0] / r3))

				demag[1] += prefactor * ((3 * dotproduct * r_vect[1] / r5) - (m[1] / r3))

				demag[2] += prefactor * ((3 * dotproduct * r_vect[2] / r5) - (m[2] / r3))

			} else {
				//	if not: add subboxes andn delete from stack
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
		nodelist[i], nodelist = nodelist[len(nodelist)-1], nodelist[:len(nodelist)-1]
	}
	return demag
}
