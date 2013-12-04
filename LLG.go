package vinamax

//Calculates the torque working on the magnisation of a Particle
func (p Particle) tau() [3]float64 {
	var m [3]float64
	m[0] = p.M[0]
	m[1] = p.M[1]
	m[2] = p.M[2]

	mxB := Cross(m, p.B_eff())
	mxmxB := Cross(m, mxB)

	mxmxB = Times(mxmxB, Alpha)
	mxB = Add(mxB, mxmxB)
	tau := Times(mxB, 1/(1+Alpha*Alpha))
	return Times(tau, -Gamma0)
}
