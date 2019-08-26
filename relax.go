package vinamax

import (
	//	"fmt"
	"log"
	"math"
)

func Relax() {
	backuptol := Errortolerance
	backuptime := T.value
	backupdt := Dt.value

	relax = true
	//minimum 5 tiny steps
	Dt.value = 1e-17
	for i := 0; i < 5; i++ {
		dopristep()
	}

	Dt.value = 1e-10
	Errortolerance = 1e-1
	for magErr > 1e-7 || magTorque > 1e-10 {
		dopristep()

		if magErr > Errortolerance {
			undobadstep()
			if Dt.value == MinDt.value {
				log.Fatal("Mindt is too small for your specified error tolerance")
			}
		}

		Dt.value = math.Min(Dt.value, 0.95*Dt.value*math.Pow(Errortolerance/magErr, (1./float64(solver.order))))

		if Dt.value < MinDt.value {
			Dt.value = MinDt.value
		}
		if Dt.value > MaxDt.value {
			Dt.value = MaxDt.value
		}
		if magErr < Errortolerance/4 {
			Errortolerance /= 1.4142
			Errortolerance = math.Max(1e-10, Errortolerance)
		}
		T.value = backuptime
	}

	Errortolerance = backuptol
	T.value = backuptime
	Dt.value = backupdt
	relax = false

}
