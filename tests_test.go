package vinamax

import (
	//	"fmt"
	"math"
	"testing"
)

func Test_damping(t *testing.T) {

	Lijst.Append(Particle{M: Vector{1., 0., 0}})

	B_ext = Vector{0, 0, 1.00}
	Dt = 1e-15
	T = 0.
	Temp = 0.
	Alpha = 1.0

	Run(1.e-9)
	if math.Abs(Lijst[0].M[0]+Lijst[0].M[1]) > 1e-12 || Lijst[0].M[2] != 1 {
		t.Error("damping did not work as expected")
	}
}

func Test_precession(t *testing.T) {
	Lijst = nil
	Lijst.Append(Particle{M: Vector{1., 0., 0}})

	B_ext = Vector{0, 0, 0.1}
	Dt = 1e-15
	T = 0.
	Temp = 0.
	Alpha = 0.0

	Run(1e-9)
	if (math.Abs(Lijst[0].M[0]-0.31) > 1e-3) || Lijst[0].M[2] != 0 {
		t.Error("precession did not work as expected")
	}
}

func Test_demag(t *testing.T) {
	Lijst = nil
	Lijst.Append(Particle{X: -2e-9, M: [3]float64{1., 0., 0}})
	Lijst.Append(Particle{X: 2e-9, M: [3]float64{0, 1., 0}})

	B_ext = Vector{0, 0, 0.01}
	Dt = 1e-14
	T = 0.
	Temp = 0.
	Alpha = 0.01
	Ku1 = 0 //10 000

	//Anisotropy_axis(0, 1, 0)
	Anisotropy_random()

	Run(3.e-9)

	if math.Abs(0.5*(Lijst[0].M[0]+Lijst[1].M[0])-0.29699) > 1e-3 || math.Abs(0.5*(Lijst[0].M[1]+Lijst[1].M[1])+0.35336) > 1e-3 || math.Abs(0.5*(Lijst[0].M[2]+Lijst[1].M[2])-0.55416) > 1e-3 {
		t.Error("demag did not work as expected")
	}
}
