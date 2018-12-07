package vinamax

import (
	//"fmt"
	"math"
	"math/rand"
)

var anisrng = rand.New(rand.NewSource(0))

//Calculates the torque working on the uniaxial anisotropy axis of a particle
//using the Langevin equation
func (p *particle) tau_u(randomv vector) vector {
	//exit condition 1 and 2
	upart := vector{0., 0., 0.}
	if Condition_1 {
		pdmdt := &p.dmdt
		dmdtxu := pdmdt.cross(p.u_anis).times(mu0 * p.msat * 4. / 3. * math.Pi * cube(p.r)/ (gamma0 * 6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h)))
		hexthdemag := p.demagnetising_field.add(p.zeeman())
		hexthdemagxm := (&hexthdemag).cross(p.m)
		hexthdemagxmxu := (&hexthdemagxm).cross(p.u_anis).times((-1) * p.msat * 4. / 3. * math.Pi * cube(p.r) / (6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h)))
		upart = dmdtxu.add(hexthdemagxmxu)
		//fmt.Println("u0:   ", upart[0])
		//fmt.Println("u1:   ", upart[1])
		//fmt.Println("u2:   ", upart[2])
		//if Test {mdotu := p.m.dot(p.u_anis)
		//	uminm := (p.u_anis.times(mdotu)).add(p.m.times(-1))
		//	upart = uminm.times((-1) * mdotu * (2* Ku1 * 4. / 3. * math.Pi * cube(p.r)) / ((6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h))*(1+(Alpha*Alpha))))
		//	pheff := &p.heff
		//	heffxm := pheff.cross(p.m)
		//	pm := &p.m
		//	mxheffxm := pm.cross(heffxm)
		//	mxheffxmxu := (&mxheffxm).cross(p.u_anis).times(mu0 * Alpha * p.msat * 4. / 3. * math.Pi * cube(p.r)/ ((6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h))*(1+(Alpha*Alpha))))
		//	upart = upart.add(mxheffxmxu)
			//fmt.Println("u0_test:   ", upart[0])
			//fmt.Println("u1_test:   ", upart[1])
			//fmt.Println("u2_test:   ", upart[2])
		//}
	} else { //this occurs when magn dynamics are much slower than rotational dynamics
		if Condition_2 {
			mdotu := p.m.dot(p.u_anis)
			uminm := (p.u_anis.times(mdotu)).add(p.m.times(-1))
			upart = uminm.times((-1) * mdotu * (2* Ku1 * 4. / 3. * math.Pi * cube(p.r)) / ((6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h))*(1+(Alpha*Alpha))))
			pheff := &p.heff
			heffxm := pheff.cross(p.m)
			pm := &p.m
			mxheffxm := pm.cross(heffxm)
			mxheffxmxu := (&mxheffxm).cross(p.u_anis).times(mu0 * Alpha * p.msat * 4. / 3. * math.Pi * cube(p.r)/ ((6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h))*(1+(Alpha*Alpha))))
			upart = upart.add(mxheffxmxu)
			//fmt.Println("u0:   ", upart[0])
			//fmt.Println("u1:   ", upart[1])
			//fmt.Println("u2:   ", upart[2])
		} else { //no conservation of angular momentum (e.g. Reeves' paper)
			mdotu := p.m.dot(p.u_anis)
			uminm := (p.u_anis.times(mdotu)).add(p.m.times(-1))
			upart = uminm.times((-1) * mdotu * (2* Ku1 * 4. / 3. * math.Pi * cube(p.r)) / (6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h)))
			//fmt.Println("u0:   ", upart[0])
			//fmt.Println("u1:   ", upart[1])
			//fmt.Println("u2:   ", upart[2])
		}
	}
	return upart.add(randomv)
	
}

//Set the randomseed for the anisotropy dynamics 
func Setrandomseed_anis(a int64) {
	randomseedcalled_anis = true
	anisrng = rand.New(rand.NewSource(a))
}

func (p *particle) calculaterandomvprefact() {
	p.randomvprefact = math.Sqrt((2. * kb * Temp) / (6. * p.eta * 4. / 3. * math.Pi * cube(p.r_h)))
}

func calculaterandomvprefacts(lijst []*particle) {
	for i := range lijst {
		lijst[i].calculaterandomvprefact()
	}
}

//Calculates the randomness working on the particles' anisotropy axis
func (p *particle) randomv() vector {
	rand_tor := vector{0., 0., 0.}
	if BrownianRotation {
		etax := anisrng.NormFloat64()
		etay := anisrng.NormFloat64()
		etaz := anisrng.NormFloat64()

		rand_tor = vector{etax, etay, etaz}
		rand_tor = rand_tor.times(p.randomvprefact/math.Sqrt(Dt))
	}
	return rand_tor
}
