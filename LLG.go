package vinamax

//Calculates the torque working on the magnisation of a Particle
func (p Particle) tau() Vector {
	var m Vector
	m[0] = p.m[0]
	m[1] = p.m[1]
	m[2] = p.m[2]

	mxB := m.cross(p.b_eff())
	mxmxB := m.cross(mxB)

	mxmxB = mxmxB.times(Alpha)
	mxB = mxB.add(mxmxB)
	tau := mxB.times(1 / (1 + Alpha*Alpha))
	return tau.times(-gamma0)
}
