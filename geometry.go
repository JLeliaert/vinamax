package vinamax

import(
	"math/rand"
)

//Adds a single particle 
func Addsingleparticle(x,y,z float64){
	Lijst.Append(Particle{X:x,Y:y,Z:z})
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
func (c Cube) Addparticles(N int){
	for i:= 0; i<N;i++ {
		px := c.x+rand.Float64()*c.s/2
		py := c.y+rand.Float64()*c.s/2
		pz := c.z+rand.Float64()*c.s/2
		Addsingleparticle(px,py,pz)
	}
}	
