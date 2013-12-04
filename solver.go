package main

//Runs the simulation for a certain time
func run(time float64) {
	for i := range lijst {
		Norm(lijst[i].m)
	}
	Write(averages(lijst))
	for j := t; t < j+time; {
		calculatedemag()
		//TODO dit variabel maken tussen euler en heun
		heunstep(lijst)
		Write(averages(lijst))
	}
}

//perform a timestep using euler forward method
func eulerstep(lijst particles) particles {
	for i := range lijst {
		lijst[i].m[0] += lijst[i].tau()[0] * dt
		lijst[i].m[1] += lijst[i].tau()[1] * dt
		lijst[i].m[2] += lijst[i].tau()[2] * dt
		lijst[i].m = Norm(lijst[i].m)

	}
	t += dt
	return lijst
}

//perform a timestep using heun method
//http://en.wikipedia.org/wiki/Heun_method
func heunstep(lijst particles) particles {
	for i := range lijst {
		taux1 := lijst[i].tau()[0]
		tauy1 := lijst[i].tau()[1]
		tauz1 := lijst[i].tau()[2]

		//tau van t+1, positie nadat met tau1 al is doorgevoerd
		lijst[i].m[0] += taux1 * dt
		lijst[i].m[1] += tauy1 * dt
		lijst[i].m[2] += tauz1 * dt

		taux2 := lijst[i].tau()[0]
		tauy2 := lijst[i].tau()[1]
		tauz2 := lijst[i].tau()[2]

		lijst[i].m[0] += ((-taux1 + taux2) * 0.5 * dt)
		lijst[i].m[1] += ((-tauy1 + tauy2) * 0.5 * dt)
		lijst[i].m[2] += ((-tauz1 + tauz2) * 0.5 * dt)

		lijst[i].m = Norm(lijst[i].m)

	}
	t += dt
	return lijst
}
