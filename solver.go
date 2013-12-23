package vinamax

//Runs the simulation for a certain time
func Run(time float64) {
	testinput()
	for i := range universe.lijst {
		norm(universe.lijst[i].m)
	}
	write(averages(universe.lijst))
	for j := T; T < j+time; {
		if Demag{
		calculatedemag()
		}
		//TODO dit variabel maken tussen euler en heun
		heunstep(universe.lijst)
		write(averages(universe.lijst))
	}
}

//perform a timestep using euler forward method
func eulerstep(Lijst []*Particle) {
	for i := range Lijst {
		Lijst[i].m[0] += Lijst[i].tau()[0] * Dt
		Lijst[i].m[1] += Lijst[i].tau()[1] * Dt
		Lijst[i].m[2] += Lijst[i].tau()[2] * Dt
		Lijst[i].m = norm(Lijst[i].m)

	}
	T += Dt
}

//perform a timestep using heun method
//http://en.wikipedia.org/wiki/Heun_method
func heunstep(Lijst []*Particle) {
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
}
