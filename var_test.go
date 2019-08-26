package vinamax

import (
	"math"
	"testing"
)

func TestVolume(t *testing.T) {
	radius := math.Pow(3./4./math.Pi, 1./3.)
	if math.Abs(volume(radius)-1.) > 1e-13 {
		t.Error("volume function gives wrong result")
	}

}
