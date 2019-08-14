package vinamax

import (
	//"fmt"
	"log"
	"math"
)

func Relax() {
	backuptol := Errortolerance
	backuptime := T.value
	backupdt := Dt.value

	relax = true
	if Demag {
		calculatedemag()
	}
	dopristep()
	Errortolerance = 1e-1
	for magErr > 5e-8 {

		if Demag {
			calculatedemag()
		}
		dopristep()

		if magErr > Errortolerance {
			undobadstep()
			if Dt.value == MinDt.value {
				log.Fatal("Mindt is too small for your specified error tolerance")
			}
		}

		Dt.value = 0.95 * Dt.value * math.Pow(Errortolerance/magErr, (1./float64(solver.order)))

		if Dt.value < MinDt.value {
			Dt.value = MinDt.value
		}
		if Dt.value > MaxDt.value {
			Dt.value = MaxDt.value
		}
		if magErr < Errortolerance/4 {
			Errortolerance /= 1.4142
		}
		T.value = backuptime
	}

	Errortolerance = backuptol
	T.value = backuptime
	Dt.value = backupdt
	relax = false

}
