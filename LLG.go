package vinamax

//Calculates the torque working on the magnisation of a Particle
func (p *Particle) tau(temp Vector) Vector {
	mxB := p.m.cross(p.b_eff(temp))
	amxmxB := p.m.cross(mxB).times(Alpha)
	mxB = mxB.add(amxmxB)
	return mxB.times(-gamma0 / (1 + Alpha*Alpha))
}
