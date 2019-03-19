package vinamax

import (
//	"fmt"
	"log"
	"math"
)

//Set the solver to use, "euler" or "heun"
func Setsolver(a string) {
	switch a {

	case "euler":
		{
			solver = "euler"
			order = 1
		}
	case "heun":
		{
			solver = "heun"
			order = 2
		}
	case "rk3":
		{
			solver = "rk3"
			order = 3
		}
	case "annelies":
		{
			solver = "annelies"
			order = 3
		}

	case "rk4":
		{
			solver = "rk4"
			order = 4
		}
	case "dopri":
		{
			solver = "dopri"
			order = 5
		}
	case "fehl56":
		{
			solver = "fehl56"
			order = 6
		}
	case "fehl67":
		{
			solver = "fehl67"
			order = 7
		}
	case "time":
		{
			solver = "time"
			order = 0
		}
	default:
		{
			log.Fatal(a, " is not a possible solver, \"euler\" or \"heun\" or \"rk3\"or \"rk4\"or \"dopri\"or \"fehl56\"or \"fehl67\"")
		}
	}
}

//Runs the simulation for a certain time

func Run(time float64) {
	gammaoveralpha = gamma0 / (1. + (Alpha * Alpha))
	testinput()
	syntaxrun()
	for i := range Universe.lijst {
		norm(Universe.lijst[i].m)
	}
	write(averagemoments(Universe.lijst),false)
	//averages is not weighted with volume, averagemoments is
	//write(averages(Universe.lijst))
	previousdemagcalc := T - demagtime
	for j := T; T < j+time; {
		if (demagevery == true) && (T-previousdemagcalc >= demagtime) {
			calculatedemag()
			previousdemagcalc = T
		}
		if Demag {
			calculatedemag()
		}
		switch solver {
		case "heun":
			{
				heunstep(Universe.lijst)
				T += Dt
			}
		case "euler":
			{
				eulerstep(Universe.lijst)
				T += Dt
			}
		case "rk3":
			{
				rk3step(Universe.lijst)
				T += Dt
			}
		case "annelies":
			{
				anneliesstep(Universe.lijst)
				T += Dt
			}
		case "rk4":
			{
				rk4step(Universe.lijst)
				T += Dt
			}
		case "dopri":
			{
				dopristep(Universe.lijst)
				T += Dt
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
					if relax==false{
					maxtauwitht = 1.e-12
					}
				}
			}
		case "fehl56":
			{
				fehl56step(Universe.lijst)
				T += Dt
				if Adaptivestep {
					if maxtauwitht > Errortolerance {
						undobadstep(Universe.lijst)
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
					maxtauwitht = 1.e-12
				}
			}
		case "fehl67":
			{
				fehl67step(Universe.lijst)
				T += Dt
				if Adaptivestep {
					if maxtauwitht > Errortolerance {
						undobadstep(Universe.lijst)
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
					//	fmt.Println("dt:   ", Dt)
					maxtauwitht = 1.e-12
				}
			}

		case "time":
			{
				T += Dt
			}
		}

		//	plotswitchtime()//EXTRA
		if Jumpnoise {
			checkallswitch(Universe.lijst)
		}
		//fmt.Println(Dt)
		//write(averages(Universe.lijst))
		write(averagemoments(Universe.lijst),false)
		if (T > j+time-Dt) && (T < j+time) {
			undobadstep(Universe.lijst)
			Dt = j + time - T + 1e-15
		}
	}

	//if suggest_timestep {
	//	printsuggestedtimestep()
	//}
}

//##################################################

//Perform a timestep using euler forward method
func eulerstep(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()
		if Condition_1 { //necessary to have noise field of anisodynamics
			p.randomvfield = p.randomv()
		}
		tau := p.tau(temp)

		if BrownianRotation { //only calculate anisodynamics when requested
			randomv := vector{0., 0., 0.}
			if Condition_1 {
				randomv = p.randomvfield
			} else { //we still need to calculate randomv
				randomv = p.randomv()
				p.randomvfield = randomv
			}

			tau_u := p.tau_u(randomv)
			p.u_anis[0] += tau_u[0] * Dt
			p.u_anis[1] += tau_u[1] * Dt
			p.u_anis[2] += tau_u[2] * Dt
			p.u_anis = norm(p.u_anis)
		}

		if p.fixed==false{
		p.m[0] += tau[0] * Dt
		p.m[1] += tau[1] * Dt
		p.m[2] += tau[2] * Dt
		}
		p.m = norm(p.m)
		//if you have to save mdotH
		p.heff = p.b_eff(temp)
	}
}

//#########################################################################
//perform a timestep using heun method
//http://en.wikipedia.org/wiki/Heun_method
func heunstep(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()
		p.tempfield = temp
		tau1 := p.tau(temp)
		p.fehlk1 = tau1

		//tau (t+1)
		if p.fixed==false{
		p.m[0] += tau1[0] * Dt
		p.m[1] += tau1[1] * Dt
		p.m[2] += tau1[2] * Dt
		}
	}
	T += Dt

	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.tempfield
		//temp := p.temp()
		tau2 := p.tau(temp)
		tau1 := p.fehlk1
		if p.fixed==false{
		p.m[0] += ((-tau1[0] + tau2[0]) * 0.5 * Dt)
		p.m[1] += ((-tau1[1] + tau2[1]) * 0.5 * Dt)
		p.m[2] += ((-tau1[2] + tau2[2]) * 0.5 * Dt)
		}

		p.m = norm(p.m)

		//	if suggest_timestep {
		//		taux := (-tau1[0] + tau2[0]) * 0.5
		//		tauy := (-tau1[1] + tau2[1]) * 0.5
		//		tauz := (-tau1[2] + tau2[2]) * 0.5
		//		torq := math.Sqrt(taux*taux + tauy*tauy + tauz*tauz)
		//		if torq > maxtauwitht {
		//			maxtauwitht = torq
		//		}
		//	}
		//if you have to save mdotH
		p.heff = p.b_eff(temp)
	}
	T -= Dt
}

