package vinamax

import (
	//"fmt"
	"math"
)

//Calculates the torque working on the magnetisation of a particle
//using the Landau Lifshitz equation
func (p *particle) tau(temp vector) vector {
	mxB := vector{0., 0., 0.}
	if Condition_1 {
		p.heff = p.b_eff(temp)
		// was originally not saved to heff
		Bxm := (&p.heff).cross(p.m)
		pm := &p.m
		amxBxm := pm.cross(Bxm).times(Alpha)
		mxB = Bxm.add(amxBxm)
		hexthdemag := p.demagnetising_field.add(p.zeeman())
		mxhexthdemag := pm.cross(hexthdemag).times((-1) * p.msat * 4. / 3. * math.Pi * cube(p.r) / (6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h)))
		extension := p.randomvfield.add(mxhexthdemag).times(Alpha / gamma0)
		test := Bxm.add(amxBxm)
		//fmt.Println("ext_0:   ", extension[0])
		//fmt.Println("ext_1:   ", extension[1])
		//fmt.Println("ext_2:   ", extension[2])
		//fmt.Println("other_0:   ", test[0])
		//fmt.Println("other_1:   ", test[1])
		//fmt.Println("other_2:   ", test[2])
		mxB = test.add(extension).times(gamma0 / (1 + (Alpha * Alpha) + ((mu0 * Alpha * p.msat * 4. / 3. * math.Pi * cube(p.r)) / (6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h) * gamma0))))
	} else { //this occurs when magn dynamics are much slower than rotational dynamics
		pm := &p.m
		//check with Jonathan
		p.heff = p.b_eff(temp)
		// was originally not saved to heff
		mxB = pm.cross(p.heff)
		amxmxB := pm.cross(mxB).times(Alpha)
		mxB = mxB.add(amxmxB).times(-gammaoveralpha)
	}
	return mxB
}
