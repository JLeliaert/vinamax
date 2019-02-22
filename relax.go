package vinamax

import (
	"fmt"
	"log"
	"math"
)



func Relax() {

				relax=true
if Demag {
					calculatedemag()
				}
				dopristep(Universe.lijst)


		for maxtauwitht>1e-9{
		fmt.Println(maxtauwitht)

				if Demag {
					calculatedemag()
				}
				dopristep(Universe.lijst)

				//fmt.Println(Dt)
				if Adaptivestep {
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
					//fmt.Println("dt:   ", Dt)
					//maxtauwitht = 1.e-12
				}
	}


	relax=false

}
