package vinamax

import (
	//	"fmt"
	"log"
	"math"
)

func Relax() {
	backuptol := Errortolerance
	backuptime := T
	backupdt := Dt

	gammaoveralpha = gamma0 / (1. + (Alpha * Alpha))
	relax = true
	if Demag {
		calculatedemag()
	}
	dopristep(Universe.lijst)
	Errortolerance = 1e-3
	for maxtauwitht > 1e-10 {

		if Demag {
			calculatedemag()
		}
		dopristep(Universe.lijst)

		if maxtauwitht > Errortolerance {
			undobadstep(Universe.lijst)
			if BrownianRotation {
				undobadstep_u_anis(Universe.lijst)
			}
			if Dt == Mindt {
				log.Fatal("mindt is too small for your specified error tolerance")
			}
		}

		Dt = 0.95 * Dt * math.Pow(Errortolerance/maxtauwitht, (1./float64(order)))

		if Dt < Mindt {
			Dt = Mindt
		}
		if Dt > Maxdt {
			Dt = Maxdt
		}
		if maxtauwitht < Errortolerance/16 {
			Errortolerance /= 1.44
		}
		T = backuptime
	}

	Errortolerance = backuptol
	T = backuptime
	Dt = backupdt
	relax = false

}
