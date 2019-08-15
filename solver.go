package vinamax

import (
	//"fmt"
	"log"
	"math"
)

type solvertype struct {
	name  string
	order int
	bt    [][]float64
	tt    []float64
}

//Set the solver to use, "euler" or "heun"
func SetSolver(a string) {
	switch a {

	case "euler":
		{
			solver.name = "euler"
			solver.order = 1
		}
	case "dopri":
		{
			solver.name = "dopri"
			solver.order = 5
			solver.tt = []float64{0., 1. / 5., 3. / 10., 4. / 5., 8. / 9., 1., 1.}

			for i := 0; i < 7; i++ {
				b := make([]float64, 7)
				solver.bt = append(solver.bt, b)
			}
			solver.bt[1][0] = 1. / 5
			solver.bt[2] = []float64{3. / 40, 9. / 40.}
			solver.bt[3] = []float64{44. / 45, -56. / 15., 32. / 9.}
			solver.bt[4] = []float64{19372. / 6561., -25360. / 2187., 64448. / 6561., -212. / 729.}
			solver.bt[5] = []float64{9017. / 3168., -355. / 33., 46732. / 5247., 49. / 176., -5103. / 18656.}
			solver.bt[6] = []float64{35. / 384., 0., 500. / 1113., 125. / 192., -2187. / 6784., 11. / 84.}
		}
	default:
		{
			log.Fatal(a, " is not a possible solver, \"euler\" or \"dopri\"")
		}
	}
}

func undobadstep() {
	for _, p := range lijst {
		p.m = p.previousm
		p.u_anis = p.previousu_anis
	}
	T.value -= Dt.value
}

//Runs the simulation for a certain time
func Run(time float64) {
	//update prefactors
	setThermPrefac()

	testinput()
	syntaxrun()
	for _, p := range lijst {
		norm(p.m)
	}

	for j := T.value; T.value < j+time; {
		if Demag {
			calculatedemag()
		}
		switch solver.name {
		case "euler":
			{
				eulerstep()
				T.value += Dt.value
			}
		case "dopri":
			{
				dopristep()
				T.value += Dt.value

				if Adaptivestep {
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
				}
			}
		case "":
			{
				log.Fatal("solver not set")
			}
		}

		//averages is not weighted with volume, averagemoments is
		write(averagemoments(), false)
	}
}

//##################################################

//Perform a timestep using euler forward method
func eulerstep() {
	for _, p := range lijst {
		p.setThermField()
		tau := p.tau()

		for q := 0; q < 3; q++ {
			p.m[q] += tau[q] * Dt.value
		}
		p.m = norm(p.m)
		//only calculate anisodynamics when necessary
		if BrownianRotation {
			randomv := p.randomv()
			p.randomvfield = randomv
			tau_u := p.tau_u(randomv)

			for q := 0; q < 3; q++ {
				p.u_anis[q] += tau_u[q] * Dt.value
			}
			p.u_anis = norm(p.u_anis)
		}
	}
}

//#########################################################################
//perform a timestep using dormand-prince

