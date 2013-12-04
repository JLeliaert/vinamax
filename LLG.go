package vinamax

//Calculates the torque working on the magnisation of a Particle
func (p Particle) tau() Vector {
	var m Vector
	m[0] = p.M[0]
	m[1] = p.M[1]
	m[2] = p.M[2]

	mxB := m.cross(p.b_eff())
	mxmxB := m.cross(mxB)

	mxmxB = mxmxB.times(Alpha)
	mxB = mxB.add(mxmxB)
	tau := mxB.times(1 / (1 + Alpha*Alpha))
	return tau.times(-gamma0)
}
