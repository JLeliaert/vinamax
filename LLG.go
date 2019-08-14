package vinamax

import (
//	"fmt"
//"math"
)

//Calculates the torque working on the magnetisation of a particle
//using the Landau Lifshitz equation
func (p *particle) tau(temp vector) vector {
	mxB := vector{0., 0., 0.}
	pm := &p.m
	p.heff = p.b_eff(temp)
	mxB = pm.cross(p.heff)
	amxmxB := pm.cross(mxB).times(Alpha.value)
	mxB = mxB.add(amxmxB).times(-gammaoveralpha)
	return mxB
}

func (p *particle) noprecess() vector {
	mxB := vector{0., 0., 0.}
	pm := &p.m
	p.heff = p.b_eff(vector{0., 0., 0.})
	mxB = pm.cross(p.heff)
	amxmxB := pm.cross(mxB).times(-1 * Alpha.value * gammaoveralpha)

	return amxmxB
}
