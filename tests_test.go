package vinamax 

import (
	"testing"
	"math"
)
func Test_damping(t *testing.T) {

	Lijst.Append(Particle{M: Vector{1., 0., 0}})

	B_ext = Vector{0, 0, 1.00}
	Dt = 1e-15
	T = 0.
	Temp = 0.
	Alpha = 1.0

	Run(1.e-9)
 if (math.Abs(Lijst[0].M[0]+Lijst[0].M[1]) > 1e-12||Lijst[0].M[2]!=1) { 
        t.Error("damping did not work as expected") 
    }

}
