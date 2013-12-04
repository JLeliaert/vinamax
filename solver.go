package vinamax

//Runs the simulation for a certain time
func Run(time float64) {
	for i := range Lijst {
		norm(Lijst[i].M)
	}
	write(averages(Lijst))
	for j := T; T < j+time; {
		calculatedemag()
		//TODO dit variabel maken tussen euler en heun
		heunstep(Lijst)
		write(averages(Lijst))
	}
}

//perform a timestep using euler forward method
func eulerstep(Lijst Particles) Particles {
	for i := range Lijst {
		Lijst[i].M[0] += Lijst[i].tau()[0] * Dt
		Lijst[i].M[1] += Lijst[i].tau()[1] * Dt
		Lijst[i].M[2] += Lijst[i].tau()[2] * Dt
		Lijst[i].M = norm(Lijst[i].M)

	}
	T += Dt
	return Lijst
}

//perform a timestep using heun method
//http://en.wikipedia.org/wiki/Heun_method
func heunstep(Lijst Particles) Particles {
	for i := range Lijst {
		taux1 := Lijst[i].tau()[0]
		tauy1 := Lijst[i].tau()[1]
		tauz1 := Lijst[i].tau()[2]

		//tau van t+1, positie nadat met tau1 al is doorgevoerd
		Lijst[i].M[0] += taux1 * Dt
		Lijst[i].M[1] += tauy1 * Dt
		Lijst[i].M[2] += tauz1 * Dt

		taux2 := Lijst[i].tau()[0]
		tauy2 := Lijst[i].tau()[1]
		tauz2 := Lijst[i].tau()[2]

		Lijst[i].M[0] += ((-taux1 + taux2) * 0.5 * Dt)
		Lijst[i].M[1] += ((-tauy1 + tauy2) * 0.5 * Dt)
		Lijst[i].M[2] += ((-tauz1 + tauz2) * 0.5 * Dt)

		Lijst[i].M = norm(Lijst[i].M)

	}
	T += Dt
	return Lijst
}
