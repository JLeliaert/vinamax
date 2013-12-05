package vinamax

import(
//	"math/rand"
)

//Adds a single particle 
func Addsingleparticle(x,y,z float64, m Vector){
	Lijst.Append(Particle{X:x,Y:y,Z:z,M:m})
}

//Deletes all particles
func Reset() {
Lijst =nil
}

type Cube struct{
	x,y,z float64 //position
	s float64 //diameter
	N int //numberofparticles
}


//Adds a lot of particles at random locations in a cubic region
//func (c Cube) Addparticles(N int){
//
//}
