package vinamax

import (
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
			solver.bt[0][0] = 1. / 5
			solver.bt[1] = []float64{3. / 40, 9. / 40.}
			solver.bt[2] = []float64{44. / 45, -56. / 15., 32. / 9.}
			solver.bt[3] = []float64{19372. / 6561., -25360. / 2187., 64448. / 6561., -212. / 729.}
			solver.bt[4] = []float64{9017. / 3168., -355. / 33., 46732. / 5247., 49. / 176., -5103. / 18656.}
			solver.bt[5] = []float64{35. / 384., 0., 500. / 1113., 125. / 192., -2187. / 6784., 11. / 84.}
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
		p.u = p.previousu
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

				if Adaptivestep && T.value < j+time {
					if totalErr > Errortolerance {
						undobadstep()
						if Dt.value == MinDt.value {
							log.Fatal("Mindt is too small for your specified error tolerance")
						}
					}

					Dt.value = 0.95 * Dt.value * math.Pow(Errortolerance/totalErr, (1./float64(solver.order)))

					if Dt.value < MinDt.value {
						Dt.value = MinDt.value
					}
					if Dt.value > MaxDt.value {
						Dt.value = MaxDt.value
					}
				}
				if Adaptivestep && T.value > j+time {
					undobadstep()
					Dt.value = j + time - T.value
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
		if BrownianRotation {
			p.setRotThermField()
			tau_u := p.tau_u()

			for q := 0; q < 3; q++ {
				p.u[q] += tau_u[q] * Dt.value
			}
			p.u = norm(p.u)
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
		p.tempu = p.u
		p.previousu = p.u
		p.setThermField()
		p.setRotThermField()
	}
	totalErr = 0.
	magTorque = 0.
	rotTorque = 0.

	//actual solver

	for q := 0; q < len(solver.tt)-1; q++ {
		T.value += solver.tt[q] * Dt.value
		if Demag {
			calculatedemag()
		}
		for _, p := range lijst {
			if MagDynamics {
				p.k[q] = p.tau()

				p.m = p.tempm
				for i := 0; i < 3; i++ {
					p.torque[i] = 0.
					for r := 0; r <= q; r++ {
						p.torque[i] += (solver.bt[q][r] * p.k[r][i] * Dt.value)
					}
					p.m[i] += p.torque[i]
				}
				p.m = norm(p.m)
			}

			if BrownianRotation {
				p.k_u[q] = p.tau_u()
				p.u = p.tempu
				for i := 0; i < 3; i++ {
					p.rotTorque[i] = 0.
					for r := 0; r <= q; r++ {
						p.rotTorque[i] += (solver.bt[q][r] * p.k_u[r][i] * Dt.value)
					}
					p.u[i] += p.rotTorque[i]
				}
				p.u = norm(p.u)
				if !MagDynamics {
					p.m = p.u
				}
			}

		}
		T.value -= solver.tt[q] * Dt.value
	}

	if MagDynamics {
		for _, p := range lijst {
			if magTorque < size(p.torque) {
				magTorque = size(p.torque)
			}
		}
	}

	if BrownianRotation {
		for _, p := range lijst {
			if rotTorque < size(p.rotTorque) {
				rotTorque = size(p.rotTorque)
			}
		}
	}

	if MagDynamics {
		for _, p := range lijst {
			p.k[6] = p.tau()
			temptorquex := ((5179/57600.*p.k[0][0] + 0.*p.k[1][0] + 7571/16695.*p.k[2][0] + 393/640.*p.k[3][0] - 92097/339200.*p.k[4][0] + 187/2100.*p.k[5][0] + 1/40.*p.k[6][0]) * Dt.value)
			temptorquey := ((5179/57600.*p.k[0][1] + 0.*p.k[1][1] + 7571/16695.*p.k[2][1] + 393/640.*p.k[3][1] - 92097/339200.*p.k[4][1] + 187/2100.*p.k[5][1] + 1/40.*p.k[6][1]) * Dt.value)
			temptorquez := ((5179/57600.*p.k[0][2] + 0.*p.k[1][2] + 7571/16695.*p.k[2][2] + 393/640.*p.k[3][2] - 92097/339200.*p.k[4][2] + 187/2100.*p.k[5][2] + 1/40.*p.k[6][2]) * Dt.value)
			p.tempm[0] += temptorquex
			p.tempm[1] += temptorquey
			p.tempm[2] += temptorquez
			p.tempm = norm(p.tempm)
			diff := math.Sqrt(sqr(p.torque[0]-temptorquex) + sqr(p.torque[1]-temptorquey) + sqr(p.torque[2]-temptorquez))
			if diff > totalErr {
				totalErr = diff
			}
		}
	}

	if BrownianRotation {

		for _, p := range lijst {
			p.k_u[6] = p.tau_u()
			temptorquex := ((5179/57600.*p.k_u[0][0] + 0.*p.k_u[1][0] + 7571/16695.*p.k_u[2][0] + 393/640.*p.k_u[3][0] - 92097/339200.*p.k_u[4][0] + 187/2100.*p.k_u[5][0] + 1/40.*p.k_u[6][0]) * Dt.value)
			temptorquey := ((5179/57600.*p.k_u[0][1] + 0.*p.k_u[1][1] + 7571/16695.*p.k_u[2][1] + 393/640.*p.k_u[3][1] - 92097/339200.*p.k_u[4][1] + 187/2100.*p.k_u[5][1] + 1/40.*p.k_u[6][1]) * Dt.value)
			temptorquez := ((5179/57600.*p.k_u[0][2] + 0.*p.k_u[1][2] + 7571/16695.*p.k_u[2][2] + 393/640.*p.k_u[3][2] - 92097/339200.*p.k_u[4][2] + 187/2100.*p.k_u[5][2] + 1/40.*p.k_u[6][2]) * Dt.value)
			p.tempu[0] += temptorquex
			p.tempu[1] += temptorquey
			p.tempu[2] += temptorquez
			p.tempu = norm(p.tempu)
			diff := math.Sqrt(sqr(p.rotTorque[0]-temptorquex) + sqr(p.rotTorque[1]-temptorquey) + sqr(p.rotTorque[2]-temptorquez))
			if diff > totalErr {
				totalErr = diff
			}
			if !MagDynamics {
				p.m = p.u
			}

		}

	}

}
