package vinamax

//Calculates the torque working on the magnetisation of a particle
//using the Landau Lifshitz equation
func (p *particle) tau(temp vector) vector {
	pm := &p.m
	mxB := pm.cross(p.b_eff(temp))
	amxmxB := pm.cross(mxB).times(Alpha)
	mxB = mxB.add(amxmxB)
	return mxB.times(-gammaoveralpha)
}
