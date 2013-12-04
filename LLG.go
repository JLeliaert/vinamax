package main

//Calculates the torque working on the magnisation of a particle
func (p particle) tau() [3]float64 {
	var m [3]float64
	m[0] = p.m[0]
	m[1] = p.m[1]
	m[2] = p.m[2]

	mxB := Cross(m, p.B_eff())
	mxmxB := Cross(m, mxB)

	mxmxB = Times(mxmxB, alpha)
	mxB = Add(mxB, mxmxB)
	tau := Times(mxB, 1/(1+alpha*alpha))
	return Times(tau, -Gamma0)
}