//#########################################################################

//perform a timestep using 3th order RK
func rk3step(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()
		p.tempfield = temp
		tau0 := p.tau(temp)
		p.fehlk1 = tau0

		//k1
		if p.fixed==false{
		p.m[0] += tau0[0] * 1 / 2. * Dt
		p.m[1] += tau0[1] * 1 / 2. * Dt
		p.m[2] += tau0[2] * 1 / 2. * Dt
		}
	}
	T += 1 / 2. * Dt

	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.tempfield
		k2 := p.tau(temp)
		p.fehlk2 = k2
		k1 := p.fehlk1
		if p.fixed==false{
		p.m[0] += ((-3/2.*k1[0] + 2*k2[0]) * Dt)
		p.m[1] += ((-3/2.*k1[1] + 2*k2[1]) * Dt)
		p.m[2] += ((-3/2.*k1[2] + 2*k2[2]) * Dt)
		}
	}
	T += 1 / 2. * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k3 := p.tau(temp)
		k1 := p.fehlk1
		k2 := p.fehlk2
		if p.fixed==false{
		p.m[0] += ((7/6.*k1[0] - 4/3.*k2[0] + 1/6.*k3[0]) * Dt)
		p.m[1] += ((7/6.*k1[1] - 4/3.*k2[1] + 1/6.*k3[1]) * Dt)
		p.m[2] += ((7/6.*k1[2] - 4/3.*k2[2] + 1/6.*k3[2]) * Dt)
		}

		p.m = norm(p.m)

		//	if suggest_timestep {
		//		taux := (7/6.*k1[0] - 4/3.*k2[0] + 1/6.*k3[0])
		//		tauy := (7/6.*k1[1] - 4/3.*k2[1] + 1/6.*k3[1])
		//		tauz := (7/6.*k1[2] - 4/3.*k2[2] + 1/6.*k3[2])
		//		torq := math.Sqrt(taux*taux + tauy*tauy + tauz*tauz)
		//		if torq > maxtauwitht {
		//			maxtauwitht = torq
		//		}
		//	}
		//if you have to save mdotH
		p.heff = p.b_eff(temp)
	}
	T -= Dt
}

//#########################################################################

//perform a timestep using 3th order anneliessolver
func anneliesstep(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()
		p.tempfield = temp
		tau0 := p.tau(temp)
		p.fehlk1 = tau0

		//k1
		if p.fixed==false{
		p.m[0] += tau0[0] * 1. / 10. * Dt
		p.m[1] += tau0[1] * 1. / 10. * Dt
		p.m[2] += tau0[2] * 1. / 10. * Dt
		}
	}
	T += 1. / 10. * Dt

	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.tempfield
		k2 := p.tau(temp)
		p.fehlk2 = k2
		k1 := p.fehlk1
		if p.fixed==false{
		p.m[0] += (((-1./10.-2189./5746.)*k1[0] + 2310./2873.*k2[0]) * Dt)
		p.m[1] += (((-1./10.-2189./5746.)*k1[1] + 2310./2873.*k2[1]) * Dt)
		p.m[2] += (((-1./10.-2189./5746.)*k1[2] + 2310./2873.*k2[2]) * Dt)
		}
	}
	T += (-1./10. + 11./26.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k3 := p.tau(temp)
		k1 := p.fehlk1
		k2 := p.fehlk2
		if p.fixed==false{
		p.m[0] += (((89./33.+2189./5746.)*k1[0] + (-475./126-2310./2873.)*k2[0] + 2873/1386.*k3[0]) * Dt)
		p.m[1] += (((89./33.+2189./5746.)*k1[1] + (-475./126-2310./2873.)*k2[1] + 2873/1386.*k3[1]) * Dt)
		p.m[2] += (((89./33.+2189./5746.)*k1[2] + (-475./126-2310./2873.)*k2[2] + 2873/1386.*k3[2]) * Dt)
		}

		p.m = norm(p.m)

		//	if suggest_timestep {
		//		taux := (7/6.*k1[0] - 4/3.*k2[0] + 1/6.*k3[0])
		//		tauy := (7/6.*k1[1] - 4/3.*k2[1] + 1/6.*k3[1])
		//		tauz := (7/6.*k1[2] - 4/3.*k2[2] + 1/6.*k3[2])
		//		torq := math.Sqrt(taux*taux + tauy*tauy + tauz*tauz)
		//		if torq > maxtauwitht {
		//			maxtauwitht = torq
		//		}
		//	}
		//if you have to save mdotH
		p.heff = p.b_eff(temp)

	}
	T += (-11/26. + 1.) * Dt
	T -= Dt
}

//#########################################################################

