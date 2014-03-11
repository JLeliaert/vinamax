package vinamax

import (
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
	default:
		{
			log.Fatal(a, " is not a possible solver, \"euler\" or \"heun\" ")
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
		if solver == "heun" {
			heunstep(universe.lijst)
		}
		if solver == "euler" {
			eulerstep(universe.lijst)
		}
		if solver == "rk3" {
			rk3step(universe.lijst)
		}
		T += Dt

		write(averages(universe.lijst))
	}
	if suggest_timestep {
		printsuggestedtimestep()
	}
}

//perform a timestep using euler forward method
func eulerstep(Lijst []*particle) {
	for _, p := range Lijst {
		temp := p.temp()

		tau := p.tau(temp)
		p.m[0] += tau[0] * Dt
		p.m[1] += tau[1] * Dt
		p.m[2] += tau[2] * Dt
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
			taux := (7/6*k1[0] - 4/3*k2[0] + 1/6*k3[0])
			tauy := (7/6*k1[1] - 4/3*k2[1] + 1/6*k3[1])
			tauz := (7/6*k1[2] - 4/3*k2[2] + 1/6*k3[2])
			torq := math.Sqrt(taux*taux + tauy*tauy + tauz*tauz)
			if torq > maxtauwitht {
				maxtauwitht = torq
			}
		}
	}
}
