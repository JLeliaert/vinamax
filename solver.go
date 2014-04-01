package vinamax

import (
	"fmt"
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

	default:
		{
			log.Fatal(a, " is not a possible solver, \"euler\" or \"heun\" or \"rk3\"or \"rk4\"or \"dopri\"")
		}
	}
}

//Runs the simulation for a certain time

func Run(time float64) {
	testinput()
	syntaxrun()
	for i := range universe.lijst {
		norm(universe.lijst[i].m)
	}
	write(averages(universe.lijst))
	for j := T; T < j+time; {
		if Demag {
			calculatedemag()
		}
		switch solver {
		case "heun":
			{
				heunstep(universe.lijst)
			}
		case "euler":
			{
				eulerstep(universe.lijst)
			}
		case "rk3":
			{
				rk3step(universe.lijst)
			}
		case "rk4":
			{
				rk4step(universe.lijst)
			}
		case "dopri":
			{
				dopristep(universe.lijst)
				if Adaptivestep {

					Dt = Dt * math.Pow(Errortolerance/maxtauwitht, (1./float64(order)))
					fmt.Println("dt:   ", Dt)
					maxtauwitht = 0
				}
			}
		}

		write(averages(universe.lijst))
	}
	if suggest_timestep {
		printsuggestedtimestep()
	}
}

//##################################################
//TODO tijd tussenin aanpassen

//Perform a timestep using euler forward method
func eulerstep(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()

		tau := p.tau(temp)
		p.m[0] += tau[0] * Dt
		p.m[1] += tau[1] * Dt
		p.m[2] += tau[2] * Dt
		T += Dt
		p.m = norm(p.m)
		if suggest_timestep {
			torq := math.Sqrt(tau[0]*tau[0] + tau[1]*tau[1] + tau[2]*tau[2])
			if torq > maxtauwitht {
				maxtauwitht = torq
			}
		}
	}
}

//perform a timestep using heun method
//http://en.wikipedia.org/wiki/Heun_method
func heunstep(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()
		tau1 := p.tau(temp)
		p.tauheun = tau1

		//tau van t+1, positie nadat met tau1 al is doorgevoerd
		p.m[0] += tau1[0] * Dt
		p.m[1] += tau1[1] * Dt
		p.m[2] += tau1[2] * Dt
		T += Dt
	}

	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.temp()
		tau2 := p.tau(temp)
		tau1 := p.tauheun
		p.m[0] += ((-tau1[0] + tau2[0]) * 0.5 * Dt)
		p.m[1] += ((-tau1[1] + tau2[1]) * 0.5 * Dt)
		p.m[2] += ((-tau1[2] + tau2[2]) * 0.5 * Dt)

		p.m = norm(p.m)

		if suggest_timestep {
			taux := (-tau1[0] + tau2[0]) * 0.5
			tauy := (-tau1[1] + tau2[1]) * 0.5
			tauz := (-tau1[2] + tau2[2]) * 0.5
			torq := math.Sqrt(taux*taux + tauy*tauy + tauz*tauz)
			if torq > maxtauwitht {
				maxtauwitht = torq
			}
		}
	}
}

//#########################################################################

//perform a timestep using 3th order RK
func rk3step(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()
		tau0 := p.tau(temp)
		p.taurk3k1 = tau0

		//k1
		p.m[0] += tau0[0] * 1 / 2. * Dt
		p.m[1] += tau0[1] * 1 / 2. * Dt
		p.m[2] += tau0[2] * 1 / 2. * Dt
		T += 1 / 2. * Dt
	}

	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.temp()
		k2 := p.tau(temp)
		p.taurk3k2 = k2
		k1 := p.taurk3k1
		p.m[0] += ((-3/2.*k1[0] + 2*k2[0]) * Dt)
		p.m[1] += ((-3/2.*k1[1] + 2*k2[1]) * Dt)
		p.m[2] += ((-3/2.*k1[2] + 2*k2[2]) * Dt)
		T += 1 / 2. * Dt
	}
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.temp()
		k3 := p.tau(temp)
		k1 := p.taurk3k1
		k2 := p.taurk3k2
		p.m[0] += ((7/6.*k1[0] - 4/3.*k2[0] + 1/6.*k3[0]) * Dt)
		p.m[1] += ((7/6.*k1[1] - 4/3.*k2[1] + 1/6.*k3[1]) * Dt)
		p.m[2] += ((7/6.*k1[2] - 4/3.*k2[2] + 1/6.*k3[2]) * Dt)

		p.m = norm(p.m)

		if suggest_timestep {
			taux := (7/6.*k1[0] - 4/3.*k2[0] + 1/6.*k3[0])
			tauy := (7/6.*k1[1] - 4/3.*k2[1] + 1/6.*k3[1])
			tauz := (7/6.*k1[2] - 4/3.*k2[2] + 1/6.*k3[2])
			torq := math.Sqrt(taux*taux + tauy*tauy + tauz*tauz)
			if torq > maxtauwitht {
				maxtauwitht = torq
			}
		}
	}
}