//perform a timestep using 4th order RK
func rk4step(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()
		tau0 := p.tau(temp)
		p.tempfield = temp
		p.fehlk1 = tau0

		//k1
		if p.fixed==false{
		p.m[0] += tau0[0] * 1 / 2. * Dt
		p.m[1] += tau0[1] * 1 / 2. * Dt
		p.m[2] += tau0[2] * 1 / 2. * Dt
		}
	}
	T += 1 / 2. * Dt

	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.tempfield
		k2 := p.tau(temp)
		p.fehlk2 = k2
		k1 := p.fehlk1
		if p.fixed==false{
		p.m[0] += ((-1/2.*k1[0] + 1/2.*k2[0]) * Dt)
		p.m[1] += ((-1/2.*k1[1] + 1/2.*k2[1]) * Dt)
		p.m[2] += ((-1/2.*k1[2] + 1/2.*k2[2]) * Dt)
		}
	}
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k3 := p.tau(temp)
		p.fehlk3 = k3
		k2 := p.fehlk2
		if p.fixed==false{
		p.m[0] += ((-1/2.*k2[0] + 1*k3[0]) * Dt)
		p.m[1] += ((-1/2.*k2[1] + 1*k3[1]) * Dt)
		p.m[2] += ((-1/2.*k2[2] + 1*k3[2]) * Dt)
		}
	}
	T += 1 / 2. * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k4 := p.tau(temp)
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		if p.fixed==false{
		p.m[0] += ((1/6.*k1[0] + 1/3.*k2[0] - 2/3.*k3[0] + 1/6.*k4[0]) * Dt)
		p.m[1] += ((1/6.*k1[1] + 1/3.*k2[1] - 2/3.*k3[1] + 1/6.*k4[1]) * Dt)
		p.m[2] += ((1/6.*k1[2] + 1/3.*k2[2] - 2/3.*k3[2] + 1/6.*k4[2]) * Dt)
		}

		//	if suggest_timestep {
		//		taux := (1/6.*k1[0] + 1/3.*k2[0] - 2/3.*k3[0] + 1/6.*k4[0])
		//		tauy := (1/6.*k1[1] + 1/3.*k2[1] - 2/3.*k3[1] + 1/6.*k4[1])
		//		tauz := (1/6.*k1[2] + 1/3.*k2[2] - 2/3.*k3[2] + 1/6.*k4[2])
		//		torq := math.Sqrt(taux*taux + tauy*tauy + tauz*tauz)

		//		if torq > maxtauwitht {
		//			maxtauwitht = torq
		//		}
		//	}
		p.m = norm(p.m)
		//if you have to save mdotH
		p.heff = p.b_eff(temp)

	}
	T -= Dt
}

//#########################################################################
//perform a timestep using dormand-prince

// Gebruik maken van de FSAL (enkel bij niet-brown noise!!!)

