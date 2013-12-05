package vinamax

//Runs the simulation for a certain time
func Run(time float64) {
	for i := range Lijst {
		norm(Lijst[i].m)
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
		Lijst[i].m[0] += Lijst[i].tau()[0] * Dt
		Lijst[i].m[1] += Lijst[i].tau()[1] * Dt
		Lijst[i].m[2] += Lijst[i].tau()[2] * Dt
		Lijst[i].m = norm(Lijst[i].m)

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
		Lijst[i].m[0] += taux1 * Dt
		Lijst[i].m[1] += tauy1 * Dt
		Lijst[i].m[2] += tauz1 * Dt

		taux2 := Lijst[i].tau()[0]
		tauy2 := Lijst[i].tau()[1]
		tauz2 := Lijst[i].tau()[2]

		Lijst[i].m[0] += ((-taux1 + taux2) * 0.5 * Dt)
		Lijst[i].m[1] += ((-tauy1 + tauy2) * 0.5 * Dt)
		Lijst[i].m[2] += ((-tauz1 + tauz2) * 0.5 * Dt)

		Lijst[i].m = norm(Lijst[i].m)

	}
	T += Dt
	return Lijst
}
