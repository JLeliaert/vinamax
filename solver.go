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
		}
	case "heun":
		{
			solver = "heun"
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
