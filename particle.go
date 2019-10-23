package vinamax

import (
	"log"
	"math"
)

//A particle essentially constains a position, magnetisation
type particle struct {
	//properties
	x, y, z float64
	m       vector
	u       vector  // Uniaxial anisotropy axis
	rc      float64 // core radius
	rh      float64 // hydrodynamic radius
	msat    float64 // Saturation magnetisation in A/m
	ku1     float64 // uniaxial anisotropy strenght
	alpha   float64 // Gilbert damping constant

	//helper variables
	demagnetising_field vector
	thermPrefac         float64
	thermRotPrefac      float64

	heff          vector //effective field
	thermField    vector
	rotThermField vector
	tempm         vector
	previousm     vector
	torque        vector
	rotTorque     vector
	tempu         vector
	previousu     vector
	k             [7]vector
	k_u           [7]vector
}

//returns the magnetization of the particle
func (p *particle) GetM() vector {
	return p.m
}
func (p *particle) GetU() vector {
	return p.u
}

func (p *particle) SetM(v vector) {
	p.m = norm(v)
}

func (p *particle) Set_msat(v float64) {
	p.msat = v
}

func (p *particle) Set_ku1(v float64) {
	p.ku1 = v
}
func (p *particle) Set_rc(v float64) {
	p.rc = v
}
func (p *particle) Set_rh(v float64) {
	p.rh = v
}

//Adds a single particle at specified coordinates with fixed spin, returns false if unsuccesfull
func addParticle(x, y, z float64) bool {

	if overlap(x, y, z, Rh.value) == true {
		return false
	}

	a := particle{x: x, y: y, z: z, rc: Rc.value, rh: Rh.value, m: M.value, alpha: Alpha.value, msat: Msat.value, u: U_anis.value, ku1: Ku1.value}
	lijst = append(lijst, &a)

	return true
}

func AddParticle(x, y, z float64) {
	if addParticle(x, y, z) == false {
		log.Fatal("Trying to add particle at overlapping locations")
	}
}

//calculates particle volume
func (p *particle) Volume() float64 {
	return 4. / 3. * math.Pi * math.Pow(p.rc, 3.)
}