//#########################################################################

//perform a timestep using 4th order RK
func rk4step(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()
		tau0 := p.tau(temp)
		p.taurk4k1 = tau0

		//k1
		p.m[0] += tau0[0] * 1 / 2. * Dt
		p.m[1] += tau0[1] * 1 / 2. * Dt
		p.m[2] += tau0[2] * 1 / 2. * Dt
		T += 1 / 2. * Dt
	}

	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.temp()
		k2 := p.tau(temp)
		p.taurk4k2 = k2
		k1 := p.taurk4k1
		p.m[0] += ((-1/2.*k1[0] + 1/2.*k2[0]) * Dt)
		p.m[1] += ((-1/2.*k1[1] + 1/2.*k2[1]) * Dt)
		p.m[2] += ((-1/2.*k1[2] + 1/2.*k2[2]) * Dt)
	}
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.temp()
		k3 := p.tau(temp)
		p.taurk4k3 = k3
		k2 := p.taurk4k2
		p.m[0] += ((-1/2.*k2[0] + 1*k3[0]) * Dt)
		p.m[1] += ((-1/2.*k2[1] + 1*k3[1]) * Dt)
		p.m[2] += ((-1/2.*k2[2] + 1*k3[2]) * Dt)
		T += 1 / 2. * Dt
	}
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.temp()
		k4 := p.tau(temp)
		k1 := p.taurk4k1
		k2 := p.taurk4k2
		k3 := p.taurk4k3
		p.m[0] += ((1/6.*k1[0] + 1/3.*k2[0] - 2/3.*k3[0] + 1/6.*k4[0]) * Dt)
		p.m[1] += ((1/6.*k1[1] + 1/3.*k2[1] - 2/3.*k3[1] + 1/6.*k4[1]) * Dt)
		p.m[2] += ((1/6.*k1[2] + 1/3.*k2[2] - 2/3.*k3[2] + 1/6.*k4[2]) * Dt)

		if suggest_timestep {
			taux := (1/6.*k1[0] + 1/3.*k2[0] - 2/3.*k3[0] + 1/6.*k4[0])
			tauy := (1/6.*k1[1] + 1/3.*k2[1] - 2/3.*k3[1] + 1/6.*k4[1])
			tauz := (1/6.*k1[2] + 1/3.*k2[2] - 2/3.*k3[2] + 1/6.*k4[2])
			torq := math.Sqrt(taux*taux + tauy*tauy + tauz*tauz)

			if torq > maxtauwitht {
				maxtauwitht = torq
			}
		}
		p.m = norm(p.m)
	}
}

//#########################################################################
//perform a timestep using dormand-prince

// Gebruik maken van de FSAL (enkel bij niet-brown noise!!!)

