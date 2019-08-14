package vinamax

import (
	"fmt"
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

//Runs the simulation for a certain time

func Run(time float64) {
	gammaoveralpha = gamma0 / (1. + (sqr(Alpha.value)))
	testinput()
	syntaxrun()
	for i := range Universe.lijst {
		norm(Universe.lijst[i].m)
	}

	for j := T.value; T.value < j+time; {
		if Demag {
			calculatedemag()
		}
		switch solver.name {
		case "euler":
			{
				eulerstep(Universe.lijst)
				T.value += Dt.value
			}
		case "dopri":
			{
				dopristep(Universe.lijst)
				T.value += Dt.value

				if Adaptivestep {
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
					// TODO WHY IS THIS HERE???
					//	if relax == false {
					//		maxtauwitht = 1.e-12
					//	}
				}
			}
		}

		//averages is not weighted with volume, averagemoments is
		write(averagemoments(Universe.lijst), false)
	}
}

//##################################################

//Perform a timestep using euler forward method
func eulerstep(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()
		tau := p.tau(temp)

		if p.fixed == false {
			for q := 0; q < 3; q++ {
				p.m[q] += tau[q] * Dt.value
			}
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

//TODO use FSAL but only at zero temperature
func dopristep(Lijst []*particle) {
	var k [7]vector
	var k_u [7]vector

	//preparations
	for _, p := range Lijst {
		p.tempm = p.m
		p.previousm = p.m
		p.tempu_anis = p.u_anis
		p.previousu_anis = p.u_anis
		p.tempfield = p.temp()
		p.randomvfield = p.randomv()
	}

	for q := 0; q < len(solver.tt); q++ {
		T.value += solver.tt[q] * Dt.value
		if Demag {
			calculatedemag()
		}
		for _, p := range Lijst {
			temp := p.tempfield
			for r := 0; r < q; r++ {
				k[r] = p.fehlk[r]
			}
			if relax == false {
				k[q] = p.tau(temp)
			}
			if relax == true {
				k[q] = p.noprecess()
			}
			p.fehlk[q] = k[q]

			p.m = p.tempm
			if p.fixed == false {
				for r := 0; r < q; r++ {
					p.m[0] += (solver.bt[q][r] * k[r][0] * Dt.value)
					p.m[1] += (solver.bt[q][r] * k[r][1] * Dt.value)
					p.m[2] += (solver.bt[q][r] * k[r][1] * Dt.value)
				}
				p.m = norm(p.m)
				//p.tempm = norm(p.tempm)
			}

			//only calculate anisodynamics when necessary
			if BrownianRotation {
				randomv := p.randomvfield
				for r := 0; r < q; r++ {
					k_u[r] = p.fehlk_u[r]
				}
				k_u[q] = p.tau_u(randomv)
				p.fehlk_u[q] = k_u[q]

				p.u_anis = p.tempu_anis
				for r := 0; r < q; r++ {
					p.u_anis[0] += (solver.bt[q][r] * k_u[r][0] * Dt.value)
					p.u_anis[1] += (solver.bt[q][r] * k_u[r][1] * Dt.value)
					p.u_anis[2] += (solver.bt[q][r] * k_u[r][2] * Dt.value)
				}
				p.u_anis = norm(p.u_anis)
				//p.tempu_anis = norm(p.tempu_anis)

			}

		}
		T.value -= solver.tt[q] * Dt.value
	}

	// THE SOLUTIONS

	var torquex, torquey, torquez, temptorquex, temptorquey, temptorquez float64
	for _, p := range Lijst {
		p.m = p.tempm
		if p.fixed == false {
			torquex = ((solver.bt[6][0]*k[0][0] + solver.bt[6][1]*k[1][0] + solver.bt[6][2]*k[2][0] + solver.bt[6][3]*k[3][0] + solver.bt[6][4]*k[4][0] + solver.bt[6][5]*k[5][0]) * Dt.value)
			torquey = ((solver.bt[6][0]*k[0][1] + solver.bt[6][1]*k[1][1] + solver.bt[6][2]*k[2][1] + solver.bt[6][3]*k[3][1] + solver.bt[6][4]*k[4][1] + solver.bt[6][5]*k[5][1]) * Dt.value)
			torquez = ((solver.bt[6][0]*k[0][2] + solver.bt[6][1]*k[1][2] + solver.bt[6][2]*k[2][2] + solver.bt[6][3]*k[3][2] + solver.bt[6][4]*k[4][2] + solver.bt[6][5]*k[5][2]) * Dt.value)
			p.m[0] += torquex
			p.m[1] += torquey
			p.m[2] += torquez
			p.m = norm(p.m)

			//and this is also the fifth order solution

			//if relax == true {
			maxtauwitht = math.Sqrt(math.Pow(torquex, 2.) + math.Pow(torquey, 2.) + math.Pow(torquez, 2.))
			//	}
		}

		if Demag {
			calculatedemag()
		}

		for _, p := range Lijst {
			temp := p.tempfield
			for r := 0; r < 7; r++ {
				k[r] = p.fehlk[r]
			}
			if relax == false {
				k[6] = p.tau(temp)
			}
			if relax == true {
				k[6] = p.noprecess()
			}
			p.fehlk[6] = k[6]

			if p.fixed == false {
				temptorquex = ((5179/57600.*k[0][0] + 0.*k[1][0] + 7571/16695.*k[2][0] + 393/640.*k[3][0] - 92097/339200.*k[4][0] + 187/2100.*k[5][0] + 1/40.*k[6][0]) * Dt.value)
				temptorquey = ((5179/57600.*k[0][1] + 0.*k[1][1] + 7571/16695.*k[2][1] + 393/640.*k[3][1] - 92097/339200.*k[4][1] + 187/2100.*k[5][1] + 1/40.*k[6][1]) * Dt.value)
				temptorquez = ((5179/57600.*k[0][2] + 0.*k[1][2] + 7571/16695.*k[2][2] + 393/640.*k[3][2] - 92097/339200.*k[4][2] + 187/2100.*k[5][2] + 1/40.*k[6][2]) * Dt.value)
				p.tempm[0] += temptorquex
				p.tempm[1] += temptorquey
				p.tempm[2] += temptorquez
				p.tempm = norm(p.tempm)
			}
			//and this is also the fourth order solution

			if BrownianRotation { //only calculate anisodynamics when requested
				randomv := p.randomvfield
				for r := 0; r < 7; r++ {
					k_u[r] = p.fehlk_u[r]
				}
				k_u[6] = p.tau_u(randomv)
				p.fehlk_u[6] = k_u[6]

				p.tempu_anis[0] += ((5179/57600.*k_u[0][0] + 0.*k_u[1][0] + 7571/16695.*k_u[2][0] + 393/640.*k_u[3][0] - 92097/339200.*k_u[4][0] + 187/2100.*k_u[5][0] + 1/40.*k_u[6][0]) * Dt.value)
				p.tempu_anis[1] += ((5179/57600.*k_u[0][1] + 0.*k_u[1][1] + 7571/16695.*k_u[2][1] + 393/640.*k_u[3][1] - 92097/339200.*k_u[4][1] + 187/2100.*k_u[5][1] + 1/40.*k_u[6][1]) * Dt.value)
				p.tempu_anis[2] += ((5179/57600.*k_u[0][2] + 0.*k_u[1][2] + 7571/16695.*k_u[2][2] + 393/640.*k_u[3][2] - 92097/339200.*k_u[4][2] + 187/2100.*k_u[5][2] + 1/40.*k_u[6][2]) * Dt.value)
				//and this is also the fourth order solution
			}

			// ERROR CONTROL

			//the error is the difference between the two solutions
			error := math.Sqrt(sqr(torquex-temptorquex) + sqr(torquey-temptorquey) + sqr(torquez-temptorquez))
			fmt.Println(error)

			//fmt.Println("error    :", error)
			if relax == false {
				maxtauwitht = error
			}

			if BrownianRotation { //only calculate anisodynamics when requested
				//the error is the difference between the two solutions
				error := math.Sqrt(sqr(p.u_anis[0]-p.tempu_anis[0]) + sqr(p.u_anis[1]-p.tempu_anis[1]) + sqr(p.u_anis[2]-p.tempu_anis[2]))

				//fmt.Println("error    :", error)
				if Adaptivestep && relax == false {
					if error > maxtauwitht { //in LLG dynamics already set to maxtauwitht if error is larger
						maxtauwitht = error
					}

				}
			}

		}
	}
}

func undobadstep(Lijst []*particle) {
	for _, p := range Lijst {
		p.m = p.previousm
		p.u_anis = p.previousu_anis
	}
	T.value -= Dt.value
}
