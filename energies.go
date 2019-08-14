package vinamax

//calculates the zeeman energy in a particle
func (p *particle) e_zeeman() float64 {
	return -1 * p.msat * p.Volume() * p.m.dot(p.zeeman())
}

//calculates the anisotropy energy in a particle
func (p *particle) e_anis() float64 {
	return -0.5 * p.msat * p.Volume() * p.m.dot(p.anis())
}

//calculates the thermal energy in a particle
func (p *particle) e_therm() float64 {
	return -1 * p.msat * p.Volume() * p.m.dot(p.thermField)
}

//calculates the demag energy in a particle
func (p *particle) e_demag() float64 {
	return -0.5 * p.msat * p.Volume() * p.m.dot(p.demag())
}

//returns total zeeman energy
func E_zeeman() float64 {
	sum := 0.
	for _, p := range lijst {
		sum += p.e_zeeman()
	}
	return sum
}

//returns total anisotropy energy
func E_anis() float64 {
	sum := 0.
	for _, p := range lijst {
		sum += p.e_anis()
	}
	return sum
}

//returns total demag energy
func E_demag() float64 {
	sum := 0.
	for _, p := range lijst {
		sum += p.e_demag()
	}
	return sum
}

//returns total thermal energy
func E_therm() float64 {
	sum := 0.
	for _, p := range lijst {
		sum += p.e_therm()
	}
	return sum
}

//returns total energy in the simulation
func E_total() float64 {
	return E_demag() + E_zeeman() + E_anis() + E_therm()
}
