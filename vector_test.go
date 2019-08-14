package vinamax

import (
	//"math"
	"testing"
)

func TestVector(t *testing.T) {
	x := vector{1., 2., 3.}
	normed := norm(x)
	if size(normed) != 1. {
		t.Error("problem in vector norm or size")
	}
	y := vector{0., 2., 0.}
	dotproduct := x.dot(y)
	if dotproduct != 4. {
		t.Error("vector dotproduct problem")
	}
	z := vector{3., 0., 1.}
	y = y.add(z)
	if y[0] != 3 || y[1] != 2 || y[2] != 1 {
		t.Error("vector add problem")
	}
	crossproduct := x.cross(y)
	if crossproduct[0] != -4 || crossproduct[1] != 8 || crossproduct[2] != -4 {
		t.Error("vector crossproduct problem")
	}
	inverse := crossproduct.times(-1)

	crossproduct.directadd(inverse)
	if size(crossproduct) != 0. {
		t.Error("vector directadd problem")
	}

	if cube(-2.) != -8 {
		t.Error("float64 cube problem")
	}
	if sqr(-2.) != 4 {
		t.Error("float64 sqr problem")
	}
}