func dopristep(Lijst []*particle) {
		var k1,k2,k3,k4,k5,k6,k7 vector
	for _, p := range Lijst {

		p.tempm = p.m
		p.previousm = p.m
		if Condition_1 { //necessary to have noise field of anisodynamics
			p.randomvfield = p.randomv()
		}

		temp := p.temp()
		if relax==false{
		k1 = p.tau(temp)
		}
		if relax==true{
		k1 = p.noprecess()
		}
		p.tempfield = temp
		p.fehlk1 = k1
		p.dmdt = k1

		if BrownianRotation { //only calculate anisodynamics when requested
			p.tempu_anis = p.u_anis
			p.previousu_anis = p.u_anis
			randomv := vector{0., 0., 0.}
			if Condition_1 {
				randomv = p.randomvfield
			} else { //we still need to calculate randomv
				randomv = p.randomv()
				p.randomvfield = randomv
			}
			k1_u := p.tau_u(randomv)
			p.fehlk1_u = k1_u
			//fmt.Println("k1_u0   ", k1_u[0])
			//fmt.Println("k1_u1   ", k1_u[1])
			//fmt.Println("k1_u2   ", k1_u[2])

			p.u_anis[0] += k1_u[0] * 1 / 5. * Dt
			p.u_anis[1] += k1_u[1] * 1 / 5. * Dt
			p.u_anis[2] += k1_u[2] * 1 / 5. * Dt
		}

		if p.fixed==false{
		p.m[0] += k1[0] * 1 / 5. * Dt
		p.m[1] += k1[1] * 1 / 5. * Dt
		p.m[2] += k1[2] * 1 / 5. * Dt
		}

	}
	T += 1 / 5. * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {

		temp := p.tempfield
		k1 := p.fehlk1
		if relax==false{
		k2 = p.tau(temp)
		}
		if relax==true{
		k2 = p.noprecess()
		}
		p.fehlk2 = k2
		p.dmdt = k2

		if BrownianRotation { //only calculate anisodynamics when requested

			randomv := p.randomvfield
			k1_u := p.fehlk1_u
			k2_u := p.tau_u(randomv)
			p.fehlk2_u = k2_u

			p.u_anis = p.tempu_anis
			p.u_anis[0] += ((3/40.*k1_u[0] + 9/40.*k2_u[0]) * Dt)
			p.u_anis[1] += ((3/40.*k1_u[1] + 9/40.*k2_u[1]) * Dt)
			p.u_anis[2] += ((3/40.*k1_u[2] + 9/40.*k2_u[2]) * Dt)
		}

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((3/40.*k1[0] + 9/40.*k2[0]) * Dt)
		p.m[1] += ((3/40.*k1[1] + 9/40.*k2[1]) * Dt)
		p.m[2] += ((3/40.*k1[2] + 9/40.*k2[2]) * Dt)
		}

	}
	T += 1 / 10. * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		if relax==false{
		k3 = p.tau(temp)
		}
		if relax==true{
		k3 = p.noprecess()
		}

		p.fehlk3 = k3
		p.dmdt = k3

		if BrownianRotation { //only calculate anisodynamics when requested
			randomv := p.randomvfield
			k1_u := p.fehlk1_u
			k2_u := p.fehlk2_u
			k3_u := p.tau_u(randomv)
			p.fehlk3_u = k3_u

			p.u_anis = p.tempu_anis
			p.u_anis[0] += ((44/45.*k1_u[0] - 56/15.*k2_u[0] + 32/9.*k3_u[0]) * Dt)
			p.u_anis[1] += ((44/45.*k1_u[1] - 56/15.*k2_u[1] + 32/9.*k3_u[1]) * Dt)
			p.u_anis[2] += ((44/45.*k1_u[2] - 56/15.*k2_u[2] + 32/9.*k3_u[2]) * Dt)
		}

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((44/45.*k1[0] - 56/15.*k2[0] + 32/9.*k3[0]) * Dt)
		p.m[1] += ((44/45.*k1[1] - 56/15.*k2[1] + 32/9.*k3[1]) * Dt)
		p.m[2] += ((44/45.*k1[2] - 56/15.*k2[2] + 32/9.*k3[2]) * Dt)
		}

	}
	T += 1 / 2. * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		if relax==false{
		k4 = p.tau(temp)
		}
		if relax==true{
		k4 = p.noprecess()
		}
		p.fehlk4 = k4
		p.dmdt = k4

		if BrownianRotation { //only calculate anisodynamics when requested
			randomv := p.randomvfield
			k1_u := p.fehlk1_u
			k2_u := p.fehlk2_u
			k3_u := p.fehlk3_u
			k4_u := p.tau_u(randomv)
			p.fehlk4_u = k4_u

			p.u_anis = p.tempu_anis
			p.u_anis[0] += ((19372/6561.*k1_u[0] - 25360/2187.*k2_u[0] + 64448/6561.*k3_u[0] - 212/729.*k4_u[0]) * Dt)
			p.u_anis[1] += ((19372/6561.*k1_u[1] - 25360/2187.*k2_u[1] + 64448/6561.*k3_u[1] - 212/729.*k4_u[1]) * Dt)
			p.u_anis[2] += ((19372/6561.*k1_u[2] - 25360/2187.*k2_u[2] + 64448/6561.*k3_u[2] - 212/729.*k4_u[2]) * Dt)
		}

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((19372/6561.*k1[0] - 25360/2187.*k2[0] + 64448/6561.*k3[0] - 212/729.*k4[0]) * Dt)
		p.m[1] += ((19372/6561.*k1[1] - 25360/2187.*k2[1] + 64448/6561.*k3[1] - 212/729.*k4[1]) * Dt)
		p.m[2] += ((19372/6561.*k1[2] - 25360/2187.*k2[2] + 64448/6561.*k3[2] - 212/729.*k4[2]) * Dt)
		}

	}
	T += (-4/5. + 8/9.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		if relax==false{
		k5 = p.tau(temp)
		}
		if relax==true{
		k5 = p.noprecess()
		}
		p.fehlk5 = k5
		p.dmdt = k5

		if BrownianRotation { //only calculate anisodynamics when requested
			randomv := p.randomvfield
			k1_u := p.fehlk1_u
			k2_u := p.fehlk2_u
			k3_u := p.fehlk3_u
			k4_u := p.fehlk4_u
			k5_u := p.tau_u(randomv)
			p.fehlk5_u = k5_u

			p.u_anis = p.tempu_anis
			p.u_anis[0] += ((9017/3168.*k1_u[0] - 355/33.*k2_u[0] + 46732/5247.*k3_u[0] + 49/176.*k4_u[0] - 5103/18656.*k5_u[0]) * Dt)
			p.u_anis[1] += ((9017/3168.*k1_u[1] - 355/33.*k2_u[1] + 46732/5247.*k3_u[1] + 49/176.*k4_u[1] - 5103/18656.*k5_u[1]) * Dt)
			p.u_anis[2] += ((9017/3168.*k1_u[2] - 355/33.*k2_u[2] + 46732/5247.*k3_u[2] + 49/176.*k4_u[2] - 5103/18656.*k5_u[2]) * Dt)
		}

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((9017/3168.*k1[0] - 355/33.*k2[0] + 46732/5247.*k3[0] + 49/176.*k4[0] - 5103/18656.*k5[0]) * Dt)
		p.m[1] += ((9017/3168.*k1[1] - 355/33.*k2[1] + 46732/5247.*k3[1] + 49/176.*k4[1] - 5103/18656.*k5[1]) * Dt)
		p.m[2] += ((9017/3168.*k1[2] - 355/33.*k2[2] + 46732/5247.*k3[2] + 49/176.*k4[2] - 5103/18656.*k5[2]) * Dt)
		}

	}
	T += 1 / 9. * Dt
	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		if relax==false{
		k6 = p.tau(temp)
		}
		if relax==true{
		k6 = p.noprecess()
		}

		p.fehlk6 = k6
		p.dmdt = k6

		if BrownianRotation { //only calculate anisodynamics when requested
			randomv := p.randomvfield
			k1_u := p.fehlk1_u
			k2_u := p.fehlk2_u
			k3_u := p.fehlk3_u
			k4_u := p.fehlk4_u
			k5_u := p.fehlk5_u
			k6_u := p.tau_u(randomv)
			p.fehlk6_u = k6_u

			p.u_anis = p.tempu_anis
			p.u_anis[0] += ((35/384.*k1_u[0] + 0.*k2_u[0] + 500/1113.*k3_u[0] + 125/192.*k4_u[0] - 2187/6784.*k5_u[0] + 11/84.*k6_u[0]) * Dt)
			p.u_anis[1] += ((35/384.*k1_u[1] + 0.*k2_u[1] + 500/1113.*k3_u[1] + 125/192.*k4_u[1] - 2187/6784.*k5_u[1] + 11/84.*k6_u[1]) * Dt)
			p.u_anis[2] += ((35/384.*k1_u[2] + 0.*k2_u[2] + 500/1113.*k3_u[2] + 125/192.*k4_u[2] - 2187/6784.*k5_u[2] + 11/84.*k6_u[2]) * Dt)
			//and this is also the fifth order solution
		}

		p.m = p.tempm
		if p.fixed==false{
		torquex:=((35/384.*k1[0] + 0.*k2[0] + 500/1113.*k3[0] + 125/192.*k4[0] - 2187/6784.*k5[0] + 11/84.*k6[0]) * Dt)
		torquey:=((35/384.*k1[1] + 0.*k2[1] + 500/1113.*k3[1] + 125/192.*k4[1] - 2187/6784.*k5[1] + 11/84.*k6[1]) * Dt)
		torquez:=((35/384.*k1[2] + 0.*k2[2] + 500/1113.*k3[2] + 125/192.*k4[2] - 2187/6784.*k5[2] + 11/84.*k6[2]) * Dt)
		p.m[0] += torquex
		p.m[1] += torquey
		p.m[2] += torquez

		//and this is also the fifth order solution
		if relax==true{
		maxtauwitht=math.Sqrt(math.Pow(torquex,2.)+math.Pow(torquey,2.)+math.Pow(torquez,2.))
		}
}

	}
	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.fehlk6
		if relax==false{
		k7 = p.tau(temp)
		}
		if relax==true{
		k7 = p.noprecess()
		}
		p.fehlk7 = k7
		p.dmdt = k7

		if p.fixed==false{
		p.tempm[0] += ((5179/57600.*k1[0] + 0.*k2[0] + 7571/16695.*k3[0] + 393/640.*k4[0] - 92097/339200.*k5[0] + 187/2100.*k6[0] + 1/40.*k7[0]) * Dt)
		p.tempm[1] += ((5179/57600.*k1[1] + 0.*k2[1] + 7571/16695.*k3[1] + 393/640.*k4[1] - 92097/339200.*k5[1] + 187/2100.*k6[1] + 1/40.*k7[1]) * Dt)
		p.tempm[2] += ((5179/57600.*k1[2] + 0.*k2[2] + 7571/16695.*k3[2] + 393/640.*k4[2] - 92097/339200.*k5[2] + 187/2100.*k6[2] + 1/40.*k7[2]) * Dt)
		}
		//and this is also the fourth order solution

		if BrownianRotation { //only calculate anisodynamics when requested
			randomv := p.randomvfield
			k1_u := p.fehlk1_u
			k2_u := p.fehlk2_u
			k3_u := p.fehlk3_u
			k4_u := p.fehlk4_u
			k5_u := p.fehlk5_u
			k6_u := p.fehlk6_u
			k7_u := p.tau_u(randomv)
			p.fehlk7_u = k7_u

			p.tempu_anis[0] += ((5179/57600.*k1_u[0] + 0.*k2_u[0] + 7571/16695.*k3_u[0] + 393/640.*k4_u[0] - 92097/339200.*k5_u[0] + 187/2100.*k6_u[0] + 1/40.*k7_u[0]) * Dt)
			p.tempu_anis[1] += ((5179/57600.*k1_u[1] + 0.*k2_u[1] + 7571/16695.*k3_u[1] + 393/640.*k4_u[1] - 92097/339200.*k5_u[1] + 187/2100.*k6_u[1] + 1/40.*k7_u[1]) * Dt)
			p.tempu_anis[2] += ((5179/57600.*k1_u[2] + 0.*k2_u[2] + 7571/16695.*k3_u[2] + 393/640.*k4_u[2] - 92097/339200.*k5_u[2] + 187/2100.*k6_u[2] + 1/40.*k7_u[2]) * Dt)
			//and this is also the fourth order solution
		}

		p.m = norm(p.m)
		p.tempm = norm(p.tempm)

		//the error is the difference between the two solutions
		error := math.Sqrt(sqr(p.m[0]-p.tempm[0]) + sqr(p.m[1]-p.tempm[1]) + sqr(p.m[2]-p.tempm[2]))

		//fmt.Println("error    :", error)
		if Adaptivestep&&relax==false {
			if error > maxtauwitht {
				maxtauwitht = error
			}
		}
		//if you have to save mdotH
		p.heff = p.b_eff(temp)

		if BrownianRotation { //only calculate anisodynamics when requested
			p.u_anis = norm(p.u_anis)
			p.tempu_anis = norm(p.tempu_anis)

			//the error is the difference between the two solutions
			error := math.Sqrt(sqr(p.u_anis[0]-p.tempu_anis[0]) + sqr(p.u_anis[1]-p.tempu_anis[1]) + sqr(p.u_anis[2]-p.tempu_anis[2]))

			//fmt.Println("error    :", error)
			if Adaptivestep&&relax==false {
				if error > maxtauwitht { //in LLG dynamics already set to maxtauwitht if error is larger
					maxtauwitht = error
				}

			}
		}

	}
	T -= Dt
}

