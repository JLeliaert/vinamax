package vinamax

import( 
	"runtime"
) 

//Runs the simulation for a certain time
func Run(time float64) {

	np := 4
	runtime.GOMAXPROCS(np)

	testinput()
	for i := range universe.lijst {
		norm(universe.lijst[i].m)
	}
	write(averages(universe.lijst))
	for j := T; T < j+time; {
		if Demag {
			calculatedemag()
		}

		//een aantal verschillende lijsten maken (met var aantal)
		lijstenmaken(np)

		//var channels maken		
		//binnen elke channel heunstep uitvoeren op een van de sublijsten

		//TODO variabel maken tussen euler en heun
		heunstep(universe.lijst)


		write(averages(universe.lijst))
	}
}

//perform a timestep using euler forward method
func eulerstep(Lijst []*Particle) {
	for _,p  := range Lijst {
		tau := p.tau()
		p.m[0] += tau[0] * Dt
		p.m[1] += tau[1] * Dt
		p.m[2] += tau[2] * Dt
		p.m = norm(p.m)

	}
	T += Dt
}

//perform a timestep using heun method
//http://en.wikipedia.org/wiki/Heun_method
func heunstep(Lijst []*Particle) {
	for _, p := range Lijst {

		tau1 := p.tau()

		//tau van t+1, positie nadat met tau1 al is doorgevoerd
		p.m[0] += tau1[0] * Dt
		p.m[1] += tau1[1] * Dt
		p.m[2] += tau1[2] * Dt

		//FOUT!! DIt moet hetzelfde termische veld zijn als in tau1
		tau2 := p.tau()

		p.m[0] += ((-tau1[0] + tau2[0]) * 0.5 * Dt)
		p.m[1] += ((-tau1[1] + tau2[1]) * 0.5 * Dt)
		p.m[2] += ((-tau1[2] + tau2[2]) * 0.5 * Dt)

		p.m = norm(p.m)

	}
	T += Dt
}
