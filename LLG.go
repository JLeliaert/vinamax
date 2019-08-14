package vinamax

//Calculates the torque working on the magnetisation of a particle
//using the Landau Lifshitz equation
func (p *particle) tau() vector {
	mxB := vector{0., 0., 0.}
	pm := &p.m
	p.heff = p.b_eff(p.thermField)
	mxB = pm.cross(p.heff)
	amxmxB := pm.cross(mxB).times(p.alpha)
	mxB = mxB.add(amxmxB).times(-gamma0 / (1. + (p.alpha * p.alpha)))
	return mxB
}

//Calculates the torque working on the magnetisation of a particle
//using only the damping factor in the Landau Lifshitz equation
func (p *particle) noprecess() vector {
	mxB := vector{0., 0., 0.}
	pm := &p.m
	p.heff = p.b_eff(vector{0., 0., 0.})
	mxB = pm.cross(p.heff)
	amxmxB := pm.cross(mxB).times(-1 * p.alpha * gamma0 / (1. + (p.alpha * p.alpha)))

	return amxmxB
}