func dopristep(Lijst []*particle) {
	for _, p := range Lijst {
		p.tempm = p.m
		fmt.Println(p.m)

		temp := p.temp()
		k1 := p.tau(temp)
		p.doprik1 = k1

		p.m[0] += k1[0] * 1 / 5. * Dt
		p.m[1] += k1[1] * 1 / 5. * Dt
		p.m[2] += k1[2] * 1 / 5. * Dt
		T += 1 / 5. * Dt

	}
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {

		temp := p.temp()
		k1 := p.doprik1
		k2 := p.tau(temp)
		p.doprik2 = k2

		p.m = p.tempm
		p.m[0] += ((3/40.*k1[0] + 9/40.*k2[0]) * Dt)
		p.m[1] += ((3/40.*k1[1] + 9/40.*k2[1]) * Dt)
		p.m[2] += ((3/40.*k1[2] + 9/40.*k2[2]) * Dt)
		T += 1 / 10. * Dt
	}
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.temp()
		k1 := p.doprik1
		k2 := p.doprik2
		k3 := p.tau(temp)
		p.doprik3 = k3

		p.m = p.tempm
		p.m[0] += ((44/45.*k1[0] - 56/15.*k2[0] + 32/9.*k3[0]) * Dt)
		p.m[1] += ((44/45.*k1[1] - 56/15.*k2[1] + 32/9.*k3[1]) * Dt)
		p.m[2] += ((44/45.*k1[2] - 56/15.*k2[2] + 32/9.*k3[2]) * Dt)
		T += 1 / 2. * Dt
	}
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.temp()
		k1 := p.doprik1
		k2 := p.doprik2
		k3 := p.doprik3
		k4 := p.tau(temp)
		p.doprik4 = k4

		p.m = p.tempm
		p.m[0] += ((19372/6561.*k1[0] - 25360/2187.*k2[0] + 64448/6561.*k3[0] - 212/729.*k4[0]) * Dt)
		p.m[1] += ((19372/6561.*k1[1] - 25360/2187.*k2[1] + 64448/6561.*k3[1] - 212/729.*k4[1]) * Dt)
		p.m[2] += ((19372/6561.*k1[2] - 25360/2187.*k2[2] + 64448/6561.*k3[2] - 212/729.*k4[2]) * Dt)
		T += (-4/5. + 8/9.) * Dt
	}
	if Demag {
		calculatedemag()
	}
	for _, p := range Lijst {
		temp := p.temp()
		k1 := p.doprik1
		k2 := p.doprik2
		k3 := p.doprik3
		k4 := p.doprik4
		k5 := p.tau(temp)
		p.doprik5 = k5

		p.m = p.tempm
		p.m[0] += ((9017/3168.*k1[0] - 355/33.*k2[0] + 46732/5247.*k3[0] + 49/176.*k4[0] - 5103/18656.*k5[0]) * Dt)
		p.m[1] += ((9017/3168.*k1[1] - 355/33.*k2[1] + 46732/5247.*k3[1] + 49/176.*k4[1] - 5103/18656.*k5[1]) * Dt)
		p.m[2] += ((9017/3168.*k1[2] - 355/33.*k2[2] + 46732/5247.*k3[2] + 49/176.*k4[2] - 5103/18656.*k5[2]) * Dt)
		T += 1 / 9. * Dt
	}
	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.temp()
		k1 := p.doprik1
		k2 := p.doprik2
		k3 := p.doprik3
		k4 := p.doprik4
		k5 := p.doprik5
		k6 := p.tau(temp)
		p.doprik6 = k6

		p.m = p.tempm
		p.m[0] += ((35/384.*k1[0] + 0.*k2[0] + 500/1113.*k3[0] + 125/192.*k4[0] - 2187/6784.*k5[0] + 11/84.*k6[0]) * Dt)
		p.m[1] += ((35/384.*k1[1] + 0.*k2[1] + 500/1113.*k3[1] + 125/192.*k4[1] - 2187/6784.*k5[1] + 11/84.*k6[1]) * Dt)
		p.m[2] += ((35/384.*k1[2] + 0.*k2[2] + 500/1113.*k3[2] + 125/192.*k4[2] - 2187/6784.*k5[2] + 11/84.*k6[2]) * Dt)
		//and this is also the fifth order solution
	}
	if Demag {
		calculatedemag()
	}

	for _, p := range Lijst {
		temp := p.temp()
		k1 := p.doprik1
		k2 := p.doprik2
		k3 := p.doprik3
		k4 := p.doprik4
		k5 := p.doprik5
		k6 := p.doprik6
		k7 := p.tau(temp)
		p.doprik7 = k7

		p.tempm[0] += ((5179/57600.*k1[0] + 0.*k2[0] + 7571/16695.*k3[0] + 393/640.*k4[0] - 92097/339200.*k5[0] + 187/2100.*k6[0] + 1/40.*k7[0]) * Dt)
		p.tempm[1] += ((5179/57600.*k1[1] + 0.*k2[1] + 7571/16695.*k3[1] + 393/640.*k4[1] - 92097/339200.*k5[1] + 187/2100.*k6[1] + 1/40.*k7[1]) * Dt)
		p.tempm[2] += ((5179/57600.*k1[2] + 0.*k2[2] + 7571/16695.*k3[2] + 393/640.*k4[2] - 92097/339200.*k5[2] + 187/2100.*k6[2] + 1/40.*k7[2]) * Dt)
		//and this is also the fourth order solution
		p.m = norm(p.m)
		p.tempm = norm(p.tempm)

		//the error is the difference between the two solutions
		error := math.Sqrt(sqr(p.m[0]-p.tempm[0]) + sqr(p.m[1]-p.tempm[1]) + sqr(p.m[2]-p.tempm[2]))

		fmt.Println(math.Sqrt(sqr(p.tempm[0]) + sqr(p.tempm[1])))
		fmt.Println(math.Sqrt(sqr(p.m[0]) + sqr(p.m[1])))

		fmt.Println("error    :", error)
		if Adaptivestep || suggest_timestep {
			if error > maxtauwitht {
				maxtauwitht = error
			}
		}
	}
}
