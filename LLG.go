package vinamax

//Calculates the torque working on the magnetisation of a particle
//using the Landau Lifshitz equation
//in case of relaxing, only return the damping term
func (p *particle) tau() vector {
	mxB := vector{0., 0., 0.}
	pm := &p.m
	if relax {
		p.heff = p.b_eff(vector{0., 0., 0.})
	} else {
		p.heff = p.b_eff(p.thermField)
	}
	mxB = pm.cross(p.heff)
	amxmxB := pm.cross(mxB).times(p.alpha)
	if relax {
		return amxmxB.times(-1 * p.alpha * gamma0 / (1. + (p.alpha * p.alpha)))
	}
	mxB = mxB.add(amxmxB).times(-gamma0 / (1. + (p.alpha * p.alpha)))
	return mxB
}
