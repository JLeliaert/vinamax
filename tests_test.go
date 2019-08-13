package vinamax

import (
	"math"
	"testing"
)

func TestDamping(t *testing.T) {

	World(0, 0, 0, 1e-8)
	Particle_radius(10e-9)
	Addsingleparticle(0, 0, 0)
	Msat(400e3)
	M_uniform(1, 0, 0)

	B_ext = func(t float64) (float64, float64, float64) { return 0., 0., 1. }
	Dt = 1e-15
	T = 0.
	Temp = 0.
	Alpha = 1.0

	Run(1.e-9)
	if math.Abs(Universe.lijst[0].m[0]+Universe.lijst[0].m[1]) > 1e-12 || Universe.lijst[0].m[2] != 1 {
		t.Error("damping did not work as expected")
	}
}