///#########################################################################
//perform a timestep using fehlberg 56 method

func fehl56step(Lijst []*particle) {
	for _, p := range Lijst {
		p.tempm = p.m
		p.previousm = p.m

		temp := p.temp()
		p.tempfield = temp
		k1 := p.tau(temp)
		p.fehlk1 = k1

		if p.fixed==false{
		p.m[0] += k1[0] * 1 / 6. * Dt
		p.m[1] += k1[1] * 1 / 6. * Dt
		p.m[2] += k1[2] * 1 / 6. * Dt
		}
	}
	T += 1 / 6. * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {

		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.tau(temp)
		p.fehlk2 = k2

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((4/75.*k1[0] + 16/75.*k2[0]) * Dt)
		p.m[1] += ((4/75.*k1[1] + 16/75.*k2[1]) * Dt)
		p.m[2] += ((4/75.*k1[2] + 16/75.*k2[2]) * Dt)
		}
	}
	T += (-1/6. + 4/15.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.tau(temp)
		p.fehlk3 = k3

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((5/6.*k1[0] - 8/3.*k2[0] + 5/2.*k3[0]) * Dt)
		p.m[1] += ((5/6.*k1[1] - 8/3.*k2[1] + 5/2.*k3[1]) * Dt)
		p.m[2] += ((5/6.*k1[2] - 8/3.*k2[2] + 5/2.*k3[2]) * Dt)
		}
	}
	T += (-4/15. + 2/3.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.tau(temp)
		p.fehlk4 = k4

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((-8/5.*k1[0] + 144/25.*k2[0] - 4.*k3[0] + 16/25.*k4[0]) * Dt)
		p.m[1] += ((-8/5.*k1[1] + 144/25.*k2[1] - 4.*k3[1] + 16/25.*k4[1]) * Dt)
		p.m[2] += ((-8/5.*k1[2] + 144/25.*k2[2] - 4.*k3[2] + 16/25.*k4[2]) * Dt)
		}
	}
	T += (-2/3. + 4/5.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.tau(temp)
		p.fehlk5 = k5

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((361/320.*k1[0] - 18/5.*k2[0] + 407/128.*k3[0] - 11/80.*k4[0] + 55/128.*k5[0]) * Dt)
		p.m[1] += ((361/320.*k1[1] - 18/5.*k2[1] + 407/128.*k3[1] - 11/80.*k4[1] + 55/128.*k5[1]) * Dt)
		p.m[2] += ((361/320.*k1[2] - 18/5.*k2[2] + 407/128.*k3[2] - 11/80.*k4[2] + 55/128.*k5[2]) * Dt)
		}
	}
	T += 1 / 5. * Dt
	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.tau(temp)
		p.fehlk6 = k6

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((-11/640.*k1[0] + 0.*k2[0] + 11/256.*k3[0] - 11/160.*k4[0] + 11/256.*k5[0] + 0.*k6[0]) * Dt)
		p.m[1] += ((-11/640.*k1[1] + 0.*k2[1] + 11/256.*k3[1] - 11/160.*k4[1] + 11/256.*k5[1] + 0.*k6[1]) * Dt)
		p.m[2] += ((-11/640.*k1[2] + 0.*k2[2] + 11/256.*k3[2] - 11/160.*k4[2] + 11/256.*k5[2] + 0.*k6[2]) * Dt)
		}
	}
	T -= 1 / 1. * Dt
	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.fehlk6
		k7 := p.tau(temp)
		p.fehlk7 = k7

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((93/640.*k1[0] + -18/5.*k2[0] + 803/256.*k3[0] - 11/160.*k4[0] + 99/256.*k5[0] + 0.*k6[0] + 1/1.*k7[0]) * Dt)
		p.m[1] += ((93/640.*k1[1] + -18/5.*k2[1] + 803/256.*k3[1] - 11/160.*k4[1] + 99/256.*k5[1] + 0.*k6[1] + 1/1.*k7[1]) * Dt)
		p.m[2] += ((93/640.*k1[2] + -18/5.*k2[2] + 803/256.*k3[2] - 11/160.*k4[2] + 99/256.*k5[2] + 0.*k6[2] + 1/1.*k7[2]) * Dt)
		}
	}
	T += 1 / 1. * Dt

	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.fehlk6
		k7 := p.fehlk7
		k8 := p.tau(temp)
		p.fehlk8 = k8

		p.m = p.tempm
		if p.fixed==false{
		p.tempm[0] += ((31/384.*k1[0] + 0.*k2[0] + 1125/2816.*k3[0] + 9/32.*k4[0] + 125/768.*k5[0] + 5/66.*k6[0] + 0/1.*k7[0]) * Dt)
		p.tempm[1] += ((31/384.*k1[1] + 0.*k2[1] + 1125/2816.*k3[1] + 9/32.*k4[1] + 125/768.*k5[1] + 5/66.*k6[1] + 0/1.*k7[1]) * Dt)
		p.tempm[2] += ((31/384.*k1[2] + 0.*k2[2] + 1125/2816.*k3[2] + 9/32.*k4[2] + 125/768.*k5[2] + 5/66.*k6[2] + 0/1.*k7[2]) * Dt)
		}
		//fifth order solution

		if p.fixed==false{
		p.m[0] = p.tempm[0] + ((-5/66.*k1[0] + -5/66.*k6[0] + 5/66.*k7[0] + 5/66.*k8[0]) * Dt)
		p.m[1] = p.tempm[1] + ((-5/66.*k1[1] + -5/66.*k6[1] + 5/66.*k7[1] + 5/66.*k8[1]) * Dt)
		p.m[2] = p.tempm[2] + ((-5/66.*k1[2] + -5/66.*k6[2] + 5/66.*k7[2] + 5/66.*k8[2]) * Dt)
		}
		//sixth order solution

		p.m = norm(p.m)
		p.tempm = norm(p.tempm)

		//the error is the difference between the two solutions
		error := math.Sqrt(sqr(p.m[0]-p.tempm[0]) + sqr(p.m[1]-p.tempm[1]) + sqr(p.m[2]-p.tempm[2]))

		//fmt.Println("error    :", error)
		if Adaptivestep {
			if error > maxtauwitht {
				maxtauwitht = error
			}
		}
		//if you have to save mdotH
		p.heff = p.b_eff(temp)

	}
	T -= Dt
}