//include FSAL but only at zero temperature
func dopristep() {

	//preparations
	for _, p := range lijst {
		p.tempm = p.m
		p.previousm = p.m
		p.tempu_anis = p.u_anis
		p.previousu_anis = p.u_anis
		p.setThermField()
		p.randomvfield = p.randomv()
	}
	magErr = 0.
	magTorque = 0.

	//actual solver

	for q := 0; q < len(solver.tt); q++ {
		T.value += solver.tt[q] * Dt.value
		if Demag {
			calculatedemag()
		}
		for _, p := range lijst {
			p.k[q] = p.tau()

			p.m = p.tempm
			for i := 0; i < 3; i++ {
				p.torque[i] = 0.
				for r := 0; r < q; r++ {
					p.torque[i] += (solver.bt[q][r] * p.k[r][i] * Dt.value)
				}
				p.m[i] += p.torque[i]
			}
			p.m = norm(p.m)

			//TODO include brownian dynamics

		}
		T.value -= solver.tt[q] * Dt.value
	}

	for _, p := range lijst {
		if magTorque < size(p.torque) {
			magTorque = size(p.torque)
		}
	}

	for _, p := range lijst {
		temptorquex := ((5179/57600.*p.k[0][0] + 0.*p.k[1][0] + 7571/16695.*p.k[2][0] + 393/640.*p.k[3][0] - 92097/339200.*p.k[4][0] + 187/2100.*p.k[5][0] + 1/40.*p.k[6][0]) * Dt.value)
		temptorquey := ((5179/57600.*p.k[0][1] + 0.*p.k[1][1] + 7571/16695.*p.k[2][1] + 393/640.*p.k[3][1] - 92097/339200.*p.k[4][1] + 187/2100.*p.k[5][1] + 1/40.*p.k[6][1]) * Dt.value)
		temptorquez := ((5179/57600.*p.k[0][2] + 0.*p.k[1][2] + 7571/16695.*p.k[2][2] + 393/640.*p.k[3][2] - 92097/339200.*p.k[4][2] + 187/2100.*p.k[5][2] + 1/40.*p.k[6][2]) * Dt.value)
		p.tempm[0] += temptorquex
		p.tempm[1] += temptorquey
		p.tempm[2] += temptorquez
		p.tempm = norm(p.tempm)
		diff := math.Sqrt(sqr(p.torque[0]-temptorquex) + sqr(p.torque[1]-temptorquey) + sqr(p.torque[2]-temptorquez))
		if diff > magErr {
			magErr = diff
		}
	}

	//and this is also the fourth order solution
	//the error is the difference between the two solutions

	//		//fmt.Println("error    :", error)

}

//TODO error control and adaptive timestep

// THE SOLUTIONS

//		if BrownianRotation { //only calculate anisodynamics when requested
//			randomv := p.randomvfield
//			for r := 0; r < 7; r++ {
//				k_u[r] = p.fehlk_u[r]
//			}
//			k_u[6] = p.tau_u(randomv)
//			p.fehlk_u[6] = k_u[6]

//			p.tempu_anis[0] += ((5179/57600.*k_u[0][0] + 0.*k_u[1][0] + 7571/16695.*k_u[2][0] + 393/640.*k_u[3][0] - 92097/339200.*k_u[4][0] + 187/2100.*k_u[5][0] + 1/40.*k_u[6][0]) * Dt.value)
//			p.tempu_anis[1] += ((5179/57600.*k_u[0][1] + 0.*k_u[1][1] + 7571/16695.*k_u[2][1] + 393/640.*k_u[3][1] - 92097/339200.*k_u[4][1] + 187/2100.*k_u[5][1] + 1/40.*k_u[6][1]) * Dt.value)
//			p.tempu_anis[2] += ((5179/57600.*k_u[0][2] + 0.*k_u[1][2] + 7571/16695.*k_u[2][2] + 393/640.*k_u[3][2] - 92097/339200.*k_u[4][2] + 187/2100.*k_u[5][2] + 1/40.*k_u[6][2]) * Dt.value)
//			//and this is also the fourth order solution
//		}

//		// ERROR CONTROL

//		if BrownianRotation { //only calculate anisodynamics when requested
//			//the error is the difference between the two solutions
//			error := math.Sqrt(sqr(p.u_anis[0]-p.tempu_anis[0]) + sqr(p.u_anis[1]-p.tempu_anis[1]) + sqr(p.u_anis[2]-p.tempu_anis[2]))

//			//fmt.Println("error    :", error)
//			if Adaptivestep { // && relax == false {
//				if error > maxtauwitht { //in LLG dynamics already set to maxtauwitht if error is larger
//					maxtauwitht = error
//				}

//			}
//		}

//	}
//}
//T.value += Dt.value
