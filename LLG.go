package vinamax

//Calculates the torque working on the magnetisation of a particle
//using the Landau Lifshitz equation
func (p *particle) tau(temp vector) vector {
	pm := &p.m
	//check with Jonathan
	p.heff =p.b_eff(temp)
	// was originally not saved to heff
	mxB := pm.cross(p.heff)
	amxmxB := pm.cross(mxB).times(Alpha)
	mxB = mxB.add(amxmxB)
	return mxB.times(-gammaoveralpha)
}