//###########################################################################################################

//perform a timestep using fehlberg 67 method

func fehl67step(Lijst []*particle) {
	for _, p := range Lijst {
		p.tempm = p.m
		p.previousm = p.m

		temp := p.temp()
		p.tempfield = temp
		k1 := p.tau(temp)
		p.fehlk1 = k1

		if p.fixed==false{
		p.m[0] += k1[0] * 2 / 27. * Dt
		p.m[1] += k1[1] * 2 / 27. * Dt
		p.m[2] += k1[2] * 2 / 27. * Dt
		}
	}
	T += 2 / 27. * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {

		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.tau(temp)
		p.fehlk2 = k2

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((1/36.*k1[0] + 1/12.*k2[0]) * Dt)
		p.m[1] += ((1/36.*k1[1] + 1/12.*k2[1]) * Dt)
		p.m[2] += ((1/36.*k1[2] + 1/12.*k2[2]) * Dt)
		}
	}
	T += (-2/27. + 1/9.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.tau(temp)
		p.fehlk3 = k3

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((1/24.*k1[0] + 0.*k2[0] + 1/8.*k3[0]) * Dt)
		p.m[1] += ((1/24.*k1[1] + 0.*k2[1] + 1/8.*k3[1]) * Dt)
		p.m[2] += ((1/24.*k1[2] + 0.*k2[2] + 1/8.*k3[2]) * Dt)
		}
	}
	T += (-1/9. + 1/6.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.tau(temp)
		p.fehlk4 = k4

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((5/12.*k1[0] + 0.*k2[0] - 25/16.*k3[0] + 25/16.*k4[0]) * Dt)
		p.m[1] += ((5/12.*k1[1] + 0.*k2[1] - 25/16.*k3[1] + 25/16.*k4[1]) * Dt)
		p.m[2] += ((5/12.*k1[2] + 0.*k2[2] - 25/16.*k3[2] + 25/16.*k4[2]) * Dt)
		}
	}
	T += (-1/6. + 5/12.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.tau(temp)
		p.fehlk5 = k5

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((1/20.*k1[0] + 0.*k2[0] + 0.*k3[0] + 1/4.*k4[0] + 1/5.*k5[0]) * Dt)
		p.m[1] += ((1/20.*k1[1] + 0.*k2[1] + 0.*k3[1] + 1/4.*k4[1] + 1/5.*k5[1]) * Dt)
		p.m[2] += ((1/20.*k1[2] + 0.*k2[2] + 0.*k3[2] + 1/4.*k4[2] + 1/5.*k5[2]) * Dt)
		}
	}
	T += (-5/12. + 1/2.) * Dt
	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.tau(temp)
		p.fehlk6 = k6

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((-25/108.*k1[0] + 0.*k2[0] + 0.*k3[0] + 125/108.*k4[0] - 65/27.*k5[0] + 125/54.*k6[0]) * Dt)
		p.m[1] += ((-25/108.*k1[1] + 0.*k2[1] + 0.*k3[1] + 125/108.*k4[1] - 65/27.*k5[1] + 125/54.*k6[1]) * Dt)
		p.m[2] += ((-25/108.*k1[2] + 0.*k2[2] + 0.*k3[2] + 125/108.*k4[2] - 65/27.*k5[2] + 125/54.*k6[2]) * Dt)
		}
	}
	T += (-1/2. + 5/6.) * Dt
	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.fehlk6
		k7 := p.tau(temp)
		p.fehlk7 = k7

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((31/300.*k1[0] + 0.*k2[0] + 0.*k3[0] + 0.*k4[0] + 61/225.*k5[0] - 2/9.*k6[0] + +13/900.*k7[0]) * Dt)
		p.m[1] += ((31/300.*k1[1] + 0.*k2[1] + 0.*k3[1] + 0.*k4[1] + 61/225.*k5[1] - 2/9.*k6[1] + +13/900.*k7[1]) * Dt)
		p.m[2] += ((31/300.*k1[2] + 0.*k2[2] + 0.*k3[2] + 0.*k4[2] + 61/225.*k5[2] - 2/9.*k6[2] + +13/900.*k7[2]) * Dt)
		}
	}
	T += (-5/6. + 1/6.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.fehlk6
		k7 := p.fehlk7
		k8 := p.tau(temp)
		p.fehlk8 = k8

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((2.*k1[0] + 0.*k2[0] + 0.*k3[0] - 53/6.*k4[0] + 704/45.*k5[0] - 107/9.*k6[0] + 67/90.*k7[0] + 3.*k8[0]) * Dt)
		p.m[1] += ((2.*k1[1] + 0.*k2[1] + 0.*k3[1] - 53/6.*k4[1] + 704/45.*k5[1] - 107/9.*k6[1] + 67/90.*k7[1] + 3.*k8[1]) * Dt)
		p.m[2] += ((2.*k1[2] + 0.*k2[2] + 0.*k3[2] - 53/6.*k4[2] + 704/45.*k5[2] - 107/9.*k6[2] + 67/90.*k7[2] + 3.*k8[2]) * Dt)
		}
	}
	T += (-1/6. + 2/3.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.fehlk6
		k7 := p.fehlk7
		k8 := p.fehlk8
		k9 := p.tau(temp)
		p.fehlk9 = k9

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((-91/108.*k1[0] + 0.*k2[0] + 0.*k3[0] + 23/108.*k4[0] - 976/135.*k5[0] + 311/54.*k6[0] - 19/60.*k7[0] + 17/6.*k8[0] - 1/12.*k9[0]) * Dt)
		p.m[1] += ((-91/108.*k1[1] + 0.*k2[1] + 0.*k3[1] + 23/108.*k4[1] - 976/135.*k5[1] + 311/54.*k6[1] - 19/60.*k7[1] + 17/6.*k8[1] - 1/12.*k9[1]) * Dt)
		p.m[2] += ((-91/108.*k1[2] + 0.*k2[2] + 0.*k3[2] + 23/108.*k4[2] - 976/135.*k5[2] + 311/54.*k6[2] - 19/60.*k7[2] + 17/6.*k8[2] - 1/12.*k9[2]) * Dt)
		}
	}
	T += (-1 / 3.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.fehlk6
		k7 := p.fehlk7
		k8 := p.fehlk8
		k9 := p.fehlk9
		k10 := p.tau(temp)
		p.fehlk10 = k10

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((2383/4100.*k1[0] + 0.*k2[0] + 0.*k3[0] - 341/164.*k4[0] + 4496/1025.*k5[0] - 301/82.*k6[0] + 2133/4100.*k7[0] + 45/82.*k8[0] + 45/164.*k9[0] + 18/41.*k10[0]) * Dt)
		p.m[1] += ((2383/4100.*k1[1] + 0.*k2[1] + 0.*k3[1] - 341/164.*k4[1] + 4496/1025.*k5[1] - 301/82.*k6[1] + 2133/4100.*k7[1] + 45/82.*k8[1] + 45/164.*k9[1] + 18/41.*k10[1]) * Dt)
		p.m[2] += ((2383/4100.*k1[2] + 0.*k2[2] + 0.*k3[2] - 341/164.*k4[2] + 4496/1025.*k5[2] - 301/82.*k6[2] + 2133/4100.*k7[2] + 45/82.*k8[2] + 45/164.*k9[2] + 18/41.*k10[2]) * Dt)
		}
	}
	T += (2 / 3.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.fehlk6
		k7 := p.fehlk7
		k8 := p.fehlk8
		k9 := p.fehlk9
		k10 := p.fehlk10
		k11 := p.tau(temp)
		p.fehlk11 = k11

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((3/205.*k1[0] + 0.*k2[0] + 0.*k3[0] + 0.*k4[0] + 0.*k5[0] - 6/41.*k6[0] - 3/205.*k7[0] - 3/41.*k8[0] + 3/41.*k9[0] + 6/41.*k10[0]) * Dt)
		p.m[1] += ((3/205.*k1[1] + 0.*k2[1] + 0.*k3[1] + 0.*k4[1] + 0.*k5[1] - 6/41.*k6[1] - 3/205.*k7[1] - 3/41.*k8[1] + 3/41.*k9[1] + 6/41.*k10[1]) * Dt)
		p.m[2] += ((3/205.*k1[2] + 0.*k2[2] + 0.*k3[2] + 0.*k4[2] + 0.*k5[2] - 6/41.*k6[2] - 3/205.*k7[2] - 3/41.*k8[2] + 3/41.*k9[2] + 6/41.*k10[2]) * Dt)
		}
	}
	T += (-1.) * Dt
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k2 := p.fehlk2
		k3 := p.fehlk3
		k4 := p.fehlk4
		k5 := p.fehlk5
		k6 := p.fehlk6
		k7 := p.fehlk7
		k8 := p.fehlk8
		k9 := p.fehlk9
		k10 := p.fehlk10
		k12 := p.tau(temp)
		p.fehlk12 = k12

		p.m = p.tempm
		if p.fixed==false{
		p.m[0] += ((-1777/4100.*k1[0] + 0.*k2[0] + 0.*k3[0] - 341/164.*k4[0] + 4496/1025.*k5[0] - 289/82.*k6[0] + 2193/4100.*k7[0] + 51/82.*k8[0] + 33/164.*k9[0] + 12/41.*k10[0] + 1.*k12[0]) * Dt)
		p.m[1] += ((-1777/4100.*k1[1] + 0.*k2[1] + 0.*k3[1] - 341/164.*k4[1] + 4496/1025.*k5[1] - 289/82.*k6[1] + 2193/4100.*k7[1] + 51/82.*k8[1] + 33/164.*k9[1] + 12/41.*k10[1] + 1.*k12[1]) * Dt)
		p.m[2] += ((-1777/4100.*k1[2] + 0.*k2[2] + 0.*k3[2] - 341/164.*k4[2] + 4496/1025.*k5[2] - 289/82.*k6[2] + 2193/4100.*k7[2] + 51/82.*k8[2] + 33/164.*k9[2] + 12/41.*k10[2] + 1.*k12[2]) * Dt)
		}
	}
	T += (1.) * Dt
	for _, p := range Lijst {
		temp := p.tempfield
		k1 := p.fehlk1
		k6 := p.fehlk6
		k7 := p.fehlk7
		k8 := p.fehlk8
		k9 := p.fehlk9
		k10 := p.fehlk10
		k11 := p.fehlk11
		k12 := p.fehlk12
		k13 := p.tau(temp)
		p.fehlk13 = k13

		p.m = p.tempm
		if p.fixed==false{
		p.tempm[0] += ((41/840.*k1[0] + 34/105.*k6[0] + 9/35.*k7[0] + 9/35.*k8[0] + 9/280.*k9[0] + 9/280.*k10[0] + 41/840.*k11[0]) * Dt)
		p.tempm[1] += ((41/840.*k1[1] + 34/105.*k6[1] + 9/35.*k7[1] + 9/35.*k8[1] + 9/280.*k9[1] + 9/280.*k10[1] + 41/840.*k11[1]) * Dt)
		p.tempm[2] += ((41/840.*k1[2] + 34/105.*k6[2] + 9/35.*k7[2] + 9/35.*k8[2] + 9/280.*k9[2] + 9/280.*k10[2] + 41/840.*k11[2]) * Dt)
		}
		//sixth order solution

		if p.fixed==false{
		p.m[0] = p.tempm[0] + ((-41/840.*k1[0] - 41/840.*k11[0] + 41/840.*k12[0] + 41/840.*k13[0]) * Dt)
		p.m[1] = p.tempm[1] + ((-41/840.*k1[1] - 41/840.*k11[1] + 41/840.*k12[1] + 41/840.*k13[1]) * Dt)
		p.m[2] = p.tempm[2] + ((-41/840.*k1[2] - 41/840.*k11[2] + 41/840.*k12[2] + 41/840.*k13[2]) * Dt)
		}
		//seventh order solution

		p.m = norm(p.m)
		p.tempm = norm(p.tempm)

		//the error is the difference between the two solutions
		error := math.Sqrt(sqr(p.m[0]-p.tempm[0]) + sqr(p.m[1]-p.tempm[1]) + sqr(p.m[2]-p.tempm[2]))

		//fmt.Println("error    :", error)
		if Adaptivestep {
			if error > maxtauwitht {
				maxtauwitht = error
			}
		}
		//if you have to save mdotH
		p.heff = p.b_eff(temp)

	}
	T -= Dt
}

//###########################################################################################################

func undobadstep(Lijst []*particle) {
	for _, p := range Lijst {
		p.m = p.previousm
	}
	T -= Dt
}

func undobadstep_u_anis(Lijst []*particle) {
	for _, p := range Lijst {
		p.u_anis = p.previousu_anis
	}
	//T -= Dt //do not repeat this! is already done for both!!!
}




