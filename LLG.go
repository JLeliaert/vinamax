package vinamax

//Calculates the torque working on the magnisation of a Particle
func (p Particle) tau() Vector {
	mxB := p.m.cross(p.b_eff())
	mxmxB := p.m.cross(mxB)
	mxmxB = mxmxB.times(Alpha)
	mxB = mxB.add(mxmxB)
	tau := mxB.times(1 / (1 + Alpha*Alpha))
	return tau.times(-gamma0)
}
