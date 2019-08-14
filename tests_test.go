package vinamax

import (
	"fmt"
	"math"
	"testing"
)

func TestDamping(test *testing.T) {

	World(0, 0, 0, 1e-8)
	Particle_radius(10e-9)
	Addsingleparticle(0, 0, 0)
	Msat(400e3)
	M_uniform(1, 0, 0)

	B_ext = func(t float64) (float64, float64, float64) { return 0.1, 0., 1. }
	Dt.value = 1e-15
	T.value = 0.
	Temp.value = 0.
	Alpha.value = 1.0
	//Relax()
	Run(1.e-9)
	fmt.Println(Universe.lijst[0].m[0], Universe.lijst[0].m[1], Universe.lijst[0].m[2])
	if math.Abs(Universe.lijst[0].m[0]+Universe.lijst[0].m[1]) > 1e-12 || Universe.lijst[0].m[2] != 1 {
		test.Error("relax did not work as expected")
	}
}
