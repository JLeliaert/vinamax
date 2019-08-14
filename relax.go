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
	dopristep(Universe.lijst)
	Errortolerance = 1e-1
	for maxtauwitht > 5e-8 {

		if Demag {
			calculatedemag()
		}
		dopristep(Universe.lijst)

		if maxtauwitht > Errortolerance {
			undobadstep(Universe.lijst)
			if Dt.value == MinDt.value {
				log.Fatal("Mindt is too small for your specified error tolerance")
			}
		}

		Dt.value = 0.95 * Dt.value * math.Pow(Errortolerance/maxtauwitht, (1./float64(solver.order)))

		if Dt.value < MinDt.value {
			Dt.value = MinDt.value
		}
		if Dt.value > MaxDt.value {
			Dt.value = MaxDt.value
		}
		//fmt.Println(maxtauwitht,"\t",Errortolerance)
		if maxtauwitht < Errortolerance/4 {
			Errortolerance /= 1.4142
		}
		T.value = backuptime
	}

	Errortolerance = backuptol
	T.value = backuptime
	Dt.value = backupdt
	relax = false

}
