package vinamax

import (
	"math"
)

func calculatedemag() {
	if FMM {
		Universe.calculatem()
	}

	for i := range Universe.lijst {
		Universe.lijst[i].demagnetising_field = Universe.lijst[i].demag()
	}
}

//Demag is calculated on a position
func demag(x, y, z float64) vector {
	prefactor := mu0 / (4 * math.Pi)
	demag := vector{0, 0, 0}

	for i := range Universe.lijst {
		if Universe.lijst[i].x != x || Universe.lijst[i].y != y || Universe.lijst[i].z != z {
			radius := Universe.lijst[i].r
			volume := 4. / 3 * math.Pi * cube(radius)
			r_vect := vector{x - Universe.lijst[i].x, y - Universe.lijst[i].y, z - Universe.lijst[i].z}
			r := Universe.lijst[i].dist(x, y, z)
			r2 := r * r
			r3 := r * r2
			r5 := r3 * r2

			dotproduct := Universe.lijst[i].m.dot(r_vect)

			demag[0] += Universe.lijst[i].msat * volume * prefactor * ((3 * dotproduct * r_vect[0] / r5) - (Universe.lijst[i].m[0] / r3))

			demag[1] += Universe.lijst[i].msat * volume * prefactor * ((3. * dotproduct * r_vect[1] / r5) - (Universe.lijst[i].m[1] / r3))

			demag[2] += Universe.lijst[i].msat * volume * prefactor * ((3 * dotproduct * r_vect[2] / r5) - (Universe.lijst[i].m[2] / r3))

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
	//put node Universe in box
	nodelist := []*node{&Universe}
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
