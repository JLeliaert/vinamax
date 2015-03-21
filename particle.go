package vinamax

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

//A particle essentially constains a position, magnetisation
type particle struct {
	x, y, z             float64
	m                   vector
	demagnetising_field vector
	u_anis              vector  // Uniaxial anisotropy axis
	c1_anis             vector  // cubic anisotropy axis
	c2_anis             vector  // cubic anisotropy axis
	c3_anis             vector  // cubic anisotropy axis
	r                   float64 // radius
	msat                float64 // Saturation magnetisation in A/m
	flip                float64 // time of next flip event
	tempnumber          float64

	heff      vector //effective field
	tempfield vector
	tempm     vector
	previousm vector
	fehlk1    vector
	fehlk2    vector
	fehlk3    vector
	fehlk4    vector
	fehlk5    vector
	fehlk6    vector
	fehlk7    vector
	fehlk8    vector
	fehlk9    vector
	fehlk10   vector
	fehlk11   vector
	fehlk12   vector
	fehlk13   vector
}

//print position and magnitisation of a particle
func (p particle) string() string {
	return fmt.Sprintf("particle@(%v, %v, %v), %v %v %v", p.x, p.y, p.z, p.m[0], p.m[1], p.m[2])
}

//Gives all particles the same specified uniaxialanisotropy-axis
func Anisotropy_axis(x, y, z float64) {
	uaniscalled = true
	a := norm(vector{x, y, z})
	for i := range universe.lijst {
		universe.lijst[i].u_anis = a
	}
}

//Gives all particles the same specified cubic1anisotropy-axis
func C1anisotropy_axis(x, y, z float64) {
	c1called = true
	a := norm(vector{x, y, z})
	for i := range universe.lijst {
		universe.lijst[i].c1_anis = a
	}
}

//Gives all particles the same specified cubic2anisotropy-axis, must be orthogonal to c1
func C2anisotropy_axis(x, y, z float64) {
	c2called = true
	a := norm(vector{x, y, z})
	for i := range universe.lijst {
		if universe.lijst[i].c1_anis.dot(a) == 0 {
			universe.lijst[i].c2_anis = a
			universe.lijst[i].c3_anis = norm(universe.lijst[i].c1_anis.cross(a))
		} else {
			log.Fatal("c1 and c2 should be orthogonal")
		}
	}
}

//Gives all particles a random anisotropy-axis
func Anisotropy_random() {
	uaniscalled = true
	for i := range universe.lijst {
		phi := rng.Float64() * (2 * math.Pi)
		theta := 2 * math.Asin(math.Sqrt(rng.Float64()))
		universe.lijst[i].u_anis = vector{math.Sin(theta) * math.Cos(phi), math.Sin(theta) * math.Sin(phi), math.Cos(theta)}
		if math.Cos(theta) < 0. {
			universe.lijst[i].u_anis = universe.lijst[i].u_anis.times(-1.)
		}
	}
}

//Gives all particles with random magnetisation orientation
func M_random() {
	magnetisationcalled = true
	for i := range universe.lijst {
		phi := rng.Float64() * (2 * math.Pi)
		theta := 2 * math.Asin(math.Sqrt(rng.Float64()))
		universe.lijst[i].m = vector{math.Sin(theta) * math.Cos(phi), math.Sin(theta) * math.Sin(phi), math.Cos(theta)}
	}
}

//Gives all particles magnetisation specified by the moment superposition model
func M_MSM(tmag, field float64) {
	r := rand.New(rand.NewSource(99))
	magnetisationcalled = true
	for i := range universe.lijst {
		volume := cube(universe.lijst[i].r) * 4. / 3. * math.Pi
		gprime := Alpha * gamma0 * mu0 / (1. + (Alpha * Alpha))
		delta := Ku1 * volume / (kb * Temp)
		msat := universe.lijst[i].msat
		hk := 2. * Ku1 / (msat * mu0)
		tau0 := gprime * hk * math.Sqrt(delta/math.Pi)
		tauN := 1. / tau0 * math.Exp(Ku1*volume/(kb*Temp)*(1.-0.82*msat*field*mu0/Ku1))
		x := volume * field * msat * mu0 / (kb * Temp)

		langevin := 1./math.Tanh(x) - 1./x

		M := langevin * (1. - math.Exp(-tmag/tauN))
		up := (2.*M + 1.) / (2.) //2.M because of random anisotropy axes
		if r.Float64() < up {
			universe.lijst[i].m = universe.lijst[i].u_anis
		} else {
			universe.lijst[i].m = universe.lijst[i].u_anis.times(-1.)
		}
	}
}

//Gives all particles a specified magnetisation direction
func M_uniform(x, y, z float64) {
	magnetisationcalled = true
	a := norm(vector{x, y, z})
	for i := range universe.lijst {
		universe.lijst[i].m = a
	}
}

//Sets the saturation magnetisation of all particles in A/m
func Msat(x float64) {
	msatcalled = true
	for i := range universe.lijst {
		universe.lijst[i].msat = x
	}
}
